package main

import (
	"fmt"

	"github.com/abhinavdahiya/go-messenger-bot"
	"github.com/abhinavdahiya/go-messenger-dispatcher"
)

type HiState struct {
	Enter, Leave dispatcher.Action

	IsMoved bool
	Chain   string
}

func (m *HiState) Name() string {
	return "hi"
}

func (m *HiState) Transit(s string) {
	m.IsMoved = true
	m.Chain = s
}

func (m *HiState) Next() string {
	return m.Chain
}

func (m *HiState) Transitor(c mbotapi.Callback) {
	if m.IsMoved {
		return
	}

	if pb := c.Postback; pb.Payload != "" {
		if pb.Payload == "HELP" {
			m.IsMoved = true
			m.Chain = "start"
			return
		}
	}
}

func (m *HiState) Flush() {
	m.IsMoved = false
	m.Chain = ""
}

func hiEnter(c mbotapi.Callback, bot *mbotapi.BotAPI) (err error) {
	gt := mbotapi.NewGenericTemplate()
	for i := 0; i < 2; i++ {
		e := mbotapi.NewElement(fmt.Sprintf("Card %d", i+1))
		if i == 1 {
			imb := mbotapi.NewURLButton("image", "http://38.media.tumblr.com/avatar_931d37a59260_128.png")
			pb := mbotapi.NewPostbackButton("ask help", "HELP")
			e.AddButton(imb, pb)
		} else {
			e.ImageURL = "http://38.media.tumblr.com/avatar_931d37a59260_128.png"
		}
		gt.AddElement(e)
	}
	_, err = bot.Send(c.Sender, gt, mbotapi.SilentNotif)
	return
}

func (m *HiState) Actions() (dispatcher.Action, dispatcher.Action) {
	return hiEnter, nil
}
