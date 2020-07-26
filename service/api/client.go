package api

import (
	v2010 "github.com/RJPearson94/twilio-sdk-go/service/api/v2010"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

// Client to manage resources that are part of the Twilio API
type API struct {
	*v2010.V2010
}

// Create a new instance of the client using session data
func New(sess *session.Session) *API {
	c := &API{}
	c.V2010 = v2010.New(sess)
	return c
}

// Create a new instance of the client with credentials
func NewWithCredentials(creds *credentials.Credentials) *API {
	return New(session.New(creds))
}
