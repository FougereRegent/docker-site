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
	saltPassword := fmt.Sprintf("%s%s", password, salt)
	hashPassword := sha256.Sum256([]byte(saltPassword))

	return HashedPassword{
		Digest: string(hashPassword[:]),
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
	saltPassword := fmt.Sprintf("%s%s", input_password, hashes.Salt)
	hashPassword := sha256.Sum256([]byte(saltPassword))

	return string(hashPassword[:]) == hashes.Digest
}
