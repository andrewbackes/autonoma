package controller

import (
	"fmt"
	"github.com/andrewbackes/autonoma/engine/controller/actions"
	"github.com/andrewbackes/autonoma/engine/sensor"
	"log"
	"time"
)

type node struct {
	parent   *node
	location sensor.Location
}

type nodeSet map[*node]struct{}

func (c *Controller) ScanArea() {
	log.Println("Scanning area.")
	for i := float64(0); i <= 360; i += 15 {
		/*
			for p := -30.0; p <= 30.0; p += 30 {
				c.send(actions.Look(p))
				time.Sleep(25 * time.Millisecond)
				c.send(actions.Read("irdistance"))
				time.Sleep(25 * time.Millisecond)
			}
		*/
		//c.send(actions.Look(0))
		//time.Sleep(25 * time.Millisecond)
		//c.send(actions.Read("irdistance"))
		//time.Sleep(25 * time.Millisecond)
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
	maneuver := c.generateManeuver(nextPath)
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
		maneuver = c.generateManeuver(nextPath)
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
		neighbors := c.mapReader.NeighboringCells(currentNode.location.X, currentNode.location.Y)
		for neighbor := range neighbors {
			neighborsNeightbors := c.mapReader.NeighboringCells(neighbor.X, neighbor.Y)
			clearPath := true
			for n := range neighborsNeightbors {
				if c.mapReader.IsUnexplored(n) && n != c.location && currentNode.location != c.location {
					log.Println("At", c.location, "Found unexplored area:", n, "Going to ", currentNode.location)
					path := generatePath(currentNode)
					return path
				}
				if !c.mapReader.IsVacant(n) {
					clearPath = false
				}
			}
			if !checked.Contains(neighbor) && c.mapReader.IsVacant(neighbor) && clearPath {
				queue <- &node{location: neighbor, parent: currentNode}
				checked.Add(neighbor)
			}
		}
	}
	return nil
}

func (c *Controller) generateManeuver(path []*node) []string {
	var maneuver []string
	lastPos := &c.location
	for _, pos := range path {
		var heading float64
		if lastPos.X < pos.location.X {
			heading = 270.0
		} else if lastPos.X > pos.location.X {
			heading = 90.0
		} else if lastPos.Y < pos.location.Y {
			heading = 180.0
		} else if lastPos.Y > pos.location.Y {
			heading = 0
		}
		rotate := actions.Rotate(heading)
		maneuver = append(maneuver, rotate)
		move := actions.Move(float64(c.mapReader.GetCellSize()), float64(pos.location.X), float64(pos.location.Y))
		maneuver = append(maneuver, move)
	}
	return maneuver
}

func generatePath(endpoint *node) []*node {
	var path []*node
	currentNode := endpoint
	for currentNode != nil && currentNode.parent != nil {
		path = append([]*node{currentNode}, path...)
		currentNode = currentNode.parent
	}
	log.Println(path)
	return path
}

func (n *node) String() string {
	return fmt.Sprintf("(%d, %d)", n.location.X, n.location.Y)
	if n.parent == nil {
		return fmt.Sprintf("{(%d, %d), parent: nil}", n.location.X, n.location.Y)
	}
	return fmt.Sprintf("{(%d, %d), parent: (%d, %d)}", n.location.X, n.location.Y, n.parent.location.X, n.parent.location.Y)
}
