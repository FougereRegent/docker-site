package helper

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
)

type HashedPassword struct {
	Digest string
	Salt   string
}

func HashPassword(password string) HashedPassword {
	salt := saltGenerator()
	salt_password := fmt.Sprintf("%s%s", password, salt)
	hash_password := sha256.Sum256([]byte(salt_password))

	return HashedPassword{
		Digest: string(hash_password[:]),
		Salt:   salt,
	}
}

func saltGenerator() string {
	var result [10]byte
	var letters = []rune("azertyuiopqsdfghjklmwxcvbnAZERTYUIOPQSDFGHJKLMWXCVBN")
	for i := 0; i < len(result); i++ {
		result[i] = byte(letters[rand.Intn(len(letters))])
	}
	return string(result[:])
}

func CheckPassword(input_password string, hashes *HashedPassword) bool {
	salt_password := fmt.Sprintf("%s%s", input_password, hashes.Salt)
	hash_password := sha256.Sum256([]byte(salt_password))

	return string(hash_password[:]) == hashes.Digest
}
