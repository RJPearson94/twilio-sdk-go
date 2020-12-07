package v1

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/trunking/v1/trunk"
	"github.com/RJPearson94/twilio-sdk-go/service/trunking/v1/trunks"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

// Trunking client is used to manage resources for Twilio Trunking
// See https://www.twilio.com/docs/sip-trunking for more details
type Trunking struct {
	client *client.Client
	Trunk  func(string) *trunk.Client
	Trunks *trunks.Client
}

// Used for testing purposes only
func (s Trunking) GetClient() *client.Client {
	return s.client
}

// New creates a new instance of the client using session data
func New(sess *session.Session) *Trunking {
	config := client.GetDefaultConfig()
	config.Beta = true
	config.SubDomain = "trunking"
	config.APIVersion = "v1"

	return NewWithClient(client.New(sess, config))
}

// NewWithClient creates a new instance of the client with a HTTP client
func NewWithClient(client *client.Client) *Trunking {
	return &Trunking{
		client: client,
		Trunks: trunks.New(client),
		Trunk: func(sid string) *trunk.Client {
			return trunk.New(client, trunk.ClientProperties{
				Sid: sid,
			})
		},
	}
}

// NewWithCredentials creates a new instance of the client with credentials
func NewWithCredentials(creds *credentials.Credentials) *Trunking {
	return New(session.New(creds))
}
