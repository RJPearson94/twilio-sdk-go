package twilio

import (
	"github.com/RJPearson94/twilio-sdk-go/service/chat"
	"github.com/RJPearson94/twilio-sdk-go/service/flex"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless"
	"github.com/RJPearson94/twilio-sdk-go/service/studio"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

type Twilio struct {
	Chat       *chat.Chat
	Flex       *flex.Flex
	Proxy      *proxy.Proxy
	Serverless *serverless.Serverless
	Studio     *studio.Studio
	TaskRouter *taskrouter.TaskRouter
}

func New(sess *session.Session) *Twilio {
	c := &Twilio{}
	c.Chat = chat.New(sess)
	c.Flex = flex.New(sess)
	c.Proxy = proxy.New(sess)
	c.Serverless = serverless.New(sess)
	c.Studio = studio.New(sess)
	c.TaskRouter = taskrouter.New(sess)
	return c
}

func NewWithCredentials(creds *credentials.Credentials) *Twilio {
	return New(session.New(creds))
}
