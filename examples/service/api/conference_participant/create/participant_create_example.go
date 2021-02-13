package main

import (
	"log"
	"os"

	"github.com/RJPearson94/twilio-sdk-go"
	v2010 "github.com/RJPearson94/twilio-sdk-go/service/api/v2010"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/conference/participants"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

var apiClient *v2010.V2010

func init() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	apiClient = twilio.NewWithCredentials(creds).API.V2010
}

func main() {
	resp, err := apiClient.
		Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Conference("CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Participants.
		Create(&participants.CreateParticipantInput{
			From: "+19876543210",
			To:   "+10123456789",
		})

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("Call SID: %s", resp.CallSid)
}
