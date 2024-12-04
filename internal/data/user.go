package data

import (
	"context"
	"errors"
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
			  VALUES ($1, $2, $3) RETURNING id, created_at`

	err := d.db.QueryRowContext(ctx, query, u.Name, u.Age, u.Password).Scan(&u.Id, &u.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (d *UserData) GetAll(ctx context.Context) ([]User, error) {
	query := `SELECT * FROM users`

	var users []User

	err := d.db.SelectContext(ctx, &users, query)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (d *UserData) GetOne(ctx context.Context, id int64) (*User, error) {
	query := `SELECT * FROM users WHERE id = $1`

	var user User

	err := d.db.GetContext(ctx, &user, query, id)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (d *UserData) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM users WHERE id = $1`

	result, err := d.db.ExecContext(ctx, query, id)

	if err != nil {
		return err
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
