// Package version contains auto-generated files. DO NOT MODIFY
package version

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// ArchiveVersionResponse defines the response fields for the archived plugin version resource
type ArchiveVersionResponse struct {
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

// Archive archives a plugin version resource
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
// This resource is currently in beta and subject to change. Please use with caution
func (c Client) Archive() (*ArchiveVersionResponse, error) {
	return c.ArchiveWithContext(context.Background())
}

// ArchiveWithContext archives a plugin version resource
// This resource is currently in beta and subject to change. Please use with caution
func (c Client) ArchiveWithContext(context context.Context) (*ArchiveVersionResponse, error) {
	op := client.Operation{
		Method: http.MethodPost,
		URI:    "/PluginService/Plugins/{pluginSid}/Versions/{sid}/Archive",
		PathParams: map[string]string{
			"pluginSid": c.pluginSid,
			"sid":       c.sid,
		},
	}

	response := &ArchiveVersionResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
