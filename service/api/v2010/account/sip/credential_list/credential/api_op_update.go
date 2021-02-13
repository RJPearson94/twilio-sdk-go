// Package credential contains auto-generated files. DO NOT MODIFY
package credential

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// UpdateCredentialInput defines input fields for updating a SIP credential
type UpdateCredentialInput struct {
	Password string `validate:"required" form:"Password"`
}

// UpdateCredentialResponse defines the response fields for the updated SIP credential
type UpdateCredentialResponse struct {
	AccountSid        string             `json:"account_sid"`
	CredentialListSid string             `json:"credential_list_sid"`
	DateCreated       utils.RFC2822Time  `json:"date_created"`
	DateUpdated       *utils.RFC2822Time `json:"date_updated,omitempty"`
	Sid               string             `json:"sid"`
	Username          string             `json:"username"`
}

// Update modifies a SIP credential resource
// See https://www.twilio.com/docs/voice/sip/api/sip-credential-resource#update-a-sip-credential-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateCredentialInput) (*UpdateCredentialResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a SIP credential resource
// See https://www.twilio.com/docs/voice/sip/api/sip-credential-resource#update-a-sip-credential-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateCredentialInput) (*UpdateCredentialResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Accounts/{accountSid}/SIP/CredentialLists/{credentialListSid}/Credentials/{sid}.json",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"accountSid":        c.accountSid,
			"credentialListSid": c.credentialListSid,
			"sid":               c.sid,
		},
	}

	if input == nil {
		input = &UpdateCredentialInput{}
	}

	response := &UpdateCredentialResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
