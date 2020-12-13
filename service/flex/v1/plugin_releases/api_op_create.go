// Package plugin_releases contains auto-generated files. DO NOT MODIFY
package plugin_releases

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreateReleaseInput defines the input fields for creating a new plugin release resource
type CreateReleaseInput struct {
	ConfigurationId string `validate:"required" form:"ConfigurationId"`
}

// CreateReleaseResponse defines the response fields for the created plugin release
type CreateReleaseResponse struct {
	AccountSid       string    `json:"account_sid"`
	ConfigurationSid string    `json:"configuration_sid"`
	DateCreated      time.Time `json:"date_created"`
	Sid              string    `json:"sid"`
	URL              string    `json:"url"`
}

// Create creates a new plugin release resource
// See https://www.twilio.com/docs/flex/developer/plugins/api/release#create-a-pluginrelease-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateReleaseInput) (*CreateReleaseResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new plugin release resource
// See https://www.twilio.com/docs/flex/developer/plugins/api/release#create-a-pluginrelease-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateReleaseInput) (*CreateReleaseResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/PluginService/Releases",
		ContentType: client.URLEncoded,
	}

	if input == nil {
		input = &CreateReleaseInput{}
	}

	response := &CreateReleaseResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
