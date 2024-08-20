package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "call-oapi-codegen-client",
	Short: "oapi-codegen-client is a client for the oapi-codegen server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, oapi-codegen-client!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
