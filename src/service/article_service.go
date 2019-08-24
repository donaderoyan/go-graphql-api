package service

import (
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/op/go-logging"
	"github.com/rs/xid"

	"github.com/donaderoyan/go-graphql-api/config"
	"github.com/donaderoyan/go-graphql-api/src/model"
)

type ArticleService struct {
	db          *sqlx.DB
	userService *UserService
	log         *logging.Logger
}
