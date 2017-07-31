package controller

import (
	"fmt"
	"github.com/andrewbackes/autonoma/engine/controller/actions"
	"github.com/andrewbackes/autonoma/engine/sensor"
	"github.com/andrewbackes/autonoma/engine/util"
	"log"
	"time"
)

var defaultDist = 5.0

type node struct {
	parent   *node
	heading  float64
	location sensor.Location
	distance float64
}

type nodeSet map[*node]struct{}

func (c *Controller) ScanArea() {
	log.Println("Scanning area.")
	for i := float64(0); i <= 360; i += 15 {
		for p := -30.0; p <= 30.0; p += 30 {
			c.send(actions.Look(p))
			time.Sleep(25 * time.Millisecond)
			c.send(actions.Read("irdistance"))
			time.Sleep(25 * time.Millisecond)
		}
		c.send(actions.Rotate(i))
		time.Sleep(25 * time.Millisecond)
		c.send(actions.Read("all"))
		time.Sleep(25 * time.Millisecond)
	}
	c.send(actions.Look(0))

}

func (c *Controller) explore() {
	log.Println("Exploring area.")
	time.Sleep(500 * time.Millisecond)
	c.ScanArea()
	time.Sleep(500 * time.Millisecond)
	nextPath := c.pathToUnexploredArea()
	maneuver := generateManeuver(nextPath)
	log.Println("Maneuver", maneuver)
	for len(nextPath) > 0 {
		for _, action := range maneuver {
			c.send(action)
			time.Sleep(250 * time.Millisecond)
		}
		c.location = nextPath[len(nextPath)-1].location
		log.Println("Location", c.location)
		c.ScanArea()
		time.Sleep(500 * time.Millisecond)
		nextPath = c.pathToUnexploredArea()
		maneuver = generateManeuver(nextPath)
		log.Println("Maneuver", maneuver)
	}
	log.Println("Area is completely explored.")
}

func (c *Controller) pathToUnexploredArea() []*node {
	maxQueueSize := 1000000
	queue := make(chan *node, maxQueueSize)
	queue <- &node{location: c.location}
	checked := sensor.NewLocationSet()
	for len(queue) != 0 {
		currentNode := <-queue
		unexplored := c.mapReader.IsUnexplored(currentNode.location)
		if unexplored {
			log.Println("Found unexplored area:", currentNode.location)
			path := generatePath(currentNode)
			return path
		}
		neighbors := c.neighbors(currentNode)
		for neighbor := range neighbors {
			if !checked.Contains(neighbor.location) && !c.mapReader.IsOccupied(neighbor.location) {
				queue <- neighbor
				checked.Add(neighbor.location)
			}
		}
	}
	return nil
}

func generateManeuver(path []*node) []string {
	var maneuver []string
	for _, pos := range path {
		rotate := actions.Rotate(pos.heading)
		maneuver = append(maneuver, rotate)
		move := actions.Move(defaultDist, float64(pos.location.X), float64(pos.location.Y))
		maneuver = append(maneuver, move)
	}
	return maneuver
}

/*
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
*/

func generatePath(endpoint *node) []*node {
	var path []*node
	currentNode := endpoint
	for currentNode != nil && currentNode.parent != nil {
		path = append([]*node{currentNode}, path...)
		currentNode = currentNode.parent
	}
	// we want to go right in front of the unexplored area
	if len(path) > 1 {
		path = path[:len(path)-1]
	}
	return path
}

func (c *Controller) neighbors(pos *node) nodeSet {
	nodes := nodeSet{}
	for i := float64(0); i <= 360; i += 45 {
		loc := util.LocationOf(pos.location, i, defaultDist)
		node := &node{parent: pos, heading: i, location: loc, distance: defaultDist}
		nodes[node] = struct{}{}
	}
	return nodes
}

func (n *node) String() string {
	return fmt.Sprintf("(%d, %d)", n.location.X, n.location.Y)
	if n.parent == nil {
		return fmt.Sprintf("{(%d, %d), parent: nil}", n.location.X, n.location.Y)
	}
	return fmt.Sprintf("{(%d, %d), parent: (%d, %d)}", n.location.X, n.location.Y, n.parent.location.X, n.parent.location.Y)
}
