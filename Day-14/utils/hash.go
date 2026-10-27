package utils

import "golang.org/x/crypto/bcrypt"

func PasswordHash(text string) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(text), 14)
	return string(password), err
}

func CheckPasswordHash(password, hashPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}
