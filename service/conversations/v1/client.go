package v1

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/conversation"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/conversations"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

type Conversations struct {
	client        *client.Client
	Conversations *conversations.Client
	Conversation  func(string) *conversation.Client
}

// Used for testing purposes only
func (s Conversations) GetClient() *client.Client {
	return s.client
}

func New(sess *session.Session) *Conversations {
	config := client.GetDefaultConfig()
	config.Beta = true
	config.SubDomain = "conversations"
	config.APIVersion = "v1"

	return NewWithClient(client.New(sess, config))
}

func NewWithClient(client *client.Client) *Conversations {
	return &Conversations{
		client:        client,
		Conversations: conversations.New(client),
		Conversation:  func(sid string) *conversation.Client { return conversation.New(client, sid) },
	}
}

func NewWithCredentials(creds *credentials.Credentials) *Conversations {
	return New(session.New(creds))
}
