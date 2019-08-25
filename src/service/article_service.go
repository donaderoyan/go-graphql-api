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

const (
	defaultListFetchSize = 10
)

type ArticleService struct {
	db          *sqlx.DB
	log         *logging.Logger
}

type relArticleUser struct {
	UserId 		string `db:"user_id"`
	ArticleId string `db:"article_id"`
}

func NewArticleService(db *sqlx.DB, log *logging.Logger) *ArticleService {
	return &ArticleService{db: db, log: log}
}

func (a *ArticleService) CreateArticle(article *model.Article, userId *string) (*model.Article, error) {
	articleId := xid.New()
	article.ID = articleId.String()
	articleSQL := `INSERT INTO articles (id, title, content) VALUES (:id, :title, :content)`
	_, err := a.db.NamedExec(articleSQL, article)
	if err != nil {
		return nil, err
	}
	//many to many relation
	rel := &relArticleUser{
		UserId 		: userId,
		ArticleId	: articleId,
	}
	articleUserSQL : `INSERT INTO rel_articles_users(user_id, article_id) VALUES(:UserId, :ArticleId)`
	_, err := a.db.NamedExec(articleUserSQL, *rel)
	if err != nil {
		return nil, err
	}
	return article, nil
}


func (a *ArticleService) ListArticles(first *int32, after *string) ([]*model.Article, error) {
	articles := make([]*model.Article, 0)
	var fetchSize int32
	if first == nil {
		fetchSize = defaultListFetchSize
	} else {
		fetchSize = *first
	}

	if after != nil {
		articleSQL := `SELECT id, title, content, created_at, modified FROM articles WHERE created_at < (SELECT created_at FROM articles WHERE id = $1) AND deleted_at IS NULL ORDER BY created_at DESC LIMIT $2;`
		decodedIndex, _ := DecodeCursor(after)
		err := a.db.Select(&articles, articleSQL, decodedIndex, fetchSize)
		if err != nil {
			return nil, err
		}
		return articles, nil
	}
	articleSQL := `SELECT id, title, content, created_at, modified FROM articles WHERE deleted_at IS NULL ORDER BY created_at DESC LIMIT $1;`
	err := u.db.Select(&articles, articleSQL, fetchSize)
	if err != nil {
		return nil, err
	}
	return articles, nil
}
