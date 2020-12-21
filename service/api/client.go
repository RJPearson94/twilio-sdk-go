package api

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	v2010 "github.com/RJPearson94/twilio-sdk-go/service/api/v2010"
	"github.com/RJPearson94/twilio-sdk-go/session"
)

// API client to manage resources that are part of the Twilio API
type API struct {
	*v2010.V2010
}

// New creates a new instance of the client using session data and config
func New(sess *session.Session, config *client.Config) *API {
	return &API{
		V2010: v2010.New(sess, config),
	}
}
