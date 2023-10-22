package token

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-jose/go-jose/v3"
	"github.com/prongbang/uam-service/pkg/common"
	"os"
	"time"
)

const (
	AES256KeySize = 32
	Expired       = "expired"
)

type AccessToken struct {
	Token string `json:"token"`
}

type Claims struct {
	Exp   int64    `json:"exp"`
	Iss   string   `json:"iss"`
	Sub   string   `json:"sub"`
	Roles []string `json:"roles"`
}

func Parse(req any) *AccessToken {
	if tk := common.AnyToStruct[AccessToken](req); tk != nil && tk.Token != "" {
		return tk
	}
	return nil
}

func GetKeyBytes() ([]byte, error) {
	return common.HexToBytes(GetKey())
}

func GetKey() string {
	return os.Getenv("JWE_SECRET")
}

func GenerateKey(keySize int) ([]byte, error) {
	key := make([]byte, keySize)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func GenerateKeyString(keySize int) (string, error) {
	key, err := GenerateKey(keySize)
	if err != nil {
		return "", err
	}
	return common.BytesToHex(key), nil
}

func Payload(jwe string) (*Claims, error) {
	key, err := GetKeyBytes()
	if err != nil {
		return nil, err
	}
	return GetPayload([]byte(jwe), key)
}

func Verification(jwe string) (*Claims, error) {
	key, err := GetKeyBytes()
	if err != nil {
		return nil, err
	}
	return Verify([]byte(jwe), key)
}

func GetPayload(jwe, key []byte) (*Claims, error) {
	payloadBytes, err := Decrypt(jwe, key)
	if err != nil {
		return nil, err
	}

	payload := Claims{}
	err = json.Unmarshal([]byte(payloadBytes), &payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}

func New(payload Claims, key []byte) (string, error) {
	// Create a new JSON Web Encryption (JWE) encrypter with the shared key.
	encrypter, err := jose.NewEncrypter(jose.A256GCM, jose.Recipient{Algorithm: jose.A256GCMKW, Key: key}, nil)
	if err != nil {
		fmt.Println("Error creating JWE encrypter:", err)
		return "", err
	}

	// Encrypt the payload.
	payloadBytes, _ := json.Marshal(payload)
	jwe, err := encrypter.Encrypt(payloadBytes)
	if err != nil {
		fmt.Println("Error encrypting payload:", err)
		return "", err
	}

	// Serialize the JWE
	jweCompact, err := jwe.CompactSerialize()
	if err != nil {
		fmt.Println("Error serializing JWE:", err)
		return "", err
	}

	return jweCompact, nil
}

func Decrypt(jweCompact, key []byte) (string, error) {
	// Parse the JWE
	jwe, err := jose.ParseEncrypted(string(jweCompact))
	if err != nil {
		fmt.Println("Error parsing JWE token:", err)
		return "", err
	}

	// Decrypt the payload using the shared symmetric key.
	decryptedPayload, err := jwe.Decrypt(key)
	if err != nil {
		fmt.Println("Error decrypting payload:", err)
		return "", err
	}

	return string(decryptedPayload), nil
}

func Verify(jweCompact, key []byte) (*Claims, error) {
	payload, err := GetPayload(jweCompact, key)
	if err != nil {
		return nil, err
	}

	// Convert the "exp" claim value to a Unix timestamp
	expirationTime := time.Unix(payload.Exp, 0)

	// Get the current time
	currentTime := time.Now()

	// Check if the token has expired
	if currentTime.After(expirationTime) {
		return nil, errors.New(Expired)
	}
	return payload, err
}
