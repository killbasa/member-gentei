package cmd

import (
	"context"
	"os"
	"strings"

	_ "github.com/lib/pq"
	zlg "github.com/mark-ignacio/zerolog-gcp"
	_ "github.com/mattn/go-sqlite3"
	"github.com/member-gentei/member-gentei/gentei/ent"
	"github.com/member-gentei/member-gentei/gentei/ent/migrate"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const (
	envNameDiscordClientID     = "DISCORD_CLIENT_ID"
	envNameDiscordClientSecret = "DISCORD_CLIENT_SECRET"
	envNameYouTubeClientID     = "YOUTUBE_CLIENT_ID"
	envNameYouTubeClientSecret = "YOUTUBE_CLIENT_SECRET"
)

var (
	flagDBEngine     string
	flagGCPProjectID string
	flagGCPLogID     string
	flagOpenDB       string
	flagVerbose      bool

	flagDiscordClientID     string
	flagDiscordClientSecret string
	flagYouTubeClientID     string
	flagYouTubeClientSecret string
	flagYouTubeRedirectURL  string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gentei",
	Short: "Everything in member-gentei that isn't the frontend.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		flagOpenDB = viper.GetString("db")
		flagDBEngine = viper.GetString("engine")
		flagGCPProjectID = viper.GetString("gcp-project")
		flagGCPLogID = viper.GetString("gcp-log-id")
		flagDiscordClientID = os.Getenv(envNameDiscordClientID)
		flagDiscordClientSecret = os.Getenv(envNameDiscordClientSecret)
		flagYouTubeClientID = os.Getenv(envNameYouTubeClientID)
		flagYouTubeClientSecret = os.Getenv(envNameYouTubeClientSecret)
		gcloWriter, err := zlg.NewCloudLoggingWriter(context.Background(), flagGCPProjectID, flagGCPLogID, zlg.CloudLoggingOptions{})
		if err != nil {
			log.Fatal().Err(err).Msg("error creating zlg.Writer")
		}
		log.Logger = log.Output(zerolog.MultiLevelWriter(
			zerolog.ConsoleWriter{Out: os.Stderr},
			gcloWriter,
		))
		if flagVerbose {
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		} else {
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	defer zlg.Flush()
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
func mustOpenDB(ctx context.Context) *ent.Client {
	logger := log.With().Str("engine", flagDBEngine).Logger()
	db, err := ent.Open(flagDBEngine, flagOpenDB)
	if err != nil {
		logger.Fatal().Err(err).Msg("error opening SQL database")
	}
	if err := db.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		logger.Fatal().Err(err).Msg("failed to create schema resources")
	}
	return db
}

func getYouTubeConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     flagYouTubeClientID,
		ClientSecret: flagYouTubeClientSecret,
		Endpoint:     google.Endpoint,
		RedirectURL:  flagYouTubeRedirectURL,
	}
}

func init() {
	persistent := rootCmd.PersistentFlags()
	persistent.BoolVarP(&flagVerbose, "verbose", "v", false, "debug/verbose logging")
	persistent.String("engine", "sqlite3", "one of: sqlite3, pgx")
	persistent.String("db", "file:ent.sqlite3?cache=shared&_fk=1", "sql connection string")
	persistent.String("gcp-project", "member-gentei", "GCP project ID")
	persistent.String("gcp-log-id", "dev", "GCP log ID")
	persistent.StringVar(&flagYouTubeRedirectURL, "youtube-redirect-url", "http://localhost:3000/login/youtube", "")
	viper.BindPFlags(persistent)
	viper.EnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()
}
