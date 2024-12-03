package data

import (
	"context"
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
	query := `INSERT INTO (title, tags, user_id)
			  VALUES ($1, $2, $3) RETURNING id, created_at`

	err := d.db.QueryRowContext(ctx, query, p.Title, pq.Array(p.Tags), p.UserId).Scan(&p.Id, &p.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}
