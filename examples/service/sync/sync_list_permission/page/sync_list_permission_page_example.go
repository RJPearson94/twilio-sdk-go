package main

import (
	"log"
	"os"

	"github.com/RJPearson94/twilio-sdk-go"
	v1 "github.com/RJPearson94/twilio-sdk-go/service/sync/v1"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/sync_list/permissions"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

var syncClient *v1.Sync

func init() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	syncClient = twilio.NewWithCredentials(creds).Sync.V1
}

func main() {
	resp, err := syncClient.
		Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		SyncList("ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Permissions.
		Page(&permissions.SyncListPermissionsPageOptions{})

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("%v sync list permission(s) found on page", len(resp.Permissions))
}
