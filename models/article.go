package models

import "time"

// Content ...
type Content struct {
	Title string `json:"title" `
	Body  string `json:"body" `
}

// Article ...
type Article struct {
	ID        string     `json:"id"`
	Content              // Promoted fields
	AuthorID  string     `json:"author_id" binding:"required"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}

// CreateArticleModel ...
type CreateArticleModel struct {
	Content         // Promoted fields
	AuthorID string `json:"author_id" binding:"required"`
}

// PackedArticleModel ...
type PackedArticleModel struct {
	ID        string     `json:"id"`
	Content              // Promoted fields
	Author    Author     `json:"author"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}

type UpdateArticleModel struct {
	ID      string `json:"id"`
	Content        // Promoted fields
}

type DeleteArticleModel struct {
	ID        string     `json:"id"`
	Content              // Promoted fields
	Author    Author     `json:"author"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
