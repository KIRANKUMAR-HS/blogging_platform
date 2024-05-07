package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/KIRANKUMAR-HS/blogging_platform/internal/model"
)

type Posts struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	AuthorID   int       `json:"author_id"`
	AuthorName string    `json:"authername"`
	CreatedAt  time.Time `json:"created_at"`
}

func (d *PsqlClient) CreatePost(newPost *model.Post) (int64, error) {
	query := `
            INSERT INTO posts (title, content, author_id)
            VALUES ($1, $2, $3)
            RETURNING id

        `

	// Insert into the database and get the ID and created_at
	err := d.db.QueryRow(query, newPost.Title, newPost.Content, newPost.Author).Scan(&newPost.ID)
	if err != nil {
		log.Err(err).Msg("Error inserting into the database, failed to add post")
		return 0, err
	}

	return newPost.ID, nil
}

// GetPosts retrieves posts with the specified limit and offset
func (p *PsqlClient) GetPosts(limit, offset int) ([]Posts, error) {

	query := `
    SELECT 
        posts.id,
        posts.title,
        posts.content,
        posts.author_id,
        users.username AS author_name,
        posts.created_at
    FROM 
        posts
    JOIN 
        users 
    ON 
        posts.author_id = users.id
    ORDER BY 
        posts.created_at DESC
    `

	rows, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Posts

	for rows.Next() {
		var post Posts
		if err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.AuthorName,
			&post.CreatedAt,
		); err != nil {
			return nil, err
		}
		fmt.Println(post)
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

// GetPosts retrieves posts with optional filtering by author and creation date, with pagination
func (p *PsqlClient) GetAllPostsByfiltering(author string, createdAfter time.Time, limit, offset int) ([]*Posts, error) {
	var posts []*Posts
	query := `
		SELECT
		    p.id,
		    p.title,
		    p.content,
		    u.username AS author,
		    p.created_at
		FROM
		    posts p
		JOIN
		    users u ON p.author_id = u.id
		WHERE
		    1=1
	`

	args := []interface{}{}
	argIndex := 1

	if author != "" {
		query += fmt.Sprintf(" AND u.username = $%d", argIndex)
		args = append(args, author)
		argIndex++
	}

	if !createdAfter.IsZero() {
		query += fmt.Sprintf(" AND p.created_at >= $%d", argIndex)
		args = append(args, createdAfter)
		argIndex++
	}

	// Apply pagination
	query += fmt.Sprintf(" ORDER BY p.created_at DESC LIMIT $%d OFFSET $%d", argIndex, argIndex+1)
	args = append(args, limit, offset)

	rows, err := p.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve posts: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var post Posts
		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorName, &post.CreatedAt); err != nil {
			return nil, fmt.Errorf("could not scan post: %w", err)
		}
		posts = append(posts, &post)
	}

	return posts, nil
}

func (p *PsqlClient) GetPostByID(Id int64) (*model.GetPost, error) {
	var post Posts

	query := `
    SELECT 
        posts.id,
        posts.title,
        posts.content,
        posts.author_id,
        users.username AS author_name,
        posts.created_at
    FROM 
        posts
    JOIN 
        users 
    ON 
        posts.author_id = users.id
    WHERE posts.id = $1
    `
	err := p.db.QueryRow(query, Id).Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID, &post.AuthorName, &post.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.Wrap(err, "row iteration error")
	}
	return &model.GetPost{
		ID:         post.ID,
		Title:      post.Title,
		Content:    post.Content,
		AuthorID:   post.AuthorID,
		AuthorName: post.AuthorName,
		CreatedAt:  post.CreatedAt,
	}, nil

}

func (p *PsqlClient) UpdatePost(post *model.Post) error {
	// updated_at = NOW()
	query := `
	UPDATE posts
	SET
	  title = COALESCE($1, title),
	  content = COALESCE($2, content)
	WHERE id = $3
	`
	_, err := p.db.Exec(query, post.Title, post.Content, post.ID)
	if err != nil {
		log.Err(err).Uint64("id", uint64(post.ID)).Msg("failed to update District")
		return err
	}
	return nil

}

func (p *PsqlClient) DeletePost(Id int64) error {

	query := `DELETE FROM posts WHERE id= $1 `

	_, err := p.db.Exec(query, Id)
	if err != nil {
		return errors.Wrap(err, "could not delete post:")
	}
	return nil

}

func (d *PsqlClient) CreateUser(newUser *model.User) (int64, error) {
	query := `
            INSERT INTO users (username, password_hash, role)
            VALUES ($1, $2, $3)
            RETURNING id

        `

	// Insert into the database and get the ID and created_at
	fmt.Println(newUser)
	err := d.db.QueryRow(query, newUser.Name, newUser.Password_hash, newUser.Role).Scan(&newUser.ID)
	if err != nil {
		log.Err(err).Msg("Error inserting into the database, failed to add post")
		return 0, err
	}

	return newUser.ID, nil
}

func (d *PsqlClient) FindByUsername(username string) (*model.User, error) {
	var user model.User
	query := `SELECT id, username, password_hash, role FROM users WHERE username = $1`
	err := d.db.QueryRow(query, username).Scan(&user.ID, &user.Name, &user.Password_hash, &user.Role)
	if err == sql.ErrNoRows {
		return &model.User{}, errors.New("user not found")
	} else if err != nil {
		return &model.User{}, err
	}

	return &user, nil
}
