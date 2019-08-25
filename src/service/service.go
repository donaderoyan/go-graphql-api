package service

import (
  "context"
  "github.com/jmoiron/sqlx"
  "github.com/op/go-logging"

  "github.com/donaderoyan/go-graphql-api/config"
)

type Service struct{
  db      *sqlx.DB
  config  *config.Configuration
  log     *logging.Logger
}

func NewService(db *sqlx.DB, config *config.Configuration) *Service {
  log := NewLogger(config)
  return &Service{db: db, config: config, log: log}
}

func (s *Service) InitServiceContext() context.Context {
  ctx := context.Background()

	roleService := NewRoleService(s.db, s.log)
	userService := NewUserService(s.db, roleService, s.log)
	authService := NewAuthService(s.config, s.log)
  articleService := NewArticleService(s.db, s.log)

	ctx = context.WithValue(ctx, "config", s.config)
	ctx = context.WithValue(ctx, "log", s.log)
	ctx = context.WithValue(ctx, "roleService", roleService)
	ctx = context.WithValue(ctx, "userService", userService)
	ctx = context.WithValue(ctx, "authService", authService)
  ctx = context.WithValue(ctx, "articleService", articleService)

  return ctx
}
