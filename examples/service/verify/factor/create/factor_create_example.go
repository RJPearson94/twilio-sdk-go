package main

import (
	"log"
	"os"

	"github.com/RJPearson94/twilio-sdk-go"
	v2 "github.com/RJPearson94/twilio-sdk-go/service/verify/v2"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/entity/factors"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

var verifySession *v2.Verify

func init() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	verifySession = twilio.NewWithCredentials(creds).Verify.V2
}

func main() {
	resp, err := verifySession.
		Service("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Entity("test").
		Factors.
		Create(&factors.CreateFactorInput{
			BindingAlg:                 "ES256",
			BindingPublicKey:           "TestKey",
			ConfigAppId:                "test",
			ConfigNotificationPlatform: "fcm",
			ConfigNotificationToken:    "notification_token",
			FactorType:                 "push",
			FriendlyName:               "test factor",
		})

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("SID: %s", resp.Sid)
}
