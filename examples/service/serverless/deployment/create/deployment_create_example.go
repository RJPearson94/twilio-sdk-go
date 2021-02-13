package main

import (
	"log"
	"os"

	"github.com/RJPearson94/twilio-sdk-go"
	v1 "github.com/RJPearson94/twilio-sdk-go/service/serverless/v1"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service/environment/deployments"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
	"github.com/RJPearson94/twilio-sdk-go/utils"
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
	resp, err := serverlessClient.
		Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Environment("ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Deployments.
		Create(&deployments.CreateDeploymentInput{
			BuildSid: utils.String("ZBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		})

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("SID: %s", resp.Sid)
}
