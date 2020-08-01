package v2

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/studio/v2/flow"
	"github.com/RJPearson94/twilio-sdk-go/service/studio/v2/flow_validation"
	"github.com/RJPearson94/twilio-sdk-go/service/studio/v2/flows"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

// Studio client is used to manage resources for Twilio Studio
// See https://www.twilio.com/docs/studio for more details
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

// New creates a new instance of the client using session data
func New(sess *session.Session) *Studio {
	config := client.GetDefaultConfig()
	config.Beta = true
	config.SubDomain = "studio"
	config.APIVersion = "v2"

	return NewWithClient(client.New(sess, config))
}

// NewWithClient creates a new instance of the client with a HTTP client
func NewWithClient(client *client.Client) *Studio {
	return &Studio{
		client: client,
		Flows:  flows.New(client),
		Flow: func(sid string) *flow.Client {
			return flow.New(client, flow.ClientProperties{
				Sid: sid,
			})
		},
		FlowValidation: flow_validation.New(client),
	}
}

// NewWithCredentials creates a new instance of the client with credentials
func NewWithCredentials(creds *credentials.Credentials) *Studio {
	return New(session.New(creds))
}
