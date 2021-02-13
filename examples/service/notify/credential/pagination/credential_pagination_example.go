package main

import (
	"log"
	"os"

	"github.com/RJPearson94/twilio-sdk-go"
	v1 "github.com/RJPearson94/twilio-sdk-go/service/notify/v1"
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
	paginator := notifyClient.
		Credentials.
		NewCredentialsPaginator()

	for paginator.Next() {
		currentPage := paginator.CurrentPage()
		log.Printf("%v credential(s) found on page %v", len(currentPage.Credentials), currentPage.Meta.Page)
	}

	if paginator.Error() != nil {
		log.Panicf("%s", paginator.Error())
	}

	log.Printf("Total number of credential(s) found: %v", len(paginator.Credentials))
}
