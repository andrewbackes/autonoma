package main

import (
	"github.com/andrewbackes/autonoma/engine"
	"log"
)

func main() {
	log.Println("Autonoma Started.")
	engine.NewEngine().Start()
}
