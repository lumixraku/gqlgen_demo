//go:generate gorunpkg github.com/99designs/gqlgen

package gqlgen

import (
	context "context"
)

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateJSConvert(ctx context.Context, input *JSConvert) (Convert, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateContainer(ctx context.Context, convertID int, convertName *string) (Convert, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Viewer(ctx context.Context) (User, error) {
	panic("not implemented")
}
func (r *queryResolver) Sites(ctx context.Context, typeArg *SiteType, skip *int, after *string, before *string, first *int, last *int) (SiteConnection, error) {
	panic("not implemented")
}
func (r *queryResolver) Converts(ctx context.Context, typeArg *SiteType, skip *int, after *string, before *string, first *int, last *int) (ConvertConnection, error) {
	panic("not implemented")
}
