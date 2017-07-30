package controller

import (
	"github.com/andrewbackes/autonoma/engine/controller/actions"
	"github.com/andrewbackes/autonoma/engine/sensor"
	"github.com/andrewbackes/autonoma/engine/util"
	"log"
	"time"
)

func (c *Controller) ScanArea() {
	for i := float64(0); i <= 360; i += 45 {
		for p := -15.0; p <= 15.0; p += 15 {
			c.send(actions.Look(p))
			time.Sleep(5 * time.Millisecond)
			c.send(actions.Read("irdistance"))
			time.Sleep(5 * time.Millisecond)
		}
		c.send(actions.Rotate(i))
		time.Sleep(5 * time.Millisecond)
		c.send(actions.Read("all"))
		time.Sleep(5 * time.Millisecond)
	}
	c.send(actions.Look(0))
}

var defaultDist = 10.0

type node struct {
	parent   *node
	heading  float64
	location sensor.Location
	distance float64
}

type nodeSet map[node]struct{}

func (c *Controller) explore() {
	log.Println("Exploring area.")
	c.ScanArea()
	time.Sleep(1000 * time.Millisecond)
	nextPath := c.pathToUnexploredArea()
	log.Println("Next path:", nextPath)
	for len(nextPath) > 0 {
		for _, action := range nextPath {
			c.send(action)
			time.Sleep(10 * time.Millisecond)
		}
		c.ScanArea()
		time.Sleep(1000 * time.Millisecond)
		nextPath = c.pathToUnexploredArea()
		log.Println("Next path:", nextPath)
	}
	log.Println("Area is completely explored.")
}

func (c *Controller) pathToUnexploredArea() []string {
	maxQueueSize := 1000000
	queue := make(chan node, maxQueueSize)
	queue <- node{location: c.location}
	checked := sensor.NewLocationSet()
	for len(queue) != 0 {
		log.Println("Queue size:", len(queue))
		node := <-queue
		if c.mapReader.IsUnexplored(node.location) {
			log.Println("New destination", node.location)
			return generateManeuver(&node)
		}
		checked.Add(node.location)
		log.Println("Checked size:", len(checked), " - ", checked)
		//log.Println("Adding checked node:", node.location)
		neighbors := c.neighbors(&node, checked)
		//log.Println("Neighbors:", neighbors)
		for neighbor := range neighbors {
			queue <- neighbor
		}
	}
	return nil
}

func generateManeuver(endpoint *node) []string {
	var maneuver []string
	node := endpoint
	for node != nil && node.parent != nil {
		rotate := actions.Rotate(node.heading)
		move := actions.Move(defaultDist, float64(node.location.X), float64(node.location.Y))
		maneuver = append([]string{rotate, move}, maneuver...)
		node = node.parent
	}
	if len(maneuver) > 2 {
		maneuver = maneuver[:len(maneuver)-2]
	}
	log.Println("Maneuver", maneuver)
	return maneuver
}

func (c *Controller) neighbors(pos *node, checked sensor.LocationSet) nodeSet {
	nodes := nodeSet{}
	for i := float64(0); i <= 360; i += 45 {
		loc := util.LocationOf(pos.location, i, defaultDist)
		if !checked.Contains(loc) && !c.mapReader.IsOccupied(loc) {
			//log.Println("Neighbors checked:", checked)
			//log.Println(loc, "Checked:", checked.Contains(loc), "Occupied:", c.mapReader.IsOccupied(loc))
			node := node{parent: pos, heading: i, location: loc, distance: defaultDist}
			nodes[node] = struct{}{}
			//log.Println("Adding neighbor:", node.location)
		}
	}
	return nodes
}
