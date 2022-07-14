package main

import (
	"log"

	"github.com/koba1108/go-mongodb/internals/domain/model"
)

func main() {
	log.Println("test")
	newUser, err := model.NewUser("ykoba", "ykoba@ykoba.com")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(newUser)
}
