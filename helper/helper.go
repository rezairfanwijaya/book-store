package helper

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type response struct {
	Meta meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type meta struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
}

func ResponseAPI(status string, code int, data interface{}) *response {
	meta := meta{
		Code:   code,
		Status: status,
	}

	responseAPI := response{
		Meta: meta,
		Data: data,
	}

	return &responseAPI
}

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
		return fmt.Errorf("password salah")
	}

	return nil
}

// format error
func ErrorFormater(err error) []string {
	var myError []string

	for _, e := range err.(validator.ValidationErrors) {
		myError = append(myError, e.Error())
	}

	return myError
}
