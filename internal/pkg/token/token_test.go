package token_test

import (
	"fmt"
	"github.com/prongbang/user-service/internal/pkg/token"
	"testing"
	"time"
)

func TestGenerateKeyString(t *testing.T) {
	key, err := token.GenerateKeyString(token.AES256KeySize)

	if err != nil {
		t.Error("Error", err)
	}
	fmt.Println(key)
}

func TestNew(t *testing.T) {
	payload := token.Claims{
		Exp: time.Now().AddDate(1000, 0, 0).Unix(),
		Sub: "d4a45b08-825b-4c5f-8a63-fa8991dd0945",
		Roles: []string{
			"1a55b471-78b8-45f6-a548-6e436a002619",
			"2ed000b0-c93d-42cb-8def-77138584778a",
		},
	}
	keyHex, _ := token.GenerateKeyString(token.AES256KeySize)
	key, _ := token.HexToBytes(keyHex)

	actual, err := token.New(payload, key)

	if err != nil {
		t.Error("Error", err)
	}
	fmt.Println(keyHex)
	fmt.Println(actual)
}

func TestVerify(t *testing.T) {
	sub := "d4a45b08-825b-4c5f-8a63-fa8991dd0945"
	key, _ := token.HexToBytes("fe5a40559cdf1a4e4b38b72acb8c601a0ecb6014d3300d3541f9b93b707fc1dd")
	jweCompact := []byte("eyJhbGciOiJBMjU2R0NNS1ciLCJlbmMiOiJBMjU2R0NNIiwiaXYiOiJ3WjZVdmViUzF5VF9PTWFMIiwidGFnIjoiX3NhYW0ySmxJcVk2WnVwZ2ZkTTlMdyJ9.jgeYY1-ZpUXBTRil-5ikZlh5SZvF-1UvfRMBAi0dqNc.uG9JcKID3yKFpVQg.wiVZZ8ht3EyYE1Q_saareNchRob9SymyNXE_xw7F6I85HAxwPUv0fIsNJQVEzO49vWz9F9YecO5k1rDRsjSNMcLdww3aUfSQBgVV7xhWK0KmELAeBnQx0w.ISc500xZjrAiBE6GE-qTxQ")

	actual, err := token.Verify(jweCompact, key)

	if err != nil || actual.Sub != sub {
		t.Error(err)
	}
	fmt.Println(sub, actual)
}
