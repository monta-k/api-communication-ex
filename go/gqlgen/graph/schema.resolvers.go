package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"api-communication-ex/gqlgen/generated"
	"api-communication-ex/gqlgen/graph/model"
	"api-communication-ex/pkg/auth"
	"context"
	"fmt"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	user := auth.UserFromContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("unauthorized")
	}
	fmt.Println("user", user)

	return &model.Todo{
		ID:   "3",
		Text: input.Text,
		Done: false,
		User: &model.User{
			ID:   input.UserID,
			Name: "user 1",
		},
	}, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	user := auth.UserFromContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("unauthorized")
	}
	fmt.Println("user", user)

	return []*model.Todo{
		{
			ID:   "1",
			Text: "todo 1",
			Done: false,
			User: &model.User{
				ID:   "1",
				Name: "user 1",
			},
		},
		{
			ID:   "2",
			Text: "todo 2",
			Done: true,
			User: &model.User{
				ID:   "2",
				Name: "user 2",
			},
		},
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
