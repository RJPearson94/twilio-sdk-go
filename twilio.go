package twilio

import (
	"github.com/RJPearson94/twilio-sdk-go/service/api"
	"github.com/RJPearson94/twilio-sdk-go/service/chat"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations"
	"github.com/RJPearson94/twilio-sdk-go/service/fax"
	"github.com/RJPearson94/twilio-sdk-go/service/flex"
	"github.com/RJPearson94/twilio-sdk-go/service/messaging"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless"
	"github.com/RJPearson94/twilio-sdk-go/service/studio"
	"github.com/RJPearson94/twilio-sdk-go/service/sync"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
	"github.com/RJPearson94/twilio-sdk-go/twiml"
)

// Twilio manage Twilio services & resources
type Twilio struct {
	// Sub client to manage api rest api resources
	API *api.API
	// Sub client to manage chat rest api resources
	Chat *chat.Chat
	// Sub client to manage conversations rest api resources
	Conversations *conversations.Conversations
	// Sub client to manage fax rest api resources
	Fax *fax.Fax
	// Sub client to manage flex rest api resources
	Flex *flex.Flex
	// Sub client to manage messaging rest api resources
	Messaging *messaging.Messaging
	// Sub client to manage proxy rest api resources
	Proxy *proxy.Proxy
	// Sub client to manage serverless rest api resources
	Serverless *serverless.Serverless
	// Sub client to manage studio rest api resources
	Studio *studio.Studio
	// Sub client to manage sync rest api resources
	Sync *sync.Sync
	// Sub client to manage task router rest api resources
	TaskRouter *taskrouter.TaskRouter
	// Sub client to manage TwiML resources
	TwiML *twiml.TwiML
}

// New create a new instance of the client using session data
func New(sess *session.Session) *Twilio {
	return &Twilio{
		API:           api.New(sess),
		Chat:          chat.New(sess),
		Conversations: conversations.New(sess),
		Fax:           fax.New(sess),
		Flex:          flex.New(sess),
		Messaging:     messaging.New(sess),
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
