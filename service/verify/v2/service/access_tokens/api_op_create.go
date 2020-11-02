// Package access_tokens contains auto-generated files. DO NOT MODIFY
package access_tokens

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreateAccessTokenInput defines the input fields for creating a new access token
type CreateAccessTokenInput struct {
	FactorType string `validate:"required" form:"FactorType"`
	Identity   string `validate:"required" form:"Identity"`
}

// CreateAccessTokenResponse defines the response fields for the created access token
type CreateAccessTokenResponse struct {
	Token string `json:"token"`
}

// Create creates an access token
// See https://www.twilio.com/docs/verify/api/access-token#create-an-accesstoken-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateAccessTokenInput) (*CreateAccessTokenResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates an access token
// See https://www.twilio.com/docs/verify/api/access-token#create-an-accesstoken-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateAccessTokenInput) (*CreateAccessTokenResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/AccessTokens",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
		},
	}

	if input == nil {
		input = &CreateAccessTokenInput{}
	}

	response := &CreateAccessTokenResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
