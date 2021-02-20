// Package plugin contains auto-generated files. DO NOT MODIFY
package plugin

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// ArchivePluginResponse defines the response fields for the archived plugin
type ArchivePluginResponse struct {
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

// Archive archives a plugin resource
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
// This resource is currently in beta and subject to change. Please use with caution
func (c Client) Archive() (*ArchivePluginResponse, error) {
	return c.ArchiveWithContext(context.Background())
}

// ArchiveWithContext archives a plugin resource
// This resource is currently in beta and subject to change. Please use with caution
func (c Client) ArchiveWithContext(context context.Context) (*ArchivePluginResponse, error) {
	op := client.Operation{
		Method: http.MethodPost,
		URI:    "/PluginService/Plugins/{sid}/Archive",
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	response := &ArchivePluginResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
