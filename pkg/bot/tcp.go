package bot

import (
	"bufio"
	"net"
)

func (b *Bot) connect() {
	conn, err := net.Dial("tcp", b.address)
	if err != nil {
		panic(err)
	}
	b.conn = conn
}

func (b *Bot) send(msg string) {
	if b.conn == nil {
		b.connect()
	}
	_, err := b.conn.Write([]byte(msg))
	if err != nil {
		panic(err)
	}
}

func (b *Bot) receive() string {
	msg, err := bufio.NewReader(b.conn).ReadString('\n')
	if err != nil {
		panic(err)
	}
	return msg
}

func (b *Bot) sendAndReceive(msg string) string {
	b.send(msg)
	return b.receive()
}
