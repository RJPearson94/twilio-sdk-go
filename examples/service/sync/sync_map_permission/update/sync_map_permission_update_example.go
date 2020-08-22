package main

import (
	"log"
	"os"

	"github.com/RJPearson94/twilio-sdk-go"
	v1 "github.com/RJPearson94/twilio-sdk-go/service/sync/v1"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/sync_map/permission"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

var syncSession *v1.Sync

func init() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	syncSession = twilio.NewWithCredentials(creds).Sync.V1
}

func main() {
	syncMap, err := syncSession.
		Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		SyncMap("MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Permission("test").
		Update(&permission.UpdateSyncMapPermissionsInput{
			Read:   true,
			Write:  false,
			Manage: false,
		})

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("Identity: %s", syncMap.Identity)
}
