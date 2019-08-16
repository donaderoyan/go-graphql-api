package resolver

import (
	getconfig "github.com/donaderoyan/go-graphql-api/config"
	"github.com/donaderoyan/go-graphql-api/src/schema"
	"github.com/donaderoyan/go-graphql-api/src/service"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/gqltesting"
	"context"
	"log"
	"testing"
)

var (
	rootSchema = graphql.MustParseSchema(schema.GetRootSchema(), &Resolver{})
	ctx        context.Context
)

func init() {
	config := getconfig.LoadConfig("../")
	db, err := getconfig.OpenDB(config)
	if err != nil {
		log.Fatalf("Unable to connect to db: %s \n", err)
	}
	log := service.NewLogger(config)
	roleService := service.NewRoleService(db, log)
	userService := service.NewUserService(db, roleService, log)
	ctx = context.WithValue(context.Background(), "userService", userService)
}

func TestBasic(t *testing.T) {
	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Context: ctx,
			Schema:  rootSchema,
			Query: `
				{
					user(email:"test@1.com") {
						id
						email
						password
					}
				}
			`,
			ExpectedResult: `
				{
					"user": {
					  "id": "1",
					  "email": "test@1.com",
					  "password": "********"
					}
				}
			`,
		},
	})
}
