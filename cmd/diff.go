package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Show differences between files",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO
		fmt.Println("Diffing functionality will go here...")
	},
}

func init() {
	rootCmd.AddCommand(diffCmd)
}
