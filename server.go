package main

import(
  "net/http"
  "log"

  graphql "github.com/graph-gophers/graphql-go"

  c "github.com/donaderoyan/go-graphql-api/config"
  h "github.com/donaderoyan/go-graphql-api/src/handler"
  "github.com/donaderoyan/go-graphql-api/src/resolver"
  "github.com/donaderoyan/go-graphql-api/src/schema"
  "github.com/donaderoyan/go-graphql-api/src/service"
  "github.com/donaderoyan/go-graphql-api/src/loader"
)

func main() {
  config := c.LoadConfig(".")
  db, err := c.OpenDB(config)
  if err != nil {
		log.Fatalf("Unable to connect to db: %s \n", err)
  }
  defer db.Close()

  s := service.NewService(db, config)
  ctx := s.InitServiceContext()

	graphqlSchema := graphql.MustParseSchema(schema.NewSchema(), &resolver.Resolver{})

	http.Handle("/login", h.AddContext(ctx, h.Login()))

	loggerHandler := &h.LoggerHandler{config.DebugMode}
	http.Handle("/query", h.AddContext(ctx, loggerHandler.Logging(h.Authenticate(&h.GraphQL{Schema: graphqlSchema, Loaders: loader.NewLoaderCollection()}))))

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "graphiql.html")
	}))

	log.Fatal(http.ListenAndServe(":3000", nil))
}
