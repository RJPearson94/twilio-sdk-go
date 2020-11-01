package v1

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/configuration"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/conversation"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/conversations"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/credential"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/credentials"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/role"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/roles"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/service"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/services"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/user"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/users"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/webhook"
	"github.com/RJPearson94/twilio-sdk-go/session"
	sessionCredentials "github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

// Conversations client is used to manage resources for Twilio Conversations
// See https://www.twilio.com/docs/conversations for more details
type Conversations struct {
	client        *client.Client
	Configuration func() *configuration.Client
	Conversations *conversations.Client
	Conversation  func(string) *conversation.Client
	Credentials   *credentials.Client
	Credential    func(string) *credential.Client
	Roles         *roles.Client
	Role          func(string) *role.Client
	Services      *services.Client
	Service       func(string) *service.Client
	Users         *users.Client
	User          func(string) *user.Client
	Webhook       func() *webhook.Client
}

// Used for testing purposes only
func (s Conversations) GetClient() *client.Client {
	return s.client
}

// New creates a new instance of the client using session data
func New(sess *session.Session) *Conversations {
	config := client.GetDefaultConfig()
	config.Beta = true
	config.SubDomain = "conversations"
	config.APIVersion = "v1"

	return NewWithClient(client.New(sess, config))
}

// NewWithClient creates a new instance of the client with a HTTP client
func NewWithClient(client *client.Client) *Conversations {
	return &Conversations{
		client: client,
		Configuration: func() *configuration.Client {
			return configuration.New(client)
		},
		Conversations: conversations.New(client),
		Conversation: func(sid string) *conversation.Client {
			return conversation.New(client, conversation.ClientProperties{
				Sid: sid,
			})
		},
		Credentials: credentials.New(client),
		Credential: func(sid string) *credential.Client {
			return credential.New(client, credential.ClientProperties{
				Sid: sid,
			})
		},
		Roles: roles.New(client),
		Role: func(sid string) *role.Client {
			return role.New(client, role.ClientProperties{
				Sid: sid,
			})
		},
		Services: services.New(client),
		Service: func(sid string) *service.Client {
			return service.New(client, service.ClientProperties{
				Sid: sid,
			})
		},
		Users: users.New(client),
		User: func(sid string) *user.Client {
			return user.New(client, user.ClientProperties{
				Sid: sid,
			})
		},
		Webhook: func() *webhook.Client { return webhook.New(client) },
	}
}

// NewWithCredentials creates a new instance of the client with credentials
func NewWithCredentials(creds *sessionCredentials.Credentials) *Conversations {
	return New(session.New(creds))
}
