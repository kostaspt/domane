package cmd

import (
	"github.com/spf13/cobra"

	"github.com/kostaspt/domane/api/internal/http/server"
)

func StartCmd() *cobra.Command {
	return &cobra.Command{
		Use: "start",
		Run: handleStartCmd,
	}
}

func handleStartCmd(cmd *cobra.Command, args []string) {
	cmd.PrintErr(server.Start())
}
