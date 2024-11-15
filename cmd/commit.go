package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Commit changes to the repository",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO
		fmt.Println("Commit functionality will go here...")
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)
}
