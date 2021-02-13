package main

import (
	"log"
	"os"

	"github.com/RJPearson94/twilio-sdk-go"
	v1 "github.com/RJPearson94/twilio-sdk-go/service/notify/v1"
	"github.com/RJPearson94/twilio-sdk-go/service/notify/v1/service/bindings"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

var notifyClient *v1.Notify

func init() {
	creds, err := credentials.New(credentials.Account{
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
		Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Bindings.
		Page(&bindings.BindingsPageOptions{})

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("%v binding(s) found on page", len(resp.Bindings))
}
