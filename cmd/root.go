package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cdiff",
	Short: "cdiff is a minimal Git clone for basic functionality, written in Go",
	Long:  `cdiff provides basic Git functionalities like init, commit, diff, etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cdiff is a minimal Git clone. Use --help for available commands.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
