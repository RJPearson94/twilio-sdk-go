// Package ip_access_control_lists contains auto-generated files. DO NOT MODIFY
package ip_access_control_lists

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// CreateIpAccessControlListInput defines input fields for creating a new IP access control list
type CreateIpAccessControlListInput struct {
	FriendlyName string `validate:"required" form:"FriendlyName"`
}

// IpAccessControlListResponse defines the response fields for creating a new IP access control list
type IpAccessControlListResponse struct {
	AccountSid   string             `json:"account_sid"`
	DateCreated  utils.RFC2822Time  `json:"date_created"`
	DateUpdated  *utils.RFC2822Time `json:"date_updated,omitempty"`
	FriendlyName string             `json:"friendly_name"`
	Sid          string             `json:"sid"`
}

// Create creates a IP access control list resource
// See https://www.twilio.com/docs/voice/sip/api/sip-ipaccesscontrollist-resource#create-a-sip-ipaccesscontrollist-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateIpAccessControlListInput) (*IpAccessControlListResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a IP access control list resource
// See https://www.twilio.com/docs/voice/sip/api/sip-ipaccesscontrollist-resource#create-a-sip-ipaccesscontrollist-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateIpAccessControlListInput) (*IpAccessControlListResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Accounts/{accountSid}/SIP/IpAccessControlLists.json",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"accountSid": c.accountSid,
		},
	}

	if input == nil {
		input = &CreateIpAccessControlListInput{}
	}

	response := &IpAccessControlListResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
