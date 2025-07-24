package tracker

import (
	"fmt"

	"github.com/AkhmadeevRus/expense-tracker.git/internal"
	"github.com/spf13/cobra"
)

var (
	addDescription string
	addAmount      float64
)

func NewAddCmd() *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "add",
		Short: "add new expense",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunAddExpenseCmd(args)
		},
	}

	addCmd.Flags().StringVarP(&addDescription, "description", "d", "", "description of the expense(required)")
	addCmd.MarkFlagRequired("description")
	addCmd.Flags().Float64VarP(&addAmount, "amount", "a", 0, "amount of the expense(required)")
	addCmd.MarkFlagRequired("amount")
	return addCmd
}

func RunAddExpenseCmd(args []string) error {
	if addAmount < 0 {
		return fmt.Errorf("amount can't be negative")
	}
	return internal.AddExpense(addAmount, addDescription)
}
