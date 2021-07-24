// Package aws_credentials contains auto-generated files. DO NOT MODIFY
package aws_credentials

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreateAWSCredentialInput defines the input fields for creating a new aws credential resource
type CreateAWSCredentialInput struct {
	AccountSid   *string `form:"AccountSid,omitempty"`
	Credentials  string  `validate:"required" form:"Credentials"`
	FriendlyName *string `form:"FriendlyName,omitempty"`
}

// CreateAWSCredentialResponse defines the response fields for the created aws credential
type CreateAWSCredentialResponse struct {
	AccountSid   string     `json:"account_sid"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	FriendlyName *string    `json:"friendly_name,omitempty"`
	Sid          string     `json:"sid"`
	URL          string     `json:"url"`
}

// Create creates a aws credential resource
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateAWSCredentialInput) (*CreateAWSCredentialResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a aws credential resource
func (c Client) CreateWithContext(context context.Context, input *CreateAWSCredentialInput) (*CreateAWSCredentialResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Credentials/AWS",
		ContentType: client.URLEncoded,
	}

	if input == nil {
		input = &CreateAWSCredentialInput{}
	}

	response := &CreateAWSCredentialResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
