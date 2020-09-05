package main

import (
	"log"
	"os"

	"github.com/RJPearson94/twilio-sdk-go"
	v1 "github.com/RJPearson94/twilio-sdk-go/service/messaging/v1"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

var messagingSession *v1.Messaging

func init() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	messagingSession = twilio.NewWithCredentials(creds).Messaging.V1
}

func main() {
	paginator := messagingSession.
		Service("MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		PhoneNumbers.
		NewPhoneNumbersPaginator()

	for paginator.Next() {
		currentPage := paginator.CurrentPage()
		log.Printf("%v phone number(s) found on page %v", len(currentPage.PhoneNumbers), currentPage.Meta.Page)
	}

	if paginator.Error() != nil {
		log.Panicf("%s", paginator.Error())
	}

	log.Printf("Total number of phone number(s) found: %v", len(paginator.PhoneNumbers))
}
