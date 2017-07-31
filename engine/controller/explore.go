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
	//log.Println("Finding path.")
	nextPath := c.pathToUnexploredArea()
	//log.Println("First path:", nextPath)
	for len(nextPath) > 0 {
		for _, action := range nextPath {
			c.send(action)
			time.Sleep(250 * time.Millisecond)
		}
		c.ScanArea()
		time.Sleep(500 * time.Millisecond)
		nextPath = c.pathToUnexploredArea()
		//log.Println("Next path:", nextPath)
	}
	log.Println("Area is completely explored.")
}

func (c *Controller) pathToUnexploredArea() []string {
	maxQueueSize := 1000000
	queue := make(chan *node, maxQueueSize)
	queue <- &node{location: c.location}
	checked := sensor.NewLocationSet()
	for len(queue) != 0 {
		//log.Println("Queue size:", len(queue))
		currentNode := <-queue
		unexplored := c.mapReader.IsUnexplored(currentNode.location)
		if unexplored {
			log.Println("Found unexplored area:", currentNode.location)
			return generateManeuver(currentNode)
		}
		//checked.Add(currentNode.location)
		neighbors := c.neighbors(currentNode)
		//log.Println("Visited:", currentNode.location, "Neighbors:", neighbors)
		for neighbor := range neighbors {
			//log.Println("Checking:", neighbor)
			//log.Println(!checked.Contains(neighbor.location) && !c.mapReader.IsOccupied(neighbor.location))
			if !checked.Contains(neighbor.location) && !c.mapReader.IsOccupied(neighbor.location) {
				//log.Println("From", currentNode.location, "Queueing", neighbor.location, "path:", generatePath(neighbor))
				queue <- neighbor
				checked.Add(neighbor.location)
			}
			//log.Println("Checked.")
		}
		//log.Println("Done checking neighbors.")
	}
	//log.Println("leaving pathToUnexploredArea()")
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

func generatePath(endpoint *node) []*node {
	var path []*node
	currentNode := endpoint
	for currentNode != nil && currentNode.parent != nil {
		path = append([]*node{currentNode}, path...)
		currentNode = currentNode.parent
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
