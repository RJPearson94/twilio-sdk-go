// This is an autogenerated file. DO NOT MODIFY
package sync_stream

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/sync_stream/messages"
)

type Client struct {
	client *client.Client

	serviceSid string
	sid        string

	// Sub client to manage messages resources
	Messages *messages.Client
}

// The properties required to manage the syncstream resources
type ClientProperties struct {
	ServiceSid string
	Sid        string
}

// Create a new instance of the client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid: properties.ServiceSid,
		sid:        properties.Sid,

		Messages: messages.New(client, messages.ClientProperties{
			ServiceSid:    properties.ServiceSid,
			SyncStreamSid: properties.Sid,
		}),
	}
}
