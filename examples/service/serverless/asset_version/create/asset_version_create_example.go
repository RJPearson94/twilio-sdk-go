package main

import (
	"log"
	"os"
	"strings"

	"github.com/RJPearson94/twilio-sdk-go"
	v1 "github.com/RJPearson94/twilio-sdk-go/service/serverless/v1"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service/asset/versions"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

var serverlessSession *v1.Serverless

func init() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	serverlessSession = twilio.NewWithCredentials(creds).Serverless.V1
}

func main() {
	resp, err := serverlessSession.
		Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Asset("ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Versions.
		Create(&versions.CreateVersionInput{
			Content: versions.CreateContentDetails{
				Body:        strings.NewReader("{}"),
				ContentType: "application/json",
				FileName:    "test.json",
			},
			Path:       "/test",
			Visibility: "private",
		})

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("SID: %s", resp.Sid)
}
