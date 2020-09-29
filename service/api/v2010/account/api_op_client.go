// Package account contains auto-generated files. DO NOT MODIFY
package account

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/balance"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/key"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/keys"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/message"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/messages"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/queue"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/queues"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/tokens"
)

// Client for managing a specific account resource
// See https://www.twilio.com/docs/iam/api/account for more details
type Client struct {
	client *client.Client

	sid string

	Balance  func() *balance.Client
	Key      func(string) *key.Client
	Keys     *keys.Client
	Message  func(string) *message.Client
	Messages *messages.Client
	Queue    func(string) *queue.Client
	Queues   *queues.Client
	Tokens   *tokens.Client
}

// ClientProperties are the properties required to manage the account resources
type ClientProperties struct {
	Sid string
}

// New creates a new instance of the account client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		sid: properties.Sid,

		Balance: func() *balance.Client {
			return balance.New(client, balance.ClientProperties{
				AccountSid: properties.Sid,
			})
		},
		Key: func(keySid string) *key.Client {
			return key.New(client, key.ClientProperties{
				AccountSid: properties.Sid,
				Sid:        keySid,
			})
		},
		Keys: keys.New(client, keys.ClientProperties{
			AccountSid: properties.Sid,
		}),
		Message: func(messageSid string) *message.Client {
			return message.New(client, message.ClientProperties{
				AccountSid: properties.Sid,
				Sid:        messageSid,
			})
		},
		Messages: messages.New(client, messages.ClientProperties{
			AccountSid: properties.Sid,
		}),
		Queue: func(queueSid string) *queue.Client {
			return queue.New(client, queue.ClientProperties{
				AccountSid: properties.Sid,
				Sid:        queueSid,
			})
		},
		Queues: queues.New(client, queues.ClientProperties{
			AccountSid: properties.Sid,
		}),
		Tokens: tokens.New(client, tokens.ClientProperties{
			AccountSid: properties.Sid,
		}),
	}
}
