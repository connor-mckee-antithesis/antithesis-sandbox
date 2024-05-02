package cmd

import (
	"github.com/formancehq/antithesis-sandbox/server"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use: "server",
	Run: func(cmd *cobra.Command, args []string) {
		s := server.NewServer()
		s.Start()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
