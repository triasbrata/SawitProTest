// This file contains the repository implementation layer.
package repository

import (
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/triasbrata/golibs/pkg/dbx"
)

func TestNewRepository(t *testing.T) {
	type args struct {
		opts NewRepositoryOptions
	}
	xbd, _ := sqlx.Open("postgres", "postgres://postgres:postgres@db:5432/database?sslmode=disable")
	mdb := dbx.New(xbd)
	tests := []struct {
		name string
		args args
		want *Repository
	}{
		{
			name: "success",
			args: args{
				opts: NewRepositoryOptions{
					Dsn: "postgres://postgres:postgres@db:5432/database?sslmode=disable",
				},
			},
			want: &Repository{
				Db: mdb,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewRepository(tt.args.opts)
			assert.EqualValues(t, tt.want.Db.Close(), got.Db.Close())
		})
	}
}
