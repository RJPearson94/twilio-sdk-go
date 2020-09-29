// Package user contains auto-generated files. DO NOT MODIFY
package user

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service/user/binding"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service/user/bindings"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service/user/channel"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service/user/channels"
)

// Client for managing a specific user resource
// See https://www.twilio.com/docs/chat/rest/user-resource for more details
type Client struct {
	client *client.Client

	serviceSid string
	sid        string

	Binding  func(string) *binding.Client
	Bindings *bindings.Client
	Channel  func(string) *channel.Client
	Channels *channels.Client
}

// ClientProperties are the properties required to manage the user resources
type ClientProperties struct {
	ServiceSid string
	Sid        string
}

// New creates a new instance of the user client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid: properties.ServiceSid,
		sid:        properties.Sid,

		Binding: func(bindingSid string) *binding.Client {
			return binding.New(client, binding.ClientProperties{
				ServiceSid: properties.ServiceSid,
				Sid:        bindingSid,
				UserSid:    properties.Sid,
			})
		},
		Bindings: bindings.New(client, bindings.ClientProperties{
			ServiceSid: properties.ServiceSid,
			UserSid:    properties.Sid,
		}),
		Channel: func(channelSid string) *channel.Client {
			return channel.New(client, channel.ClientProperties{
				ServiceSid: properties.ServiceSid,
				Sid:        channelSid,
				UserSid:    properties.Sid,
			})
		},
		Channels: channels.New(client, channels.ClientProperties{
			ServiceSid: properties.ServiceSid,
			UserSid:    properties.Sid,
		}),
	}
}
