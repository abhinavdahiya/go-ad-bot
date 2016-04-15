package main

import (
	"fmt"
	"log"
	"os"

	"github.com/abhinavdahiya/go-messenger-bot"
	"github.com/codegangsta/negroni"
)

func main() {
	bot := mbotapi.NewBotAPI("CAAOmDyEhPfwBAJ9JJpRLvZB3i2dLdJXIZBZCEYYTNJkpUnzHs6KYKN5I1vySuwTmF3uxRrwCHDLQBJJiMHZAXgHiZAV87ugaLR3iYnW1k5OD9KPmvaMAP0rEZCSfUHz0lolyoUelwNJQJdYZAfLElMfhvlGn1ZCPrTXcNNepItZC1509LIGdMK9ZAwZANI9Six9Sp1EB6Ozu0CRtgZDZD", "my_voice_is_my_password")

	callbacks, mux := bot.SetWebhook("/webhook")
	n := negroni.Classic()
	n.UseHandler(mux)

	bind := fmt.Sprintf("%s:%s", os.Getenv("OPENSHIFT_GO_IP"), os.Getenv("OPENSHIFT_GO_PORT"))
	go n.Run(bind)

	log.Printf("%#v\n", callbacks)
	for callback := range callbacks {
		log.Printf("[%#v] %s\n", callback.Sender, callback.Message.Text)

		msg := mbotapi.NewMessage(callback.Message.Text)
		resp, err := bot.Send(callback.Sender, msg, mbotapi.RegularNotif)
		log.Printf("%#v (%s)", resp, err)
	}
}
