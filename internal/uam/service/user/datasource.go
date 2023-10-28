package user

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/prongbang/uam-service/internal/localizations"
	"github.com/prongbang/uam-service/internal/uam/database"
	"github.com/prongbang/uam-service/internal/uam/service/permissions"
	"github.com/prongbang/uam-service/internal/uam/service/role"
	"github.com/prongbang/uam-service/pkg/core"
	"github.com/prongbang/uam-service/pkg/cryptox"
	"github.com/uptrace/bun"
)

type DataSource interface {
	Count(params Params) int64
	GetList(params Params) []User
	GetById(params ParamsGetById) User
	GetLevelById(userId string) int
	GetByEmail(email string) User
	GetByUsername(username string) User
	Add(data *CreateUser) error
	AddTx(data *CreateUser) (*bun.Tx, error)
	Update(data *UpdateUser) error
	UpdatePassword(userId string, password string) error
	Delete(id string) error
}

type dataSource struct {
	Driver database.Drivers
}

func (d *dataSource) Count(params Params) int64 {
	db := d.Driver.GetPqDB()
	ctx := context.Background()

	var count int64 = 0
	sql := `
	SELECT
		COUNT(u.id)
	FROM users u
	%s
	WHERE u.flag = ?
`
	joins := ""
	args := []any{}
	if params.Permission != permissions.All {
		args = []any{params.Payload.UserID, params.Payload.UserID}
		joins = "INNER JOIN users_creators uc ON uc.user_id = u.id AND (uc.created_by = ? OR uc.user_id = ?)"
	}

	args = append(args, core.FlagAvailable)

	sql = fmt.Sprintf(sql, joins)

	err := db.NewRaw(sql, args...).Scan(ctx, &count)
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
		COALESCE((SELECT JSON_AGG(JSON_BUILD_OBJECT('id', r.id, 'name', r.name)) FROM roles r INNER JOIN users_roles ur ON ur.role_id = r.id AND ur.user_id = u.id), '[]') AS roles_json
	FROM users u
	%s
	WHERE u.flag = ?
	ORDER BY u.created_at DESC
	`

	joins := ""
	args := []any{}
	if params.Permission != permissions.All {
		args = []any{params.Payload.UserID, params.Payload.UserID}
		joins = "INNER JOIN users_creators uc ON uc.user_id = u.id AND (uc.created_by = ? OR uc.user_id = ?)"
	}

	args = append(args, core.FlagAvailable)

	if params.LimitNo > 0 && params.OffsetNo >= 0 {
		sql += " LIMIT ? OFFSET ?"
		args = append(args, params.LimitNo, params.OffsetNo)
	}

	sql = fmt.Sprintf(sql, joins)

	var rows []User
	err := db.NewRaw(sql, args...).Scan(ctx, &rows)
	if err != nil {
		fmt.Println(err)
		return []User{}
	}
	if len(rows) > 0 {

		// Parse json to list
		for i, u := range rows {
			r := []role.Role{}
			_ = json.Unmarshal([]byte(u.RolesJson), &r)
			rows[i].Roles = r
		}

		return rows
	}
	return []User{}
}

func (d *dataSource) GetById(params ParamsGetById) User {
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
		COALESCE((SELECT JSON_AGG(JSON_BUILD_OBJECT('id', r.id, 'name', r.name)) FROM roles r INNER JOIN users_roles ur ON ur.role_id = r.id AND ur.user_id = u.id), '[]') AS roles_json
	FROM users u
	%s
	WHERE u.id = ? AND u.flag = ?
	LIMIT 1
	`

	joins := ""
	args := []any{}
	if params.Permission != permissions.All {
		args = []any{params.Payload.UserID, params.Payload.UserID}
		joins = "INNER JOIN users_creators uc ON uc.user_id = u.id AND (uc.created_by = ? OR uc.user_id = ?)"
	}
	args = append(args, params.ID, core.FlagAvailable)

	sql = fmt.Sprintf(sql, joins)

	var rows []User
	err := db.NewRaw(sql, args...).Scan(ctx, &rows)
	if err != nil {
		fmt.Println(err)
		return User{}
	}
	if len(rows) > 0 {

		// Parse json to list
		for i, u := range rows {
			r := []role.Role{}
			_ = json.Unmarshal([]byte(u.RolesJson), &r)
			rows[i].Roles = r
		}

		return rows[0]
	}
	return User{}
}

func (d *dataSource) GetLevelById(userId string) int {
	db := d.Driver.GetPqDB()
	ctx := context.Background()

	sql := `
	SELECT 
	    r.level 
	FROM roles r 
	INNER JOIN users_roles ur ON ur.role_id = r.id 
	WHERE ur.user_id = ?
	LIMIT 1
	`

	args := []any{userId}

	var level = 0
	err := db.NewRaw(sql, args...).Scan(ctx, &level)
	if err == nil {
		return level
	}
	return 0
}

func (d *dataSource) GetByEmail(email string) User {
	db := d.Driver.GetPqDB()
	ctx := context.Background()

	var rows []User
	err := db.NewSelect().Model(&rows).Where("u.email = ?", email).Limit(1).Scan(ctx)
	if err != nil {
		fmt.Println(err)
		return User{}
	}
	if len(rows) > 0 {
		return rows[0]
	}
	return User{}
}

func (d *dataSource) GetByUsername(username string) User {
	db := d.Driver.GetPqDB()
	ctx := context.Background()

	var rows []User
	err := db.NewSelect().Model(&rows).Where("u.username = ?", username).Limit(1).Scan(ctx)
	if err != nil {
		fmt.Println(err)
		return User{}
	}
	if len(rows) > 0 {
		return rows[0]
	}
	return User{}
}

func (d *dataSource) Add(data *CreateUser) error {
	db := d.Driver.GetPqDB()
	ctx := context.Background()

	// Hash password
	data.Password = cryptox.HashPassword(data.Password)

	id := new(string)
	_ = db.NewInsert().Model(data).Returning("id").Scan(ctx, id)
	if *id != "" {
		data.ID = id
		return nil
	}
	return errors.New(localizations.CommonCannotAddData)
}

func (d *dataSource) AddTx(data *CreateUser) (*bun.Tx, error) {
	db := d.Driver.GetPqDB()
	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		return &tx, err
	}

	// Hash password
	data.Password = cryptox.HashPassword(data.Password)

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

func (d *dataSource) Update(data *UpdateUser) error {
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
	if data.FirstName != "" {
		value["first_name"] = data.FirstName
	}
	if data.LastName != "" {
		value["last_name"] = data.LastName
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
	if data.LastLogin != nil && !data.LastLogin.IsZero() {
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
	db := d.Driver.GetPqDB()
	ctx := context.Background()
	rs, err := db.NewDelete().
		Table("users").
		Where("id = ?", id).
		Exec(ctx)
	if err == nil {
		if row, e := rs.RowsAffected(); e == nil && row > 0 {
			return nil
		}
	}
	return errors.New(localizations.CommonCannotDeleteData)
}

func NewDataSource(
	dbDriver database.Drivers,
) DataSource {
	return &dataSource{
		Driver: dbDriver,
	}
}
