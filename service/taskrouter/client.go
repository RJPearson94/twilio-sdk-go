package taskrouter

import (
	v1 "github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

type TaskRouter struct {
	V1 *v1.TaskRouter
}

func New(sess *session.Session) *TaskRouter {
	c := &TaskRouter{}
	c.V1 = v1.New(sess)
	return c
}

func NewWithCredentials(creds *credentials.Credentials) *TaskRouter {
	return New(session.New(creds))
}
