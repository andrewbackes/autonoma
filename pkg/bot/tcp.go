package bot

import (
	"bufio"
	"fmt"
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
	_, err := t.conn.Write([]byte(msg))
	if err != nil {
		panic(err)
	}
	fmt.Println("sent:", msg)
}

func (t *tcpSendReceiver) receive() string {
	fmt.Println(t.conn)
	msg, err := bufio.NewReader(t.conn).ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Println("recieved:", msg)
	return msg
}

func (t *tcpSendReceiver) sendAndReceive(msg string) string {
	t.send(msg)
	return t.receive()
}
