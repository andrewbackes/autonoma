package control

import (
	"github.com/andrewbackes/autonoma/pkg/planning"
)

type bot interface {
	Send(Commands)
}

func Execute(m *planning.Motions, b bot) {
	return
}
