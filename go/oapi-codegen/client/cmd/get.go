package cmd

import (
	"api-communication-ex/oapi-codegen/adapters"
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a pet by ID",
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := cmd.PersistentFlags().GetString("id")
		if err != nil {
			return err
		}

		client, err := adapters.NewClientWithResponses("http://localhost:8080")
		if err != nil {
			return err
		}

		ctx := context.Background()

		response, err := client.ShowPetByIdWithResponse(ctx, id)
		if err != nil {
			return err
		}

		fmt.Println("http response code:", response.StatusCode())
		fmt.Println("Successfully got pet name:", response.JSON200)

		return nil
	},
}

func init() {
	GetCmd.PersistentFlags().StringP("id", "", "", "ID of the pet to get")
	GetCmd.MarkFlagRequired("id")
	rootCmd.AddCommand(GetCmd)
}
