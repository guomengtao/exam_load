package utils

import "golang.org/x/crypto/bcrypt"

 
func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed), err
}

func CheckPassword(input string, hashed string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(input))
    return err == nil
}