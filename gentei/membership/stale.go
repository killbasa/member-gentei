package membership

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/member-gentei/member-gentei/gentei/apis"
	"github.com/member-gentei/member-gentei/gentei/ent"
	"github.com/member-gentei/member-gentei/gentei/ent/guild"
	"github.com/member-gentei/member-gentei/gentei/ent/user"
	"github.com/rs/zerolog/log"
	"golang.org/x/oauth2"
)

type CheckStaleOptions struct {
	// StaleThreshold is used in a <= comparison to the last stored membership check time.
	StaleThreshold time.Duration
	// MembershipChangeHook gets called whnen a user experiences a change in channel membership.
	MembershipChangeHook func(userID uint64, results *CheckResultSet) error
}

var DefaultCheckStaleOptions = &CheckStaleOptions{
	StaleThreshold: time.Hour * 12,
}

func CheckStale(ctx context.Context, db *ent.Client, youtubeConfig *oauth2.Config, options *CheckStaleOptions) error {
	if options == nil {
		options = DefaultCheckStaleOptions
	}
	staleThreshold := options.StaleThreshold
	if options.StaleThreshold > 0 {
		staleThreshold *= -1
	}
	for {
		staleUserIDs, err := db.User.Query().
			Where(
				user.HasGuildsWith(
					guild.HasYoutubeTalents(),
				),
				user.LastCheckLTE(time.Now().Add(staleThreshold)),
			).
			Limit(1000).
			IDs(ctx)
		if err != nil {
			return err
		}
		if len(staleUserIDs) == 0 {
			break
		}
		for _, userID := range staleUserIDs {
			// TODO: https://github.com/member-gentei/member-gentei/issues/92
			results, err := CheckForUser(ctx, db, youtubeConfig, userID, nil)
			if err != nil {
				return fmt.Errorf("error checking memberships for user '%d': %w", userID, err)
			}
			if options.MembershipChangeHook != nil && (len(results.Lost) > 0 || len(results.Gained) > 0) {
				err = options.MembershipChangeHook(userID, results)
				if err != nil {
					return fmt.Errorf("error calling MembershipChangeHook for user '%d': %w", userID, err)
				}
			}
			err = db.User.UpdateOneID(userID).
				SetLastCheck(time.Now()).
				Exec(ctx)
			if err != nil {
				return fmt.Errorf("error saving LastCheck for user '%d': %w", userID, err)
			}
		}
	}
	return nil
}

// RefreshAllUserGuildEdges refreshes guild edges for all registered users. Returns a slice of userIDs that could not be refreshed.
func RefreshAllUserGuildEdges(ctx context.Context, db *ent.Client, discordConfig *oauth2.Config) error {
	// refresh everyone's tokens
	var (
		userTokensInvalid []uint64
		after             uint64
	)
	const pageSize = 1000
	for {
		userIDs, err := db.User.Query().
			Where(
				user.IDGT(after),
			).
			Limit(pageSize).
			IDs(ctx)
		if err != nil {
			return fmt.Errorf("error paginating user IDs: %w", err)
		}
		for _, userID := range userIDs {
			logger := log.With().Str("userID", strconv.FormatUint(userID, 10)).Logger()
			ts, err := apis.GetRefreshingDiscordTokenSource(ctx, db, discordConfig, userID)
			if err != nil {
				logger.Warn().Err(err).Msg("error creating Discord TokenSource, skipping")
				continue
			}
			token, err := ts.Token()
			if err != nil {
				logger.Warn().Err(err).Msg("error getting Discord token for user")
				// TODO: enforce and delete
				userTokensInvalid = append(userTokensInvalid, userID)
				continue
			}
			added, removed, err := RefreshUserGuildEdges(ctx, db, token, userID)
			if err != nil {
				logger.Err(err).Msg("error refreshing guilds for user")
				return err
			}
			if len(added)+len(removed) > 0 {
				logger.Info().
					Strs("addedGuildIDs", uints64ToStrs(added)).
					Strs("removedGuildIDs", uints64ToStrs(removed)).
					Msg("refreshed with changes")
			}
		}
		if len(userIDs) < pageSize {
			break
		}
	}
	if len(userTokensInvalid) > 0 {
		log.Info().Int("count", len(userTokensInvalid)).
			Msg("failed to refresh some Discord tokens")
	}
	return nil
}

// Refreshes guilds for all registered users.
func RefreshUserGuildEdges(
	ctx context.Context,
	db *ent.Client,
	token *oauth2.Token,
	userID uint64,
) (added []uint64, removed []uint64, err error) {
	svc, err := discordgo.New(fmt.Sprintf("Bearer %s", token.AccessToken))
	if err != nil {
		err = fmt.Errorf("error creating discordgo.Session: %w", err)
		return
	}
	userGuilds, err := svc.UserGuilds(0, "", "")
	if err != nil {
		err = fmt.Errorf("error getting UserGuilds: %w", err)
		return
	}
	guildIDs := make([]uint64, len(userGuilds))
	for i, dg := range userGuilds {
		guildID, convErr := strconv.ParseUint(dg.ID, 10, 64)
		if convErr != nil {
			err = convErr
			return
		}
		guildIDs[i] = guildID
	}
	// add guilds
	addGuildIDs, err := db.Guild.Query().
		Where(
			guild.IDIn(guildIDs...),
			guild.Not(
				guild.HasMembersWith(user.ID(userID)),
			),
		).
		IDs(ctx)
	if err != nil {
		err = fmt.Errorf("error getting Guilds to add: %w", err)
		return
	}
	// remove guilds
	removeGuildIDs, err := db.Guild.Query().
		Where(
			guild.IDNotIn(guildIDs...),
			guild.HasMembersWith(user.ID(userID)),
		).
		IDs(ctx)
	if err != nil {
		err = fmt.Errorf("error getting Guilds to remove: %w", err)
		return
	}
	// actually do it
	err = db.User.UpdateOneID(userID).
		AddGuildIDs(addGuildIDs...).
		RemoveGuildIDs(removeGuildIDs...).
		Exec(ctx)
	return addGuildIDs, removeGuildIDs, err
}

func uints64ToStrs(input []uint64) []string {
	output := make([]string, len(input))
	for i, n := range input {
		output[i] = strconv.FormatUint(n, 10)
	}
	return output
}
