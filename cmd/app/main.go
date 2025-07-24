package main

import (
	"fmt"

	"github.com/AkhmadeevRus/expense-tracker.git/cmd/tracker"
)

func main() {
	cmd := tracker.NewRootCmd()

	if err := cmd.Execute(); err != nil {
		fmt.Println("error execute:", err)
	}
}
