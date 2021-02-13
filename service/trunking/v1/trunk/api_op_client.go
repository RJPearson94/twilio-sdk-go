// Package trunk contains auto-generated files. DO NOT MODIFY
package trunk

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/trunking/v1/trunk/credential_list"
	"github.com/RJPearson94/twilio-sdk-go/service/trunking/v1/trunk/credential_lists"
	"github.com/RJPearson94/twilio-sdk-go/service/trunking/v1/trunk/ip_access_control_list"
	"github.com/RJPearson94/twilio-sdk-go/service/trunking/v1/trunk/ip_access_control_lists"
	"github.com/RJPearson94/twilio-sdk-go/service/trunking/v1/trunk/origination_url"
	"github.com/RJPearson94/twilio-sdk-go/service/trunking/v1/trunk/origination_urls"
	"github.com/RJPearson94/twilio-sdk-go/service/trunking/v1/trunk/phone_number"
	"github.com/RJPearson94/twilio-sdk-go/service/trunking/v1/trunk/phone_numbers"
	"github.com/RJPearson94/twilio-sdk-go/service/trunking/v1/trunk/recording"
)

// Client for managing a specific trunk resource
// See https://www.twilio.com/docs/sip-trunking/api/trunk-resource for more details
type Client struct {
	client *client.Client

	sid string

	CredentialList       func(string) *credential_list.Client
	CredentialLists      *credential_lists.Client
	IpAccessControlList  func(string) *ip_access_control_list.Client
	IpAccessControlLists *ip_access_control_lists.Client
	OriginationURL       func(string) *origination_url.Client
	OriginationURLs      *origination_urls.Client
	PhoneNumber          func(string) *phone_number.Client
	PhoneNumbers         *phone_numbers.Client
	Recording            func() *recording.Client
}

// ClientProperties are the properties required to manage the trunk resources
type ClientProperties struct {
	Sid string
}

// New creates a new instance of the trunk client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		sid: properties.Sid,

		CredentialList: func(credentialListSid string) *credential_list.Client {
			return credential_list.New(client, credential_list.ClientProperties{
				Sid:      credentialListSid,
				TrunkSid: properties.Sid,
			})
		},
		CredentialLists: credential_lists.New(client, credential_lists.ClientProperties{
			TrunkSid: properties.Sid,
		}),
		IpAccessControlList: func(ipAccessControlListSid string) *ip_access_control_list.Client {
			return ip_access_control_list.New(client, ip_access_control_list.ClientProperties{
				Sid:      ipAccessControlListSid,
				TrunkSid: properties.Sid,
			})
		},
		IpAccessControlLists: ip_access_control_lists.New(client, ip_access_control_lists.ClientProperties{
			TrunkSid: properties.Sid,
		}),
		OriginationURL: func(originationURLSid string) *origination_url.Client {
			return origination_url.New(client, origination_url.ClientProperties{
				Sid:      originationURLSid,
				TrunkSid: properties.Sid,
			})
		},
		OriginationURLs: origination_urls.New(client, origination_urls.ClientProperties{
			TrunkSid: properties.Sid,
		}),
		PhoneNumber: func(phoneNumberSid string) *phone_number.Client {
			return phone_number.New(client, phone_number.ClientProperties{
				Sid:      phoneNumberSid,
				TrunkSid: properties.Sid,
			})
		},
		PhoneNumbers: phone_numbers.New(client, phone_numbers.ClientProperties{
			TrunkSid: properties.Sid,
		}),
		Recording: func() *recording.Client {
			return recording.New(client, recording.ClientProperties{
				TrunkSid: properties.Sid,
			})
		},
	}
}
