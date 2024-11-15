package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show the status of files",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO
		fmt.Println("Status functionality will go here...")
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
