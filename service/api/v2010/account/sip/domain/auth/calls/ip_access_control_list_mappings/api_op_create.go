// Package ip_access_control_list_mappings contains auto-generated files. DO NOT MODIFY
package ip_access_control_list_mappings

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// CreateIpAccessControlListMappingInput defines input fields for creating a new IP access control list mapping
type CreateIpAccessControlListMappingInput struct {
	IpAccessControlListSid string `validate:"required" form:"IpAccessControlListSid"`
}

// IpAccessControlListMappingResponse defines the response fields for creating a new IP access control list mapping
type IpAccessControlListMappingResponse struct {
	AccountSid   string             `json:"account_sid"`
	DateCreated  utils.RFC2822Time  `json:"date_created"`
	DateUpdated  *utils.RFC2822Time `json:"date_updated,omitempty"`
	FriendlyName string             `json:"friendly_name"`
	Sid          string             `json:"sid"`
}

// Create creates a IP access control list mapping resource
// See https://www.twilio.com/docs/voice/sip/api/sip-ipaccesscontrollistmapping-resource#create-a-sip-ipaccesscontrollistmapping-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateIpAccessControlListMappingInput) (*IpAccessControlListMappingResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a IP access control list mapping resource
// See https://www.twilio.com/docs/voice/sip/api/sip-ipaccesscontrollistmapping-resource#create-a-sip-ipaccesscontrollistmapping-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateIpAccessControlListMappingInput) (*IpAccessControlListMappingResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Accounts/{accountSid}/SIP/Domains/{domainSid}/Auth/Calls/IpAccessControlListMappings.json",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"accountSid": c.accountSid,
			"domainSid":  c.domainSid,
		},
	}

	if input == nil {
		input = &CreateIpAccessControlListMappingInput{}
	}

	response := &IpAccessControlListMappingResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
