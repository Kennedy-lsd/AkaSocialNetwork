package data

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Data struct {
	User interface {
		Create(context.Context, *User) error
	}
	Post interface {
		Create(context.Context, *Post) error
	}
}

func NewPostgresData(db *sqlx.DB) Data {
	return Data{
		User: &UserData{db},
		Post: &PostData{db},
	}
}
