package user

import (
	"github.com/DanielLiu1123/xcli/pkg/cmd/user/lookup"
	"github.com/DanielLiu1123/xcli/pkg/model"
	"github.com/spf13/cobra"
)

func NewCmdUser(globalOpt *model.GlobalOpt) *cobra.Command {

	c := &cobra.Command{
		Use:   "user",
		Short: "Manage users",
		Example: `  # Lookup users by ID
  $ xcli user lookup --by-id <user_id>,<user_id>

  # Lookup users by username
  $ xcli user lookup --by-username <username>,<username>
`,
	}

	c.AddCommand(lookup.NewCmdUserLookup(globalOpt))

	return c
}
