// Package plugin_configurations contains auto-generated files. DO NOT MODIFY
package plugin_configurations

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreateConfigurationInput defines the input fields for creating a new plugin configuration resource
type CreateConfigurationInput struct {
	Description *string   `form:"Description,omitempty"`
	Name        string    `validate:"required" form:"Name"`
	Plugins     *[]string `form:"Plugins,omitempty"`
}

// CreateConfigurationResponse defines the response fields for the created plugin configuration
type CreateConfigurationResponse struct {
	AccountSid  string    `json:"account_sid"`
	DateCreated time.Time `json:"date_created"`
	Description *string   `json:"description,omitempty"`
	Disabled    bool      `json:"disabled"`
	Name        string    `json:"name"`
	Sid         string    `json:"sid"`
	URL         string    `json:"url"`
}

// Create creates a new plugin configuration resource
// See https://www.twilio.com/docs/flex/developer/plugins/api/plugin-configuration#create-a-pluginconfiguration-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateConfigurationInput) (*CreateConfigurationResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new plugin configuration resource
// See https://www.twilio.com/docs/flex/developer/plugins/api/plugin-configuration#create-a-pluginconfiguration-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateConfigurationInput) (*CreateConfigurationResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/PluginService/Configurations",
		ContentType: client.URLEncoded,
	}

	if input == nil {
		input = &CreateConfigurationInput{}
	}

	response := &CreateConfigurationResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
