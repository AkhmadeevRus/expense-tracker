package tracker

import "github.com/spf13/cobra"

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "expense-tracker",
	}

	cmd.AddCommand(NewAddCmd())
	cmd.AddCommand(NewDeleteCmd())
	cmd.AddCommand(NewUpdateCmd())
	cmd.AddCommand(NewListCmd())
	cmd.AddCommand(NewSummaryCmd())
	cmd.AddCommand(NewSummaryByMonthCmd())

	return cmd
}
