package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed)
}

func CheckPassword(input string, hashed string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(input))
    return err == nil
}