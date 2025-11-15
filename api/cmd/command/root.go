package command

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "Firemap",
	Short: "Application for mapping fire incidents",
	Long:  `Firemap is an application designed to map and track fire incidents in real-time.`,
}

func Execute() error {
	return rootCmd.Execute()
}
