package main

import (
	"log"
	"os"

	"github.com/RJPearson94/twilio-sdk-go"
	v1 "github.com/RJPearson94/twilio-sdk-go/service/flex/v1"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

var flexSession *v1.Flex

func init() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	flexSession = twilio.NewWithCredentials(creds).Flex.V1
}

func main() {
	paginator := flexSession.
		PluginConfigurations.
		NewConfigurationsPaginator()

	for paginator.Next() {
		currentPage := paginator.CurrentPage()
		log.Printf("%v plugin configuration(s) found on page %v", len(currentPage.Configurations), currentPage.Meta.Page)
	}

	if paginator.Error() != nil {
		log.Panicf("%s", paginator.Error())
	}

	log.Printf("Total number of plugin configuration(s) found: %v", len(paginator.Configurations))
}