package v1

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/service"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/services"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

type Proxy struct {
	client   *client.Client
	Service  func(string) *service.Client
	Services *services.Client
}

// Used for testing purposes only
func (s Proxy) GetClient() *client.Client {
	return s.client
}

func New(sess *session.Session) *Proxy {
	config := client.GetDefaultConfig()
	config.Beta = true
	config.SubDomain = "proxy"
	config.APIVersion = "v1"

	return NewWithClient(client.New(sess, config))
}

func NewWithClient(client *client.Client) *Proxy {
	return &Proxy{
		client:   client,
		Services: services.New(client),
		Service: func(sid string) *service.Client {
			return service.New(client, service.ClientProperties{
				Sid: sid,
			})
		},
	}
}

func NewWithCredentials(creds *credentials.Credentials) *Proxy {
	return New(session.New(creds))
}
