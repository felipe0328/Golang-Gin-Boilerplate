package utils

import (
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func ConvertStringToHash(input string) (hashedInput string, err error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	hashedInput = string(hashedBytes)

	return
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(password), []byte(hashedPassword))
}

func RemoveSpacesFromInput(input string)(output string){
	return html.EscapeString(strings.TrimSpace(input))
}