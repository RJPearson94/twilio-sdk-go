package main

import (
	"log"
	"os"

	"github.com/RJPearson94/twilio-sdk-go"
	v1 "github.com/RJPearson94/twilio-sdk-go/service/conversations/v1"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/credentials"
	sessionCredentials "github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

var conversationSession *v1.Conversations

func init() {
	creds, err := sessionCredentials.New(sessionCredentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	conversationSession = twilio.NewWithCredentials(creds).Conversations.V1
}

func main() {
	resp, err := conversationSession.
		Credentials.
		Page(&credentials.CredentialsPageOptions{})

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("%v credential(s) found on page", len(resp.Credentials))
}
