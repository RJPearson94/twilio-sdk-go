// Package ip_address contains auto-generated files. DO NOT MODIFY
package ip_address

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// UpdateIpAddressInput defines input fields for updating a IP address
type UpdateIpAddressInput struct {
	CidrPrefixLength *int    `form:"CidrPrefixLength,omitempty"`
	FriendlyName     *string `form:"FriendlyName,omitempty"`
	IpAddress        *string `form:"IpAddress,omitempty"`
}

// UpdateIpAddressResponse defines the response fields for the updated IP address
type UpdateIpAddressResponse struct {
	AccountSid             string             `json:"account_sid"`
	CidrPrefixLength       int                `json:"cidr_prefix_length"`
	DateCreated            utils.RFC2822Time  `json:"date_created"`
	DateUpdated            *utils.RFC2822Time `json:"date_updated,omitempty"`
	FriendlyName           string             `json:"friendly_name"`
	IpAccessControlListSid string             `json:"ip_access_control_list_sid"`
	IpAddress              string             `json:"ip_address"`
	Sid                    string             `json:"sid"`
}

// Update modifies a IP address resource
// See https://www.twilio.com/docs/voice/sip/api/sip-ipaddress-resource#update-a-sip-ipaddress-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateIpAddressInput) (*UpdateIpAddressResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a IP address resource
// See https://www.twilio.com/docs/voice/sip/api/sip-ipaddress-resource#update-a-sip-ipaddress-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateIpAddressInput) (*UpdateIpAddressResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Accounts/{accountSid}/SIP/IpAccessControlLists/{ipAccessControlListSid}/IpAddresses/{sid}.json",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"accountSid":             c.accountSid,
			"ipAccessControlListSid": c.ipAccessControlListSid,
			"sid":                    c.sid,
		},
	}

	if input == nil {
		input = &UpdateIpAddressInput{}
	}

	response := &UpdateIpAddressResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
