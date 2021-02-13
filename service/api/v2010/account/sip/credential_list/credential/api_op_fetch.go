// Package credential contains auto-generated files. DO NOT MODIFY
package credential

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// FetchCredentialResponse defines the response fields for retrieving a SIP credential
type FetchCredentialResponse struct {
	AccountSid        string             `json:"account_sid"`
	CredentialListSid string             `json:"credential_list_sid"`
	DateCreated       utils.RFC2822Time  `json:"date_created"`
	DateUpdated       *utils.RFC2822Time `json:"date_updated,omitempty"`
	Sid               string             `json:"sid"`
	Username          string             `json:"username"`
}

// Fetch retrieves a SIP credential resource
// See https://www.twilio.com/docs/voice/sip/api/sip-credential-resource#fetch-a-sip-credential-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchCredentialResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a SIP credential resource
// See https://www.twilio.com/docs/voice/sip/api/sip-credential-resource#fetch-a-sip-credential-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchCredentialResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Accounts/{accountSid}/SIP/CredentialLists/{credentialListSid}/Credentials/{sid}.json",
		PathParams: map[string]string{
			"accountSid":        c.accountSid,
			"credentialListSid": c.credentialListSid,
			"sid":               c.sid,
		},
	}

	response := &FetchCredentialResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
