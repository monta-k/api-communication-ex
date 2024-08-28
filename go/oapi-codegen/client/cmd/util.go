package cmd

import (
	"context"
	"net/http"
)

func AddAuthHeader(value string) func(ctx context.Context, req *http.Request) error {
	return func(ctx context.Context, req *http.Request) error {
		req.Header.Add("Authorization", "Bearer "+value)
		return nil
	}
}
