// Package credential_list contains auto-generated files. DO NOT MODIFY
package credential_list

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// FetchCredentialListResponse defines the response fields for retrieving a credential list
type FetchCredentialListResponse struct {
	AccountSid   string             `json:"account_sid"`
	DateCreated  utils.RFC2822Time  `json:"date_created"`
	DateUpdated  *utils.RFC2822Time `json:"date_updated,omitempty"`
	FriendlyName string             `json:"friendly_name"`
	Sid          string             `json:"sid"`
}

// Fetch retrieves a credential list resource
// See https://www.twilio.com/docs/voice/sip/api/sip-credentiallist-resource#fetch-a-sip-credentiallist-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchCredentialListResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a credential list resource
// See https://www.twilio.com/docs/voice/sip/api/sip-credentiallist-resource#fetch-a-sip-credentiallist-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchCredentialListResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Accounts/{accountSid}/SIP/CredentialLists/{sid}.json",
		PathParams: map[string]string{
			"accountSid": c.accountSid,
			"sid":        c.sid,
		},
	}

	response := &FetchCredentialListResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
