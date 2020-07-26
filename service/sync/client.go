package sync

import (
	v1 "github.com/RJPearson94/twilio-sdk-go/service/sync/v1"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

type Sync struct {
	V1 *v1.Sync
}

func New(sess *session.Session) *Sync {
	c := &Sync{}
	c.V1 = v1.New(sess)
	return c
}

func NewWithCredentials(creds *credentials.Credentials) *Sync {
	return New(session.New(creds))
}
