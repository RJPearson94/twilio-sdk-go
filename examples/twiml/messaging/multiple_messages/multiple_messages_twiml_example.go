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
	response := twimlClient.MessagingResponse()
	response.Message(utils.String("Hello World"))
	response.Message(utils.String("Hello World Again"))
	twiML, err := response.ToTwiML()

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("TwiML: %s", *twiML)
}
