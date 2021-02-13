package main

import (
	"log"
	"os"

	"github.com/RJPearson94/twilio-sdk-go"
	v1 "github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/field_type/field_values"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

var autopilotClient *v1.Autopilot

func init() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	autopilotClient = twilio.NewWithCredentials(creds).Autopilot.V1
}

func main() {
	resp, err := autopilotClient.
		Assistant("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		FieldType("UBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		FieldValues.
		Create(&field_values.CreateFieldValueInput{
			Language: "en-US",
			Value:    "test",
		})

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("SID: %s", resp.Sid)
}
