// Package credential_list_mapping contains auto-generated files. DO NOT MODIFY
package credential_list_mapping

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// FetchCredentialListMappingResponse defines the response fields for retrieving a credential list mapping
type FetchCredentialListMappingResponse struct {
	AccountSid   string             `json:"account_sid"`
	DateCreated  utils.RFC2822Time  `json:"date_created"`
	DateUpdated  *utils.RFC2822Time `json:"date_updated,omitempty"`
	FriendlyName string             `json:"friendly_name"`
	Sid          string             `json:"sid"`
}

// Fetch retrieves a credential list mapping resource
// See https://www.twilio.com/docs/voice/sip/api/sip-domain-registration-credentiallistmapping-resource#fetch-a-sip-domain-registration-credentiallistmapping-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchCredentialListMappingResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a credential list mapping resource
// See https://www.twilio.com/docs/voice/sip/api/sip-domain-registration-credentiallistmapping-resource#fetch-a-sip-domain-registration-credentiallistmapping-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchCredentialListMappingResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Accounts/{accountSid}/SIP/Domains/{domainSid}/Auth/Registrations/CredentialListMappings/{sid}.json",
		PathParams: map[string]string{
			"accountSid": c.accountSid,
			"domainSid":  c.domainSid,
			"sid":        c.sid,
		},
	}

	response := &FetchCredentialListMappingResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
