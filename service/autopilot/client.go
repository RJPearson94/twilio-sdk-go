package autopilot

import (
	v1 "github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

// Autopilot client is used to manage versioned resources for Twilio Autopilot
// See https://www.twilio.com/docs/autopilot for more details on the API
// See https://www.twilio.com/autopilot for more details on the product
type Autopilot struct {
	V1 *v1.Autopilot
}

// New creates a new instance of the client using session data
func New(sess *session.Session) *Autopilot {
	return &Autopilot{
		V1: v1.New(sess),
	}
}

// NewWithCredentials creates a new instance of the client with credentials
func NewWithCredentials(creds *credentials.Credentials) *Autopilot {
	return New(session.New(creds))
}
