package loader

import (
	"fmt"
	"github.com/donaderoyan/go-graphql-api/src/model"
	"github.com/donaderoyan/go-graphql-api/src/service"
	"context"
	"gopkg.in/nicksrandall/dataloader.v5"
	"sync"
)

// FilmLoader contains the client required to load film resources.
type userLoader struct {
}

func newUserLoader() dataloader.BatchFunc {
	return userLoader{}.loadBatch
}

func (ldr userLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var (
		n       = len(keys)
		results = make([]*dataloader.Result, n)
		wg      sync.WaitGroup
	)

	wg.Add(n)

	for i, key := range keys {
		go func(i int, key dataloader.Key) {
			defer wg.Done()
			user, err := ctx.Value("userService").(*service.UserService).FindByEmail(key.String())
			results[i] = &dataloader.Result{Data: user, Error: err}
		}(i, key)
	}

	wg.Wait()

	return results
}

func LoadUser(ctx context.Context, key string) (*model.User, error) {
	var user *model.User

	data, err := loadOne(ctx, userLoaderKey, key)
	if err != nil {
		return nil, err
	}
	
	user, ok := data.(*model.User)
	if !ok {
		return nil, fmt.Errorf("wrong type: the expected type is %T but got %T", user, data)
	}

	return user, nil
}
