package main

import (
	"log"
	"os"

	"github.com/RJPearson94/twilio-sdk-go"
	v1 "github.com/RJPearson94/twilio-sdk-go/service/accounts/v1"
	"github.com/RJPearson94/twilio-sdk-go/service/accounts/v1/credentials/aws_credential"
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
	resp, err := accountsClient.
		Credentials.
		AWSCredential("CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Update(&aws_credential.UpdateAWSCredentialInput{})

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("SID: %s", resp.Sid)
}
