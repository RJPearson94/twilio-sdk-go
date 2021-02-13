// Package ip_address contains auto-generated files. DO NOT MODIFY
package ip_address

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// FetchIpAddressResponse defines the response fields for retrieving a IP address
type FetchIpAddressResponse struct {
	AccountSid             string             `json:"account_sid"`
	CidrPrefixLength       int                `json:"cidr_prefix_length"`
	DateCreated            utils.RFC2822Time  `json:"date_created"`
	DateUpdated            *utils.RFC2822Time `json:"date_updated,omitempty"`
	FriendlyName           string             `json:"friendly_name"`
	IpAccessControlListSid string             `json:"ip_access_control_list_sid"`
	IpAddress              string             `json:"ip_address"`
	Sid                    string             `json:"sid"`
}

// Fetch retrieves a IP address resource
// See https://www.twilio.com/docs/voice/sip/api/sip-ipaddress-resource#fetch-a-sip-ipaddress-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchIpAddressResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a IP address resource
// See https://www.twilio.com/docs/voice/sip/api/sip-ipaddress-resource#fetch-a-sip-ipaddress-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchIpAddressResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Accounts/{accountSid}/SIP/IpAccessControlLists/{ipAccessControlListSid}/IpAddresses/{sid}.json",
		PathParams: map[string]string{
			"accountSid":             c.accountSid,
			"ipAccessControlListSid": c.ipAccessControlListSid,
			"sid":                    c.sid,
		},
	}

	response := &FetchIpAddressResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
