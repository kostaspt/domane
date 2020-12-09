package main

import (
	"fmt"
	"os"

	"github.com/kostaspt/domane/api/cmd"
	"github.com/kostaspt/domane/api/config"
)

func main() {
	config.Load()

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
