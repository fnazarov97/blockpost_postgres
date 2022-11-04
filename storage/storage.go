package storage

import "article/models"

type StorageI interface {
	AddArticle(id string, entity models.CreateArticleModel) error
	GetArticleByID(id string) (models.PackedArticleModel, error)
	GetArticleList(offset, limit int, search string) (resp []models.Article, err error)
	UpdateArticle(entered models.UpdateArticleModel) error
	DeleteArticle(id string) error

	AddAuthor(id string, entity models.CreateAuthorModel) error
	GetAuthorByID(id string) (models.AuthorWithArticles, error)
	GetArticlesByAuthorID(id string) (resp []models.Article, err error)
	GetAuthorList(offset, limit int, search string) (resp []models.Author, err error)
	UpdateAuthor(req models.UpdateAuthorModel) error
	DeleteAuthor(id string) error
}
