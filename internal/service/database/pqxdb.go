package database

import (
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
	"os"
)

type PqxDriver interface {
	Connect() *bun.DB
}

type pqxDriver struct {
}

// Connect implements PqxDriver.
func (p *pqxDriver) Connect() *bun.DB {
	config, err := pgx.ParseConfig(os.Getenv("PQ_CONNECTION"))
	if err != nil {
		panic(err)
	}
	config.PreferSimpleProtocol = true

	sqldb := stdlib.OpenDB(*config)
	db := bun.NewDB(sqldb, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	return db
}

func NewPqDriver() PqxDriver {
	return &pqxDriver{}
}
