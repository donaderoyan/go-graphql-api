package loader

import (
  "fmt"
  "github.com/donaderoyan/go-graphql-api/src/model"
  "github.com/donaderoyan/go-graphql-api/src/service"
  "context"
  "gopkg.in/nicksrandall/dataloader.v5"
  "sync"
)

type articleLoader struct {
}

func newArticleLoader() dataloader.BatchFunc {
	return articleLoader{}.loadBatch
}

func (ldr articleLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
  var (
		n       = len(keys)
		results = make([]*dataloader.Result, n)
		//wg      sync.WaitGroup
	)

  wg.Add(n)

  for i, key := range keys {
    go func (i int, key dataloader.Key)  {
      defer wg.Done()
      author, err := ctx.Value("articleService").(*service.ArticleService).FindArticleByUser(keys)
      results[i] = &dataloader.Result{Data: author, Error: err}
    }(i, key)
  }

  wg.Wait()
  return results
}

func LoadArticle(ctx context.Context, userId string) (*model.User, error) {
  var user *model.User

	data, err := loadOne(ctx, articleLoaderKey, userId)
	if err != nil {
		return nil, err
	}

	user, ok := data.(*model.User)
	if !ok {
		return nil, fmt.Errorf("wrong type: the expected type is %T but got %T", user, data)
	}

	return user, nil
}
