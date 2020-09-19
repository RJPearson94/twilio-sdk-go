package twilio

import (
	"github.com/RJPearson94/twilio-sdk-go/service/api"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot"
	"github.com/RJPearson94/twilio-sdk-go/service/chat"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations"
	"github.com/RJPearson94/twilio-sdk-go/service/fax"
	"github.com/RJPearson94/twilio-sdk-go/service/flex"
	"github.com/RJPearson94/twilio-sdk-go/service/messaging"
	"github.com/RJPearson94/twilio-sdk-go/service/monitor"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless"
	"github.com/RJPearson94/twilio-sdk-go/service/studio"
	"github.com/RJPearson94/twilio-sdk-go/service/sync"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
	"github.com/RJPearson94/twilio-sdk-go/twiml"
)

// Twilio clients manage all the available Twilio services & resources within the SDK
type Twilio struct {
	API           *api.API
	Autopilot     *autopilot.Autopilot
	Chat          *chat.Chat
	Conversations *conversations.Conversations
	Fax           *fax.Fax
	Flex          *flex.Flex
	Messaging     *messaging.Messaging
	Monitor       *monitor.Monitor
	Proxy         *proxy.Proxy
	Serverless    *serverless.Serverless
	Studio        *studio.Studio
	Sync          *sync.Sync
	TaskRouter    *taskrouter.TaskRouter
	TwiML         *twiml.TwiML
}

// New create a new instance of the client using session data
func New(sess *session.Session) *Twilio {
	return &Twilio{
		API:           api.New(sess),
		Autopilot:     autopilot.New(sess),
		Chat:          chat.New(sess),
		Conversations: conversations.New(sess),
		Fax:           fax.New(sess),
		Flex:          flex.New(sess),
		Messaging:     messaging.New(sess),
		Monitor:       monitor.New(sess),
		Proxy:         proxy.New(sess),
		Serverless:    serverless.New(sess),
		Studio:        studio.New(sess),
		Sync:          sync.New(sess),
		TaskRouter:    taskrouter.New(sess),
		TwiML:         twiml.New(),
	}
}

// NewWithCredentials creates a new instance of the client with credentials
func NewWithCredentials(creds *credentials.Credentials) *Twilio {
	return New(session.New(creds))
}
