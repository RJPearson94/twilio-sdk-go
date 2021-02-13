package main

import (
	"log"
	"os"

	"github.com/google/uuid"

	"github.com/RJPearson94/twilio-sdk-go"
	v1 "github.com/RJPearson94/twilio-sdk-go/service/conversations/v1"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/credentials"
	sessionCredentials "github.com/RJPearson94/twilio-sdk-go/session/credentials"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

var conversationClient *v1.Conversations

func init() {
	creds, err := sessionCredentials.New(sessionCredentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	conversationClient = twilio.NewWithCredentials(creds).Conversations.V1
}

func main() {
	resp, err := conversationClient.
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
