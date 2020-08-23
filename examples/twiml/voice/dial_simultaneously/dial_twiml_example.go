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
	response := twimlClient.VoiceResponse()
	dial := response.Dial(nil)
	dial.Number("415-123-9876")
	dial.Number("415-123-9877")
	dial.Number("415-123-9878")
	twiML, err := response.ToTwiML()

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("TwiML: %s", *twiML)
}
