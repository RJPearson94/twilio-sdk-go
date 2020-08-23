package main

import (
	"log"

	"github.com/RJPearson94/twilio-sdk-go/twiml"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

var twimlClient *twiml.TwiML

func init() {
	twimlClient = twiml.New()
}

func main() {
	response := twimlClient.VoiceResponse()
	response.Dial(utils.String("415-123-9876"))
	twiML, err := response.ToTwiML()

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("TwiML: %s", *twiML)
}
