package tracker

import (
	"github.com/AkhmadeevRus/expense-tracker.git/internal"
	"github.com/spf13/cobra"
)

var (
	updateExpenseId   int64
	updateDescription string
	updateAmount      float64
)

func NewUpdateCmd() *cobra.Command {
	updateCmd := cobra.Command{
		Use:   "update",
		Short: "update expense",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunUpdateExpenseCmd(args)
		},
	}
	updateCmd.Flags().Int64VarP(&updateExpenseId, "id", "i", 0, "id for update")
	updateCmd.MarkFlagRequired("id")
	updateCmd.Flags().StringVarP(&updateDescription, "description", "d", "", "description for update")
	updateCmd.Flags().Float64VarP(&updateAmount, "amount", "a", 0., "amount for update")
	return &updateCmd
}

func RunUpdateExpenseCmd(args []string) error {
	return internal.UpdateExpense(updateExpenseId, updateDescription, updateAmount)
}
