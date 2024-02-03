package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
)

// func errorHandler(err error) error {

// }

func LoadFile[T any](filePath string, loadInto *T) error {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal("Could not open file")
		return errors.New("could not open file")
	}

	defer file.Close()

	bytes, err := io.ReadAll(file)

	if err != nil {
		log.Fatal("Could not open file")
		return errors.New("could not open file")
	}

	err = json.Unmarshal(bytes, loadInto)

	if err != nil {
		log.Fatal("Could not open file")
		return errors.New("could not open file")
	}

	return nil

}
