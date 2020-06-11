package serverless

import (
	v1 "github.com/RJPearson94/twilio-sdk-go/service/serverless/v1"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

type Serverless struct {
	V1 *v1.Serverless
}

func New(sess *session.Session) *Serverless {
	c := &Serverless{}
	c.V1 = v1.New(sess)
	return c
}

func NewWithCredentials(creds *credentials.Credentials) *Serverless {
	return New(session.New(creds))
}
