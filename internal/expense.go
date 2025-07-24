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
	CreatedAt   time.Time `json:"createAt"`
	UpdateAt    time.Time `json:"UpdateAt"`
}

func NewExpense(id int64, description string, amount float64) *Expense {
	return &Expense{
		ID:          id,
		Description: description,
		Amount:      amount,
		CreatedAt:   time.Now(),
		UpdateAt:    time.Now(),
	}
}

func AddExpense(amount float64, description string) error {
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

	newExpense := NewExpense(newExpenseId, description, amount)
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

func ListExpenses() error {
	expenses, err := ReadExpenseFromFile()
	if err != nil {
		return err
	}

	if len(expenses) == 0 {
		fmt.Println("expenses not found")
		return nil
	}

	for _, expense := range expenses {
		fmt.Printf("\nid:%s  %s  %.f$\n\n", strconv.FormatInt(expense.ID, 10), expense.Description, expense.Amount)
	}
	return nil
}

func UpdateExpense(id int64, description string, amount float64) error {
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
