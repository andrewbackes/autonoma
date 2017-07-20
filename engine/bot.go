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
	conn        net.Conn
	moveHistory []BotMove
}

func NewBot() *Bot {
	return &Bot{
		moveHistory: make([]BotMove, 0),
	}
}

type BotMove struct {
	dir string
	t   time.Duration
	pwr int64
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

func (b *Bot) Make(m BotMove) {
	if b.conn == nil {
		b.connect()
	}
	msg := format(m)
	_, err := b.conn.Write(msg)
	if err != nil {
		fmt.Println(err)
	}
	b.updateHistory(m)
}

func (b *Bot) updateHistory(m BotMove) {
	b.moveHistory = append(b.moveHistory, m)
	if len(b.moveHistory) > 1000 {
		b.moveHistory = b.moveHistory[len(b.moveHistory)-1000:]
	}
}

func format(m BotMove) []byte {
	sec := m.t.Seconds()
	s := strconv.FormatFloat(sec, 'E', -1, 64)
	p := strconv.FormatInt(m.pwr, 10)
	msg := fmt.Sprintf("move %s %s %s ", m.dir, s, p)
	return []byte(msg)
}
