package bot

import (
	"fmt"
	"testing"
)

type mock struct {
	sent     []string
	received string
}

func (m *mock) send(msg string) {
	m.sent = append(m.sent, msg)
}

func (m *mock) receive() string {
	return `{"heading":0}`
}

func TestRotate(t *testing.T) {
	m := &mock{sent: make([]string, 0)}
	bot := Bot{sendReceiver: m}
	bot.Rotate(0.0)
	fmt.Println(m.sent)
}
