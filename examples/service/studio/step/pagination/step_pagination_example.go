package main

import (
	"log"
	"os"

	"github.com/RJPearson94/twilio-sdk-go"
	v2 "github.com/RJPearson94/twilio-sdk-go/service/studio/v2"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

var studioClient *v2.Studio

func init() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	studioClient = twilio.NewWithCredentials(creds).Studio.V2
}

func main() {
	paginator := studioClient.
		Flow("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Execution("FNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Steps.
		NewStepsPaginator()

	for paginator.Next() {
		currentPage := paginator.CurrentPage()
		log.Printf("%v step(s) found on page %v", len(currentPage.Steps), currentPage.Meta.Page)
	}

	if paginator.Error() != nil {
		log.Panicf("%s", paginator.Error())
	}

	log.Printf("Total number of step(s) found: %v", len(paginator.Steps))
}
