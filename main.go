package main

import (
	"TDD-GoAPI/config"
)

func main() {
	config.LoadEnv()
	err := config.StartServer()
	if err != nil {
		return
	}
}
