package main

import (
	"fmt"
	"log"
	"math/rand"
)

func OnErrorPanic(err error, helpText string) {
	if err != nil {
		log.Fatalf("%s: \n, %v", helpText, err)
		panic(fmt.Sprintf("%s: \n, %v", helpText, err))
	}
}

func SimulateDownServer(lb *LoadBalancer) bool {

	if !lb.Config.RandomServerOff {
		return true
	}

	trueFalseArr := []bool{true, false}
	idx := rand.Intn(len(trueFalseArr))

	return trueFalseArr[idx]
}
