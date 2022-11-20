package models

import "time"

// Author ...
type Author struct {
	ID        string     `json:"id"`
	Fullname  string     `json:"fullname" binding:"required" minLength:"2" maxLength:"255" example:"John Doe Steve"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}

// AuthorWithArticles ...
type AuthorWithArticles struct {
	ID        string     `json:"id"`
	Fullname  string     `json:"fullname" binding:"required" minLength:"2" maxLength:"255" example:"John Doe Steve"`
	Articles  []Article  `json:"articles"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}

// CreateAuthorModel ...
type CreateAuthorModel struct {
	Fullname string `json:"fullname" binding:"required" minLength:"2" maxLength:"255" example:"John Doe Steve"`
}

// UpdateAuthorModel ...
type UpdateAuthorModel struct {
	ID       string `json:"id"`
	Fullname string `json:"fullname" binding:"required" minLength:"2" maxLength:"255" example:"John Doe Steve"`
}
