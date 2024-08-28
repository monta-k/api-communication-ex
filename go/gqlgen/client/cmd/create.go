package cmd

import (
	"api-communication-ex/gqlgen/graph/model"
	"context"
	"fmt"

	"github.com/shurcooL/graphql"
	"github.com/spf13/cobra"
)

var mutation struct {
	CreateTodo struct {
		ID   graphql.ID
		Text graphql.String
		Done graphql.Boolean
		User struct {
			ID   graphql.ID
			Name graphql.String
		}
	} `graphql:"createTodo(input: $input)"`
}

var variables = map[string]any{
	"input": model.NewTodo{
		Text:   "Buy milk",
		UserID: "1",
	},
}

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create todo",
	RunE: func(cmd *cobra.Command, args []string) error {
		httpClient := NewHTTPClientWithAuthHeader("token")
		client := graphql.NewClient("http://localhost:8080/graphql", httpClient)

		ctx := context.Background()

		if err := client.Mutate(ctx, &mutation, variables); err != nil {
			return err
		}

		fmt.Println("Successfully create todo:", mutation.CreateTodo)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(CreateCmd)
}
