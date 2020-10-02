package main

import (
	"log"
	"os"

	"github.com/RJPearson94/twilio-sdk-go"
	v2010 "github.com/RJPearson94/twilio-sdk-go/service/api/v2010"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

var apiSession *v2010.V2010

func init() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	apiSession = twilio.NewWithCredentials(creds).API.V2010
}

func main() {
	paginator := apiSession.
		Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Call("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Recordings.
		NewRecordingsPaginator()

	for paginator.Next() {
		currentPage := paginator.CurrentPage()
		log.Printf("%v recording(s) found on page %v", len(currentPage.Recordings), currentPage.Page)
	}

	if paginator.Error() != nil {
		log.Panicf("%s", paginator.Error())
	}

	log.Printf("Total number of recording(s) found: %v", len(paginator.Recordings))
}
