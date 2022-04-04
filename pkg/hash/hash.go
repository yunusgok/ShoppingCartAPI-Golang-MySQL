//Package hash provides hashing operations for password hashing and controlling
package hash

import (
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

//charset contains characters to create random string as salt
const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

//seededRand random generator created with seed which is current unix time
var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

// CreateSalt creates random string with predefined charset
func CreateSalt() string {
	b := make([]byte, bcrypt.MaxCost)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// HashPassword created hashed string with given password
// bcrypt used as hashing algorithm
// return a string with a size of 31 as a max size
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash compares given password and hash
// returns if they match
func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
