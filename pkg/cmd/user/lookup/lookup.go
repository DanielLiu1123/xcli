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

type lookupOpt struct {
	byIds       []string
	byUsernames []string
	byAuth      bool
}

func NewCmdUserLookup(globalOpt *model.GlobalOpt) *cobra.Command {

	opt := &lookupOpt{}

	c := &cobra.Command{
		Use:   "lookup",
		Short: "Lookup users",
		PreRun: func(cmd *cobra.Command, args []string) {
			validate(opt)
		},
		Run: func(cmd *cobra.Command, args []string) {
			lookupUsers(globalOpt, opt)
		},
		Example: `  # Lookup users by user IDs
  $ xcli user lookup --by-id <user_id>,<user_id>
  
  # Lookup users by usernames
  $ xcli user lookup --by-username <username>,<username>

  # Lookup the authenticated user
  $ xcli user lookup --by-auth
`,
	}

	c.Flags().StringSliceVar(&opt.byIds, "by-id", []string{}, "Lookup users by user IDs")
	c.Flags().StringSliceVar(&opt.byUsernames, "by-username", []string{}, "Lookup users by usernames")
	c.Flags().BoolVar(&opt.byAuth, "by-auth", false, "Lookup the authenticated user")

	return c
}

func validate(opt *lookupOpt) {
	if len(opt.byIds) == 0 && len(opt.byUsernames) == 0 && !opt.byAuth {
		log.Fatal("Must specify at least one of --by-id, --by-username, or --by-auth")
	}
	if len(opt.byIds) > 0 && len(opt.byUsernames) > 0 {
		log.Fatal("Cannot specify both --by-ids and --by-usernames at the same time")
	}
}

func lookupUsers(globalOpt *model.GlobalOpt, opt *lookupOpt) {
	x := client.NewXClient(globalOpt.Auth)

	if len(opt.byIds) > 0 {
		lookupByIds(x, opt.byIds)
	} else if len(opt.byUsernames) > 0 {
		lookupByUsernames(x, opt.byUsernames)
	} else if opt.byAuth {
		lookupByAuth(x)
	}
}

func lookupByAuth(x *twitter.Client) {
	resp, err := x.AuthUserLookup(context.Background(), twitter.UserLookupOpts{})
	if err != nil {
		log.Fatal(err)
	}

	enc, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(enc))
}

func lookupByUsernames(x *twitter.Client, usernames []string) {
	resp, err := x.UserNameLookup(context.Background(), usernames, twitter.UserLookupOpts{})
	if err != nil {
		log.Fatal(err)
	}

	enc, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(enc))
}

func lookupByIds(x *twitter.Client, ids []string) {
	resp, err := x.UserLookup(context.Background(), ids, twitter.UserLookupOpts{})
	if err != nil {
		log.Fatal(err)
	}

	enc, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(enc))
}
