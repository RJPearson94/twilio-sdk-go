// Package sync_list contains auto-generated files. DO NOT MODIFY
package sync_list

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/sync_list/item"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/sync_list/items"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/sync_list/permission"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/sync_list/permissions"
)

// Client for managing a specific list resource
// See https://www.twilio.com/docs/sync/api/list-resource for more details
type Client struct {
	client *client.Client

	serviceSid string
	sid        string

	Item        func(int) *item.Client
	Items       *items.Client
	Permission  func(string) *permission.Client
	Permissions *permissions.Client
}

// ClientProperties are the properties required to manage the synclist resources
type ClientProperties struct {
	ServiceSid string
	Sid        string
}

// New creates a new instance of the synclist client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid: properties.ServiceSid,
		sid:        properties.Sid,

		Item: func(index int) *item.Client {
			return item.New(client, item.ClientProperties{
				Index:       index,
				ServiceSid:  properties.ServiceSid,
				SyncListSid: properties.Sid,
			})
		},
		Items: items.New(client, items.ClientProperties{
			ServiceSid:  properties.ServiceSid,
			SyncListSid: properties.Sid,
		}),
		Permission: func(identity string) *permission.Client {
			return permission.New(client, permission.ClientProperties{
				Identity:    identity,
				ServiceSid:  properties.ServiceSid,
				SyncListSid: properties.Sid,
			})
		},
		Permissions: permissions.New(client, permissions.ClientProperties{
			ServiceSid:  properties.ServiceSid,
			SyncListSid: properties.Sid,
		}),
	}
}
