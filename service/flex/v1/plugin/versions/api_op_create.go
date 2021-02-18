// Package versions contains auto-generated files. DO NOT MODIFY
package versions

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreateVersionInput defines the input fields for creating a new plugin version resource
type CreateVersionInput struct {
	Changelog *string `form:"Changelog,omitempty"`
	PluginURL string  `validate:"required" form:"PluginUrl"`
	Private   *bool   `form:"Private,omitempty"`
	Version   string  `validate:"required" form:"Version"`
}

// CreateVersionResponse defines the response fields for the created plugin version
type CreateVersionResponse struct {
	AccountSid  string    `json:"account_sid"`
	Archived    bool      `json:"archived"`
	Changelog   *string   `json:"changelog,omitempty"`
	DateCreated time.Time `json:"date_created"`
	PluginSid   string    `json:"plugin_sid"`
	PluginURL   string    `json:"plugin_url"`
	Private     bool      `json:"private"`
	Sid         string    `json:"sid"`
	URL         string    `json:"url"`
	Version     string    `json:"version"`
}

// Create creates a new plugin version resource
// See https://www.twilio.com/docs/flex/developer/plugins/api/plugin-version#create-a-pluginversion-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
// This resource is currently in beta and subject to change. Please use with caution
func (c Client) Create(input *CreateVersionInput) (*CreateVersionResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new plugin version resource
// See https://www.twilio.com/docs/flex/developer/plugins/api/plugin-version#create-a-pluginversion-resource for more details
// This resource is currently in beta and subject to change. Please use with caution
func (c Client) CreateWithContext(context context.Context, input *CreateVersionInput) (*CreateVersionResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/PluginService/Plugins/{pluginSid}/Versions",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"pluginSid": c.pluginSid,
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
