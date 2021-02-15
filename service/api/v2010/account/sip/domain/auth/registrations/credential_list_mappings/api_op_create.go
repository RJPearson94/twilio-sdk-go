// Package credential_list_mappings contains auto-generated files. DO NOT MODIFY
package credential_list_mappings

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// CreateCredentialListMappingInput defines input fields for creating a new credential list mapping
type CreateCredentialListMappingInput struct {
	CredentialListSid string `validate:"required" form:"CredentialListSid"`
}

// CreateCredentialListMappingResponse defines the response fields for creating a new credential list mapping
type CreateCredentialListMappingResponse struct {
	AccountSid   string             `json:"account_sid"`
	DateCreated  utils.RFC2822Time  `json:"date_created"`
	DateUpdated  *utils.RFC2822Time `json:"date_updated,omitempty"`
	FriendlyName string             `json:"friendly_name"`
	Sid          string             `json:"sid"`
}

// Create creates a credential list mapping resource
// See https://www.twilio.com/docs/voice/sip/api/sip-domain-registration-credentiallistmapping-resource#create-a-sip-domain-registration-credentiallistmapping-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateCredentialListMappingInput) (*CreateCredentialListMappingResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a credential list mapping resource
// See https://www.twilio.com/docs/voice/sip/api/sip-domain-registration-credentiallistmapping-resource#create-a-sip-domain-registration-credentiallistmapping-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateCredentialListMappingInput) (*CreateCredentialListMappingResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Accounts/{accountSid}/SIP/Domains/{domainSid}/Auth/Registrations/CredentialListMappings.json",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"accountSid": c.accountSid,
			"domainSid":  c.domainSid,
		},
	}

	if input == nil {
		input = &CreateCredentialListMappingInput{}
	}

	response := &CreateCredentialListMappingResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
