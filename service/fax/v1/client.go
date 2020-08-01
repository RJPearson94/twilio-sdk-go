package v1

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/fax/v1/fax"
	"github.com/RJPearson94/twilio-sdk-go/service/fax/v1/faxes"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

// Fax client is used to manage resources for Programmable Fax
// See https://www.twilio.com/docs/fax for more details
type Fax struct {
	client *client.Client
	Faxes  *faxes.Client
	Fax    func(string) *fax.Client
}

// Used for testing purposes only
func (s Fax) GetClient() *client.Client {
	return s.client
}

// New creates a new instance of the client using session data
func New(sess *session.Session) *Fax {
	config := client.GetDefaultConfig()
	config.Beta = false
	config.SubDomain = "fax"
	config.APIVersion = "v1"

	return NewWithClient(client.New(sess, config))
}

// NewWithClient creates a new instance of the client with a HTTP client
func NewWithClient(client *client.Client) *Fax {
	return &Fax{
		client: client,
		Faxes:  faxes.New(client),
		Fax: func(sid string) *fax.Client {
			return fax.New(client, fax.ClientProperties{
				Sid: sid,
			})
		},
	}
}

// NewWithCredentials creates a new instance of the client with credentials
func NewWithCredentials(creds *credentials.Credentials) *Fax {
	return New(session.New(creds))
}
