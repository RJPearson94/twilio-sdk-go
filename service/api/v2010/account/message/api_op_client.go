// This is an autogenerated file. DO NOT MODIFY
package message

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/message/feedback"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/message/media"
)

// Client for managing a specific message resource
// See https://www.twilio.com/docs/sms/api/message-resource for more details
type Client struct {
	client *client.Client

	accountSid string
	sid        string

	Feedback *feedback.Client
	Media    func(string) *media.Client
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

		Feedback: feedback.New(client, feedback.ClientProperties{
			AccountSid: properties.AccountSid,
			MessageSid: properties.Sid,
		}),
		Media: func(mediaSid string) *media.Client {
			return media.New(client, media.ClientProperties{
				AccountSid: properties.AccountSid,
				MessageSid: properties.Sid,
				Sid:        mediaSid,
			})
		},
	}
}
