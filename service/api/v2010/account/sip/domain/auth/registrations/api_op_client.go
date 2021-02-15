// Package registrations contains auto-generated files. DO NOT MODIFY
package registrations

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/sip/domain/auth/registrations/credential_list_mapping"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/sip/domain/auth/registrations/credential_list_mappings"
)

// Client for managing SIP domain auth registrations resources
type Client struct {
	client *client.Client

	accountSid string
	domainSid  string

	CredentialListMapping  func(string) *credential_list_mapping.Client
	CredentialListMappings *credential_list_mappings.Client
}

// ClientProperties are the properties required to manage the registrations resources
type ClientProperties struct {
	AccountSid string
	DomainSid  string
}

// New creates a new instance of the registrations client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		accountSid: properties.AccountSid,
		domainSid:  properties.DomainSid,

		CredentialListMapping: func(credentialListMappingSid string) *credential_list_mapping.Client {
			return credential_list_mapping.New(client, credential_list_mapping.ClientProperties{
				AccountSid: properties.AccountSid,
				DomainSid:  properties.DomainSid,
				Sid:        credentialListMappingSid,
			})
		},
		CredentialListMappings: credential_list_mappings.New(client, credential_list_mappings.ClientProperties{
			AccountSid: properties.AccountSid,
			DomainSid:  properties.DomainSid,
		}),
	}
}
