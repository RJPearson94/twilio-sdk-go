package taskrouter

import (
	"log"
	"os"

	"github.com/RJPearson94/twilio-sdk-go"
	v1 "github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1"
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
	err := taskrouterSession.
		Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		TaskChannel("TCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Delete()

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("Task Channel Deleted")
}
