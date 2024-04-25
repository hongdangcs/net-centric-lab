package lab3

import (
	"crypto/sha256"
	"encoding/base64"
)

func EncryptPassword(password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password))
	return base64.StdEncoding.EncodeToString(hasher.Sum(nil))
}
