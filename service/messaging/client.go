package messaging

import (
	v1 "github.com/RJPearson94/twilio-sdk-go/service/messaging/v1"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

type Messaging struct {
	V1 *v1.Messaging
}

func New(sess *session.Session) *Messaging {
	c := &Messaging{}
	c.V1 = v1.New(sess)
	return c
}

func NewWithCredentials(creds *credentials.Credentials) *Messaging {
	return New(session.New(creds))
}
