// Package version contains auto-generated files. DO NOT MODIFY
package version

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchVersionResponse defines the response fields for the retrieved plugin version resource
type FetchVersionResponse struct {
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

// Fetch retrieves a plugin version resource
// See https://www.twilio.com/docs/flex/developer/plugins/api/plugin-version#fetch-a-pluginversion-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchVersionResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a plugin version resource
// See https://www.twilio.com/docs/flex/developer/plugins/api/plugin-version#fetch-a-pluginversion-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchVersionResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/PluginService/Plugins/{pluginSid}/Versions/{sid}",
		PathParams: map[string]string{
			"pluginSid": c.pluginSid,
			"sid":       c.sid,
		},
	}

	response := &FetchVersionResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
