package data

import (
	"context"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type Post struct {
	Id        int64     `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	Tags      []string  `json:"tags" db:"tags"`
	UserId    int64     `json:"user_id" db:"user_id"`
}

type PostData struct {
	db *sqlx.DB
}

func (d *PostData) Create(ctx context.Context, p *Post) error {
	query := `INSERT INTO posts (title, tags, user_id)
			  VALUES ($1, $2, $3) RETURNING id, created_at`

	err := d.db.QueryRowContext(ctx, query, p.Title, pq.Array(p.Tags), p.UserId).Scan(&p.Id, &p.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (d *PostData) GetAll(ctx context.Context) ([]Post, error) {
	query := `SELECT * FROM posts`

	rows, err := d.db.QueryxContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		var tags pq.StringArray
		err := rows.Scan(&post.Id, &post.Title, &post.CreatedAt, &tags, &post.UserId)
		if err != nil {
			return nil, err
		}
		post.Tags = tags
		posts = append(posts, post)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return posts, nil
}

func (d *PostData) GetOne(ctx context.Context, id int64) (*Post, error) {
	query := `SELECT * FROM posts WHERE id = $1`

	var post Post
	var tags pq.StringArray

	err := d.db.QueryRowxContext(ctx, query, id).Scan(
		&post.Id,
		&post.Title,
		&post.CreatedAt,
		&tags,
		&post.UserId,
	)
	if err != nil {
		return nil, err
	}

	post.Tags = tags

	return &post, nil
}

func (d *PostData) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM posts WHERE id = $1`

	result, err := d.db.ExecContext(ctx, query, id)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("Not Found")
	}

	return nil
}

func (d *PostData) Update(ctx context.Context, id int64, p *Post) error {
	query := `
			UPDATE posts
			SET title = $1, tags = $2
			WHERE id = $3
		`

	result, err := d.db.ExecContext(ctx, query, p.Title, pq.Array(p.Tags), id)
	if err != nil {
		return errors.New("failed to update post")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("Not Found")
	}

	return nil
}
