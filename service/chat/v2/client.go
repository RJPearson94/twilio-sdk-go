package v2

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/services"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

type Chat struct {
	client   *client.Client
	Services *services.Client
	Service  func(string) *service.Client
}

// Used for testing purposes only
func (s Chat) GetClient() *client.Client {
	return s.client
}

func New(sess *session.Session) *Chat {
	config := client.GetDefaultConfig()
	config.Beta = false
	config.SubDomain = "chat"
	config.APIVersion = "v2"

	return NewWithClient(client.New(sess, config))
}

func NewWithClient(client *client.Client) *Chat {
	return &Chat{
		client:   client,
		Services: services.New(client),
		Service:  func(sid string) *service.Client { return service.New(client, sid) },
	}
}

func NewWithCredentials(creds *credentials.Credentials) *Chat {
	return New(session.New(creds))
}
