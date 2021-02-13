package main

import (
	"log"
	"os"

	"github.com/RJPearson94/twilio-sdk-go"
	v2 "github.com/RJPearson94/twilio-sdk-go/service/chat/v2"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

var chatClient *v2.Chat

func init() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	chatClient = twilio.NewWithCredentials(creds).Chat.V2
}

func main() {
	paginator := chatClient.
		Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		User("USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Channels.
		NewUserChannelsPaginator()

	for paginator.Next() {
		currentPage := paginator.CurrentPage()
		log.Printf("%v user channel(s) found on page %v", len(currentPage.Channels), currentPage.Meta.Page)
	}

	if paginator.Error() != nil {
		log.Panicf("%s", paginator.Error())
	}

	log.Printf("Total number of user channel(s) found: %v", len(paginator.Channels))
}
