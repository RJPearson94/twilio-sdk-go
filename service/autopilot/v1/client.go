package v1

import (
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistants"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/session"
)

type Autopilot struct {
	client     *client.Client
	Assistant  func(string) *assistant.Client
	Assistants *assistants.Client
}

// Used for testing purposes only
func (s Autopilot) GetClient() *client.Client {
	return s.client
}

func New(sess *session.Session) *Autopilot {
	config := client.GetDefaultConfig()
	config.Beta = false
	config.SubDomain = "autopilot"
	config.APIVersion = "v1"

	return NewWithClient(client.New(sess, config))
}

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

func NewWithCredentials(creds *credentials.Credentials) *Autopilot {
	return New(session.New(creds))
}
