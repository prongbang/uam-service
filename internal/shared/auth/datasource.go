package auth

import (
	"context"
	"fmt"
	"github.com/prongbang/uam-service/internal/service/database"
	"github.com/prongbang/uam-service/internal/shared/user"
)

type DataSource interface {
	GetByUsername(username string) user.User
	GetByEmail(email string) user.User
}

type dataSource struct {
	Driver database.Drivers
}

func (d *dataSource) GetByUsername(username string) user.User {
	db := d.Driver.GetPqDB()
	ctx := context.Background()

	item := user.User{}
	err := db.NewSelect().Model(&item).Where("u.username = ?", username).Scan(ctx)
	if err != nil {
		fmt.Println(err)
	}

	return item
}

func (d *dataSource) GetByEmail(email string) user.User {
	db := d.Driver.GetPqDB()
	ctx := context.Background()

	item := user.User{}
	err := db.NewSelect().Model(&item).Where("u.email = ?", email).Scan(ctx)
	if err != nil {
		fmt.Println(err)
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
