// Package service contains auto-generated files. DO NOT MODIFY
package service

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/document"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/documents"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/sync_list"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/sync_lists"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/sync_map"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/sync_maps"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/sync_stream"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/sync_streams"
)

// Client for managing a specific service resource
// See https://www.twilio.com/docs/sync/api/service for more details
type Client struct {
	client *client.Client

	sid string

	Document    func(string) *document.Client
	Documents   *documents.Client
	SyncList    func(string) *sync_list.Client
	SyncLists   *sync_lists.Client
	SyncMap     func(string) *sync_map.Client
	SyncMaps    *sync_maps.Client
	SyncStream  func(string) *sync_stream.Client
	SyncStreams *sync_streams.Client
}

// ClientProperties are the properties required to manage the service resources
type ClientProperties struct {
	Sid string
}

// New creates a new instance of the service client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		sid: properties.Sid,

		Document: func(documentSid string) *document.Client {
			return document.New(client, document.ClientProperties{
				ServiceSid: properties.Sid,
				Sid:        documentSid,
			})
		},
		Documents: documents.New(client, documents.ClientProperties{
			ServiceSid: properties.Sid,
		}),
		SyncList: func(syncListSid string) *sync_list.Client {
			return sync_list.New(client, sync_list.ClientProperties{
				ServiceSid: properties.Sid,
				Sid:        syncListSid,
			})
		},
		SyncLists: sync_lists.New(client, sync_lists.ClientProperties{
			ServiceSid: properties.Sid,
		}),
		SyncMap: func(syncMapSid string) *sync_map.Client {
			return sync_map.New(client, sync_map.ClientProperties{
				ServiceSid: properties.Sid,
				Sid:        syncMapSid,
			})
		},
		SyncMaps: sync_maps.New(client, sync_maps.ClientProperties{
			ServiceSid: properties.Sid,
		}),
		SyncStream: func(syncStreamSid string) *sync_stream.Client {
			return sync_stream.New(client, sync_stream.ClientProperties{
				ServiceSid: properties.Sid,
				Sid:        syncStreamSid,
			})
		},
		SyncStreams: sync_streams.New(client, sync_streams.ClientProperties{
			ServiceSid: properties.Sid,
		}),
	}
}
