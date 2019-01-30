package control

import (
	"github.com/andrewbackes/autonoma/pkg/planning"
)

type sender interface {
	Send(Commands)
}

func Execute(m *planning.Motions, s sender) {
	return
}
