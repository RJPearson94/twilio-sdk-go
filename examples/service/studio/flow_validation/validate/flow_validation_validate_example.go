package main

import (
	"log"
	"os"

	"github.com/RJPearson94/twilio-sdk-go"
	v2 "github.com/RJPearson94/twilio-sdk-go/service/studio/v2"
	"github.com/RJPearson94/twilio-sdk-go/service/studio/v2/flow_validation"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

var studioSession *v2.Studio

func init() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	studioSession = twilio.NewWithCredentials(creds).Studio.V2
}

func main() {
	resp, err := studioSession.
		FlowValidation.
		Validate(&flow_validation.ValidateFlowInput{
			FriendlyName: "test",
			Status:       "draft",
			Definition: `
{
	"description": "A New Flow",
	"flags": {
		"allow_concurrent_calls": true
	},
	"initial_state": "Trigger",
	"states": [
		{
			"name": "Trigger",
			"properties": {
				"offset": {
					"x": 0,
					"y": 0
				}
			},
			"transitions": [],
			"type": "trigger"
		}
	]
}`,
		})

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("Valid: %v", resp.Valid)
}
