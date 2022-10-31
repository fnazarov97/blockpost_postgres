package postgres

import (
	"article/models"
	"errors"
)

// im.Db.InMemoryArticleData ...

// AddArticle ...
func (p *Postgres) AddArticle(id string, entity models.CreateArticleModel) error {
	// var article models.Article
	// article.ID = id
	// article.Content = entity.Content
	// foundAuthor, err := im.GetAuthorByID(entity.AuthorID)
	// if err != nil {
	// 	return err
	// }
	// article.AuthorID = foundAuthor.ID
	// article.CreatedAt = time.Now()

	// im.Db.InMemoryArticleData = append(im.Db.InMemoryArticleData, article)

	return nil
}

// GetArticleByID ...
func (p *Postgres) GetArticleByID(id string) (models.PackedArticleModel, error) {
	var result models.PackedArticleModel
	// for _, v := range im.Db.InMemoryArticleData {
	// 	if v.ID == id && v.DeletedAt == nil {
	// 		author, err := im.GetAuthorByID(v.AuthorID)
	// 		if err != nil {
	// 			return result, err
	// 		}

	// 		result.ID = v.ID
	// 		result.Content = v.Content
	// 		result.Author = author
	// 		result.CreatedAt = v.CreatedAt
	// 		result.UpdatedAt = v.UpdatedAt
	// 		result.DeletedAt = v.DeletedAt
	// 		return result, nil
	// 	}
	// }
	return result, errors.New("article not found")
}

// GetArticleList ...
func (p *Postgres) GetArticleList(offset, limit int, search string) (resp []models.Article, err error) {
	// var off, c int
	// for _, v := range im.Db.InMemoryArticleData {
	// 	if v.DeletedAt == nil &&
	// 		strings.Contains(strings.ToLower(v.Title+v.Body), strings.ToLower(search)) {
	// 		if offset <= off {
	// 			c++
	// 			resp = append(resp, v)
	// 		}
	// 		if c >= limit {
	// 			break
	// 		}
	// 		off++
	// 	}
	// }
	return resp, err
}

//UpdateArticle ...
func (p *Postgres) UpdateArticle(entered models.UpdateArticleModel) error {
	// for i, v := range im.Db.InMemoryArticleData {
	// 	if v.ID == entered.ID && v.DeletedAt == nil {
	// 		v.Body = entered.Body
	// 		v.Title = entered.Title
	// 		t := time.Now()
	// 		v.UpdatedAt = &t
	// 		im.Db.InMemoryArticleData[i] = v
	// 		return nil
	// 	}
	// }
	return errors.New("article not found!")
}

//DeleteArticle ...
func (p *Postgres) DeleteArticle(id string) error {
	// for i, v := range im.Db.InMemoryArticleData {
	// 	if v.ID == id {
	// 		if v.DeletedAt != nil {
	// 			return errors.New("Article alredy deleted")
	// 		}
	// 		t := time.Now()
	// 		v.DeletedAt = &t
	// 		im.Db.InMemoryArticleData[i] = v
	// 		return nil
	// 	}
	// }
	return errors.New("Not found for delete")
}
