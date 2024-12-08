package cmd

import (
	"github.com/DanielLiu1123/xcli/pkg/cmd/tweet"
	"github.com/DanielLiu1123/xcli/pkg/cmd/user"
	"github.com/DanielLiu1123/xcli/pkg/model"
	"github.com/spf13/cobra"
	"os"
)

func NewCmdRoot(buildInfo *model.BuildInfo) *cobra.Command {

	opt := &model.GlobalOpt{
		Auth: &model.Auth{},
	}

	c := &cobra.Command{
		Use:     "xcli",
		Version: buildInfo.Version,
		Short:   "X CLI",
		Long:    "Command line for X (Twitter)",
		Example: `  # Create a tweet
  $ xcli tweet create --text "Hello, world"
  
  # Delete tweets
  $ xcli tweet delete <tweet_id> <tweet_id>

  # Lookup tweets
  $ xcli tweet lookup <tweet_id> <tweet_id>

  # Lookup users by ID
  $ xcli user lookup --by-id <user_id>,<user_id>

  # Lookup users by username
  $ xcli user lookup --by-username <username>,<username>
`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			overrideWithEnv(opt)
		},
	}

	c.PersistentFlags().StringVar(&opt.Auth.APIKey, "api-key", "", "API key")
	c.PersistentFlags().StringVar(&opt.Auth.APISecret, "api-secret", "", "API secret")
	c.PersistentFlags().StringVar(&opt.Auth.AccessToken, "access-token", "", "Access token")
	c.PersistentFlags().StringVar(&opt.Auth.AccessSecret, "access-secret", "", "Access secret")

	c.AddCommand(tweet.NewCmdTweet(opt))
	c.AddCommand(user.NewCmdUser(opt))

	return c
}

func overrideWithEnv(opt *model.GlobalOpt) {
	if val := os.Getenv("API_KEY"); val != "" {
		opt.Auth.APIKey = val
	}
	if val := os.Getenv("API_SECRET"); val != "" {
		opt.Auth.APISecret = val
	}
	if val := os.Getenv("ACCESS_TOKEN"); val != "" {
		opt.Auth.AccessToken = val
	}
	if val := os.Getenv("ACCESS_SECRET"); val != "" {
		opt.Auth.AccessSecret = val
	}
}
