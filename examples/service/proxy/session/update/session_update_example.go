package main

import (
	"log"
	"os"

	"github.com/RJPearson94/twilio-sdk-go"
	v1 "github.com/RJPearson94/twilio-sdk-go/service/proxy/v1"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/service/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

var proxyClient *v1.Proxy

func init() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	proxyClient = twilio.NewWithCredentials(creds).Proxy.V1
}

func main() {
	resp, err := proxyClient.
		Service("KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Session("KCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Update(&session.UpdateSessionInput{
			Status: utils.String("in-progress"),
		})

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("SID: %s", resp.Sid)
}
