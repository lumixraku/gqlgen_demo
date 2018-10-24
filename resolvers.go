package gqlgen

import context "context"

type Resolver struct {
	Meme []string
}

func New() *Resolver {
	return &Resolver{
		Meme: []string{},
	}
}

func (r *Resolver) Mutation_createTodo(ctx context.Context, input NewTodo) (Todo, error) {
	return Todo{
		ID: "ccss",
	}, nil
}

func (r *Resolver) Query_todos(ctx context.Context) ([]Todo, error) {
	return []Todo{{
		ID: "ccss",
	}}, nil
}

func (r *Resolver) Query_todo(ctx context.Context) (Todo, error) {
	return Todo{
		ID: "ccss",
	}, nil
}

func (r *Resolver) Todo_user(ctx context.Context, obj *Todo) (User, error) {
	return User{
		ID: "uuuuuuu",
	}, nil
}
