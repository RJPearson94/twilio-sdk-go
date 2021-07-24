package main

import (
	"log"
	"os"

	"github.com/RJPearson94/twilio-sdk-go"
	v1 "github.com/RJPearson94/twilio-sdk-go/service/accounts/v1"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

var accountsClient *v1.Accounts

func init() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	accountsClient = twilio.NewWithCredentials(creds).Accounts.V1
}

func main() {
	paginator := accountsClient.
		Credentials.
		PublicKeys.
		NewPublicKeysPaginator()

	for paginator.Next() {
		currentPage := paginator.CurrentPage()
		log.Printf("%v credential(s) found on page %v", len(currentPage.Credentials), currentPage.Meta.Page)
	}

	if paginator.Error() != nil {
		log.Panicf("%s", paginator.Error())
	}

	log.Printf("Total number of credential(s) found: %v", len(paginator.Credentials))
}
