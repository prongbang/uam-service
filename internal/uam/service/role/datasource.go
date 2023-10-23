package role

import (
	"context"
	"errors"
	"fmt"
	"github.com/prongbang/uam-service/internal/localizations"
	"github.com/prongbang/uam-service/internal/uam/database"
	"github.com/prongbang/uam-service/pkg/core"
)

const (
	Level1 = 1
)

type DataSource interface {
	Count(params Params) int64
	GetList(params Params) []Role
	GetListByUnderRoles(roles []string) []Role
	GetById(params ParamsGetById) Role
	GetLevelById(userId string) int
	GetByName(name string) Role
	GetListByUserId(userId string) []Role
	Add(data *CreateRole) error
	Update(data *UpdateRole) error
	Delete(id string) error
}

type dataSource struct {
	Driver database.Drivers
}

func (d *dataSource) Count(params Params) int64 {
	db := d.Driver.GetPqDB()
	ctx := context.Background()

	sql := `
	SELECT COUNT(r.id) FROM (
		SELECT r.id, r.name, r.level FROM roles r
		INNER JOIN users_roles ur ON ur.role_id = r.id 
		WHERE r.level >= (SELECT r.level FROM roles r INNER JOIN users_roles ur ON ur.role_id = r.id WHERE ur.user_id = ? LIMIT 1)
		GROUP BY r.id
	) AS r
	ORDER BY r.level
	`

	args := []any{params.UserID}

	var id int64 = 0
	err := db.NewRaw(sql, args...).Scan(ctx, &id)
	if err == nil {
		return id
	}
	return 0
}

func (d *dataSource) GetList(params Params) []Role {
	db := d.Driver.GetPqDB()
	ctx := context.Background()

	sql := `
	SELECT r.id, r.name FROM (
		SELECT r.id, r.name, r.level FROM roles r
		INNER JOIN users_roles ur ON ur.role_id = r.id 
		WHERE r.level >= (SELECT r.level FROM roles r INNER JOIN users_roles ur ON ur.role_id = r.id WHERE ur.user_id = ? LIMIT 1)
		GROUP BY r.id
	) AS r
	ORDER BY r.level
	`

	args := []any{params.UserID}
	if params.LimitNo > 0 && params.OffsetNo >= 0 {
		sql += " LIMIT ? OFFSET ?"
		args = append(args, params.LimitNo, params.OffsetNo)
	}

	var rows []Role
	err := db.NewRaw(sql, args...).Scan(ctx, &rows)
	if err != nil {
		fmt.Println(err)
		return []Role{}
	}
	if len(rows) > 0 {
		return rows
	}
	return []Role{}
}

func (d *dataSource) GetListByUnderRoles(roles []string) []Role {
	db := d.Driver.GetPqDB()
	ctx := context.Background()
	sql := `
	SELECT 
		DISTINCT r.id, r.name
	FROM roles AS r
	INNER JOIN (SELECT level FROM roles WHERE id IN (%s)) AS q ON r.level >= q.level`

	var rows []Role
	args := []any{}
	for _, r := range roles {
		args = append(args, r)
	}
	sql = fmt.Sprintf(sql, core.Commas(args))

	err := db.NewRaw(sql, args...).Scan(ctx, &rows)
	if err != nil {
		fmt.Println(err)
		return []Role{}
	}
	if len(rows) > 0 {
		return rows
	}
	return []Role{}
}

func (d *dataSource) GetById(params ParamsGetById) Role {
	db := d.Driver.GetPqDB()
	ctx := context.Background()

	sql := `
	SELECT r.id, r.name FROM roles r
	INNER JOIN users_roles ur ON ur.role_id = r.id 
	WHERE r.level >= (SELECT r.level FROM roles r INNER JOIN users_roles ur ON ur.role_id = r.id WHERE ur.user_id = ? LIMIT 1)
	AND r.id = ?
	`

	args := []any{params.UserID, params.ID}

	var rows []Role
	err := db.NewRaw(sql, args...).Scan(ctx, &rows)
	if err != nil {
		fmt.Println(err)
		return Role{}
	}
	if len(rows) > 0 {
		return rows[0]
	}
	return Role{}
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

func (d *dataSource) GetByName(name string) Role {
	db := d.Driver.GetPqDB()
	ctx := context.Background()

	var rows []Role
	err := db.NewSelect().
		Model(&rows).
		ColumnExpr("r.id, r.name").
		Where("UPPER(r.name) = UPPER(?)", name).
		Limit(1).
		Scan(ctx)
	if err != nil {
		fmt.Println(err)
		return Role{}
	}
	if len(rows) > 0 {
		return rows[0]
	}
	return Role{}
}

func (d *dataSource) GetListByUserId(userId string) []Role {
	db := d.Driver.GetPqDB()
	ctx := context.Background()

	var rows []Role
	err := db.NewSelect().
		Model(&rows).
		ColumnExpr("r.id, r.name").
		Join("JOIN users_roles AS ur").
		JoinOn("ur.role_id = r.id").
		Where("ur.user_id = ?", userId).
		Scan(ctx)
	if err != nil {
		fmt.Println(err)
		return []Role{}
	}
	return rows
}

func (d *dataSource) Add(data *CreateRole) error {
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

func (d *dataSource) Update(data *UpdateRole) error {
	db := d.Driver.GetPqDB()
	ctx := context.Background()

	value := map[string]interface{}{}
	if data.Name != "" {
		value["name"] = data.Name
	}
	if data.Level > 0 {
		value["level"] = data.Level
	}
	if len(value) == 0 {
		return errors.New(localizations.CommonThereIsNoDataUpdate)
	}

	_, err := db.NewUpdate().
		Table("roles").
		Model(&value).
		Where("id = ?", data.ID).
		Exec(ctx)
	if err == nil {
		return nil
	}
	return errors.New(localizations.CommonCannotAddData)
}

func (d *dataSource) Delete(id string) error {
	db := d.Driver.GetPqDB()
	ctx := context.Background()
	rs, err := db.NewDelete().
		Table("roles").
		Where("id = ?", id).
		Exec(ctx)
	if err == nil {
		if row, e := rs.RowsAffected(); e == nil && row > 0 {
			return nil
		}
	}
	return errors.New(localizations.CommonCannotDeleteData)
}

func NewDataSource(driver database.Drivers) DataSource {
	return &dataSource{
		Driver: driver,
	}
}
