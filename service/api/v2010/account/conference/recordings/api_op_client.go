// Package recordings contains auto-generated files. DO NOT MODIFY
package recordings

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing recording resources
// See https://www.twilio.com/docs/voice/api/recording for more details
type Client struct {
	client *client.Client

	accountSid    string
	conferenceSid string
}

// ClientProperties are the properties required to manage the recordings resources
type ClientProperties struct {
	AccountSid    string
	ConferenceSid string
}

// New creates a new instance of the recordings client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		accountSid:    properties.AccountSid,
		conferenceSid: properties.ConferenceSid,
	}
}
