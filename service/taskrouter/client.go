package taskrouter

import (
	v1 "github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

// TaskRouter client is used to manage versioned resources for Twilio TaskRouter
// See https://www.twilio.com/docs/taskrouter for more details on the API
// See https://www.twilio.com/taskrouter for more details on the product
type TaskRouter struct {
	V1 *v1.TaskRouter
}

// New creates a new instance of the client using session data
func New(sess *session.Session) *TaskRouter {
	return &TaskRouter{
		V1: v1.New(sess),
	}
}

// NewWithCredentials creates a new instance of the client with credentials
func NewWithCredentials(creds *credentials.Credentials) *TaskRouter {
	return New(session.New(creds))
}
