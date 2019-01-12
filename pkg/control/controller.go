package control

import (
	"github.com/andrewbackes/autonoma/pkg/planning"
)

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) Execute(*planning.Motions) *Commands {
	return &Commands{}
}
