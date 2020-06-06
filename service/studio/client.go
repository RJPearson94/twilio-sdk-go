package studio

import (
	v2 "github.com/RJPearson94/twilio-sdk-go/service/studio/v2"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

type Studio struct {
	V2 *v2.Studio
}

func New(sess *session.Session) *Studio {
	c := &Studio{}
	c.V2 = v2.New(sess)
	return c
}

func NewWithCredentials(creds *credentials.Credentials) *Studio {
	return New(session.New(creds))
}
