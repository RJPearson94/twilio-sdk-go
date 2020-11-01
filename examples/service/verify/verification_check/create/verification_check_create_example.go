package main

import (
	"log"
	"os"

	"github.com/RJPearson94/twilio-sdk-go"
	v2 "github.com/RJPearson94/twilio-sdk-go/service/verify/v2"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/verification_check"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
	"github.com/RJPearson94/twilio-sdk-go/utils"
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
		VerificationCheck.
		Create(&verification_check.CreateVerificationCheckInput{
			To:   utils.String(os.Getenv("DESTINATION_PHONE_NUMBER")),
			Code: os.Getenv("TWILIO_VERIFY_CODE"),
		})

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("SID: %s", resp.Sid)
}
