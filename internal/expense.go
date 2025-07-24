package internal

import (
	"fmt"
	"strconv"
	"time"
)

type Expense struct {
	ID          int64     `json:"id"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	Category    string    `json:"category"`
	CreatedAt   time.Time `json:"createAt"`
	UpdateAt    time.Time `json:"UpdateAt"`
}

func NewExpense(id int64, description string, amount float64, category string) *Expense {
	return &Expense{
		ID:          id,
		Description: description,
		Amount:      amount,
		Category:    category,
		CreatedAt:   time.Now(),
		UpdateAt:    time.Now(),
	}
}

func AddExpense(amount float64, description string, category string) error {
	expenses, err := ReadExpenseFromFile()
	if err != nil {
		return err
	}

	var newExpenseId int64
	if len(expenses) > 0 {
		newExpenseId = expenses[len(expenses)-1].ID + 1
	} else {
		newExpenseId = 1
	}

	newExpense := NewExpense(newExpenseId, description, amount, category)
	expenses = append(expenses, *newExpense)
	return WriteExpenseToFile(expenses)
}

func DeleteExpense(id int64) error {
	expenses, err := ReadExpenseFromFile()
	if err != nil {
		return err
	}

	var found bool
	for _, expense := range expenses {
		if expense.ID == id {
			found = true
			expenses = append(expenses[:id], expenses[id+1:]...)
			break
		}
	}

	if !found {
		return fmt.Errorf("expense was not found with id:%d", id)
	}
	return WriteExpenseToFile(expenses)
}

func ListExpenses(category string) error {
	expenses, err := ReadExpenseFromFile()
	if err != nil {
		return err
	}

	if len(expenses) == 0 {
		fmt.Println("expenses not found")
		return nil
	}
	if category == "" {
		for _, expense := range expenses {
			fmt.Printf("\nid:%s  %s  %.f$\n\n", strconv.FormatInt(expense.ID, 10), expense.Description, expense.Amount)
		}
	} else {
		for _, expense := range expenses {
			if expense.Category == category {
				fmt.Printf("\nid:%s  %s  %.f$ categoty:%s\n\n", strconv.FormatInt(expense.ID, 10), expense.Description, expense.Amount, expense.Category)
			}
		}
	}
	return nil
}

func UpdateExpense(id int64, description string, amount float64, category string) error {
	expenses, err := ReadExpenseFromFile()
	if err != nil {
		return err
	}

	var found bool
	var updateExpenses []Expense
	for _, expense := range expenses {
		if expense.ID == id {
			found = true
			expense.Description = description
			expense.Amount = amount
			expense.Category = category
			expense.UpdateAt = time.Now()
		}
		updateExpenses = append(updateExpenses, expense)
	}

	if !found {
		fmt.Printf("\nexpense by id:%d not found\n\n", id)
		return nil
	}

	return WriteExpenseToFile(updateExpenses)
}

func SummaryExpensesByMonth(month int) error {
	expenses, err := ReadExpenseFromFile()
	if err != nil {
		return err
	}

	if len(expenses) == 0 {
		fmt.Println("expenses not found")
		return nil
	}

	var sumAmountByMonth float64

	if month == 0 {
		for _, expense := range expenses {
			sumAmountByMonth += expense.Amount
		}
	} else {
		for _, expense := range expenses {
			if int(expense.CreatedAt.Month()) == month {
				sumAmountByMonth += expense.Amount
			}
		}
	}
	fmt.Printf("\ntotal expenses for the %dth month: %.f$\n\n", month, sumAmountByMonth)
	return nil
}
