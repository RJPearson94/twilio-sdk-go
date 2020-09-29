// Package field_values contains auto-generated files. DO NOT MODIFY
package field_values

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing field value resources
// See https://www.twilio.com/docs/autopilot/api/field-value for more details
type Client struct {
	client *client.Client

	assistantSid string
	fieldTypeSid string
}

// ClientProperties are the properties required to manage the field values resources
type ClientProperties struct {
	AssistantSid string
	FieldTypeSid string
}

// New creates a new instance of the field values client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		assistantSid: properties.AssistantSid,
		fieldTypeSid: properties.FieldTypeSid,
	}
}
