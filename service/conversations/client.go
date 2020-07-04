package conversations

import (
	v1 "github.com/RJPearson94/twilio-sdk-go/service/conversations/v1"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

type Conversations struct {
	V1 *v1.Conversations
}

func New(sess *session.Session) *Conversations {
	c := &Conversations{}
	c.V1 = v1.New(sess)
	return c
}

func NewWithCredentials(creds *credentials.Credentials) *Conversations {
	return New(session.New(creds))
}
