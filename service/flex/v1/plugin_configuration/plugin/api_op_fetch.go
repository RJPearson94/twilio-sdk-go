// Package plugin contains auto-generated files. DO NOT MODIFY
package plugin

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchPluginResponse defines the response fields for the retrieved plugin resource
type FetchPluginResponse struct {
	AccountSid       string    `json:"account_sid"`
	ConfigurationSid string    `json:"configuration_sid"`
	DateCreated      time.Time `json:"date_created"`
	Phase            int       `json:"phase"`
	PluginSid        string    `json:"plugin_sid"`
	PluginURL        string    `json:"plugin_url"`
	PluginVersionSid string    `json:"plugin_version_sid"`
	Private          bool      `json:"private"`
	URL              string    `json:"url"`
	UniqueName       string    `json:"unique_name"`
	Version          string    `json:"version"`
}

// Fetch retrieves a plugin resource
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
// This resource is currently in beta and subject to change. Please use with caution
func (c Client) Fetch() (*FetchPluginResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a plugin resource
// This resource is currently in beta and subject to change. Please use with caution
func (c Client) FetchWithContext(context context.Context) (*FetchPluginResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/PluginService/Configurations/{configurationSid}/Plugins/{sid}",
		PathParams: map[string]string{
			"configurationSid": c.configurationSid,
			"sid":              c.sid,
		},
	}

	response := &FetchPluginResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
