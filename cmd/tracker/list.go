package tracker

import (
	"github.com/AkhmadeevRus/expense-tracker.git/internal"
	"github.com/spf13/cobra"
)

func NewListCmd() *cobra.Command {
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "list all expenses",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunListExpensesCmd(args)
		},
	}
	return listCmd
}

func RunListExpensesCmd(args []string) error {
	return internal.ListExpenses()
}
