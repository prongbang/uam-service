package role

import (
	"context"
	"fmt"
	"github.com/prongbang/user-service/internal/service/database"
	"github.com/prongbang/user-service/pkg/core"
)

type DataSource interface {
	Count() int64
	GetList(filter Filter) []Role
	GetListByUnderLevel(level int) []Role
	GetListByUnderRoles(roles []string) []Role
	GetById(id string) Role
	GetByName(name string) Role
	Add(data *Role) error
	Update(data *Role) error
	Delete(id string) error
}

type dataSource struct {
	Driver database.Drivers
}

func (d *dataSource) Count() int64 {
	db := d.Driver.GetPqDB()
	ctx := context.Background()

	var id int64 = 0
	err := db.NewRaw("SELECT count(id) FROM roless").Scan(ctx, &id)
	if err == nil {
		return id
	}
	return 0
}

func (d *dataSource) GetList(filter Filter) []Role {
	db := d.Driver.GetPqDB()
	ctx := context.Background()

	sql := "SELECT id, name, level FROM roless LIMIT ? OFFSET ?"
	var rows []Role
	err := db.NewRaw(sql, filter.LimitNo, filter.OffsetNo).Scan(ctx, &rows)
	if err != nil {
		fmt.Println(err)
		return []Role{}
	}
	if len(rows) > 0 {
		return rows
	}
	return []Role{}
}

func (d *dataSource) GetListByUnderLevel(level int) []Role {
	db := d.Driver.GetPqDB()
	ctx := context.Background()

	sql := `SELECT id, name, level FROM roles WHERE level >= ?`
	var rows []Role
	err := db.NewRaw(sql, level).Scan(ctx, &rows)
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
		distinct r.id, r.name, r.level
	FROM roles AS r
	INNER JOIN (
		SELECT level
		FROM roles
		WHERE id IN (%s)
	) AS q ON r.level >= q.level`

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

func (d *dataSource) GetById(id string) Role {
	db := d.Driver.GetPqDB()
	ctx := context.Background()

	var rows []Role
	err := db.NewSelect().Table("roles").Model(&rows).Where("id = ?", id).Limit(1).Scan(ctx)
	if err != nil {
		fmt.Println(err)
		return Role{}
	}
	if len(rows) > 0 {
		return rows[0]
	}
	return Role{}
}

func (d *dataSource) GetByName(name string) Role {
	db := d.Driver.GetPqDB()
	ctx := context.Background()

	var rows []Role
	err := db.NewSelect().Table("roles").Model(&rows).Where("UPPER(name) = UPPER(?)", name).Limit(1).Scan(ctx)
	if err != nil {
		fmt.Println(err)
		return Role{}
	}
	if len(rows) > 0 {
		return rows[0]
	}
	return Role{}
}

func (d *dataSource) Add(data *Role) error {
	db := d.Driver.GetPqDB()
	ctx := context.Background()

	id := ""
	_, err := db.NewInsert().Table("roles").Model(data).Returning("id", &id).Exec(ctx)
	if err == nil {
		if id == "" {
			data.ID = id
			return nil
		}
	}
	return fmt.Errorf("%s", "Cannot add or update a child row")
}

func (d *dataSource) Update(data *Role) error {
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
		return fmt.Errorf("%s", "Is not data to update")
	}

	rs, err := db.NewUpdate().Model(&value).Table("roles").Where("id = ?", data.ID).Exec(ctx)
	if err == nil {
		if row, e := rs.LastInsertId(); e == nil {
			if row > -1 {
				return nil
			}
		}
	}
	return fmt.Errorf("%s", "Cannot add or update a child row")
}

func (d *dataSource) Delete(id string) error {
	db := d.Driver.GetPqDB()
	ctx := context.Background()
	rs, err := db.NewDelete().Table("roles").Where("id = ?", id).Exec(ctx)
	if err == nil {
		if row, e := rs.RowsAffected(); e == nil && row > 0 {
			return nil
		}
	}
	return fmt.Errorf("%s", "Cannot delete a child row")
}

func NewDataSource(driver database.Drivers) DataSource {
	return &dataSource{
		Driver: driver,
	}
}
