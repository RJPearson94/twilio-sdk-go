package proxy

import (
	v1 "github.com/RJPearson94/twilio-sdk-go/service/proxy/v1"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

type Proxy struct {
	V1 *v1.Proxy
}

func New(sess *session.Session) *Proxy {
	c := &Proxy{}
	c.V1 = v1.New(sess)
	return c
}

func NewWithCredentials(creds *credentials.Credentials) *Proxy {
	return New(session.New(creds))
}
