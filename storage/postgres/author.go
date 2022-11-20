package postgres

import (
	"article/models"
	"errors"
)

// AddAuthor ...
func (p Postgres) AddAuthor(id string, entity models.CreateAuthorModel) error {
	_, err := p.DB.Exec(`Insert into author(id, fullname, created_at) 
							VALUES($1,$2,now())`, id, entity.Fullname)
	if err != nil {
		return err
	}
	return nil
}

// GetAuthorByID ...
func (p Postgres) GetAuthorByID(id string) (models.AuthorWithArticles, error) {
	var result models.AuthorWithArticles
	row := p.DB.QueryRow("SELECT * FROM author WHERE deleted_at is null and id=$1", id)
	err := row.Scan(&result.ID, &result.CreatedAt, &result.UpdatedAt, &result.DeletedAt, &result.Fullname)
	if err != nil {
		return result, err
	}
	result.Articles, err = p.GetArticlesByAuthorID(id)
	if err != nil {
		return result, err
	}
	return result, nil
}

// GetArticlesByAuthorID ...
func (p Postgres) GetArticlesByAuthorID(id string) (resp []models.Article, err error) {
	rows, err := p.DB.Queryx(`SELECT 
									 id, 
									 title, 
									 body, 
									 author_id, 
									 created_at, 
									 updated_at,
									 deleted_at  
							FROM article
							WHERE deleted_at is NULL and author_id = $1 `, id)
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

// GetAuthorList ...
func (p Postgres) GetAuthorList(offset, limit int, search string) (resp []models.Author, err error) {
	rows, err := p.DB.Queryx(`SELECT
	id,
	fullname,
	created_at,
	updated_at,
	deleted_at 
	FROM author WHERE deleted_at IS NULL AND (fullname ILIKE '%' || $1 || '%')
	LIMIT $2
	OFFSET $3
	`, search, limit, offset)
	if err != nil {
		return resp, err
	}

	for rows.Next() {
		var a models.Author

		err := rows.Scan(
			&a.ID,
			&a.Fullname,
			&a.CreatedAt,
			&a.UpdatedAt,
			&a.DeletedAt,
		)
		if err != nil {
			return resp, err
		}
		resp = append(resp, a)
	}

	return resp, err
}

// UpdateAuthor ...
func (p Postgres) UpdateAuthor(req models.UpdateAuthorModel) error {
	res, err := p.DB.NamedExec("UPDATE author  SET fullname=:f, updated_at=now() WHERE deleted_at IS NULL AND id=:id", map[string]interface{}{
		"id": req.ID,
		"f":  req.Fullname,
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
	return errors.New("author not found")
}

// DeleteAuthor ...
func (p Postgres) DeleteAuthor(id string) error {
	res, err := p.DB.Exec("UPDATE author  SET deleted_at=now() WHERE id=$1 AND deleted_at IS NULL", id)
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

	return errors.New("author had been deleted already")
}
