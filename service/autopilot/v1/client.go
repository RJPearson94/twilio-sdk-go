package v1

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistants"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

// Autopilot client is used to manage resources for Twilio Autopilot
// See https://www.twilio.com/docs/autopilot for more details
type Autopilot struct {
	client     *client.Client
	Assistant  func(string) *assistant.Client
	Assistants *assistants.Client
}

// Used for testing purposes only
func (s Autopilot) GetClient() *client.Client {
	return s.client
}

// New creates a new instance of the client using session data
func New(sess *session.Session) *Autopilot {
	config := client.GetDefaultConfig()
	config.Beta = false
	config.SubDomain = "autopilot"
	config.APIVersion = "v1"

	return NewWithClient(client.New(sess, config))
}

// NewWithClient creates a new instance of the client with a HTTP client
func NewWithClient(client *client.Client) *Autopilot {
	return &Autopilot{
		client: client,
		Assistant: func(sid string) *assistant.Client {
			return assistant.New(client, assistant.ClientProperties{
				Sid: sid,
			})
		},
		Assistants: assistants.New(client),
	}
}

// NewWithCredentials creates a new instance of the client with credentials
func NewWithCredentials(creds *credentials.Credentials) *Autopilot {
	return New(session.New(creds))
}
