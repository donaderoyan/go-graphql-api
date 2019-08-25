package resolver

import (
	"github.com/donaderoyan/go-graphql-api/src/model"
	"github.com/donaderoyan/go-graphql-api/src/service"
)

type articlesConnectionResolver struct {
	users      []*model.Article
	totalCount int
	from       *string
	to         *string
}

func (r *articlesConnectionResolver) TotalCount() int32 {
	return int32(r.totalCount)
}

func (r *articlesConnectionResolver) Edges() *[]*usersEdgeResolver {
	l := make([]*usersEdgeResolver, len(r.users))
	for i := range l {
		l[i] = &usersEdgeResolver{
			cursor: service.EncodeCursor(&(r.users[i].ID)),
			model:  r.users[i],
		}
	}
	return &l
}

func (r *articlesConnectionResolver) PageInfo() *pageInfoResolver {
	return &pageInfoResolver{
		startCursor: service.EncodeCursor(r.from),
		endCursor:   service.EncodeCursor(r.to),
		hasNextPage: false,
	}
}
