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

var defaultDist = 15.0

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
	time.Sleep(500 * time.Millisecond)
	nextPath := c.pathToUnexploredArea()
	log.Println("Next path:", nextPath)
	for len(nextPath) > 0 {
		for _, action := range nextPath {
			c.send(action)
			time.Sleep(10 * time.Millisecond)
		}
		c.ScanArea()
		time.Sleep(500 * time.Millisecond)
		nextPath = c.pathToUnexploredArea()
		log.Println("Next path:", nextPath)
	}
	log.Println("Area is completely explored.")
}

func (c *Controller) pathToUnexploredArea() []string {
	maxQueueSize := 100000
	queue := make(chan node, maxQueueSize)
	queue <- node{location: c.location}
	visited := sensor.NewLocationSet()
	for len(queue) != 0 {
		node := <-queue
		if c.mapReader.IsUnexplored(node.location) {
			log.Println("New destination", node.location)
			return generateManeuver(&node)
		}
		visited.Add(node.location)
		neighbors := c.neighbors(&node, visited)
		log.Println("Neighbors:", neighbors)
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
		log.Println("Maneuver", maneuver)
		node = node.parent
	}
	return maneuver
}

func (c *Controller) neighbors(pos *node, visited sensor.LocationSet) nodeSet {
	nodes := nodeSet{}
	for i := float64(0); i <= 360; i += 45 {
		loc := util.LocationOf(pos.location, i, defaultDist)
		log.Println(loc, "Contains:", visited.Contains(loc), "Unexplored:", c.mapReader.IsUnexplored(loc))
		if !visited.Contains(loc) && c.mapReader.IsUnexplored(loc) {
			node := node{parent: pos, heading: i, location: loc, distance: defaultDist}
			nodes[node] = struct{}{}
		}
	}
	return nodes
}
