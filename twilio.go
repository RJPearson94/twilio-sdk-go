package twilio

import (
	"github.com/RJPearson94/twilio-sdk-go/service/serverless"
	"github.com/RJPearson94/twilio-sdk-go/service/studio"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

type Twilio struct {
	Serverless *serverless.Serverless
	Studio     *studio.Studio
	TaskRouter *taskrouter.TaskRouter
}

func New(sess *session.Session) *Twilio {
	c := &Twilio{}
	c.Serverless = serverless.New(sess)
	c.Studio = studio.New(sess)
	c.TaskRouter = taskrouter.New(sess)
	return c
}

func NewWithCredentials(creds *credentials.Credentials) *Twilio {
	return New(session.New(creds))
}
