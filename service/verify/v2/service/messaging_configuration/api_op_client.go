// Package messaging_configuration contains auto-generated files. DO NOT MODIFY
package messaging_configuration

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific messaging configuration resource
type Client struct {
	client *client.Client

	countryCode string
	serviceSid  string
}

// ClientProperties are the properties required to manage the messaging configuration resources
type ClientProperties struct {
	CountryCode string
	ServiceSid  string
}

// New creates a new instance of the messaging configuration client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		countryCode: properties.CountryCode,
		serviceSid:  properties.ServiceSid,
	}
}
