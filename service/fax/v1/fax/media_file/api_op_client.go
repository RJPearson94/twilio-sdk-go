// Package media_file contains auto-generated files. DO NOT MODIFY
package media_file

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific media resource
// See https://www.twilio.com/docs/fax/api/fax-media-resource for more details
type Client struct {
	client *client.Client

	faxSid string
	sid    string
}

// ClientProperties are the properties required to manage the media file resources
type ClientProperties struct {
	FaxSid string
	Sid    string
}

// New creates a new instance of the media file client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		faxSid: properties.FaxSid,
		sid:    properties.Sid,
	}
}
