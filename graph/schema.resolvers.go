package graph

import (
	"context"
	"go-graph/graph/model"
	"math/rand"
	"strconv"
)

var todos []*model.Todo // In-memory storage for todos

// CreateTodo is the resolver for the createTodo mutation.
func (r *mutationResolver) CreateTodo(ctx context.Context, text string) (*model.Todo, error) {
	newTodo := &model.Todo{
		ID:   strconv.Itoa(rand.Int()), // Generate a random ID
		Text: text,
		Done: false,
	}
	todos = append(todos, newTodo)
	return newTodo, nil
}

// Todos is the resolver for the query todos request.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return todos, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }
