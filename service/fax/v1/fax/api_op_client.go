// Package fax contains auto-generated files. DO NOT MODIFY
package fax

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/fax/v1/fax/media_file"
	"github.com/RJPearson94/twilio-sdk-go/service/fax/v1/fax/media_files"
)

// Client for managing a specific fax resource
// See https://www.twilio.com/docs/fax/api/fax-resource for more details
type Client struct {
	client *client.Client

	sid string

	MediaFile  func(string) *media_file.Client
	MediaFiles *media_files.Client
}

// ClientProperties are the properties required to manage the fax resources
type ClientProperties struct {
	Sid string
}

// New creates a new instance of the fax client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		sid: properties.Sid,

		MediaFile: func(mediaSid string) *media_file.Client {
			return media_file.New(client, media_file.ClientProperties{
				FaxSid: properties.Sid,
				Sid:    mediaSid,
			})
		},
		MediaFiles: media_files.New(client, media_files.ClientProperties{
			FaxSid: properties.Sid,
		}),
	}
}
