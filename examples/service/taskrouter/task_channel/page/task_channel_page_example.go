package main

import (
	"log"
	"os"

	"github.com/RJPearson94/twilio-sdk-go"
	v1 "github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/task_channels"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

var taskrouterClient *v1.TaskRouter

func init() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	taskrouterClient = twilio.NewWithCredentials(creds).TaskRouter.V1
}

func main() {
	resp, err := taskrouterClient.
		Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		TaskChannels.
		Page(&task_channels.TaskChannelsPageOptions{})

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("%v task channel(s) found on page", len(resp.TaskChannels))
}
