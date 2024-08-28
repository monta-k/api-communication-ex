package cmd

import (
	"api-communication-ex/oapi-codegen/adapters"
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var PostCmd = &cobra.Command{
	Use:   "post",
	Short: "Post a pet",
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := cmd.PersistentFlags().GetInt64("id")
		if err != nil {
			return err
		}
		name, err := cmd.PersistentFlags().GetString("name")
		if err != nil {
			return err
		}
		tag, err := cmd.PersistentFlags().GetString("tag")
		if err != nil {
			return err
		}

		client, err := adapters.NewClientWithResponses("http://localhost:8080", adapters.WithRequestEditorFn(AddAuthHeader("token")))
		if err != nil {
			return err
		}

		ctx := context.Background()

		body := adapters.CreatePetsJSONRequestBody{
			Id:   id,
			Name: name,
			Tag:  &tag,
		}

		response, err := client.CreatePetsWithResponse(ctx, body)
		if err != nil {
			return err
		}

		fmt.Println("http response code:", response.StatusCode())

		return nil
	},
}

func init() {
	PostCmd.PersistentFlags().Int64P("id", "", 1, "ID of the pet to post")
	PostCmd.PersistentFlags().StringP("name", "", "Dog", "Name of the pet to post")
	PostCmd.PersistentFlags().StringP("tag", "", "cute", "Tag of the pet to post")
	PostCmd.MarkFlagRequired("id")
	PostCmd.MarkFlagRequired("name")
	rootCmd.AddCommand(PostCmd)
}
