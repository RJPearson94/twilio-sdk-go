package chat

import (
	v2 "github.com/RJPearson94/twilio-sdk-go/service/chat/v2"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

type Chat struct {
	V2 *v2.Chat
}

func New(sess *session.Session) *Chat {
	c := &Chat{}
	c.V2 = v2.New(sess)
	return c
}

func NewWithCredentials(creds *credentials.Credentials) *Chat {
	return New(session.New(creds))
}
