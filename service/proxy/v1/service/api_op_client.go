// Package service contains auto-generated files. DO NOT MODIFY
package service

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/service/phone_number"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/service/phone_numbers"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/service/session"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/service/sessions"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/service/short_code"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/service/short_codes"
)

// Client for managing a specific service resource
// See https://www.twilio.com/docs/proxy/api/service for more details
type Client struct {
	client *client.Client

	sid string

	PhoneNumber  func(string) *phone_number.Client
	PhoneNumbers *phone_numbers.Client
	Session      func(string) *session.Client
	Sessions     *sessions.Client
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

		PhoneNumber: func(phoneNumberSid string) *phone_number.Client {
			return phone_number.New(client, phone_number.ClientProperties{
				ServiceSid: properties.Sid,
				Sid:        phoneNumberSid,
			})
		},
		PhoneNumbers: phone_numbers.New(client, phone_numbers.ClientProperties{
			ServiceSid: properties.Sid,
		}),
		Session: func(sessionSid string) *session.Client {
			return session.New(client, session.ClientProperties{
				ServiceSid: properties.Sid,
				Sid:        sessionSid,
			})
		},
		Sessions: sessions.New(client, sessions.ClientProperties{
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
