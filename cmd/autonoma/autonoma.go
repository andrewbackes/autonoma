package main

import (
	"github.com/andrewbackes/autonoma/pkg/autonoma"
	"github.com/andrewbackes/autonoma/pkg/bot/comm"
	"github.com/andrewbackes/autonoma/pkg/bot/specs"
	"github.com/andrewbackes/autonoma/pkg/os"
	"github.com/andrewbackes/autonoma/pkg/ui"
)

const (
	//botAddress = "192.168.86.74:9091"
	botAddress = "localhost:9091"
)

func main2() {
	// Define communications to the bot.
	com := comm.New(botAddress)
	spec := specs.Spec{}

	botOS := os.New(com, spec)
	webUI := ui.New(botOS)
	go webUI.ListenAndServe()
	botOS.Start()
}

func main() {
	a := autonoma.New()

	a.Start()
}

/*

Select a mission
mission selects the next target destination
pathing router determines the sequence of coordinates
motion planner determines how to manipulate the motors to get to the next coordinate
scans map world and locates bot
reroute

--

## Forward to wall Mission:
input: none
scan (manuever)
locate
update world
until 12 inches from wall:
	move forward small amount (manuever)
loop


## Manual Navigation Mission:
input: manual commands
scan (manuever)
locate
update world
until done:
	execute manual moves (manuever)
loop


## Go to destination Mission:
input: destination point
scan (manuever)
locate
update world
until close to destination point:
	find a route
	rotate toward next point (manuever)
	move forward to next point (manuever)



## Map the whole world
input: none
scan (manuever)
locate
update world
until world is all identified
	identify point to scan unidentified area
	run mission	 "go to destination"


Mission
 NextManuever(World, position, orientation)

VehicleController
 PerformManuever(m)


*/
