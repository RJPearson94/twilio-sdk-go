// Package aws_credential contains auto-generated files. DO NOT MODIFY
package aws_credential

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchAWSCredentialResponse defines the response fields for the retrieved aws credential resource
type FetchAWSCredentialResponse struct {
	AccountSid   string     `json:"account_sid"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	FriendlyName *string    `json:"friendly_name,omitempty"`
	Sid          string     `json:"sid"`
	URL          string     `json:"url"`
}

// Fetch retrieves a aws credential resource
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchAWSCredentialResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a aws credential resource
func (c Client) FetchWithContext(context context.Context) (*FetchAWSCredentialResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Credentials/AWS/{sid}",
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	response := &FetchAWSCredentialResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
