// Package v1 contains auto-generated files. DO NOT MODIFY
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
)

// Conversations client is used to manage resources for Twilio Conversations
// See https://www.twilio.com/docs/conversations for more details
// This client is currently in beta and subject to change. Please use with caution
type Conversations struct {
	client *client.Client

	Configuration func() *configuration.Client
	Conversation  func(string) *conversation.Client
	Conversations *conversations.Client
	Credential    func(string) *credential.Client
	Credentials   *credentials.Client
	Role          func(string) *role.Client
	Roles         *roles.Client
	Service       func(string) *service.Client
	Services      *services.Client
	User          func(string) *user.Client
	Users         *users.Client
	Webhook       func() *webhook.Client
}

// NewWithClient creates a new instance of the client with a HTTP client
func NewWithClient(client *client.Client) *Conversations {
	return &Conversations{
		client: client,

		Configuration: func() *configuration.Client { return configuration.New(client) },
		Conversation: func(conversationSid string) *conversation.Client {
			return conversation.New(client, conversation.ClientProperties{
				Sid: conversationSid,
			})
		},
		Conversations: conversations.New(client),
		Credential: func(credentialSid string) *credential.Client {
			return credential.New(client, credential.ClientProperties{
				Sid: credentialSid,
			})
		},
		Credentials: credentials.New(client),
		Role: func(roleSid string) *role.Client {
			return role.New(client, role.ClientProperties{
				Sid: roleSid,
			})
		},
		Roles: roles.New(client),
		Service: func(serviceSid string) *service.Client {
			return service.New(client, service.ClientProperties{
				Sid: serviceSid,
			})
		},
		Services: services.New(client),
		User: func(userSid string) *user.Client {
			return user.New(client, user.ClientProperties{
				Sid: userSid,
			})
		},
		Users:   users.New(client),
		Webhook: func() *webhook.Client { return webhook.New(client) },
	}
}

// GetClient is used for testing purposes only
func (s Conversations) GetClient() *client.Client {
	return s.client
}

// New creates a new instance of the client using session data and config
func New(sess *session.Session, clientConfig *client.Config) *Conversations {
	config := client.NewAPIClientConfig(clientConfig)
	config.Beta = true
	config.SubDomain = "conversations"
	config.APIVersion = "v1"

	return NewWithClient(client.New(sess, config))
}
