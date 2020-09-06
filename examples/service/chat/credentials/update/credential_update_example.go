package main

import (
	"log"
	"os"

	"github.com/RJPearson94/twilio-sdk-go"
	v2 "github.com/RJPearson94/twilio-sdk-go/service/chat/v2"
	chatCredential "github.com/RJPearson94/twilio-sdk-go/service/chat/v2/credential"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

var chatSession *v2.Chat

func init() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	chatSession = twilio.NewWithCredentials(creds).Chat.V2
}

func main() {
	resp, err := chatSession.
		Credential("CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Update(&chatCredential.UpdateCredentialInput{
			Secret: utils.String("new secret"),
		})

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("SID: %s", resp.Sid)
}