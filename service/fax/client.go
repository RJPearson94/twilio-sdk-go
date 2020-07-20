package fax

import (
	v1 "github.com/RJPearson94/twilio-sdk-go/service/fax/v1"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

type Fax struct {
	V1 *v1.Fax
}

func New(sess *session.Session) *Fax {
	c := &Fax{}
	c.V1 = v1.New(sess)
	return c
}

func NewWithCredentials(creds *credentials.Credentials) *Fax {
	return New(session.New(creds))
}
