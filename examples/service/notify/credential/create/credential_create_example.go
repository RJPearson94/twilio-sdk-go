package main

import (
	"log"
	"os"

	"github.com/google/uuid"

	"github.com/RJPearson94/twilio-sdk-go"
	v1 "github.com/RJPearson94/twilio-sdk-go/service/notify/v1"
	"github.com/RJPearson94/twilio-sdk-go/service/notify/v1/credentials"
	sessionCredentials "github.com/RJPearson94/twilio-sdk-go/session/credentials"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

var notifyClient *v1.Notify

func init() {
	creds, err := sessionCredentials.New(sessionCredentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	notifyClient = twilio.NewWithCredentials(creds).Notify.V1
}

func main() {
	resp, err := notifyClient.
		Credentials.
		Create(&credentials.CreateCredentialInput{
			Type:   "fcm",
			Secret: utils.String(uuid.New().String()),
		})

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("SID: %s", resp.Sid)
}
