package utils

import (
	"errors"
	"fmt"
	"log"
)

var (
	ErrInvalidUserObject     = errors.New("invalid user object")
	ErrUnableToHashPassword  = errors.New("unable to create hash password, please verify")
	ErrUnableToRegisterUser  = errors.New("unable to register user in database, please verify")
	ErrInvalidUserData       = errors.New("invalid user login data, please verify")
	ErrUnableToCreateUser    = errors.New("invalid to create user")
	ErrInvalidLogin          = errors.New("invalid login")
	ErrUserNotFound          = errors.New("user not found")
	ErrInvalidPassword       = errors.New("invalid password")
	ErrUnableToGenerateToken = errors.New("unable to generate token")
)

func StringWrapError(errorMessage, err error) string {
	return fmt.Sprintf("%s: %s", errorMessage.Error(), err.Error())
}

func WrapError(errorMessage, err error) error {
	return fmt.Errorf("%s: %s", errorMessage.Error(), err.Error())
}

func CheckPanic(err error) {
	if err != nil {
		LogError(err)
		panic(err)
	}
}

func CheckFatal(err error) {
	if err != nil {
		LogError(err)
		log.Fatal(err)
	}
}
