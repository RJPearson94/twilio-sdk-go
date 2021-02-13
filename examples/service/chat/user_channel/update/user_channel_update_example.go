package main

import (
	"log"
	"os"

	"github.com/RJPearson94/twilio-sdk-go"
	v2 "github.com/RJPearson94/twilio-sdk-go/service/chat/v2"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service/user/channel"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
	"github.com/RJPearson94/twilio-sdk-go/utils"
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
	resp, err := chatClient.
		Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		User("USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Channel("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Update(&channel.UpdateUserChannelInput{
			NotificationLevel: utils.String("default"),
		})

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("Account SID: %s", resp.AccountSid)
}
