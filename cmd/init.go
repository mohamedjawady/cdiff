package cmd

import (
	"fmt"

	"github.com/mohamedjawady/cdiff/commands"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new repository",
	Run: func(cmd *cobra.Command, args []string) {
		if err := commands.InitRepo(); err != nil {
			fmt.Println("Error initializing repository:", err)
		} else {
			fmt.Println("Repository initialized successfully!")
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
