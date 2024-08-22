package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "call-graphql",
	Short: "graphql is a client for the graphql server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, graphql!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
