package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func PasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	password = hex.EncodeToString(hash.Sum(nil))
	return string(password)
}

func PasswordVerify(hashedPassword, password string) bool {
	return hashedPassword == PasswordHash(password)
}
