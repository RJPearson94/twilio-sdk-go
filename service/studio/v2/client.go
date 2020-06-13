package v2

import (
	"github.com/RJPearson94/twilio-sdk-go/service/studio/v2/flow"
	"github.com/RJPearson94/twilio-sdk-go/service/studio/v2/flow_validation"
	"github.com/RJPearson94/twilio-sdk-go/service/studio/v2/flows"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/session"
)

type Studio struct {
	client         *client.Client
	Flow           func(string) *flow.Client
	Flows          *flows.Client
	FlowValidation *flow_validation.Client
}

// Used for testing purposes only
func (s Studio) GetClient() *client.Client {
	return s.client
}

func New(sess *session.Session) *Studio {
	config := client.GetDefaultConfig()
	config.Beta = true
	config.SubDomain = "studio"
	config.APIVersion = "v2"

	return NewWithClient(client.New(sess, config))
}

func NewWithClient(client *client.Client) *Studio {
	return &Studio{
		client:         client,
		Flows:          flows.New(client),
		Flow:           func(sid string) *flow.Client { return flow.New(client, sid) },
		FlowValidation: flow_validation.New(client),
	}
}

func NewWithCredentials(creds *credentials.Credentials) *Studio {
	return New(session.New(creds))
}
