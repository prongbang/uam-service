package user

import (
	"context"
	"fmt"
	"github.com/prongbang/user-service/internal/service/database"
)

type DataSource interface {
	FindByUsername(username string) User
	FindByEmail(email string) User
}

type dataSource struct {
	Driver database.Drivers
}

func (d *dataSource) FindByUsername(username string) User {
	db := d.Driver.GetPqDB()
	ctx := context.Background()

	item := User{}
	err := db.NewSelect().Model(&item).Where("username = ?", username).Scan(ctx)
	if err != nil {
		fmt.Println("FindByUsername error:", err)
	}

	return item
}

func (d *dataSource) FindByEmail(email string) User {
	db := d.Driver.GetPqDB()
	ctx := context.Background()

	item := User{}
	err := db.NewSelect().Model(&item).Where("email = ?", email).Scan(ctx)
	if err != nil {
		fmt.Println("FindByEmail error:", err)
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
