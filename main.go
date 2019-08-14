package main

import(
  "github.com/donaderoyan/graphql-api/app"
  gconfig "github.com/donaderoyan/graphql-api/config"
)

func main() {
  config := gconfig.LoadConfig(".")
  app.Initialize(config)
}
