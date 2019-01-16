package v3

import (
	"github.com/andrewbackes/autonoma/pkg/perception/signal"
)

type SignalHandler func(*signal.Signal)

func (b *Bot) Listen(h SignalHandler) {}
