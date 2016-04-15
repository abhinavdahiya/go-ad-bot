package main

import (
	"regexp"

	"github.com/abhinavdahiya/go-messenger-bot"
	"github.com/abhinavdahiya/go-messenger-dispatcher"
)

type StartState struct {
	Enter, Leave dispatcher.Action

	IsMoved bool
	Chain   string
}

func (m *StartState) Name() string {
	return "start"
}

func (m *StartState) Transit(s string) {
	m.IsMoved = true
	m.Chain = s
}

func (m *StartState) Next() string {
	return m.Chain
}

func (m *StartState) Transitor(c mbotapi.Callback) {
	if m.IsMoved {
		return
	}

	if msg := c.Message; msg.Text != "" {
		if match, _ := regexp.MatchString("(?i).*hi.*", msg.Text); match {
			m.IsMoved = true
			m.Chain = "hi"
			return
		}
	}
}

func (m *StartState) Flush() {
	m.IsMoved = false
	m.Chain = ""
}

func startEnter(c mbotapi.Callback, bot *mbotapi.BotAPI) error {
	msg := mbotapi.NewMessage("HI boss!!")
	bot.Send(c.Sender, msg, mbotapi.RegularNotif)
	return nil
}

func startLeave(c mbotapi.Callback, bot *mbotapi.BotAPI) error {
	msg := mbotapi.NewMessage("To the next step boss!!")
	bot.Send(c.Sender, msg, mbotapi.RegularNotif)
	return nil
}

func (m *StartState) Actions() (dispatcher.Action, dispatcher.Action) {
	return startEnter, startLeave
}
