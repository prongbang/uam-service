package auth

import (
	"github.com/prongbang/user-service/internal/shared/role"
)

type Token struct {
}

type Credential struct {
	Token       string            `json:"token"`
	Role        []role.Role       `json:"roles"`
	Permissions []role.Permission `json:"permissions"`
}
