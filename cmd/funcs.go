package main

import (
	"fmt"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
)

const (
	longTaskCommand = "/long"
)

type controller struct {
	b *tg.BotAPI
}

func (c controller) handleMSG(msg *tg.Message) {
	if msg.IsCommand() {
		c.handleCommand(msg)
	} else {

	}
}

func (c controller) handleCommand(msg *tg.Message) {
	switch {
	case strings.Contains(msg.Text, longTaskCommand):
		split := strings.Split(msg.Text, " ")
		delay := time.Second
		if len(split) < 2 {
			c.longTask(delay, msg.Chat.ID)
			return
		}
		number, err := strconv.ParseInt(split[1], 10, 64)
		if err != nil {
			c.longTask(delay, msg.Chat.ID)
			return
		}
		delay, err = time.ParseDuration(fmt.Sprintf("%d", number) + "s")
		if err != nil {
			c.longTask(delay, msg.Chat.ID)
			return
		}

		c.longTask(delay, msg.Chat.ID)
	}
}

func (c controller) longTask(d time.Duration, chatID int64) {
	time.Sleep(d)
	_, err := c.b.Send(tg.NewMessage(chatID, fmt.Sprintf("time: %s", time.Now())))
	if err != nil {
		log.Debugf("err send msg: %s", err)
	}
}
