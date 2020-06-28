package api

import (
	v2010 "github.com/RJPearson94/twilio-sdk-go/service/api/v2010"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

type API struct {
	*v2010.V2010
}

func New(sess *session.Session) *API {
	c := &API{}
	c.V2010 = v2010.New(sess)
	return c
}

func NewWithCredentials(creds *credentials.Credentials) *API {
	return New(session.New(creds))
}
