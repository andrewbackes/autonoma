package main

import (
	"github.com/andrewbackes/autonoma/pkg/bot/comm"
	"github.com/andrewbackes/autonoma/pkg/bot/specs"
	"github.com/andrewbackes/autonoma/pkg/os"
	"github.com/andrewbackes/autonoma/pkg/ui"
)

const (
	//botAddress = "192.168.86.74:9091"
	botAddress = "localhost:9091"
)

func main() {
	// Define communications to the bot.
	com := comm.New(botAddress)
	spec := specs.Spec{}

	botOS := os.New(com, spec)
	webUI := ui.New(botOS)
	go webUI.ListenAndServe()
	botOS.Start()
}
