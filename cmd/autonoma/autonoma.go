package main

import (
	"github.com/andrewbackes/autonoma/pkg/bot/v3"
	"github.com/andrewbackes/autonoma/pkg/os"
	"github.com/andrewbackes/autonoma/pkg/ui"
)

func main() {
	b := &v3.Bot{}
	botOS := os.New(b)
	webUI := ui.New(botOS)
	go webUI.ListenAndServe()
	botOS.Start()
}
