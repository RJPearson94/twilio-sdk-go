// Package toll_free contains auto-generated files. DO NOT MODIFY
package toll_free

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing available toll free phone number resources
// See https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-tollfree-resource for more details
type Client struct {
	client *client.Client

	accountSid  string
	countryCode string
}

// ClientProperties are the properties required to manage the toll free resources
type ClientProperties struct {
	AccountSid  string
	CountryCode string
}

// New creates a new instance of the toll free client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		accountSid:  properties.AccountSid,
		countryCode: properties.CountryCode,
	}
}
