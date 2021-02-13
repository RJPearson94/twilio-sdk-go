package main

import (
	"log"
	"os"

	"github.com/RJPearson94/twilio-sdk-go"
	v1 "github.com/RJPearson94/twilio-sdk-go/service/trunking/v1"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

var trunkingClient *v1.Trunking

func init() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	trunkingClient = twilio.NewWithCredentials(creds).Trunking.V1
}

func main() {
	paginator := trunkingClient.
		Trunk("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		OriginationURLs.
		NewOriginationURLsPaginator()

	for paginator.Next() {
		currentPage := paginator.CurrentPage()
		log.Printf("%v origination url(s) found on page %v", len(currentPage.OriginationURLs), currentPage.Meta.Page)
	}

	if paginator.Error() != nil {
		log.Panicf("%s", paginator.Error())
	}

	log.Printf("Total number of origination url(s) found: %v", len(paginator.OriginationURLs))
}
