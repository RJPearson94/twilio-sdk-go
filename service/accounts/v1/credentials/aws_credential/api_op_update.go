// Package aws_credential contains auto-generated files. DO NOT MODIFY
package aws_credential

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// UpdateAWSCredentialInput defines input fields for updating a aws credential resource
type UpdateAWSCredentialInput struct {
	FriendlyName *string `form:"FriendlyName,omitempty"`
}

// UpdateAWSCredentialResponse defines the response fields for the updated aws credential resource
type UpdateAWSCredentialResponse struct {
	AccountSid   string     `json:"account_sid"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	FriendlyName *string    `json:"friendly_name,omitempty"`
	Sid          string     `json:"sid"`
	URL          string     `json:"url"`
}

// Update modifies a aws credential resource
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateAWSCredentialInput) (*UpdateAWSCredentialResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a aws credential resource
func (c Client) UpdateWithContext(context context.Context, input *UpdateAWSCredentialInput) (*UpdateAWSCredentialResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Credentials/AWS/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	if input == nil {
		input = &UpdateAWSCredentialInput{}
	}

	response := &UpdateAWSCredentialResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
