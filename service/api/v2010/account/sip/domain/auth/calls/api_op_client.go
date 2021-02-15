// Package calls contains auto-generated files. DO NOT MODIFY
package calls

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/sip/domain/auth/calls/credential_list_mapping"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/sip/domain/auth/calls/credential_list_mappings"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/sip/domain/auth/calls/ip_access_control_list_mapping"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/sip/domain/auth/calls/ip_access_control_list_mappings"
)

// Client for managing SIP domain auth call resources
type Client struct {
	client *client.Client

	accountSid string
	domainSid  string

	CredentialListMapping       func(string) *credential_list_mapping.Client
	CredentialListMappings      *credential_list_mappings.Client
	IpAccessControlListMapping  func(string) *ip_access_control_list_mapping.Client
	IpAccessControlListMappings *ip_access_control_list_mappings.Client
}

// ClientProperties are the properties required to manage the calls resources
type ClientProperties struct {
	AccountSid string
	DomainSid  string
}

// New creates a new instance of the calls client
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
		IpAccessControlListMapping: func(ipAccessControlListMappingSid string) *ip_access_control_list_mapping.Client {
			return ip_access_control_list_mapping.New(client, ip_access_control_list_mapping.ClientProperties{
				AccountSid: properties.AccountSid,
				DomainSid:  properties.DomainSid,
				Sid:        ipAccessControlListMappingSid,
			})
		},
		IpAccessControlListMappings: ip_access_control_list_mappings.New(client, ip_access_control_list_mappings.ClientProperties{
			AccountSid: properties.AccountSid,
			DomainSid:  properties.DomainSid,
		}),
	}
}
