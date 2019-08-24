package resolver

import (
	"github.com/donaderoyan/go-graphql-api/src/model"
	"github.com/donaderoyan/go-graphql-api/src/service"
	"github.com/op/go-logging"
	"context"
)

func (r *Resolver) CreateArticle(ctx context.Context, args *struct {
	Title    string
	Content string
}) (*articleResolver, error) {

  if isAuthorized := ctx.Value("is_authorized").(bool); !isAuthorized {
    return nil, errors.New(c.CredentialsError)
  }

  userId := ctx.Value("user_id").(*string)

  article := &model.Article{
    Title : args.Title,
    Content : args.Content,
  }

  article, err := ctx.Value("userService").(*service.ArticleService).CreateArticle(article, userId)
  if err != nil {
    ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
  }
  ctx.Value("log").(*logging.Logger).Debugf("Created article : %v by %s", *article, *userId)
  return &userResolver{article}, nil
}
