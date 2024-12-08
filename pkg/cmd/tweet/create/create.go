package create

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/DanielLiu1123/xcli/pkg/client"
	"github.com/DanielLiu1123/xcli/pkg/model"
	"github.com/g8rswimmer/go-twitter/v2"
	"github.com/spf13/cobra"
	"log"
)

type createOpt struct {
	Text                  string
	ForSuperFollowersOnly bool
}

func NewCmdTweetCreate(globalOpt *model.GlobalOpt) *cobra.Command {
	opt := &createOpt{}

	c := &cobra.Command{
		Use:   "create",
		Short: "Create a tweet",
		Run: func(cmd *cobra.Command, args []string) {
			createTweet(globalOpt, opt)
		},
		Example: `  # Create a tweet
  $ xcli tweet create --text "Hello, world"
`,
	}

	c.Flags().StringVar(&opt.Text, "text", "", "Text of the tweet")
	_ = c.MarkFlagRequired("text")
	c.Flags().BoolVar(&opt.ForSuperFollowersOnly, "for-super-followers-only", false, "For super followers only")

	return c
}

func createTweet(globalOpt *model.GlobalOpt, opt *createOpt) {
	x := client.NewXClient(globalOpt.Auth)

	req := twitter.CreateTweetRequest{
		Text:                  opt.Text,
		ForSuperFollowersOnly: opt.ForSuperFollowersOnly,
	}
	tweetResponse, err := x.CreateTweet(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	enc, err := json.MarshalIndent(tweetResponse, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(enc))
}
