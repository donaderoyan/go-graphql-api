package resolver

import (
	"github.com/donaderoyan/go-graphql-api/src/model"
	"github.com/graph-gophers/graphql-go"
)

type roleResolver struct {
	role *model.Role
}

func (r *roleResolver) ID() graphql.ID {
	return graphql.ID(r.role.ID)
}

func (r *roleResolver) Name() *string {
	return &r.role.Name
}
