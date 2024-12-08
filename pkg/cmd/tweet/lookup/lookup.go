package lookup

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

func NewCmdTweetLookup(globalOpt *model.GlobalOpt) *cobra.Command {

	c := &cobra.Command{
		Use:   "lookup",
		Short: "Lookup tweets",
		PreRun: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				log.Fatal("Please provide at least one tweet ID")
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			lookupTweets(globalOpt, args)
		},
		Example: `  # Lookup tweets
  $ xcli tweet lookup <tweet_id> <tweet_id>
`,
	}

	return c
}

func lookupTweets(globalOpt *model.GlobalOpt, ids []string) {
	x := client.NewXClient(globalOpt.Auth)

	resp, err := x.TweetLookup(context.Background(), ids, twitter.TweetLookupOpts{})
	if err != nil {
		log.Fatal(err)
	}

	enc, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(enc))
}
