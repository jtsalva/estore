package auth

import (
	"crypto/rand"

	"golang.org/x/crypto/argon2" // Requires golang.org/x/crypto/blake2b
)

const (
	SALTLENGTH int    = 25
	TIME       uint32 = 3
	MEMORY     uint32 = 32 * 1024
	THREADS    uint8  = 1
	KEYLENGTH  uint32 = 64
)

func generateSalt() ([]byte, error) {
	s := make([]byte, SALTLENGTH)
	_, err := rand.Read(s)
	return s, err
}

func generateKey(password string, salt []byte) []byte {
	return argon2.IDKey([]byte(password), salt, TIME, MEMORY, THREADS, KEYLENGTH)
}

func HashPassword(password string) (string, error) {
	var key []byte

	salt, err := generateSalt()
	if err != nil {
		return string(key), err
	}

	key = generateKey(password, salt)

	// Concatenate key and salt
	hash := string(append(key, salt...))

	return hash, nil
}
