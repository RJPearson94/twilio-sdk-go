// Package domain contains auto-generated files. DO NOT MODIFY
package domain

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/sip/domain/auth"
)

// Client for managing a specific SIP domain resource
type Client struct {
	client *client.Client

	accountSid string
	sid        string

	Auth *auth.Client
}

// ClientProperties are the properties required to manage the domain resources
type ClientProperties struct {
	AccountSid string
	Sid        string
}

// New creates a new instance of the domain client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		accountSid: properties.AccountSid,
		sid:        properties.Sid,

		Auth: auth.New(client, auth.ClientProperties{
			AccountSid: properties.AccountSid,
			DomainSid:  properties.Sid,
		}),
	}
}
