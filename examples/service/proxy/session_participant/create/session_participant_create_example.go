package main

import (
	"log"
	"os"

	"github.com/RJPearson94/twilio-sdk-go"
	v1 "github.com/RJPearson94/twilio-sdk-go/service/proxy/v1"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/service/session/participants"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

var proxySession *v1.Proxy

func init() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	proxySession = twilio.NewWithCredentials(creds).Proxy.V1
}

func main() {
	resp, err := proxySession.
		Service("KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Session("KCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Participants.
		Create(&participants.CreateParticipantInput{
			Identifier: "+123456789",
		})

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("SID: %s", resp.Sid)
}