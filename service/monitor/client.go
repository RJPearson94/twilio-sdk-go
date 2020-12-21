package monitor

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	v1 "github.com/RJPearson94/twilio-sdk-go/service/monitor/v1"
	"github.com/RJPearson94/twilio-sdk-go/session"
)

// Monitor client is used to manage versioned resources for Twilio Monitor
type Monitor struct {
	V1 *v1.Monitor
}

// New creates a new instance of the client using session data and config
func New(sess *session.Session, config *client.Config) *Monitor {
	return &Monitor{
		V1: v1.New(sess, config),
	}
}
