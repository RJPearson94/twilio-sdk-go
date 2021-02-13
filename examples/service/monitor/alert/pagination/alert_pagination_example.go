package main

import (
	"log"
	"os"

	"github.com/RJPearson94/twilio-sdk-go"
	v1 "github.com/RJPearson94/twilio-sdk-go/service/monitor/v1"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

var monitorClient *v1.Monitor

func init() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	monitorClient = twilio.NewWithCredentials(creds).Monitor.V1
}

func main() {
	paginator := monitorClient.
		Alerts.
		NewAlertsPaginator()

	for paginator.Next() {
		currentPage := paginator.CurrentPage()
		log.Printf("%v alert(s) found on page %v", len(currentPage.Alerts), currentPage.Meta.Page)
	}

	if paginator.Error() != nil {
		log.Panicf("%s", paginator.Error())
	}

	log.Printf("Total number of alert(s) found: %v", len(paginator.Alerts))
}
