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

func (r *mutationResolver) CreateJSConvert(ctx context.Context, input *JSCreateInput) (Convert, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateJSConvert(ctx context.Context, input *JSUpdateInput, where *ConvertWhereInput) (Convert, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteJSConvert(ctx context.Context, where *ConvertWhereInput) (*Convert, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Converts(ctx context.Context, where *ConvertWhereInput, orderBy *ConvertOrderByInput, skip *int, after *string, before *string, first *int, last *int) (ConvertConnection, error) {
	// panic("not implemented")
	return ConvertConnection{
		Edges: []ConvertEdge{
			ConvertEdge{
				Node: AppConvert{
					ID:                "213123",
					DownloadURLStatus: 5,
					AppID:             "appid",
				},
			},
			ConvertEdge{
				Node: XpathConvert{
					ID:              "243242",
					ConvertXpathURL: "www.go.com",
				},
			},
		},
	}, nil
}
