// Package credential_list contains auto-generated files. DO NOT MODIFY
package credential_list

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// UpdateCredentialListInput defines input fields for updating a credential list
type UpdateCredentialListInput struct {
	FriendlyName string `validate:"required" form:"FriendlyName"`
}

// UpdateCredentialListResponse defines the response fields for the updated credential list
type UpdateCredentialListResponse struct {
	AccountSid   string             `json:"account_sid"`
	DateCreated  utils.RFC2822Time  `json:"date_created"`
	DateUpdated  *utils.RFC2822Time `json:"date_updated,omitempty"`
	FriendlyName string             `json:"friendly_name"`
	Sid          string             `json:"sid"`
}

// Update modifies a credential list resource
// See https://www.twilio.com/docs/voice/sip/api/sip-credentiallist-resource#update-a-sip-credentiallist-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateCredentialListInput) (*UpdateCredentialListResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a credential list resource
// See https://www.twilio.com/docs/voice/sip/api/sip-credentiallist-resource#update-a-sip-credentiallist-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateCredentialListInput) (*UpdateCredentialListResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Accounts/{accountSid}/SIP/CredentialLists/{sid}.json",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"accountSid": c.accountSid,
			"sid":        c.sid,
		},
	}

	if input == nil {
		input = &UpdateCredentialListInput{}
	}

	response := &UpdateCredentialListResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
