// Package credential contains auto-generated files. DO NOT MODIFY
package credential

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// UpdateCredentialInput defines input fields for updating a credential resource
type UpdateCredentialInput struct {
	ApiKey       *string `form:"ApiKey,omitempty"`
	Certificate  *string `form:"Certificate,omitempty"`
	FriendlyName *string `form:"FriendlyName,omitempty"`
	PrivateKey   *string `form:"PrivateKey,omitempty"`
	Sandbox      *bool   `form:"Sandbox,omitempty"`
	Secret       *string `form:"Secret,omitempty"`
	Type         *string `form:"Type,omitempty"`
}

// UpdateCredentialResponse defines the response fields for the updated credential
type UpdateCredentialResponse struct {
	AccountSid   string     `json:"account_sid"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	FriendlyName *string    `json:"friendly_name,omitempty"`
	Sandbox      *string    `json:"sandbox,omitempty"`
	Sid          string     `json:"sid"`
	Type         string     `json:"type"`
	URL          string     `json:"url"`
}

// Update modifies a credentials resource
// See https://www.twilio.com/docs/conversations/api/credential-resource#update-a-credential-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateCredentialInput) (*UpdateCredentialResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a credentials resource
// See https://www.twilio.com/docs/conversations/api/credential-resource#update-a-credential-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateCredentialInput) (*UpdateCredentialResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Credentials/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	if input == nil {
		input = &UpdateCredentialInput{}
	}

	response := &UpdateCredentialResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
