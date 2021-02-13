// Package ip_access_control_list contains auto-generated files. DO NOT MODIFY
package ip_access_control_list

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/sip/ip_access_control_list/ip_address"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/sip/ip_access_control_list/ip_addresses"
)

// Client for managing a specific IP Access Control List resource
type Client struct {
	client *client.Client

	accountSid string
	sid        string

	IpAddress   func(string) *ip_address.Client
	IpAddresses *ip_addresses.Client
}

// ClientProperties are the properties required to manage the ip access control list resources
type ClientProperties struct {
	AccountSid string
	Sid        string
}

// New creates a new instance of the ip access control list client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		accountSid: properties.AccountSid,
		sid:        properties.Sid,

		IpAddress: func(ipAddressSid string) *ip_address.Client {
			return ip_address.New(client, ip_address.ClientProperties{
				AccountSid:             properties.AccountSid,
				IpAccessControlListSid: properties.Sid,
				Sid:                    ipAddressSid,
			})
		},
		IpAddresses: ip_addresses.New(client, ip_addresses.ClientProperties{
			AccountSid:             properties.AccountSid,
			IpAccessControlListSid: properties.Sid,
		}),
	}
}
