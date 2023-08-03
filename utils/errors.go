package utils

import "log"

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