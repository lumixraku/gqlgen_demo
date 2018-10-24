//go:generate gorunpkg github.com/99designs/gqlgen

package gqlgen

import (
	"context"
	"fmt"
	"code.byted.org/gopkg/logs"
	"math/rand"
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
func (r *Resolver) Todo() TodoResolver {
	return &todoResolver{r}
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
	return r.todos, nil
}
func (r *queryResolver) Todo(ctx context.Context, id string) (Todo, error) {
	if len(r.todos) > 0 {
		logs.Debug("%+v", r.todos)
		return r.todos[0], nil
	} else {
		return Todo{ID: "1111"}, nil
	}

}

type todoResolver struct{ *Resolver }

func (r *todoResolver) User(ctx context.Context, obj *Todo, ID int) (User, error) {
	//panic("not implemented")
	return User {
		ID: "get UUUU",
	}, nil
}
