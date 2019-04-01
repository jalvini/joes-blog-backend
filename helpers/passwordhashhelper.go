package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(hash)
}

func ComparePassword(hash, password string) bool {
	checker := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return (checker == nil)
}
