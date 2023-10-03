package cryptox_test

import (
	"fmt"
	"github.com/prongbang/user-service/pkg/cryptox"
	"testing"
)

func TestHashPassword(t *testing.T) {
	// Given
	password := "examplePassword"

	// When
	actual := cryptox.HashPassword(password)

	// Then
	if actual == "" {
		t.Error("Error hash password:", actual)
	}
	fmt.Println(actual)
}

func TestVerifyPassword(t *testing.T) {
	// Given
	password := "examplePassword"
	hashedPassword := "SUY1e1Km00wK4YgsuzGme81nEBwdwIwmHQ8ZzVZ3j+M:O9ViWvfaRHedpI5JrZbzmQ"

	// When
	actual := cryptox.VerifyPassword(password, hashedPassword)

	// Then
	if !actual {
		t.Error("Error verify password:", actual)
	}
}
