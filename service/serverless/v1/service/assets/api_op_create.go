// Package assets contains auto-generated files. DO NOT MODIFY
package assets

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreateAssetInput defines the input fields for creating a new asset resource
type CreateAssetInput struct {
	FriendlyName string `validate:"required" form:"FriendlyName"`
}

// CreateAssetResponse defines the response fields for the created asset
type CreateAssetResponse struct {
	AccountSid   string     `json:"account_sid"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	FriendlyName string     `json:"friendly_name"`
	ServiceSid   string     `json:"service_sid"`
	Sid          string     `json:"sid"`
	URL          string     `json:"url"`
}

// Create creates a new asset
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/asset#create-an-asset-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateAssetInput) (*CreateAssetResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new asset
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/asset#create-an-asset-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateAssetInput) (*CreateAssetResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Assets",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
		},
	}

	if input == nil {
		input = &CreateAssetInput{}
	}

	response := &CreateAssetResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
