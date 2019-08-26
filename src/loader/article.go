package loader

import (
  "fmt"
  "github.com/donaderoyan/go-graphql-api/src/model"
  "github.com/donaderoyan/go-graphql-api/src/service"
  "context"
  "gopkg.in/nicksrandall/dataloader.v5"
  //"sync"
)

type articlesLoader struct {
}

func newArticleLoader() dataloader.BatchFunc {
	return articlesLoader{}.loadBatch
}

func (ldr articlesLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
  var (
		n       = len(keys)
		results = make([]*dataloader.Result, n)
		//wg      sync.WaitGroup
	)
  data, err := ctx.Value("articleService").(*service.ArticleService).FindArticlesByUser(keys)

  //wg.Add(n)

  for i, v := range data {
    results[i] = &dataloader.Result{Data: v}
  }


  //wg.Wait()
  return results
}




func LoadArticlesByUser(ctx context.Context, userId string) ([]*model.Article, error) {
	data, err := loadOne(ctx, articlesLoaderKey, userId)
	if err != nil {
		return nil, err
	}

	article, ok := data.([]*model.Article)
	if !ok {
		return nil, fmt.Errorf("wrong type: the expected type is %T but got %T", user, data)
	}

	return article, nil
}
