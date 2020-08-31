package main

import (
	"log"
	"os"

	"github.com/RJPearson94/twilio-sdk-go"
	v2 "github.com/RJPearson94/twilio-sdk-go/service/studio/v2"
	"github.com/RJPearson94/twilio-sdk-go/service/studio/v2/flow/execution"
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
		Flow("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Execution("FNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Update(&execution.UpdateExecutionInput{
			Status: "ended",
		})

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("SID: %s", resp.Sid)
}
