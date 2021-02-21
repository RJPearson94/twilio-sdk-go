package main

import (
	"log"
	"os"

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
	paginator := videoClient.
		Room("RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Participants.
		NewParticipantsPaginator()

	for paginator.Next() {
		currentPage := paginator.CurrentPage()
		log.Printf("%v participant(s) found on page %v", len(currentPage.Participants), currentPage.Meta.Page)
	}

	if paginator.Error() != nil {
		log.Panicf("%s", paginator.Error())
	}

	log.Printf("Total number of participant(s) found: %v", len(paginator.Participants))
}
