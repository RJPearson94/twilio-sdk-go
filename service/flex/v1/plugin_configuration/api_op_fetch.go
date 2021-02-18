// Package plugin_configuration contains auto-generated files. DO NOT MODIFY
package plugin_configuration

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchConfigurationResponse defines the response fields for the retrieved plugin configuration resource
type FetchConfigurationResponse struct {
	AccountSid  string    `json:"account_sid"`
	Archived    bool      `json:"archived"`
	DateCreated time.Time `json:"date_created"`
	Description *string   `json:"description,omitempty"`
	Name        string    `json:"name"`
	Sid         string    `json:"sid"`
	URL         string    `json:"url"`
}

// Fetch retrieves a plugin configuration resource
// See https://www.twilio.com/docs/flex/developer/plugins/api/plugin-configuration#fetch-a-pluginconfiguration-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchConfigurationResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a plugin configuration resource
// See https://www.twilio.com/docs/flex/developer/plugins/api/plugin-configuration#fetch-a-pluginconfiguration-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchConfigurationResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/PluginService/Configurations/{sid}",
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	response := &FetchConfigurationResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
