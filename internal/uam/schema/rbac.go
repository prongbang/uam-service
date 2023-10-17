package schema

import (
	"context"
	"fmt"
	"github.com/prongbang/uam-service/internal/uam/database"
	"github.com/uptrace/bun"
)

const tableRbac = "rbacs"

type RBAC struct {
	ID    string `json:"id" bun:"id,pk,type:uuid,default:uuid_generate_v4()"`
	PType string `json:"pType" bun:"p_type"`
	V0    string `json:"v0" bun:"v0"`
	V1    string `json:"v1" bun:"v1"`
	V2    string `json:"v2" bun:"v2"`
	V3    string `json:"v3" bun:"v3"`
	V4    string `json:"v4" bun:"v4"`
	V5    string `json:"v5" bun:"v5"`
}

var _ bun.AfterCreateTableHook = (*RBAC)(nil)

func (*RBAC) AfterCreateTable(ctx context.Context, query *bun.CreateTableQuery) error {
	_, err := query.DB().NewCreateIndex().
		Model((*RBAC)(nil)).Table(tableRbac).IfNotExists().
		Index("p_type_idx").Column("p_type").
		Index("v0_idx").Column("v0").
		Index("v1_idx").Column("v1").
		Index("v2_idx").Column("v2").
		Index("v3_idx").Column("v3").
		Index("v4_idx").Column("v4").
		Index("v5_idx").Column("v5").
		Exec(ctx)
	return err
}

type RBACSchema interface {
	Initial()
}

type rbacSchema struct {
	DbDriver database.Drivers
}

func (u *rbacSchema) Initial() {
	ctx := context.Background()
	db := u.DbDriver.GetPqDB()

	_, err := db.NewCreateTable().Model((*RBAC)(nil)).Table(tableRbac).IfNotExists().Exec(ctx)
	if err != nil {
		fmt.Println("Can't create table", tableRbac, err)
	} else {
		fmt.Println(fmt.Sprintf("Table %s created", tableRbac))
	}
}

func NewRBACSchema(dbDriver database.Drivers) RBACSchema {
	return &rbacSchema{
		DbDriver: dbDriver,
	}
}
