// Package v1 contains auto-generated files. DO NOT MODIFY
package v1

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/trunking/v1/trunk"
	"github.com/RJPearson94/twilio-sdk-go/service/trunking/v1/trunks"
	"github.com/RJPearson94/twilio-sdk-go/session"
)

// Trunking client is used to manage resources for Twilio SIP Trunking
// See https://www.twilio.com/docs/sip-trunking for more details
// This client is currently in beta and subject to change. Please use with caution
type Trunking struct {
	client *client.Client

	Trunk  func(string) *trunk.Client
	Trunks *trunks.Client
}

// NewWithClient creates a new instance of the client with a HTTP client
func NewWithClient(client *client.Client) *Trunking {
	return &Trunking{
		client: client,

		Trunk: func(trunkSid string) *trunk.Client {
			return trunk.New(client, trunk.ClientProperties{
				Sid: trunkSid,
			})
		},
		Trunks: trunks.New(client),
	}
}

// GetClient is used for testing purposes only
func (s Trunking) GetClient() *client.Client {
	return s.client
}

// New creates a new instance of the client using session data and config
func New(sess *session.Session, clientConfig *client.Config) *Trunking {
	config := client.NewAPIClientConfig(clientConfig)
	config.Beta = true
	config.SubDomain = "trunking"
	config.APIVersion = "v1"

	return NewWithClient(client.New(sess, config))
}
