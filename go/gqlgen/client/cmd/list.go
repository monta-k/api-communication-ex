package cmd

import (
	"context"
	"fmt"

	"github.com/shurcooL/graphql"
	"github.com/spf13/cobra"
)

var query struct {
	Todos []struct {
		ID   graphql.ID
		Text graphql.String
		Done graphql.Boolean
		User struct {
			ID   graphql.ID
			Name graphql.String
		}
	}
}

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all todos",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := graphql.NewClient("http://localhost:8080/graphql", nil)

		ctx := context.Background()

		if err := client.Query(ctx, &query, nil); err != nil {
			return err
		}

		fmt.Println("Successfully got todos:", query.Todos)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(ListCmd)
}
