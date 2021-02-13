package main

import (
	"log"
	"os"

	"github.com/RJPearson94/twilio-sdk-go"
	v1 "github.com/RJPearson94/twilio-sdk-go/service/trunking/v1"
	"github.com/RJPearson94/twilio-sdk-go/service/trunking/v1/trunk/origination_urls"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
	"github.com/google/uuid"
)

var trunkingClient *v1.Trunking

func init() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	trunkingClient = twilio.NewWithCredentials(creds).Trunking.V1
}

func main() {
	resp, err := trunkingClient.
		Trunk("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		OriginationURLs.
		Create(&origination_urls.CreateOriginationURLInput{
			Weight:       0,
			Priority:     0,
			Enabled:      false,
			FriendlyName: uuid.New().String(),
			SipURL:       "sip:test@test.com",
		})

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("SID: %s", resp.Sid)
}
