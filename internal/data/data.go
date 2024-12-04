package data

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Data struct {
	User interface {
		Create(context.Context, *User) error
		GetAll(context.Context) ([]User, error)
		GetOne(context.Context, int64) (*User, error)
		Delete(context.Context, int64) error
	}
	Post interface {
		Create(context.Context, *Post) error
		GetAll(context.Context) ([]Post, error)
		GetOne(context.Context, int64) (*Post, error)
		Delete(context.Context, int64) error
		Update(context.Context, int64, *Post) error
	}
}

func NewPostgresData(db *sqlx.DB) Data {
	return Data{
		User: &UserData{db},
		Post: &PostData{db},
	}
}
