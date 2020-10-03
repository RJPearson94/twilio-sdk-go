// Package message contains auto-generated files. DO NOT MODIFY
package message

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/message/feedbacks"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/message/media_attachment"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/message/media_attachments"
)

// Client for managing a specific message resource
// See https://www.twilio.com/docs/sms/api/message-resource for more details
type Client struct {
	client *client.Client

	accountSid string
	sid        string

	Feedbacks        *feedbacks.Client
	MediaAttachment  func(string) *media_attachment.Client
	MediaAttachments *media_attachments.Client
}

// ClientProperties are the properties required to manage the message resources
type ClientProperties struct {
	AccountSid string
	Sid        string
}

// New creates a new instance of the message client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		accountSid: properties.AccountSid,
		sid:        properties.Sid,

		Feedbacks: feedbacks.New(client, feedbacks.ClientProperties{
			AccountSid: properties.AccountSid,
			MessageSid: properties.Sid,
		}),
		MediaAttachment: func(mediaSid string) *media_attachment.Client {
			return media_attachment.New(client, media_attachment.ClientProperties{
				AccountSid: properties.AccountSid,
				MessageSid: properties.Sid,
				Sid:        mediaSid,
			})
		},
		MediaAttachments: media_attachments.New(client, media_attachments.ClientProperties{
			AccountSid: properties.AccountSid,
			MessageSid: properties.Sid,
		}),
	}
}
