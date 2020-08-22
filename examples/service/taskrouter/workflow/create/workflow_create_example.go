package main

import (
	"log"
	"os"

	"github.com/RJPearson94/twilio-sdk-go"
	v1 "github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/workflows"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

var taskrouterSession *v1.TaskRouter

func init() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	taskrouterSession = twilio.NewWithCredentials(creds).TaskRouter.V1
}

func main() {
	workflow, err := taskrouterSession.
		Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Workflows.
		Create(&workflows.CreateWorkflowInput{
			FriendlyName:  "Test 2",
			Configuration: `{ "task_routing": { "default_filter": { "queue": "WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX" } } }`,
		})

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("SID: %s", workflow.Sid)
}
