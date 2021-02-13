package main

import (
	"log"
	"os"

	"github.com/RJPearson94/twilio-sdk-go"
	v1 "github.com/RJPearson94/twilio-sdk-go/service/messaging/v1"
	"github.com/RJPearson94/twilio-sdk-go/service/messaging/v1/service/phone_numbers"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

var messagingClient *v1.Messaging

func init() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	messagingClient = twilio.NewWithCredentials(creds).Messaging.V1
}

func main() {
	resp, err := messagingClient.
		Service("MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		PhoneNumbers.
		Page(&phone_numbers.PhoneNumbersPageOptions{})

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("%v phone numbers(s) found on page", len(resp.PhoneNumbers))
}
