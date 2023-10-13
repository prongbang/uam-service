package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/prongbang/uam-service/internal/localizations"
	"github.com/prongbang/uam-service/internal/service/database"
	"github.com/prongbang/uam-service/pkg/core"
	"github.com/prongbang/uam-service/pkg/cryptox"
	"github.com/uptrace/bun"
)

type DataSource interface {
	Count(params Params) int64
	CountByUnderUserId(userId string, params Params) int64
	GetList(params Params) []User
	GetListByUnderUserId(userId string, params Params) []User
	GetById(id string) User
	Add(data *User) error
	Update(data *User) error
	UpdatePassword(userId string, password string) error
	Delete(id string) error
	DeleteTx(id string) (*bun.Tx, error)
}

type dataSource struct {
	Driver database.Drivers
}

func (d *dataSource) Count(params Params) int64 {
	db := d.Driver.GetPqDB()
	ctx := context.Background()

	var count int64 = 0
	err := db.NewRaw("SELECT count(id) FROM roles").Scan(ctx, &count)
	if err == nil {
		return count
	}
	return 0
}

func (d *dataSource) CountByUnderUserId(userId string, params Params) int64 {
	db := d.Driver.GetPqDB()
	ctx := context.Background()

	var count int64 = 0
	sql := `
	SELECT
		COUNT(u.id)
	FROM (
		SELECT r.level FROM users u 
		INNER JOIN user_roles ur ON ur.user_id = u.id 
		INNER JOIN roles r ON ur.role_id = r.id
		WHERE u.flag = ? AND u.id = ?
	) AS us
	INNER JOIN users u ON u.flag = ?
	INNER JOIN user_roles ur ON ur.user_id = u.id
	INNER JOIN roles r ON ur.role_id = r.id
	WHERE r.level >= us.level`
	err := db.NewRaw(sql, core.FlagAvailable, userId, core.FlagAvailable).Scan(ctx, &count)
	if err == nil {
		return count
	}
	return 0
}

func (d *dataSource) GetList(params Params) []User {
	db := d.Driver.GetPqDB()
	ctx := context.Background()

	sql := `
	SELECT
		u.id, 
		u.username, 
		u.email, 
		u.first_name, 
		u.last_name, 
		u.avatar, 
		u.mobile, 
		u.flag, 
		u.last_login, 
		u.created_at, 
		u.updated_at,
		r.id AS role_id,
		r.name AS role_name
	FROM users u ON u.flag = ?
	INNER JOIN user_roles ur ON ur.user_id = u.id
	INNER JOIN roles r ON ur.role_id = r.id
	LIMIT ? OFFSET ?`
	var rows []User
	err := db.NewRaw(sql, core.FlagAvailable, params.LimitNo, params.OffsetNo).Scan(ctx, &rows)
	if err != nil {
		fmt.Println(err)
		return []User{}
	}
	if len(rows) > 0 {
		return rows
	}
	return []User{}
}

func (d *dataSource) GetListByUnderUserId(userId string, params Params) []User {
	db := d.Driver.GetPqDB()
	ctx := context.Background()

	var rows []User
	sql := `
	SELECT
		u.id, 
		u.username, 
		u.email, 
		u.first_name, 
		u.last_name, 
		u.avatar, 
		u.mobile, 
		u.flag, 
		u.last_login, 
		u.created_at, 
		u.updated_at,
		r.id AS role_id,
		r.name AS role_name
	FROM (
		SELECT r.level FROM users u 
		INNER JOIN user_roles ur ON ur.user_id = u.id 
		INNER JOIN roles r ON ur.role_id = r.id
		WHERE u.flag = ? AND u.id = ?
	) AS us
	INNER JOIN users u ON u.flag = ?
	INNER JOIN user_roles ur ON ur.user_id = u.id
	INNER JOIN roles r ON ur.role_id = r.id
	WHERE r.level >= us.level
	LIMIT ? OFFSET ?`
	err := db.NewRaw(sql, core.FlagAvailable, userId, core.FlagAvailable, params.LimitNo, params.OffsetNo).Scan(ctx, &rows)
	if err != nil {
		fmt.Println(err)
		return []User{}
	}
	if len(rows) > 0 {
		return rows
	}
	return []User{}
}

func (d *dataSource) GetById(id string) User {
	db := d.Driver.GetPqDB()
	ctx := context.Background()

	var rows []User
	err := db.NewSelect().Model(&rows).Where("u.id = ?", id).Limit(1).Scan(ctx)
	if err != nil {
		fmt.Println(err)
		return User{}
	}
	if len(rows) > 0 {
		return rows[0]
	}
	return User{}
}

func (d *dataSource) Add(data *User) error {
	db := d.Driver.GetPqDB()
	ctx := context.Background()

	id := new(string)
	_ = db.NewInsert().Model(data).Returning("id").Scan(ctx, id)
	if *id != "" {
		data.ID = id
		return nil
	}
	return errors.New(localizations.CommonCannotAddData)
}

func (d *dataSource) Update(data *User) error {
	db := d.Driver.GetPqDB()
	ctx := context.Background()

	value := map[string]interface{}{}
	if data.Username != "" {
		value["username"] = data.Username
	}
	if data.Email != "" {
		value["email"] = data.Email
	}
	if data.Password != "" {
		value["password"] = cryptox.HashPassword(data.Password)
	}
	if data.Firstname != "" {
		value["first_name"] = data.Firstname
	}
	if data.Lastname != "" {
		value["last_name"] = data.Lastname
	}
	if data.Avatar != "" {
		value["avatar"] = data.Avatar
	}
	if data.Mobile != "" {
		value["mobile"] = data.Mobile
	}
	if data.Flag > core.FlagIgnore {
		value["flag"] = data.Flag
	}
	if !data.LastLogin.IsZero() {
		value["last_login"] = data.LastLogin
	}
	if len(value) == 0 {
		return errors.New(localizations.CommonThereIsNoDataUpdate)
	}

	_, err := db.NewUpdate().Table("users").Model(&value).Where("id = ?", data.ID).Exec(ctx)
	if err == nil {
		return nil
	}
	return errors.New(localizations.CommonCannotAddData)
}

func (d *dataSource) UpdatePassword(userId string, password string) error {
	db := d.Driver.GetPqDB()
	ctx := context.Background()

	value := map[string]interface{}{}
	if password != "" {
		value["password"] = cryptox.HashPassword(password)
	}
	if len(value) == 0 {
		return errors.New(localizations.CommonThereIsNoDataUpdate)
	}

	_, err := db.NewUpdate().Table("users").Model(&value).Where("id = ?", userId).Exec(ctx)
	if err == nil {
		return nil
	}
	return errors.New(localizations.CommonCannotAddData)
}

func (d *dataSource) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}

func (d *dataSource) DeleteTx(id string) (*bun.Tx, error) {
	//TODO implement me
	panic("implement me")
}

func NewDataSource(
	dbDriver database.Drivers,
) DataSource {
	return &dataSource{
		Driver: dbDriver,
	}
}
