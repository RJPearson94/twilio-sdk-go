// Package service contains auto-generated files. DO NOT MODIFY
package service

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/messaging/v1/service/alpha_sender"
	"github.com/RJPearson94/twilio-sdk-go/service/messaging/v1/service/alpha_senders"
	"github.com/RJPearson94/twilio-sdk-go/service/messaging/v1/service/phone_number"
	"github.com/RJPearson94/twilio-sdk-go/service/messaging/v1/service/phone_numbers"
	"github.com/RJPearson94/twilio-sdk-go/service/messaging/v1/service/short_code"
	"github.com/RJPearson94/twilio-sdk-go/service/messaging/v1/service/short_codes"
)

// Client for managing a specific service resource
// See https://www.twilio.com/docs/sms/services/api for more details
type Client struct {
	client *client.Client

	sid string

	AlphaSender  func(string) *alpha_sender.Client
	AlphaSenders *alpha_senders.Client
	PhoneNumber  func(string) *phone_number.Client
	PhoneNumbers *phone_numbers.Client
	ShortCode    func(string) *short_code.Client
	ShortCodes   *short_codes.Client
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

		AlphaSender: func(alphaSenderSid string) *alpha_sender.Client {
			return alpha_sender.New(client, alpha_sender.ClientProperties{
				ServiceSid: properties.Sid,
				Sid:        alphaSenderSid,
			})
		},
		AlphaSenders: alpha_senders.New(client, alpha_senders.ClientProperties{
			ServiceSid: properties.Sid,
		}),
		PhoneNumber: func(phoneNumberSid string) *phone_number.Client {
			return phone_number.New(client, phone_number.ClientProperties{
				ServiceSid: properties.Sid,
				Sid:        phoneNumberSid,
			})
		},
		PhoneNumbers: phone_numbers.New(client, phone_numbers.ClientProperties{
			ServiceSid: properties.Sid,
		}),
		ShortCode: func(shortCodeSid string) *short_code.Client {
			return short_code.New(client, short_code.ClientProperties{
				ServiceSid: properties.Sid,
				Sid:        shortCodeSid,
			})
		},
		ShortCodes: short_codes.New(client, short_codes.ClientProperties{
			ServiceSid: properties.Sid,
		}),
	}
}
