// Package account contains auto-generated files. DO NOT MODIFY
package account

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/address"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/addresses"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/application"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/applications"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/available_phone_number"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/available_phone_numbers"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/balance"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/call"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/calls"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/conference"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/conferences"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/key"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/keys"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/message"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/messages"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/queue"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/queues"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/recording"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/recordings"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/tokens"
)

// Client for managing a specific account resource
// See https://www.twilio.com/docs/iam/api/account for more details
type Client struct {
	client *client.Client

	sid string

	Address               func(string) *address.Client
	Addresses             *addresses.Client
	Application           func(string) *application.Client
	Applications          *applications.Client
	AvailablePhoneNumber  func(string) *available_phone_number.Client
	AvailablePhoneNumbers *available_phone_numbers.Client
	Balance               func() *balance.Client
	Call                  func(string) *call.Client
	Calls                 *calls.Client
	Conference            func(string) *conference.Client
	Conferences           *conferences.Client
	Key                   func(string) *key.Client
	Keys                  *keys.Client
	Message               func(string) *message.Client
	Messages              *messages.Client
	Queue                 func(string) *queue.Client
	Queues                *queues.Client
	Recording             func(string) *recording.Client
	Recordings            *recordings.Client
	Tokens                *tokens.Client
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

		Address: func(addressSid string) *address.Client {
			return address.New(client, address.ClientProperties{
				AccountSid: properties.Sid,
				Sid:        addressSid,
			})
		},
		Addresses: addresses.New(client, addresses.ClientProperties{
			AccountSid: properties.Sid,
		}),
		Application: func(applicationSid string) *application.Client {
			return application.New(client, application.ClientProperties{
				AccountSid: properties.Sid,
				Sid:        applicationSid,
			})
		},
		Applications: applications.New(client, applications.ClientProperties{
			AccountSid: properties.Sid,
		}),
		AvailablePhoneNumber: func(countryCode string) *available_phone_number.Client {
			return available_phone_number.New(client, available_phone_number.ClientProperties{
				AccountSid:  properties.Sid,
				CountryCode: countryCode,
			})
		},
		AvailablePhoneNumbers: available_phone_numbers.New(client, available_phone_numbers.ClientProperties{
			AccountSid: properties.Sid,
		}),
		Balance: func() *balance.Client {
			return balance.New(client, balance.ClientProperties{
				AccountSid: properties.Sid,
			})
		},
		Call: func(callSid string) *call.Client {
			return call.New(client, call.ClientProperties{
				AccountSid: properties.Sid,
				Sid:        callSid,
			})
		},
		Calls: calls.New(client, calls.ClientProperties{
			AccountSid: properties.Sid,
		}),
		Conference: func(conferenceSid string) *conference.Client {
			return conference.New(client, conference.ClientProperties{
				AccountSid: properties.Sid,
				Sid:        conferenceSid,
			})
		},
		Conferences: conferences.New(client, conferences.ClientProperties{
			AccountSid: properties.Sid,
		}),
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
		Recording: func(recordingSid string) *recording.Client {
			return recording.New(client, recording.ClientProperties{
				AccountSid: properties.Sid,
				Sid:        recordingSid,
			})
		},
		Recordings: recordings.New(client, recordings.ClientProperties{
			AccountSid: properties.Sid,
		}),
		Tokens: tokens.New(client, tokens.ClientProperties{
			AccountSid: properties.Sid,
		}),
	}
}
