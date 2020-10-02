// Package participants contains auto-generated files. DO NOT MODIFY
package participants

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing participant resources
// See https://www.twilio.com/docs/voice/api/conference-participant-resource for more details
type Client struct {
	client *client.Client

	accountSid     string
	conferencesSid string
}

// ClientProperties are the properties required to manage the participants resources
type ClientProperties struct {
	AccountSid     string
	ConferencesSid string
}

// New creates a new instance of the participants client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		accountSid:     properties.AccountSid,
		conferencesSid: properties.ConferencesSid,
	}
}
