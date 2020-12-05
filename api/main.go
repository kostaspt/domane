package main

import (
	"fmt"
	"os"

	"github.com/kostaspt/domane-server/cmd"
	"github.com/kostaspt/domane-server/config"
)

func main() {
	config.Load()

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
