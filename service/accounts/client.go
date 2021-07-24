package accounts

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	v1 "github.com/RJPearson94/twilio-sdk-go/service/accounts/v1"
	"github.com/RJPearson94/twilio-sdk-go/session"
)

// Accounts client is used to manage versioned resources for the Twilio account
type Accounts struct {
	V1 *v1.Accounts
}

// New creates a new instance of the client using session data and config
func New(sess *session.Session, config *client.Config) *Accounts {
	return &Accounts{
		V1: v1.New(sess, config),
	}
}
