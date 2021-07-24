// Package public_keys contains auto-generated files. DO NOT MODIFY
package public_keys

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreatePublicKeyInput defines the input fields for creating a new public key resource
type CreatePublicKeyInput struct {
	AccountSid   *string `form:"AccountSid,omitempty"`
	FriendlyName *string `form:"FriendlyName,omitempty"`
	PublicKey    string  `validate:"required" form:"PublicKey"`
}

// CreatePublicKeyResponse defines the response fields for the created public key
type CreatePublicKeyResponse struct {
	AccountSid   string     `json:"account_sid"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	FriendlyName *string    `json:"friendly_name,omitempty"`
	Sid          string     `json:"sid"`
	URL          string     `json:"url"`
}

// Create creates a public key resource
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreatePublicKeyInput) (*CreatePublicKeyResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a public key resource
func (c Client) CreateWithContext(context context.Context, input *CreatePublicKeyInput) (*CreatePublicKeyResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Credentials/PublicKeys",
		ContentType: client.URLEncoded,
	}

	if input == nil {
		input = &CreatePublicKeyInput{}
	}

	response := &CreatePublicKeyResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
