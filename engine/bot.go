package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

const (
	botAddr = "10.0.0.14:9091"
)

type Bot struct {
	conn net.Conn
}

func NewBot() *Bot {
	return &Bot{}
}

func (b *Bot) connect() error {
	if b.conn != nil {
		b.conn.Close()
	}
	c, err := net.DialTimeout("tcp", botAddr, 15*time.Second)
	if err != nil {
		fmt.Println("Could't connect to bot")
		return err
	}
	b.conn = c
	return nil
}

func (b *Bot) Move(dir string, t time.Duration, pwr int64) {
	if b.conn == nil {
		b.connect()
	}
	msg := formatMove(dir, t, pwr)
	_, err := b.conn.Write(msg)
	if err != nil {
		fmt.Println(err)
	}
}

func formatMove(dir string, t time.Duration, pwr int64) []byte {
	sec := t.Seconds()
	s := strconv.FormatFloat(sec, 'E', -1, 64)
	p := strconv.FormatInt(pwr, 10)
	msg := fmt.Sprintf("move %s %s %s ", dir, s, p)
	return []byte(msg)
}
