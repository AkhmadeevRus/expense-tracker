package internal

type Budget struct {
	Month  int64   `json:"month"`
	Amount float64 `json:"amount"`
}

func NewBudget(month int64, amount float64) *Budget {
	return &Budget{
		Month:  month,
		Amount: amount,
	}
}

func BudgetMonth(month int64, amount float64) error {
	budgets, err := ReadBudgetFromFile()
	if err != nil {
		return err
	}

	for i, budget := range budgets {
		if budget.Month == month {
			budgets[i].Amount = amount
			return WriteBudgetToFile(budgets)
		}
	}

	budget := Budget{Month: month, Amount: amount}
	budgets = append(budgets, budget)
	return WriteBudgetToFile(budgets)
}

func GetMonthlyBudget(month int64) (float64, error) {
	budgets, err := ReadBudgetFromFile()
	if err != nil {
		return 0, err
	}

	for _, budget := range budgets {
		if budget.Month == month {
			return budget.Amount, nil
		}
	}

	return 0, nil
}
