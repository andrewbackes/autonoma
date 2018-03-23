package bot

import (
	"bufio"
	log "github.com/sirupsen/logrus"
	"net"
)

type sendReceiver interface {
	receive() string
	send(string)
}

type tcpSendReceiver struct {
	conn    net.Conn
	address string
}

func (t *tcpSendReceiver) connect() {
	log.Info("Connecting to ", t.address)
	conn, err := net.Dial("tcp", t.address)
	if err != nil {
		panic(err)
	}
	t.conn = conn
}

func (t *tcpSendReceiver) send(msg string) {
	if t.conn == nil {
		t.connect()
	}
	t.ready()
	_, err := t.conn.Write([]byte(msg + "\n"))
	if err != nil {
		panic(err)
	}
	log.Info("Sent:", msg)
}

func (t *tcpSendReceiver) ready() {
	log.Info("Ready?")
	_, err := t.conn.Write([]byte(`{"command":"isready"}\n`))
	if err != nil {
		panic(err)
	}
	resp := t.receive()
	if resp != "readyok" {
		panic("got " + resp + " wanted 'readyok'")
	}
}

func (t *tcpSendReceiver) receive() string {
	msg, err := bufio.NewReader(t.conn).ReadString('\n')
	if err != nil {
		panic(err)
	}
	log.Info("Recieved:", msg)
	return msg
}
