// Package plugins contains auto-generated files. DO NOT MODIFY
package plugins

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreatePluginInput defines the input fields for creating a new plugin resource
type CreatePluginInput struct {
	Description  *string `form:"Description,omitempty"`
	FriendlyName *string `form:"FriendlyName,omitempty"`
	UniqueName   string  `validate:"required" form:"UniqueName"`
}

// CreatePluginResponse defines the response fields for the created plugin
type CreatePluginResponse struct {
	AccountSid   string     `json:"account_sid"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	Description  *string    `json:"description,omitempty"`
	Disabled     bool       `json:"disabled"`
	FriendlyName string     `json:"friendly_name"`
	Sid          string     `json:"sid"`
	URL          string     `json:"url"`
	UniqueName   string     `json:"unique_name"`
}

// Create creates a new plugin resource
// See https://www.twilio.com/docs/flex/developer/plugins/api/plugin#create-a-plugin-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreatePluginInput) (*CreatePluginResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new plugin resource
// See https://www.twilio.com/docs/flex/developer/plugins/api/plugin#create-a-plugin-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreatePluginInput) (*CreatePluginResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/PluginService/Plugins",
		ContentType: client.URLEncoded,
	}

	if input == nil {
		input = &CreatePluginInput{}
	}

	response := &CreatePluginResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
