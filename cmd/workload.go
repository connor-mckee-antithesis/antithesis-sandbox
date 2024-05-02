package cmd

import (
	"github.com/formancehq/antithesis-sandbox/workload"
	"github.com/spf13/cobra"
)

var workloadCmd = &cobra.Command{
	Use: "workload",
	Run: func(cmd *cobra.Command, args []string) {
		w := workload.Workload{}
		w.Execute()
	},
}

func init() {
	rootCmd.AddCommand(workloadCmd)
}
