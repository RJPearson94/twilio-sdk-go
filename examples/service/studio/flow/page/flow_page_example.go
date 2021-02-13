package main

import (
	"log"
	"os"

	"github.com/RJPearson94/twilio-sdk-go"
	v2 "github.com/RJPearson94/twilio-sdk-go/service/studio/v2"
	"github.com/RJPearson94/twilio-sdk-go/service/studio/v2/flows"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

var studioClient *v2.Studio

func init() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	studioClient = twilio.NewWithCredentials(creds).Studio.V2
}

func main() {
	resp, err := studioClient.
		Flows.
		Page(&flows.FlowsPageOptions{})

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("%v flow(s) found on page", len(resp.Flows))
}
