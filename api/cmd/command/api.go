package command

import (
	"firemap/internal/infrastructure/di"

	"github.com/spf13/cobra"
)

func init() {
	var apiCmd = &cobra.Command{
		Use:   "api",
		Short: "Run Firemap app api",
		Long:  `Run api for Firemap application`,
		Run: func(cmd *cobra.Command, args []string) {
			runApi()
		},
	}

	rootCmd.AddCommand(apiCmd)
}

func runApi() {
	pm := di.InitializeProcessManager()
	pm.RunHTTPServer()
}
