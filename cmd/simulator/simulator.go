package main

import (
	"fmt"
	"github.com/andrewbackes/autonoma/pkg/simulator"
	"os"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Simulator started.")
	if len(os.Args) < 2 {
		fmt.Println("You must provide a sequence file path.")
		os.Exit(1)
	}
	sequenceFile := os.Args[1]
	sequenceDelay := time.Second
	if len(os.Args) >= 3 {
		d, err := strconv.Atoi(os.Args[2])
		if err != nil {
			panic(err)
		}
		sequenceDelay = time.Duration(d) * time.Second
	}
	sim := simulator.New(sequenceFile, sequenceDelay)
	sim.Listen()
}
