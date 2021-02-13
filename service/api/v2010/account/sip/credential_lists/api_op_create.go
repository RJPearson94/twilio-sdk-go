// Package credential_lists contains auto-generated files. DO NOT MODIFY
package credential_lists

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// CreateCredentialListInput defines input fields for creating a new credential list
type CreateCredentialListInput struct {
	FriendlyName string `validate:"required" form:"FriendlyName"`
}

// CreateCredentialListResponse defines the response fields for creating a new credential list
type CreateCredentialListResponse struct {
	AccountSid   string             `json:"account_sid"`
	DateCreated  utils.RFC2822Time  `json:"date_created"`
	DateUpdated  *utils.RFC2822Time `json:"date_updated,omitempty"`
	FriendlyName string             `json:"friendly_name"`
	Sid          string             `json:"sid"`
}

// Create creates a credential list resource
// See https://www.twilio.com/docs/voice/sip/api/sip-credentiallist-resource#create-a-sip-credentiallist-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateCredentialListInput) (*CreateCredentialListResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a credential list resource
// See https://www.twilio.com/docs/voice/sip/api/sip-credentiallist-resource#create-a-sip-credentiallist-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateCredentialListInput) (*CreateCredentialListResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Accounts/{accountSid}/SIP/CredentialLists.json",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"accountSid": c.accountSid,
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
