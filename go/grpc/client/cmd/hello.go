package cmd

import (
	hellopb "api-communication-ex/grpc/pkg/grpc"
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

var HelloCmd = &cobra.Command{
	Use:   "hello",
	Short: "call hello grpc",
	RunE: func(cmd *cobra.Command, args []string) error {
		name, err := cmd.PersistentFlags().GetString("name")
		if err != nil {
			return err
		}

		conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return err
		}
		defer conn.Close()

		md := metadata.New(map[string]string{"authorization": "Bearer token"})
		ctx := metadata.NewOutgoingContext(context.Background(), md)
		client := hellopb.NewGreetingServiceClient(conn)

		res, err := client.Hello(ctx, &hellopb.HelloRequest{
			Name: name,
		})
		if err != nil {
			return err
		}
		fmt.Println(res.GetMessage())

		return nil
	},
}

func init() {
	HelloCmd.PersistentFlags().StringP("name", "", "john", "Name to say hello")
	HelloCmd.MarkFlagRequired("name")
	rootCmd.AddCommand(HelloCmd)
}
