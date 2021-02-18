// Package plugin_release contains auto-generated files. DO NOT MODIFY
package plugin_release

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchReleaseResponse defines the response fields for the retrieved plugin release resource
type FetchReleaseResponse struct {
	AccountSid       string    `json:"account_sid"`
	ConfigurationSid string    `json:"configuration_sid"`
	DateCreated      time.Time `json:"date_created"`
	Sid              string    `json:"sid"`
	URL              string    `json:"url"`
}

// Fetch retrieves a plugin release resource
// See https://www.twilio.com/docs/flex/developer/plugins/api/release#fetch-a-pluginrelease-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
// This resource is currently in beta and subject to change. Please use with caution
func (c Client) Fetch() (*FetchReleaseResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a plugin release resource
// See https://www.twilio.com/docs/flex/developer/plugins/api/release#fetch-a-pluginrelease-resource for more details
// This resource is currently in beta and subject to change. Please use with caution
func (c Client) FetchWithContext(context context.Context) (*FetchReleaseResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/PluginService/Releases/{sid}",
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	response := &FetchReleaseResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
