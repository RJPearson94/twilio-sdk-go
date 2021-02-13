package main

import (
	"log"
	"os"

	"github.com/RJPearson94/twilio-sdk-go"
	v1 "github.com/RJPearson94/twilio-sdk-go/service/messaging/v1"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

var messagingClient *v1.Messaging

func init() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	messagingClient = twilio.NewWithCredentials(creds).Messaging.V1
}

func main() {
	paginator := messagingClient.
		Service("MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		ShortCodes.
		NewShortCodesPaginator()

	for paginator.Next() {
		currentPage := paginator.CurrentPage()
		log.Printf("%v short code(s) found on page %v", len(currentPage.ShortCodes), currentPage.Meta.Page)
	}

	if paginator.Error() != nil {
		log.Panicf("%s", paginator.Error())
	}

	log.Printf("Total number of short code(s) found: %v", len(paginator.ShortCodes))
}
