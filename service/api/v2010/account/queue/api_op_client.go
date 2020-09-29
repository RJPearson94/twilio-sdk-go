// Package queue contains auto-generated files. DO NOT MODIFY
package queue

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/queue/member"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/queue/members"
)

// Client for managing a specific queue resource
// See https://www.twilio.com/docs/voice/api/queue-resource for more details
type Client struct {
	client *client.Client

	accountSid string
	sid        string

	Member  func(string) *member.Client
	Members *members.Client
}

// ClientProperties are the properties required to manage the queue resources
type ClientProperties struct {
	AccountSid string
	Sid        string
}

// New creates a new instance of the queue client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		accountSid: properties.AccountSid,
		sid:        properties.Sid,

		Member: func(sid string) *member.Client {
			return member.New(client, member.ClientProperties{
				AccountSid: properties.AccountSid,
				QueueSid:   properties.Sid,
				Sid:        sid,
			})
		},
		Members: members.New(client, members.ClientProperties{
			AccountSid: properties.AccountSid,
			QueueSid:   properties.Sid,
		}),
	}
}
