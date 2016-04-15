package main

import (
	"fmt"
	"log"
	"os"

	"github.com/abhinavdahiya/go-messenger-bot"
	"github.com/codegangsta/negroni"
)

func main() {
	bot, err := mbotapi.NewBotAPI("CAAOmDyEhPfwBAJ9JJpRLvZB3i2dLdJXIZBZCEYYTNJkpUnzHs6KYKN5I1vySuwTmF3uxRrwCHDLQBJJiMHZAXgHiZAV87ugaLR3iYnW1k5OD9KPmvaMAP0rEZCSfUHz0lolyoUelwNJQJdYZAfLElMfhvlGn1ZCPrTXcNNepItZC1509LIGdMK9ZAwZANI9Six9Sp1EB6Ozu0CRtgZDZD", "my_voice_is_my_password")
	if err != nil {
		log.Panic(err)
	}

	callbacks, mux := bot.SetWebhook("/webhook")
	n := negroni.Classic()
	n.UseHandler(mux)

	bind := fmt.Sprintf("%s:%s", os.Getenv("OPENSHIFT_GO_IP"), os.Getenv("OPENSHIFT_GO_PORT"))
	go n.Run(bind)

	for callback := range callbacks {
		log.Printf("[%#v] %s", callback.Sender, callback.Message.Text)

		msg := mbotapi.NewMessage(callback.Message.Text)
		bot.Send(callback.Sender, msg, mbotapi.RegularNotif)
	}
}
