// Package ip_addresses contains auto-generated files. DO NOT MODIFY
package ip_addresses

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// CreateIpAddressInput defines input fields for creating a new IP address resource
type CreateIpAddressInput struct {
	CidrPrefixLength *int   `form:"CidrPrefixLength,omitempty"`
	FriendlyName     string `validate:"required" form:"FriendlyName"`
	IpAddress        string `validate:"required" form:"IpAddress"`
}

// CreateIpAddressResponse defines the response fields for creating a new IP address resource
type CreateIpAddressResponse struct {
	AccountSid             string             `json:"account_sid"`
	CidrPrefixLength       int                `json:"cidr_prefix_length"`
	DateCreated            utils.RFC2822Time  `json:"date_created"`
	DateUpdated            *utils.RFC2822Time `json:"date_updated,omitempty"`
	FriendlyName           string             `json:"friendly_name"`
	IpAccessControlListSid string             `json:"ip_access_control_list_sid"`
	IpAddress              string             `json:"ip_address"`
	Sid                    string             `json:"sid"`
}

// Create creates a IP address resource
// See https://www.twilio.com/docs/voice/sip/api/sip-ipaddress-resource#create-a-sip-ipaddress-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateIpAddressInput) (*CreateIpAddressResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a IP address resource
// See https://www.twilio.com/docs/voice/sip/api/sip-ipaddress-resource#create-a-sip-ipaddress-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateIpAddressInput) (*CreateIpAddressResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Accounts/{accountSid}/SIP/IpAccessControlLists/{ipAccessControlListSid}/IpAddresses.json",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"accountSid":             c.accountSid,
			"ipAccessControlListSid": c.ipAccessControlListSid,
		},
	}

	if input == nil {
		input = &CreateIpAddressInput{}
	}

	response := &CreateIpAddressResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
