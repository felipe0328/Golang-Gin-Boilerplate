package utils

import (
	"errors"
	"log"
)

var (
	InvalidUserObject = errors.New("invalid user object")
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