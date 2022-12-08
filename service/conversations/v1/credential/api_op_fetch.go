// Package credential contains auto-generated files. DO NOT MODIFY
package credential

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchCredentialResponse defines the response fields for the retrieved credential
type FetchCredentialResponse struct {
	AccountSid   string     `json:"account_sid"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	FriendlyName *string    `json:"friendly_name,omitempty"`
	Sandbox      string     `json:"sandbox,omitempty"`
	Sid          string     `json:"sid"`
	Type         string     `json:"type"`
	URL          string     `json:"url"`
}

// Fetch retrieves a credential resource
// See https://www.twilio.com/docs/conversations/api/credential-resource#fetch-a-credential-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchCredentialResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a credential resource
// See https://www.twilio.com/docs/conversations/api/credential-resource#fetch-a-credential-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchCredentialResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Credentials/{sid}",
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	response := &FetchCredentialResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
