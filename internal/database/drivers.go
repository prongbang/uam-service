package database

import (
	"github.com/uptrace/bun"
)

type Drivers interface {
	GetPqDB() *bun.DB
}

type drivers struct {
	PqDB *bun.DB
}

func (d *drivers) GetPqDB() *bun.DB {
	return d.PqDB
}

func NewDrivers(
	pqDB PqDriver,
) Drivers {
	return &drivers{
		PqDB: pqDB.Connect(),
	}
}
