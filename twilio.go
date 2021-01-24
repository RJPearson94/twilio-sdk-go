package twilio

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/api"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot"
	"github.com/RJPearson94/twilio-sdk-go/service/chat"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations"
	"github.com/RJPearson94/twilio-sdk-go/service/fax"
	"github.com/RJPearson94/twilio-sdk-go/service/flex"
	"github.com/RJPearson94/twilio-sdk-go/service/lookups"
	"github.com/RJPearson94/twilio-sdk-go/service/messaging"
	"github.com/RJPearson94/twilio-sdk-go/service/monitor"
	"github.com/RJPearson94/twilio-sdk-go/service/notify"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless"
	"github.com/RJPearson94/twilio-sdk-go/service/studio"
	"github.com/RJPearson94/twilio-sdk-go/service/sync"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter"
	"github.com/RJPearson94/twilio-sdk-go/service/trunking"
	"github.com/RJPearson94/twilio-sdk-go/service/verify"
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
	Lookups       *lookups.Lookups
	Messaging     *messaging.Messaging
	Monitor       *monitor.Monitor
	Notify        *notify.Notify
	Proxy         *proxy.Proxy
	Serverless    *serverless.Serverless
	Studio        *studio.Studio
	Sync          *sync.Sync
	TaskRouter    *taskrouter.TaskRouter
	Trunking      *trunking.Trunking
	TwiML         *twiml.TwiML
	Verify        *verify.Verify
}

// New create a new instance of the client using session data
func New(sess *session.Session) *Twilio {
	return NewWithConfig(sess, nil)
}

// NewWithConfig create a new instance of the client using session data and config
func NewWithConfig(sess *session.Session, config *client.Config) *Twilio {
	return &Twilio{
		API:           api.New(sess, config),
		Autopilot:     autopilot.New(sess, config),
		Chat:          chat.New(sess, config),
		Conversations: conversations.New(sess, config),
		Fax:           fax.New(sess, config),
		Flex:          flex.New(sess, config),
		Lookups:       lookups.New(sess, config),
		Messaging:     messaging.New(sess, config),
		Monitor:       monitor.New(sess, config),
		Notify:        notify.New(sess, config),
		Proxy:         proxy.New(sess, config),
		Serverless:    serverless.New(sess, config),
		Studio:        studio.New(sess, config),
		Sync:          sync.New(sess, config),
		TaskRouter:    taskrouter.New(sess, config),
		Trunking:      trunking.New(sess, config),
		TwiML:         twiml.New(),
		Verify:        verify.New(sess, config),
	}
}

// NewWithCredentials creates a new instance of the client with credentials
func NewWithCredentials(creds *credentials.Credentials) *Twilio {
	return New(session.New(creds))
}
