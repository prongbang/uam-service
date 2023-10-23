package user_creator

import "github.com/uptrace/bun"

type CreateUserCreator struct {
	bun.BaseModel `bun:"table:users_creators,alias:u"`
	ID            *string `json:"id" bun:"id,pk,type:uuid"`
	UserID        string  `json:"userId" bun:"user_id,type:uuid"`
	CreatedBy     string  `json:"createdBy" bun:"created_by,type:uuid"`
}
