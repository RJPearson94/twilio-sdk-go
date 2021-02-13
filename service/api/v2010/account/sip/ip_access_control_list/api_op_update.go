// Package ip_access_control_list contains auto-generated files. DO NOT MODIFY
package ip_access_control_list

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// UpdateIpAccessControlListInput defines input fields for updating a IP control access list
type UpdateIpAccessControlListInput struct {
	FriendlyName string `validate:"required" form:"FriendlyName"`
}

// UpdateIpAccessControlListResponse defines the response fields for the updated IP control access list
type UpdateIpAccessControlListResponse struct {
	AccountSid   string             `json:"account_sid"`
	DateCreated  utils.RFC2822Time  `json:"date_created"`
	DateUpdated  *utils.RFC2822Time `json:"date_updated,omitempty"`
	FriendlyName string             `json:"friendly_name"`
	Sid          string             `json:"sid"`
}

// Update modifies a IP control access list resource
// See https://www.twilio.com/docs/voice/sip/api/sip-ipaccesscontrollist-resource#update-a-sip-ipaccesscontrollist-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateIpAccessControlListInput) (*UpdateIpAccessControlListResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a IP control access list resource
// See https://www.twilio.com/docs/voice/sip/api/sip-ipaccesscontrollist-resource#update-a-sip-ipaccesscontrollist-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateIpAccessControlListInput) (*UpdateIpAccessControlListResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Accounts/{accountSid}/SIP/IpAccessControlLists/{sid}.json",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"accountSid": c.accountSid,
			"sid":        c.sid,
		},
	}

	if input == nil {
		input = &UpdateIpAccessControlListInput{}
	}

	response := &UpdateIpAccessControlListResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
