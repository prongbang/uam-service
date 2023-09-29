package database

import (
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

type PqDriver interface {
	Connect() *bun.DB
}

type pqDriver struct {
}

// Connect implements PqDriver.
func (p *pqDriver) Connect() *bun.DB {
	config, err := pgx.ParseConfig("postgres://postgres:@localhost:5432/test?sslmode=disable")
	if err != nil {
		panic(err)
	}
	config.PreferSimpleProtocol = true

	sqldb := stdlib.OpenDB(*config)
	db := bun.NewDB(sqldb, pgdialect.New())

	return db
}

func NewPqDriver() PqDriver {
	return &pqDriver{}
}
