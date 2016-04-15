package main

import (
	"log"

	"github.com/abhinavdahiya/go-messenger-bot"
	"github.com/codegangsta/negroni"
)

func main() {
	bot, err := mbotapi.NewBotAPI("ACCESS_TOKEN", "VERIFY_TOKEN")
	if err != nil {
		log.Panic(err)
	}

	callbacks, mux := bot.SetWebhook("/webhook")
	n := negroni.Classic()
	n.UseHandler(mux)
	//go

	for callback := range callbacks {
		log.Printf("[%#v] %s", callback.Sender, callback.Message.Text)

		msg := mbotapi.NewMessage(callback.Message.Text)
		bot.Send(callback.Sender, msg, mbotapi.RegularNotif)
	}
}
