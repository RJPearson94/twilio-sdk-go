// Package v1 contains auto-generated files. DO NOT MODIFY
package v1

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/services"
	"github.com/RJPearson94/twilio-sdk-go/session"
)

// Serverless client is used to manage resources for Twilio Serverless/ Runtime
// See https://www.twilio.com/docs/runtime for more details
type Serverless struct {
	client *client.Client

	Service  func(string) *service.Client
	Services *services.Client
}

// NewWithClient creates a new instance of the client with a HTTP client
func NewWithClient(client *client.Client) *Serverless {
	return &Serverless{
		client: client,

		Service: func(serviceSid string) *service.Client {
			return service.New(client, service.ClientProperties{
				Sid: serviceSid,
			})
		},
		Services: services.New(client),
	}
}

// GetClient is used for testing purposes only
func (s Serverless) GetClient() *client.Client {
	return s.client
}

// New creates a new instance of the client using session data and config
func New(sess *session.Session, clientConfig *client.Config) *Serverless {
	config := client.NewAPIClientConfig(clientConfig)
	config.Beta = false
	config.SubDomain = "serverless"
	config.APIVersion = "v1"

	return NewWithClient(client.New(sess, config))
}
