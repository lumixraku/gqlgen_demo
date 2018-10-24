//go:generate gorunpkg github.com/99designs/gqlgen

package gqlgen

import (
	context "context"
	"fmt"
	"math/rand"

	"code.byted.org/gopkg/logs"
)

type Resolver struct {
	todos         []Todo
	todosTomorrow []Todo
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTodo(ctx context.Context, input NewTodo) (Todo, error) {
	newTodo := Todo{
		User: User{
			ID:   input.UserID,
			Name: "name ::" + input.Text,
		},
		Text: input.Text,
		ID:   fmt.Sprintf("T%d", rand.Int()),
	}

	//相同于r.Resolver.todos = append(r.Resolver.todos, newTodo)
	r.todos = append(r.todos, newTodo)
	return newTodo, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todos(ctx context.Context) ([]Todo, error) {
	fmt.Printf("%+v", r.todos)
	return r.todos, nil
}
func (r *queryResolver) Todo(ctx context.Context, id string) (Todo, error) {
	// panic("not implemented")
	if len(r.todos) > 0 {
		logs.Debug("%+v", r.todos)
		return r.todos[0], nil
	} else {
		return Todo{ID: "1111"}, nil
	}
}
