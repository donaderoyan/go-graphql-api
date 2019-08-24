package resolver


import (
  "github.com/donaderoyan/go-graphql-api/src/model"
  graphql "github.com/graph-gophers/graphql-go"
  "time"
)
// articleResolver resolves article properties
type articleResolver struct {
	u *model.Article
}

func (r *articleResolver) ID() graphql.ID {
  return graphql.ID(r.u.ID)
}

func (r *articleResolver) Title() *string {
  return &r.u.Title
}

func (r *articleResolver) Content() *string {
  return &r.u.Content
}

func (r *userResolver) CreatedAt() (*graphql.Time, error) {
	if r.u.CreatedAt == "" {
		return nil, nil
	}

	t, err := time.Parse(time.RFC3339, r.u.CreatedAt)
	return &graphql.Time{Time: t}, err
}

func (r *userResolver) Modified() (*graphql.Time, error) {
	if r.u.Modified == "" {
		return nil, nil
	}

	t, err := time.Parse(time.RFC3339, r.u.Modified)
	return &graphql.Time{Time: t}, err
}

func (r *userResolver) Author(ctx context.Context) (*userResolver, error) {
  return newUserArticle(ctx, r.u.ID)
}
