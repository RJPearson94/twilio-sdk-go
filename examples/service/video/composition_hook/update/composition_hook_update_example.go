package main

import (
	"log"
	"os"

	"github.com/RJPearson94/twilio-sdk-go/service/video/v1/composition_hook"
	"github.com/RJPearson94/twilio-sdk-go/utils"

	"github.com/RJPearson94/twilio-sdk-go"
	v1 "github.com/RJPearson94/twilio-sdk-go/service/video/v1"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

var videoClient *v1.Video

func init() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	videoClient = twilio.NewWithCredentials(creds).Video.V1
}

func main() {
	resp, err := videoClient.
		CompositionHook("HKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Update(&composition_hook.UpdateCompositionHookInput{
			FriendlyName: "Test",
			AudioSources: &[]string{"*"},
			Format:       utils.String("mp4"),
		})

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("SID: %s", resp.Sid)
}
