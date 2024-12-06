package tweet

import (
	"github.com/DanielLiu1123/xcli/pkg/cmd/tweet/create"
	"github.com/DanielLiu1123/xcli/pkg/cmd/tweet/delete"
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

	return c
}
