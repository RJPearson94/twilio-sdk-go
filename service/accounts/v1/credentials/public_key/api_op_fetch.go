// Package public_key contains auto-generated files. DO NOT MODIFY
package public_key

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchPublicKeyResponse defines the response fields for the retrieved public key resource
type FetchPublicKeyResponse struct {
	AccountSid   string     `json:"account_sid"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	FriendlyName *string    `json:"friendly_name,omitempty"`
	Sid          string     `json:"sid"`
	URL          string     `json:"url"`
}

// Fetch retrieves a public key resource
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchPublicKeyResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a public key resource
func (c Client) FetchWithContext(context context.Context) (*FetchPublicKeyResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Credentials/PublicKeys/{sid}",
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	response := &FetchPublicKeyResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
