package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "call-connect-go",
	Short: "connect-go is a client for the connect-go server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, connect-go!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
