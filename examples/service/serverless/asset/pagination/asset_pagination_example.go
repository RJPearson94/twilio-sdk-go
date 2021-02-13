package main

import (
	"log"
	"os"

	"github.com/RJPearson94/twilio-sdk-go"
	v1 "github.com/RJPearson94/twilio-sdk-go/service/serverless/v1"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

var serverlessClient *v1.Serverless

func init() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	serverlessClient = twilio.NewWithCredentials(creds).Serverless.V1
}

func main() {
	paginator := serverlessClient.
		Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Assets.
		NewAssetsPaginator()

	for paginator.Next() {
		currentPage := paginator.CurrentPage()
		log.Printf("%v asset(s) found on page %v", len(currentPage.Assets), currentPage.Meta.Page)
	}

	if paginator.Error() != nil {
		log.Panicf("%s", paginator.Error())
	}

	log.Printf("Total number of asset(s) found: %v", len(paginator.Assets))
}
