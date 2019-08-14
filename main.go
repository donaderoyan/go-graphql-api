package main

import(
  "github.com/donaderoyan/go-graphql-api/app"
  getconfig "github.com/donaderoyan/go-graphql-api/config"
)

func main() {
  config := getconfig.LoadConfig(".")
  app.Initialize(config)
}
