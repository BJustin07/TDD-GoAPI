package main

import (
	"TDD-GoAPI/config"
	"log"
)

func main() {
	config.LoadEnv()
	err := config.ConnectToDatabase()
	if err != nil {
		log.Fatal(err)
	}

}
