package tracker

import (
	"github.com/AkhmadeevRus/expense-tracker.git/internal"
	"github.com/spf13/cobra"
)

var month int

func NewSummaryCmd() *cobra.Command {
	summaryCmd := &cobra.Command{
		Use:   "summary",
		Short: "summary expenses by month",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunSummaryExpensesCmd(args)
		},
	}

	summaryCmd.Flags().IntVarP(&month, "month", "m", 0, "summary by month")
	return summaryCmd
}

func RunSummaryExpensesCmd(args []string) error {
	return internal.SummaryExpensesByMonth(month)
}
