// This file contains the repository implementation layer.
package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/triasbrata/golibs/pkg/dbx"
)

type Repository struct {
	Db dbx.DB
}

type NewRepositoryOptions struct {
	Dsn string
}

func NewRepository(opts NewRepositoryOptions) *Repository {
	db, err := sqlx.Open("postgres", opts.Dsn)
	if err != nil {
		panic(err)
	}
	return &Repository{
		Db: dbx.New(db),
	}
}
