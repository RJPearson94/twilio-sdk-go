// Package ip_access_control_lists contains auto-generated files. DO NOT MODIFY
package ip_access_control_lists

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreateIpAccessControlListInput defines the input fields for creating a new IP access control list resource
type CreateIpAccessControlListInput struct {
	IpAccessControlListSid string `validate:"required" form:"IpAccessControlListSid"`
}

// CreateIpAccessControlListResponse defines the response fields for the created IP access control list resource
type CreateIpAccessControlListResponse struct {
	AccountSid   string     `json:"account_sid"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	FriendlyName string     `json:"friendly_name"`
	Sid          string     `json:"sid"`
	TrunkSid     string     `json:"trunk_sid"`
	URL          string     `json:"url"`
}

// Create associates an IP access control list resource with the SIP trunk
// See https://www.twilio.com/docs/sip-trunking/api/ipaccesscontrollist-resource#create-an-ipaccesscontrollist-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateIpAccessControlListInput) (*CreateIpAccessControlListResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext associates an IP access control list resource with the SIP trunk
// See https://www.twilio.com/docs/sip-trunking/api/ipaccesscontrollist-resource#create-an-ipaccesscontrollist-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateIpAccessControlListInput) (*CreateIpAccessControlListResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Trunks/{trunkSid}/IpAccessControlLists",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"trunkSid": c.trunkSid,
		},
	}

	if input == nil {
		input = &CreateIpAccessControlListInput{}
	}

	response := &CreateIpAccessControlListResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
