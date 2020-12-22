// Package versions contains auto-generated files. DO NOT MODIFY
package versions

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

type CreateContentDetails struct {
	Body        io.ReadSeeker `validate:"required" mapstructure:"Body"`
	ContentType string        `validate:"required" mapstructure:"ContentType"`
	FileName    string        `validate:"required" mapstructure:"FileName"`
}

// CreateVersionInput defines the input fields for creating a new asset version resource
type CreateVersionInput struct {
	Content    CreateContentDetails `validate:"required" mapstructure:"Content"`
	Path       string               `validate:"required" mapstructure:"Path"`
	Visibility string               `validate:"required" mapstructure:"Visibility"`
}

// CreateVersionResponse defines the response fields for the created asset version
type CreateVersionResponse struct {
	AccountSid  string    `json:"account_sid"`
	AssetSid    string    `json:"asset_sid"`
	DateCreated time.Time `json:"date_created"`
	Path        string    `json:"path"`
	ServiceSid  string    `json:"service_sid"`
	Sid         string    `json:"sid"`
	URL         string    `json:"url"`
	Visibility  string    `json:"visibility"`
}

// Create creates a new asset version
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/asset#create-an-asset-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateVersionInput) (*CreateVersionResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new asset version
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/asset#create-an-asset-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateVersionInput) (*CreateVersionResponse, error) {
	op := client.Operation{
		OverrideBaseURL: utils.String(client.CreateBaseURL("serverless-upload", "v1", nil, nil)),
		Method:          http.MethodPost,
		URI:             "/Services/{serviceSid}/Assets/{assetSid}/Versions",
		ContentType:     client.FormData,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"assetSid":   c.assetSid,
		},
	}

	if input == nil {
		input = &CreateVersionInput{}
	}

	response := &CreateVersionResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
