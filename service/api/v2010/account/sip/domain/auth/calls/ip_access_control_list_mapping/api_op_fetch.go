// Package ip_access_control_list_mapping contains auto-generated files. DO NOT MODIFY
package ip_access_control_list_mapping

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// FetchIpAccessControlListMappingResponse defines the response fields for retrieving a credential list mapping
type FetchIpAccessControlListMappingResponse struct {
	AccountSid   string             `json:"account_sid"`
	DateCreated  utils.RFC2822Time  `json:"date_created"`
	DateUpdated  *utils.RFC2822Time `json:"date_updated,omitempty"`
	FriendlyName string             `json:"friendly_name"`
	Sid          string             `json:"sid"`
}

// Fetch retrieves a IP control access list mapping resource
// See https://www.twilio.com/docs/voice/sip/api/sip-ipaccesscontrollistmapping-resource#fetch-a-sip-ipaccesscontrollistmapping-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchIpAccessControlListMappingResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a IP control access list mapping resource
// See https://www.twilio.com/docs/voice/sip/api/sip-ipaccesscontrollistmapping-resource#fetch-a-sip-ipaccesscontrollistmapping-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchIpAccessControlListMappingResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Accounts/{accountSid}/SIP/Domains/{domainSid}/Auth/Calls/IpAccessControlListMappings/{sid}.json",
		PathParams: map[string]string{
			"accountSid": c.accountSid,
			"domainSid":  c.domainSid,
			"sid":        c.sid,
		},
	}

	response := &FetchIpAccessControlListMappingResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
