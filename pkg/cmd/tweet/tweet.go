package tweet

import (
	"github.com/DanielLiu1123/xcli/pkg/cmd/tweet/create"
	"github.com/DanielLiu1123/xcli/pkg/cmd/tweet/delete"
	"github.com/DanielLiu1123/xcli/pkg/cmd/tweet/lookup"
	"github.com/DanielLiu1123/xcli/pkg/model"
	"github.com/spf13/cobra"
)

func NewCmdTweet(globalOpt *model.GlobalOpt) *cobra.Command {

	c := &cobra.Command{
		Use:   "tweet",
		Short: "Manage tweets",
	}

	c.AddCommand(create.NewCmdTweetCreate(globalOpt))
	c.AddCommand(delete.NewCmdTweetDelete(globalOpt))
	c.AddCommand(lookup.NewCmdTweetLookup(globalOpt))

	return c
}
