package v1

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/conversation"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/conversations"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/role"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/roles"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/webhook"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

// Conversations client is used to manage resources for Twilio Conversations
// See https://www.twilio.com/docs/conversations for more details
type Conversations struct {
	client        *client.Client
	Conversations *conversations.Client
	Conversation  func(string) *conversation.Client
	Roles         *roles.Client
	Role          func(string) *role.Client
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
		client:        client,
		Conversations: conversations.New(client),
		Conversation: func(sid string) *conversation.Client {
			return conversation.New(client, conversation.ClientProperties{
				Sid: sid,
			})
		},
		Roles: roles.New(client),
		Role: func(sid string) *role.Client {
			return role.New(client, role.ClientProperties{
				Sid: sid,
			})
		},
		Webhook: func() *webhook.Client { return webhook.New(client) },
	}
}

// NewWithCredentials creates a new instance of the client with credentials
func NewWithCredentials(creds *credentials.Credentials) *Conversations {
	return New(session.New(creds))
}
