package cmd

import (
	"api-communication-ex/oapi-codegen/adapters"
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all pets",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := adapters.NewClientWithResponses("http://localhost:8080", adapters.WithRequestEditorFn(AddAuthHeader("token")))
		if err != nil {
			return err
		}

		ctx := context.Background()

		response, err := client.ListPetsWithResponse(ctx, nil)
		if err != nil {
			return err
		}

		fmt.Println("http response code:", response.StatusCode())
		fmt.Println("Successfully got pets:", response.JSON200)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(ListCmd)
}
