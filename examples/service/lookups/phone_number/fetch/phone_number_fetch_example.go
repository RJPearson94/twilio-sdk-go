package main

import (
	"log"
	"os"

	"github.com/RJPearson94/twilio-sdk-go"
	v1 "github.com/RJPearson94/twilio-sdk-go/service/lookups/v1"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

var lookupSession *v1.Lookups

func init() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	lookupSession = twilio.NewWithCredentials(creds).Lookups.V1
}

func main() {
	resp, err := lookupSession.
		PhoneNumber("+10123456789").
		Fetch(nil)

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("National Format: %s", resp.NationalFormat)
	log.Printf("URL: %s", resp.URL)
}
