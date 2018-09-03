package auth

import (
		"bytes"
)

func PasswordMatchesHash(password string, hash string) bool {
	key := []byte(hash[:KEYLENGTH])
	salt := []byte(hash[KEYLENGTH:])

	testKey := generateKey(password, salt)

	if bytes.Equal(key, testKey) {
		return true
	} else {
		return false
	}
}