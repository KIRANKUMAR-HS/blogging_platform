package model

import "time"

type Post struct {
	// gorm.Model
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Author    int64     `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
}

type GetPost struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	AuthorID   int       `json:"author_id"`
	AuthorName string    `json:"authername"` //users.username
	CreatedAt  time.Time `json:"created_at"`
}

type User struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	Role          string `json:"role"`
	Password_hash string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"Username"`
	Password string `json:"password"`
}
