package cmd

import (
	greetv1 "api-communication-ex/connect-go/gen/greet/v1"
	"api-communication-ex/connect-go/gen/greet/v1/greetv1connect"
	"context"
	"log"
	"net/http"

	myconnect "api-communication-ex/pkg/connect"

	"connectrpc.com/connect"
	"github.com/spf13/cobra"
)

var ConnectCmd = &cobra.Command{
	Use:   "connect",
	Short: "use connect",
	RunE: func(cmd *cobra.Command, args []string) error {
		interceptors := connect.WithInterceptors(myconnect.NewAuthClientInterceptor("token"))
		client := greetv1connect.NewGreetServiceClient(http.DefaultClient, "http://localhost:8080", interceptors)

		res, err := client.Greet(
			context.Background(),
			connect.NewRequest(&greetv1.GreetRequest{Name: "John"}),
		)
		if err != nil {
			return err
		}

		log.Println("Response:", res)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(ConnectCmd)
}
