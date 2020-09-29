// Package permissions contains auto-generated files. DO NOT MODIFY
package permissions

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing list permission resources
// See https://www.twilio.com/docs/sync/api/sync-list-permission-resource for more details
type Client struct {
	client *client.Client

	serviceSid  string
	syncListSid string
}

// ClientProperties are the properties required to manage the permissions resources
type ClientProperties struct {
	ServiceSid  string
	SyncListSid string
}

// New creates a new instance of the permissions client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid:  properties.ServiceSid,
		syncListSid: properties.SyncListSid,
	}
}
