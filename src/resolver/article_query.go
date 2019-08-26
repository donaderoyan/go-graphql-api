package resolver

import (
	"errors"
	c "github.com/donaderoyan/go-graphql-api/config"
	"github.com/donaderoyan/go-graphql-api/src/loader"
	"github.com/donaderoyan/go-graphql-api/src/service"
	"github.com/op/go-logging"
	"context"
)

// func (r *Resolver) listArticles(ctx context.Context, args struct {
//   First *int32
//   }) (*articlesConnectionResolver, error) {
//
// }
