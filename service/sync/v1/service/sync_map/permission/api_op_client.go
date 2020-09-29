// Package permission contains auto-generated files. DO NOT MODIFY
package permission

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing map item permission resources
// See https://www.twilio.com/docs/sync/api/sync-map-permission-resource for more details
type Client struct {
	client *client.Client

	identity   string
	serviceSid string
	syncMapSid string
}

// ClientProperties are the properties required to manage the permission resources
type ClientProperties struct {
	Identity   string
	ServiceSid string
	SyncMapSid string
}

// New creates a new instance of the permission client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		identity:   properties.Identity,
		serviceSid: properties.ServiceSid,
		syncMapSid: properties.SyncMapSid,
	}
}
