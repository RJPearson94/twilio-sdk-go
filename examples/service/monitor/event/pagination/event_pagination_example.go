package main

import (
	"log"
	"os"

	"github.com/RJPearson94/twilio-sdk-go"
	v1 "github.com/RJPearson94/twilio-sdk-go/service/monitor/v1"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

var monitorSession *v1.Monitor

func init() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	monitorSession = twilio.NewWithCredentials(creds).Monitor.V1
}

func main() {
	paginator := monitorSession.
		Events.
		NewEventsPaginator()

	for paginator.Next() {
		currentPage := paginator.CurrentPage()
		log.Printf("%v event(s) found on page %v", len(currentPage.Events), currentPage.Meta.Page)
	}

	if paginator.Error() != nil {
		log.Panicf("%s", paginator.Error())
	}

	log.Printf("Total number of event(s) found: %v", len(paginator.Events))
}
