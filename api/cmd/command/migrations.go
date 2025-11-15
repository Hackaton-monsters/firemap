package command

import (
	"firemap/internal/infrastructure/di"

	"github.com/spf13/cobra"
)

func init() {
	var migrateCmd = &cobra.Command{
		Use:   "migrate",
		Short: "Run migration",
		Long:  `Execute migrations`,
		Run: func(cmd *cobra.Command, args []string) {
			migrate()
		},
	}

	rootCmd.AddCommand(migrateCmd)
}

func migrate() {
	pm := di.InitializeProcessManager()
	pm.Migrate()
}
