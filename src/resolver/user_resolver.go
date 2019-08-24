package resolver

import (
	"github.com/donaderoyan/go-graphql-api/src/model"
	graphql "github.com/graph-gophers/graphql-go"
	"time"
	"github.com/donaderoyan/go-graphql-api/src/loader"
)

// userResolver resolves user properties
type userResolver struct {
	u *model.User
}

// ID resolves user ID
func (r *userResolver) ID() graphql.ID {
	return graphql.ID(r.u.ID)
}

// Email resolves user Email
func (r *userResolver) Email() *string {
	return &r.u.Email
}

// Password resolves user Password
func (r *userResolver) Password() *string {
	maskedPassword := "********"
	return &maskedPassword
}

// IPAddress resolves user IPAddress
func (r *userResolver) IPAddress() *string {
	return &r.u.IPAddress
}

// CreatedAt resolves user CreatedAt
func (r *userResolver) CreatedAt() (*graphql.Time, error) {
	if r.u.CreatedAt == "" {
		return nil, nil
	}

	t, err := time.Parse(time.RFC3339, r.u.CreatedAt)
	return &graphql.Time{Time: t}, err
}

func (r *userResolver) Roles() *[]*roleResolver {
	l := make([]*roleResolver, len(r.u.Roles))
	for i := range l {
		l[i] = &roleResolver{
			role: r.u.Roles[i],
		}
	}
	return &l
}

func newUserArticle(ctx context.Context, articleID string) (*roleResolver, error) {
	author, err := loader.LoadAuthor(ctx, articleID)
	if err != nil {
		return nil, err
	}

	return author, nil
}
