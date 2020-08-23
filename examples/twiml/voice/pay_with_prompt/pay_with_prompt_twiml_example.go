package main

import (
	"log"

	"github.com/RJPearson94/twilio-sdk-go/twiml"
	"github.com/RJPearson94/twilio-sdk-go/twiml/voice/verbs"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

var twimlClient *twiml.TwiML

func init() {
	twimlClient = twiml.New()
}

func main() {
	response := twimlClient.VoiceResponse()
	pay := response.Pay()
	prompt := pay.PromptWithAttributes(verbs.PromptAttributes{
		For:     utils.String("cvv"),
		Attempt: utils.Int(1),
	})
	prompt.Say("Please enter your cvv which is on the back of your card")
	twiML, err := response.ToTwiML()

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("TwiML: %s", *twiML)
}
