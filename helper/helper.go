package helper

import (
	"fmt"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func GetENV(path string) (map[string]string, error) {
	env, err := godotenv.Read(path)
	if err != nil {
		return env, fmt.Errorf("error get env : %v", err.Error())
	}

	return env, nil
}

func EcryptPassword(password string) (string, error) {
	// encrypt password
	passENC, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", fmt.Errorf("failed encrypt password : %v", err.Error())
	}

	return string(passENC), nil
}

func VerifyPassword(rawPassword, passwordHash string) error {
	// cocokan password input dan password dari db
	if err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(rawPassword)); err != nil {
		return fmt.Errorf("wrong password")
	}

	return nil
}
