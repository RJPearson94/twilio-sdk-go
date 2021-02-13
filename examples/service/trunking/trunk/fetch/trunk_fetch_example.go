package main

import (
	"log"
	"os"

	"github.com/RJPearson94/twilio-sdk-go"
	v1 "github.com/RJPearson94/twilio-sdk-go/service/trunking/v1"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

var trunkingClient *v1.Trunking

func init() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	trunkingClient = twilio.NewWithCredentials(creds).Trunking.V1
}

func main() {
	resp, err := trunkingClient.
		Trunk("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Fetch()

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("SID: %s", resp.Sid)
}
