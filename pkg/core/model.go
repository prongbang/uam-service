package core

import "time"

type Model struct {
	CreatedAt time.Time `json:"createdAt" db:"created_at" bun:"created_at,default:current_timestamp"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at" bun:"updated_at"`
}
