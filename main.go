package main

import (
	"fmt"
	"github.com/andrewbackes/autonoma/engine"
)

func main() {
	fmt.Println("Autonoma Started.")
	engine.NewEngine().Start()
}
