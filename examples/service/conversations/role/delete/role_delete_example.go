package main

import (
	"log"
	"os"

	"github.com/RJPearson94/twilio-sdk-go"
	v1 "github.com/RJPearson94/twilio-sdk-go/service/conversations/v1"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

var conversationSession *v1.Conversations

func init() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	conversationSession = twilio.NewWithCredentials(creds).Conversations.V1
}

func main() {
	err := conversationSession.
		Role("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Delete()

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("Role successfully deleted")
}
