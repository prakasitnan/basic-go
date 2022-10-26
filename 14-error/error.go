package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
)

func ReadFile(name string) (string, error) {
	// read file
	var data string
	var err error

	randNum := rand.Intn(100)
	log.Print(randNum)
	if randNum == 1 {
		data = "file......."
		err = nil
	} else {
		data = ""
		err = errors.New("file not found")
	}
	return data, err
}

func main() {
	data, err := ReadFile("profile.txt")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("read file success :", data)
}
