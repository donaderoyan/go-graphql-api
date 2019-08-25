package resolver

import (
	"github.com/donaderoyan/go-graphql-api/src/model"
	"github.com/graph-gophers/graphql-go"
)

type articlesEdgeResolver struct {
	cursor graphql.ID
	model  *model.Article
}

func (r *articlesEdgeResolver) Cursor() graphql.ID {
	return r.cursor
}

func (r *articlesEdgeResolver) Node() *articleResolver {
	return &articleResolver{u: r.model}
}
