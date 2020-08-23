package main

import (
	"log"

	"github.com/RJPearson94/twilio-sdk-go/twiml"
)

var twimlClient *twiml.TwiML

func init() {
	twimlClient = twiml.New()
}

func main() {
	response := twimlClient.MessagingResponse()
	message := response.Message(nil)
	message.Body("Hello world")
	message.Media("https://localhost/media")
	twiML, err := response.ToTwiML()

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("TwiML: %s", *twiML)
}
