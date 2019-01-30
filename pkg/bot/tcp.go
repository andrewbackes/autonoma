package bot

import (
	"bufio"
	log "github.com/sirupsen/logrus"
	"io"
	"net"
	"strings"
	"time"
)

type sendReceiver interface {
	receive() (string, error)
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
	log.Debug("Sent:", msg)
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

func (t *tcpSendReceiver) receive() (string, error) {
	t.conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	msg, err := bufio.NewReader(t.conn).ReadString('\n')
	if err != nil {
		log.Error(err)
		t.conn.Close()
		t.conn = nil
		return "", err
	}
	log.Debug("Recieved:", msg[:len(msg)-1])
	return msg, nil
}
