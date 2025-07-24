package tracker

import (
	"github.com/AkhmadeevRus/expense-tracker.git/internal"
	"github.com/spf13/cobra"
)

var month int

func NewSummaryByMonthCmd() *cobra.Command {
	summaryMonthCmd := &cobra.Command{
		Use:   "monthSummary",
		Short: "summary all expenses by month",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunSummaryExpensesByMonthCmd(args)
		},
	}
	summaryMonthCmd.Flags().IntVarP(&month, "monthSummary", "m", 0, "Month for which to show the summary")
	summaryMonthCmd.MarkFlagRequired("monthSummary")
	return summaryMonthCmd
}

func RunSummaryExpensesByMonthCmd(args []string) error {
	return internal.SummaryExpensesByMonth(month)
}
