package v1

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/services"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

type Serverless struct {
	client   *client.Client
	Service  func(string) *service.Client
	Services *services.Client
}

// Used for testing purposes only
func (s Serverless) GetClient() *client.Client {
	return s.client
}

func New(sess *session.Session) *Serverless {
	config := client.GetDefaultConfig()
	config.Beta = false
	config.SubDomain = "serverless"
	config.APIVersion = "v1"

	return NewWithClient(client.New(sess, config))
}

func NewWithClient(client *client.Client) *Serverless {
	return &Serverless{
		client:   client,
		Service:  func(sid string) *service.Client { return service.New(client, sid) },
		Services: services.New(client),
	}
}

func NewWithCredentials(creds *credentials.Credentials) *Serverless {
	return New(session.New(creds))
}
