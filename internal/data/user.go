package data

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

type User struct {
	Id        int64     `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Age       int64     `json:"age" db:"age"`
	Password  string    `json:"-" db:"password"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type UserData struct {
	db *sqlx.DB
}

func (d *UserData) Create(ctx context.Context, u *User) error {
	query := `INSERT INTO users (name, age, password)
			  VALUES ($1, $2, $3) RETURNING id created_at`

	err := d.db.QueryRowContext(ctx, query, u.Name, u.Age, u.Password).Scan(&u.Id, &u.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}
