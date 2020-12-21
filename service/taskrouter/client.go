package taskrouter

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	v1 "github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1"
	"github.com/RJPearson94/twilio-sdk-go/session"
)

// TaskRouter client is used to manage versioned resources for Twilio TaskRouter
// See https://www.twilio.com/docs/taskrouter for more details on the API
// See https://www.twilio.com/taskrouter for more details on the product
type TaskRouter struct {
	V1 *v1.TaskRouter
}

// New creates a new instance of the client using session data and config
func New(sess *session.Session, config *client.Config) *TaskRouter {
	return &TaskRouter{
		V1: v1.New(sess, config),
	}
}
