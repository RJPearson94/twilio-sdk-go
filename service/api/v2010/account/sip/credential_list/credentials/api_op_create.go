// Package credentials contains auto-generated files. DO NOT MODIFY
package credentials

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// CreateCredentialInput defines input fields for creating a new SIP credential
type CreateCredentialInput struct {
	Password string `validate:"required" form:"Password"`
	Username string `validate:"required" form:"Username"`
}

// CreateCredentialResponse defines the response fields for creating a new SIP credential
type CreateCredentialResponse struct {
	AccountSid        string             `json:"account_sid"`
	CredentialListSid string             `json:"credential_list_sid"`
	DateCreated       utils.RFC2822Time  `json:"date_created"`
	DateUpdated       *utils.RFC2822Time `json:"date_updated,omitempty"`
	Sid               string             `json:"sid"`
	Username          string             `json:"username"`
}

// Create creates a SIP credential resource
// See https://www.twilio.com/docs/voice/sip/api/sip-credential-resource#create-a-sip-credential-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateCredentialInput) (*CreateCredentialResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a SIP credential resource
// See https://www.twilio.com/docs/voice/sip/api/sip-credential-resource#create-a-sip-credential-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateCredentialInput) (*CreateCredentialResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Accounts/{accountSid}/SIP/CredentialLists/{credentialListSid}/Credentials.json",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"accountSid":        c.accountSid,
			"credentialListSid": c.credentialListSid,
		},
	}

	if input == nil {
		input = &CreateCredentialInput{}
	}

	response := &CreateCredentialResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
