package utils

import (
	"errors"
	"log"
)

var (
	ErrInvalidUserObject    = errors.New("invalid user object")
	ErrUnableToHashPasswrod = errors.New("unable to create hash password, please verify")
	ErrUnableToRegisterUser = errors.New("unable to register user in database, please verify")
	ErrInvalidUserData      = errors.New("invalid user login data, please verify")
)

func CheckPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func CheckFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
