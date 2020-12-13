// Package v1 contains auto-generated files. DO NOT MODIFY
package v1

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/monitor/v1/alert"
	"github.com/RJPearson94/twilio-sdk-go/service/monitor/v1/alerts"
	"github.com/RJPearson94/twilio-sdk-go/service/monitor/v1/event"
	"github.com/RJPearson94/twilio-sdk-go/service/monitor/v1/events"
	"github.com/RJPearson94/twilio-sdk-go/session"
	sessionCredentials "github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

// Monitor client is used to manage resources for Twilio Monitor
type Monitor struct {
	client *client.Client

	Alert  func(string) *alert.Client
	Alerts *alerts.Client
	Event  func(string) *event.Client
	Events *events.Client
}

// NewWithClient creates a new instance of the client with a HTTP client
func NewWithClient(client *client.Client) *Monitor {
	return &Monitor{
		client: client,

		Alert: func(alertSid string) *alert.Client {
			return alert.New(client, alert.ClientProperties{
				Sid: alertSid,
			})
		},
		Alerts: alerts.New(client),
		Event: func(eventSid string) *event.Client {
			return event.New(client, event.ClientProperties{
				Sid: eventSid,
			})
		},
		Events: events.New(client),
	}
}

// GetClient is used for testing purposes only
func (s Monitor) GetClient() *client.Client {
	return s.client
}

// NewWithCredentials creates a new instance of the client with credentials
func NewWithCredentials(creds *sessionCredentials.Credentials) *Monitor {
	return New(session.New(creds))
}

// New creates a new instance of the client using session data
func New(sess *session.Session) *Monitor {
	config := client.GetDefaultConfig()
	config.Beta = false
	config.SubDomain = "monitor"
	config.APIVersion = "v1"

	return NewWithClient(client.New(sess, config))
}
