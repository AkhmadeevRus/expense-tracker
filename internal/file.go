package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

func getExpenseFilePath() string {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("error getting current working 	directory:", err)
		return ""
	}
	return path.Join(cwd, "expenses.json")
}

func ReadExpenseFromFile() ([]Expense, error) {
	filePath := getExpenseFilePath()
	_, err := os.Stat(filePath)

	if os.IsNotExist(err) {
		fmt.Println("file dosen't exists. Creating new file")
		file, err := os.Create(filePath)
		if err != nil {
			return nil, err
		}
		defer file.Close()
		os.WriteFile(filePath, []byte("[]"), os.ModeAppend.Perm())
		return []Expense{}, err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tasks []Expense
	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func WriteExpenseToFile(tasks []Expense) error {
	filePath := getExpenseFilePath()
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(tasks)
	if err != nil {
		return err
	}

	return nil
}
