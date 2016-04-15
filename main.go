package main

import (
	"fmt"
	"log"
	"os"

	"github.com/abhinavdahiya/go-messenger-bot"
	"github.com/abhinavdahiya/go-messenger-dispatcher"
	"github.com/codegangsta/negroni"
)

func main() {
	bot := mbotapi.NewBotAPI("CAAOmDyEhPfwBAJ9JJpRLvZB3i2dLdJXIZBZCEYYTNJkpUnzHs6KYKN5I1vySuwTmF3uxRrwCHDLQBJJiMHZAXgHiZAV87ugaLR3iYnW1k5OD9KPmvaMAP0rEZCSfUHz0lolyoUelwNJQJdYZAfLElMfhvlGn1ZCPrTXcNNepItZC1509LIGdMK9ZAwZANI9Six9Sp1EB6Ozu0CRtgZDZD", "my_voice_is_my_password")
	bot.Debug = true

	callbacks, mux := bot.SetWebhook("/webhook")
	n := negroni.Classic()
	n.UseHandler(mux)

	bind := fmt.Sprintf("%s:%s", os.Getenv("OPENSHIFT_GO_IP"), os.Getenv("OPENSHIFT_GO_PORT"))
	go n.Run(bind)

	dp := dispatcher.NewDispatcher()
	dp.Debug = true
	dp.AddState(&StartState{})
	dp.AddState(&HiState{})

	for callback := range callbacks {
		log.Printf("%s", dp.Process(callback, bot))
	}
}
