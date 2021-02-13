// Package credential_lists contains auto-generated files. DO NOT MODIFY
package credential_lists

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreateCredentialListInput defines the input fields for creating a new credential list resource
type CreateCredentialListInput struct {
	CredentialListSid string `validate:"required" form:"CredentialListSid"`
}

// CreateCredentialListResponse defines the response fields for the created credential list resource
type CreateCredentialListResponse struct {
	AccountSid   string     `json:"account_sid"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	FriendlyName string     `json:"friendly_name"`
	Sid          string     `json:"sid"`
	TrunkSid     string     `json:"trunk_sid"`
	URL          string     `json:"url"`
}

// Create associates a credential list resource with the SIP trunk
// See https://www.twilio.com/docs/sip-trunking/api/credentiallist-resource#create-a-credentiallist-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateCredentialListInput) (*CreateCredentialListResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext associates a credential list resource with the SIP trunk
// See https://www.twilio.com/docs/sip-trunking/api/credentiallist-resource#create-a-credentiallist-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateCredentialListInput) (*CreateCredentialListResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Trunks/{trunkSid}/CredentialLists",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"trunkSid": c.trunkSid,
		},
	}

	if input == nil {
		input = &CreateCredentialListInput{}
	}

	response := &CreateCredentialListResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
