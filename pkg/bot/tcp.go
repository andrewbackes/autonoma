package bot

import (
	"bufio"
	log "github.com/sirupsen/logrus"
	"io"
	"net"
	"strings"
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
	_, err := t.conn.Write([]byte(`{"command": "isready"}` + "\n"))
	if err != nil {
		panic(err)
	}
	resp, err := bufio.NewReader(t.conn).ReadString('\n')
	if err == io.EOF {
		t.connect()
		t.ready()
		return
	}
	if err != nil {
		panic(err)
	}
	if strings.TrimSpace(resp) != `{"status":"readyok"}` {
		panic("got " + resp + ` wanted '{"status":"readyok"}'`)
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
