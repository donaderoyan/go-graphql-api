package app

import(
  "net/http"
  "context"
  "log"
  //"fmt"

  graphql "github.com/graph-gophers/graphql-go"

  gconfig "github.com/donaderoyan/go-graphql-api/config"
  h "github.com/donaderoyan/go-graphql-api/handler"
  "github.com/donaderoyan/go-graphql-api/resolver"
  "github.com/donaderoyan/go-graphql-api/schema"
  "github.com/donaderoyan/go-graphql-api/service"
  "github.com/donaderoyan/go-graphql-api/loader"
)

func Initialize(config *gconfig.Configuration) {
  db, err := gconfig.OpenDB(config)
  if err != nil {
		log.Fatalf("Unable to connect to db: %s \n", err)
  }

  defer db.Close()
  ctx := context.Background()
	log := service.NewLogger(config)
	roleService := service.NewRoleService(db, log)
	userService := service.NewUserService(db, roleService, log)
	authService := service.NewAuthService(config, log)

	ctx = context.WithValue(ctx, "config", config)
	ctx = context.WithValue(ctx, "log", log)
	ctx = context.WithValue(ctx, "roleService", roleService)
	ctx = context.WithValue(ctx, "userService", userService)
	ctx = context.WithValue(ctx, "authService", authService)

	graphqlSchema := graphql.MustParseSchema(schema.GetRootSchema(), &resolver.Resolver{})

	http.Handle("/login", h.AddContext(ctx, h.Login()))

	loggerHandler := &h.LoggerHandler{config.DebugMode}
	http.Handle("/query", h.AddContext(ctx, loggerHandler.Logging(h.Authenticate(&h.GraphQL{Schema: graphqlSchema, Loaders: loader.NewLoaderCollection()}))))

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "graphiql.html")
	}))

	log.Fatal(http.ListenAndServe(":3000", nil))
}
