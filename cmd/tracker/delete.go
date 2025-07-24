package tracker

import (
	"github.com/AkhmadeevRus/expense-tracker.git/internal"
	"github.com/spf13/cobra"
)

var deleteExpenseId int64

func NewDeleteCmd() *cobra.Command {
	deleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "delete an expense",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunDeleteExpenseCmd(args)
		},
	}
	deleteCmd.Flags().Int64VarP(&deleteExpenseId, "id", "i", 0, "id of the expense to delete")
	return deleteCmd
}

func RunDeleteExpenseCmd(args []string) error {
	return internal.DeleteExpense(deleteExpenseId)
}
