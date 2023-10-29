package user_creator

import (
	"context"
	"errors"
	"github.com/prongbang/uam-service/internal/localizations"
	"github.com/prongbang/uam-service/internal/uam/bunx"
	"github.com/prongbang/uam-service/internal/uam/database"
	"github.com/uptrace/bun"
)

type DataSource interface {
	AddTx(data *CreateUserCreator) (*bun.Tx, error)
	DeleteTx(data *DeleteUserCreator) (*bun.Tx, error)
}

type dataSource struct {
	Driver database.Drivers
}

func (d *dataSource) AddTx(data *CreateUserCreator) (*bun.Tx, error) {
	db := d.Driver.GetPqDB()
	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		return &tx, err
	}

	id := new(string)
	err = tx.NewInsert().Model(data).Returning("id").Scan(ctx, id)
	if err == nil {
		if *id != "" {
			data.ID = id
			return &tx, nil
		}
		return &tx, errors.New(localizations.CommonCannotAddData)
	}
	return &tx, errors.New(localizations.CommonCannotAddData)
}

func (d *dataSource) DeleteTx(data *DeleteUserCreator) (*bun.Tx, error) {
	db := d.Driver.GetPqDB()

	wheres := "user_id = ? AND created_by = ?"
	args := []any{
		data.UserID,
		data.CreatedBy,
	}

	tx, err := bunx.DeleteTx(db, "users_creators", wheres, args)
	if err == nil {
		return tx, nil
	}
	return tx, errors.New(localizations.CommonCannotDeleteData)
}

func NewDataSource(
	dbDriver database.Drivers,
) DataSource {
	return &dataSource{
		Driver: dbDriver,
	}
}
