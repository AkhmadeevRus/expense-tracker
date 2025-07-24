package tracker

import (
	"github.com/AkhmadeevRus/expense-tracker.git/internal"
	"github.com/spf13/cobra"
)

func NewSummaryCmd() *cobra.Command {
	summaryCmd := &cobra.Command{
		Use:   "summary",
		Short: "summary all expenses",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunSummaryExpensesCmd(args)
		},
	}
	return summaryCmd
}

func RunSummaryExpensesCmd(args []string) error {
	return internal.SummaryExpenses()
}
