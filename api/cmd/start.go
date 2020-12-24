package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/kostaspt/domane/api/internal/net/server"
)

func StartCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "start",
		Run: handleStartCmd,
	}

	cmd.PersistentFlags().IntP("port", "p", 4000, "Server's port")
	viper.BindPFlag("port", cmd.PersistentFlags().Lookup("port"))

	return cmd
}

func handleStartCmd(cmd *cobra.Command, args []string) {
	cmd.PrintErr(server.Start())
}
