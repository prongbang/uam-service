package forgot

import (
	"github.com/prongbang/uam-service/internal/service/database"
)

type DataSource interface {
}

type dataSource struct {
	Driver database.Drivers
}

func NewDataSource(
	dbDriver database.Drivers,
) DataSource {
	return &dataSource{
		Driver: dbDriver,
	}
}
