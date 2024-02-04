package main

import (
	"fmt"
	"log"
)

func OnErrorPanic(err error, helpText string) {
	if err != nil {
		log.Fatalf("%s: \n, %v", helpText, err)
		panic(fmt.Sprintf("%s: \n, %v", helpText, err))
	}
}
