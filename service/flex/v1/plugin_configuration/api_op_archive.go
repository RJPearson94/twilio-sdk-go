// Package plugin_configuration contains auto-generated files. DO NOT MODIFY
package plugin_configuration

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// ArchiveConfigurationResponse defines the response fields for the archived plugin configuration resource
type ArchiveConfigurationResponse struct {
	AccountSid  string    `json:"account_sid"`
	Archived    bool      `json:"archived"`
	DateCreated time.Time `json:"date_created"`
	Description *string   `json:"description,omitempty"`
	Name        string    `json:"name"`
	Sid         string    `json:"sid"`
	URL         string    `json:"url"`
}

// Archive archives a plugin configuration resource
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
// This resource is currently in beta and subject to change. Please use with caution
func (c Client) Archive() (*ArchiveConfigurationResponse, error) {
	return c.ArchiveWithContext(context.Background())
}

// ArchiveWithContext archives a plugin configuration resource
// This resource is currently in beta and subject to change. Please use with caution
func (c Client) ArchiveWithContext(context context.Context) (*ArchiveConfigurationResponse, error) {
	op := client.Operation{
		Method: http.MethodPost,
		URI:    "/PluginService/Configurations/{sid}/Archive",
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	response := &ArchiveConfigurationResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
