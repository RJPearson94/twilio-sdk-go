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
	AccountSid   string     `json:"account_sid"`
	Archived     bool       `json:"archived"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	Description  *string    `json:"description,omitempty"`
	FriendlyName string     `json:"friendly_name"`
	Sid          string     `json:"sid"`
	URL          string     `json:"url"`
	UniqueName   string     `json:"unique_name"`
}

// Fetch retrieves a plugin resource
// See https://www.twilio.com/docs/flex/developer/plugins/api/plugin#fetch-a-plugin-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
// This resource is currently in beta and subject to change. Please use with caution
func (c Client) Fetch() (*FetchPluginResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a plugin resource
// See https://www.twilio.com/docs/flex/developer/plugins/api/plugin#fetch-a-plugin-resource for more details
// This resource is currently in beta and subject to change. Please use with caution
func (c Client) FetchWithContext(context context.Context) (*FetchPluginResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/PluginService/Plugins/{sid}",
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	response := &FetchPluginResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
