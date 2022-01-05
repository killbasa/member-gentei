// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/member-gentei/member-gentei/gentei/ent/guild"
	"github.com/member-gentei/member-gentei/gentei/ent/schema"
)

// Guild is the model entity for the Guild schema.
type Guild struct {
	config `json:"-"`
	// ID of the ent.
	// Discord guild ID
	ID uint64 `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	// Discord guild name
	Name string `json:"name,omitempty"`
	// IconHash holds the value of the "icon_hash" field.
	// Discord guild icon hash
	IconHash string `json:"icon_hash,omitempty"`
	// AuditChannel holds the value of the "audit_channel" field.
	// Audit log channel ID
	AuditChannel uint64 `json:"audit_channel,omitempty"`
	// Language holds the value of the "language" field.
	// IETF BCP 47 language tag
	Language guild.Language `json:"language,omitempty"`
	// AdminSnowflakes holds the value of the "admin_snowflakes" field.
	// Discord snowflakes of users and groups that can modify server settings. The first snowflake is always the server owner.
	AdminSnowflakes []uint64 `json:"admin_snowflakes,omitempty"`
	// ModeratorSnowflakes holds the value of the "moderator_snowflakes" field.
	// Discord snowflakes of users and groups that can read server settings
	ModeratorSnowflakes []uint64 `json:"moderator_snowflakes,omitempty"`
	// Settings holds the value of the "settings" field.
	Settings *schema.GuildSettings `json:"settings,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GuildQuery when eager-loading is set.
	Edges GuildEdges `json:"edges"`
}

// GuildEdges holds the relations/edges for other nodes in the graph.
type GuildEdges struct {
	// Members holds the value of the members edge.
	Members []*User `json:"members,omitempty"`
	// Admins holds the value of the admins edge.
	Admins []*User `json:"admins,omitempty"`
	// Roles holds the value of the roles edge.
	Roles []*GuildRole `json:"roles,omitempty"`
	// YoutubeTalents holds the value of the youtube_talents edge.
	YoutubeTalents []*YouTubeTalent `json:"youtube_talents,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// MembersOrErr returns the Members value or an error if the edge
// was not loaded in eager-loading.
func (e GuildEdges) MembersOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.Members, nil
	}
	return nil, &NotLoadedError{edge: "members"}
}

// AdminsOrErr returns the Admins value or an error if the edge
// was not loaded in eager-loading.
func (e GuildEdges) AdminsOrErr() ([]*User, error) {
	if e.loadedTypes[1] {
		return e.Admins, nil
	}
	return nil, &NotLoadedError{edge: "admins"}
}

// RolesOrErr returns the Roles value or an error if the edge
// was not loaded in eager-loading.
func (e GuildEdges) RolesOrErr() ([]*GuildRole, error) {
	if e.loadedTypes[2] {
		return e.Roles, nil
	}
	return nil, &NotLoadedError{edge: "roles"}
}

// YoutubeTalentsOrErr returns the YoutubeTalents value or an error if the edge
// was not loaded in eager-loading.
func (e GuildEdges) YoutubeTalentsOrErr() ([]*YouTubeTalent, error) {
	if e.loadedTypes[3] {
		return e.YoutubeTalents, nil
	}
	return nil, &NotLoadedError{edge: "youtube_talents"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Guild) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case guild.FieldAdminSnowflakes, guild.FieldModeratorSnowflakes, guild.FieldSettings:
			values[i] = new([]byte)
		case guild.FieldID, guild.FieldAuditChannel:
			values[i] = new(sql.NullInt64)
		case guild.FieldName, guild.FieldIconHash, guild.FieldLanguage:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Guild", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Guild fields.
func (gu *Guild) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case guild.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			gu.ID = uint64(value.Int64)
		case guild.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				gu.Name = value.String
			}
		case guild.FieldIconHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field icon_hash", values[i])
			} else if value.Valid {
				gu.IconHash = value.String
			}
		case guild.FieldAuditChannel:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field audit_channel", values[i])
			} else if value.Valid {
				gu.AuditChannel = uint64(value.Int64)
			}
		case guild.FieldLanguage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field language", values[i])
			} else if value.Valid {
				gu.Language = guild.Language(value.String)
			}
		case guild.FieldAdminSnowflakes:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field admin_snowflakes", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &gu.AdminSnowflakes); err != nil {
					return fmt.Errorf("unmarshal field admin_snowflakes: %w", err)
				}
			}
		case guild.FieldModeratorSnowflakes:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field moderator_snowflakes", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &gu.ModeratorSnowflakes); err != nil {
					return fmt.Errorf("unmarshal field moderator_snowflakes: %w", err)
				}
			}
		case guild.FieldSettings:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field settings", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &gu.Settings); err != nil {
					return fmt.Errorf("unmarshal field settings: %w", err)
				}
			}
		}
	}
	return nil
}

// QueryMembers queries the "members" edge of the Guild entity.
func (gu *Guild) QueryMembers() *UserQuery {
	return (&GuildClient{config: gu.config}).QueryMembers(gu)
}

// QueryAdmins queries the "admins" edge of the Guild entity.
func (gu *Guild) QueryAdmins() *UserQuery {
	return (&GuildClient{config: gu.config}).QueryAdmins(gu)
}

// QueryRoles queries the "roles" edge of the Guild entity.
func (gu *Guild) QueryRoles() *GuildRoleQuery {
	return (&GuildClient{config: gu.config}).QueryRoles(gu)
}

// QueryYoutubeTalents queries the "youtube_talents" edge of the Guild entity.
func (gu *Guild) QueryYoutubeTalents() *YouTubeTalentQuery {
	return (&GuildClient{config: gu.config}).QueryYoutubeTalents(gu)
}

// Update returns a builder for updating this Guild.
// Note that you need to call Guild.Unwrap() before calling this method if this Guild
// was returned from a transaction, and the transaction was committed or rolled back.
func (gu *Guild) Update() *GuildUpdateOne {
	return (&GuildClient{config: gu.config}).UpdateOne(gu)
}

// Unwrap unwraps the Guild entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gu *Guild) Unwrap() *Guild {
	tx, ok := gu.config.driver.(*txDriver)
	if !ok {
		panic("ent: Guild is not a transactional entity")
	}
	gu.config.driver = tx.drv
	return gu
}

// String implements the fmt.Stringer.
func (gu *Guild) String() string {
	var builder strings.Builder
	builder.WriteString("Guild(")
	builder.WriteString(fmt.Sprintf("id=%v", gu.ID))
	builder.WriteString(", name=")
	builder.WriteString(gu.Name)
	builder.WriteString(", icon_hash=")
	builder.WriteString(gu.IconHash)
	builder.WriteString(", audit_channel=")
	builder.WriteString(fmt.Sprintf("%v", gu.AuditChannel))
	builder.WriteString(", language=")
	builder.WriteString(fmt.Sprintf("%v", gu.Language))
	builder.WriteString(", admin_snowflakes=")
	builder.WriteString(fmt.Sprintf("%v", gu.AdminSnowflakes))
	builder.WriteString(", moderator_snowflakes=")
	builder.WriteString(fmt.Sprintf("%v", gu.ModeratorSnowflakes))
	builder.WriteString(", settings=")
	builder.WriteString(fmt.Sprintf("%v", gu.Settings))
	builder.WriteByte(')')
	return builder.String()
}

// Guilds is a parsable slice of Guild.
type Guilds []*Guild

func (gu Guilds) config(cfg config) {
	for _i := range gu {
		gu[_i].config = cfg
	}
}
