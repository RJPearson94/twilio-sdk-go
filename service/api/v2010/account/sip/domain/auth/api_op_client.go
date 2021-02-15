// Package auth contains auto-generated files. DO NOT MODIFY
package auth

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/sip/domain/auth/calls"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/sip/domain/auth/registrations"
)

// Client for managing SIP domain auth resources
type Client struct {
	client *client.Client

	accountSid string
	domainSid  string

	Calls         *calls.Client
	Registrations *registrations.Client
}

// ClientProperties are the properties required to manage the auth resources
type ClientProperties struct {
	AccountSid string
	DomainSid  string
}

// New creates a new instance of the auth client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		accountSid: properties.AccountSid,
		domainSid:  properties.DomainSid,

		Calls: calls.New(client, calls.ClientProperties{
			AccountSid: properties.AccountSid,
			DomainSid:  properties.DomainSid,
		}),
		Registrations: registrations.New(client, registrations.ClientProperties{
			AccountSid: properties.AccountSid,
			DomainSid:  properties.DomainSid,
		}),
	}
}
