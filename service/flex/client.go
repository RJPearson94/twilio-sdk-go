package flex

import (
	v1 "github.com/RJPearson94/twilio-sdk-go/service/flex/v1"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

type Flex struct {
	V1 *v1.Flex
}

func New(sess *session.Session) *Flex {
	c := &Flex{}
	c.V1 = v1.New(sess)
	return c
}

func NewWithCredentials(creds *credentials.Credentials) *Flex {
	return New(session.New(creds))
}
