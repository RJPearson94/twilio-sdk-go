package autopilot

import (
	v1 "github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

type Autopilot struct {
	V1 *v1.Autopilot
}

func New(sess *session.Session) *Autopilot {
	c := &Autopilot{}
	c.V1 = v1.New(sess)
	return c
}

func NewWithCredentials(creds *credentials.Credentials) *Autopilot {
	return New(session.New(creds))
}
