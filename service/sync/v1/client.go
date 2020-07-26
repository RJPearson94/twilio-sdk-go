package v1

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/services"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

type Sync struct {
	client *client.Client

	Service  func(string) *service.Client
	Services *services.Client
}

// Used for testing purposes only
func (s Sync) GetClient() *client.Client {
	return s.client
}

func New(sess *session.Session) *Sync {
	config := client.GetDefaultConfig()
	config.Beta = false
	config.SubDomain = "sync"
	config.APIVersion = "v1"

	return NewWithClient(client.New(sess, config))
}

func NewWithClient(client *client.Client) *Sync {
	return &Sync{
		client: client,
		Service: func(sid string) *service.Client {
			return service.New(client, service.ClientProperties{
				Sid: sid,
			})
		},
		Services: services.New(client),
	}
}

func NewWithCredentials(creds *credentials.Credentials) *Sync {
	return New(session.New(creds))
}
