// Package v2 contains auto-generated files. DO NOT MODIFY
package v2

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/studio/v2/flow"
	"github.com/RJPearson94/twilio-sdk-go/service/studio/v2/flow_validation"
	"github.com/RJPearson94/twilio-sdk-go/service/studio/v2/flows"
	"github.com/RJPearson94/twilio-sdk-go/session"
	sessionCredentials "github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

// Studio client is used to manage resources for Twilio Studio
// See https://www.twilio.com/docs/studio for more details
// This client is currently in beta and subject to change. Please use with caution
type Studio struct {
	client *client.Client

	Flow           func(string) *flow.Client
	FlowValidation *flow_validation.Client
	Flows          *flows.Client
}

// NewWithClient creates a new instance of the client with a HTTP client
func NewWithClient(client *client.Client) *Studio {
	return &Studio{
		client: client,

		Flow: func(flowSid string) *flow.Client {
			return flow.New(client, flow.ClientProperties{
				Sid: flowSid,
			})
		},
		FlowValidation: flow_validation.New(client),
		Flows:          flows.New(client),
	}
}

// GetClient is used for testing purposes only
func (s Studio) GetClient() *client.Client {
	return s.client
}

// NewWithCredentials creates a new instance of the client with credentials
func NewWithCredentials(creds *sessionCredentials.Credentials) *Studio {
	return New(session.New(creds))
}

// New creates a new instance of the client using session data
func New(sess *session.Session) *Studio {
	config := client.GetDefaultConfig()
	config.Beta = true
	config.SubDomain = "studio"
	config.APIVersion = "v2"

	return NewWithClient(client.New(sess, config))
}
