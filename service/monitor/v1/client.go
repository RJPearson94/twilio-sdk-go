package v1

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/monitor/v1/alert"
	"github.com/RJPearson94/twilio-sdk-go/service/monitor/v1/alerts"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

// Monitor client is used to manage resources for Twilio Monitor
type Monitor struct {
	client *client.Client

	Alert  func(string) *alert.Client
	Alerts *alerts.Client
}

// Used for testing purposes only
func (s Monitor) GetClient() *client.Client {
	return s.client
}

// New creates a new instance of the client using session data
func New(sess *session.Session) *Monitor {
	config := client.GetDefaultConfig()
	config.Beta = true
	config.SubDomain = "monitor"
	config.APIVersion = "v1"

	return NewWithClient(client.New(sess, config))
}

// NewWithClient creates a new instance of the client with a HTTP client
func NewWithClient(client *client.Client) *Monitor {
	return &Monitor{
		client: client,
		Alert: func(sid string) *alert.Client {
			return alert.New(client, alert.ClientProperties{
				Sid: sid,
			})
		},
		Alerts: alerts.New(client),
	}
}

// NewWithCredentials creates a new instance of the client with credentials
func NewWithCredentials(creds *credentials.Credentials) *Monitor {
	return New(session.New(creds))
}
