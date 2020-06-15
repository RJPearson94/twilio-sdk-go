package v1

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/configuration"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

type Flex struct {
	client        *client.Client
	Configuration func() *configuration.Client
}

// Used for testing purposes only
func (s Flex) GetClient() *client.Client {
	return s.client
}

func New(sess *session.Session) *Flex {
	config := client.GetDefaultConfig()
	config.Beta = false
	config.SubDomain = "flex-api"
	config.APIVersion = "v1"

	return NewWithClient(client.New(sess, config))
}

func NewWithClient(client *client.Client) *Flex {
	return &Flex{
		client:        client,
		Configuration: func() *configuration.Client { return configuration.New(client) },
	}
}

func NewWithCredentials(creds *credentials.Credentials) *Flex {
	return New(session.New(creds))
}
