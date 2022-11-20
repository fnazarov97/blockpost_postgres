package postgres

import (
	"article/models"
	"errors"
	"fmt"
)

// im.Db.InMemoryArticleData ...

// AddArticle ...
func (p *Postgres) AddArticle(id string, entity models.CreateArticleModel) error {
	_, err := p.DB.Exec(`Insert into article(id, title, body, author_id, created_at) 
						 VALUES($1, $2, $3, $4, now())
						`, id, entity.Content.Title, entity.Content.Body, entity.AuthorID)
	if err != nil {
		return err
	}
	return nil
}

// GetArticleByID ...
func (p *Postgres) GetArticleByID(id string) (models.PackedArticleModel, error) {
	var result models.PackedArticleModel
	err := p.DB.QueryRow(`	SELECT 
								ar.id, 
								ar.title, 
								ar.body, 
								ar.author_id,
								ar.created_at, 
								ar.updated_at,
								ar.deleted_at,
								au.fullname,
								au.created_at,
								au.updated_at, 
								au.deleted_at
							FROM article as ar join author as au on ar.author_id = au.id
							WHERE ar.deleted_at is NULL and ar.id=$1`, id).Scan(
		&result.ID, &result.Content.Title,
		&result.Content.Body, &result.Author.ID,
		&result.CreatedAt, &result.UpdatedAt, &result.DeletedAt,
		&result.Author.Fullname, &result.Author.CreatedAt,
		&result.Author.UpdatedAt, &result.Author.DeletedAt)
	fmt.Println(err)
	if err != nil {
		return result, err
	}
	return result, nil
}

// GetArticleList ...
func (p *Postgres) GetArticleList(offset, limit int, search string) (resp []models.Article, err error) {
	rows, err := p.DB.Queryx(`SELECT 
									 id, 
									 title, 
									 body, 
									 author_id, 
									 created_at, 
									 updated_at,
									 deleted_at  
								FROM article
							WHERE deleted_at is NULL and (title || ' ' || body ILIKE '%' || $1 || '%') 
							LIMIT $2 OFFSET $3`, search, limit, offset)
	if err != nil {
		return resp, err
	}
	for rows.Next() {
		var row models.Article
		err := rows.Scan(&row.ID, &row.Content.Title, &row.Content.Body,
			&row.AuthorID, &row.CreatedAt, &row.UpdatedAt, &row.DeletedAt)
		if err != nil {
			return resp, err
		}
		resp = append(resp, row)
	}
	return resp, err
}

//UpdateArticle ...
func (p *Postgres) UpdateArticle(entered models.UpdateArticleModel) error {
	res, err := p.DB.NamedExec(`UPDATE article SET title = :t, body = :b, updated_at = now()
	WHERE deleted_at IS NULL and id = :id`, map[string]interface{}{
		"id": entered.ID,
		"t":  entered.Content.Title,
		"b":  entered.Content.Body,
	})
	if err != nil {
		return err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n > 0 {
		return nil
	}
	return errors.New("article not found")
}

//DeleteArticle ...
func (p *Postgres) DeleteArticle(id string) error {
	res, err := p.DB.Exec("UPDATE article  SET deleted_at = now() WHERE id=$1 AND deleted_at IS NULL", id)
	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if n > 0 {
		return nil
	}

	return errors.New("article had been deleted already")
}
