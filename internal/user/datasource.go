package user

import "github.com/prongbang/user-service/internal/database"

type DataSource interface {
}

type dataSource struct {
	DbDriver database.Drivers
}

func NewDataSource(
	dbDriver database.Drivers,
) DataSource {
	return &dataSource{
		DbDriver: dbDriver,
	}
}
