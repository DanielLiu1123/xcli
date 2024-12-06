package delete

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

func NewCmdTweetDelete(globalOpt *model.GlobalOpt) *cobra.Command {

	c := &cobra.Command{
		Use:   "delete",
		Short: "Delete tweets",
		Run: func(cmd *cobra.Command, args []string) {
			deleteTweets(globalOpt, args)
		},
	}

	return c
}

func deleteTweets(globalOpt *model.GlobalOpt, ids []string) {
	x := client.NewXClient(globalOpt.Auth)

	for _, id := range ids {
		deleteTweet(x, id)
	}
}

func deleteTweet(x *twitter.Client, id string) {
	tweetResponse, err := x.DeleteTweet(context.Background(), id)
	if err != nil {
		log.Fatal(err)
	}

	enc, err := json.MarshalIndent(tweetResponse, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(enc))
}
