// Package public_key contains auto-generated files. DO NOT MODIFY
package public_key

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// UpdatePublicKeyInput defines input fields for updating a public key resource
type UpdatePublicKeyInput struct {
	FriendlyName *string `form:"FriendlyName,omitempty"`
}

// UpdatePublicKeyResponse defines the response fields for the updated public key resource
type UpdatePublicKeyResponse struct {
	AccountSid   string     `json:"account_sid"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	FriendlyName *string    `json:"friendly_name,omitempty"`
	Sid          string     `json:"sid"`
	URL          string     `json:"url"`
}

// Update modifies a public key resource
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdatePublicKeyInput) (*UpdatePublicKeyResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a public key resource
func (c Client) UpdateWithContext(context context.Context, input *UpdatePublicKeyInput) (*UpdatePublicKeyResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Credentials/PublicKeys/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	if input == nil {
		input = &UpdatePublicKeyInput{}
	}

	response := &UpdatePublicKeyResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
