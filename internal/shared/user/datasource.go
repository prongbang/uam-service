package user

import (
	"context"
	"fmt"
	"github.com/prongbang/user-service/internal/service/database"
)

type DataSource interface {
	GetByUsername(username string) User
	GetByEmail(email string) User
}

type dataSource struct {
	Driver database.Drivers
}

func (d *dataSource) GetByUsername(username string) User {
	db := d.Driver.GetPqDB()
	ctx := context.Background()

	item := User{}
	err := db.NewSelect().Model(&item).Where("username = ?", username).Scan(ctx)
	if err != nil {
		fmt.Println("GetByUsername error:", err)
	}

	return item
}

func (d *dataSource) GetByEmail(email string) User {
	db := d.Driver.GetPqDB()
	ctx := context.Background()

	item := User{}
	err := db.NewSelect().Model(&item).Where("email = ?", email).Scan(ctx)
	if err != nil {
		fmt.Println("GetByEmail error:", err)
	}

	return item
}

func NewDataSource(
	dbDriver database.Drivers,
) DataSource {
	return &dataSource{
		Driver: dbDriver,
	}
}
