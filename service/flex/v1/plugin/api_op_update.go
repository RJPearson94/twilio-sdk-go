// Package plugin contains auto-generated files. DO NOT MODIFY
package plugin

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// UpdatePluginInput defines input fields for updating a plugin resource
type UpdatePluginInput struct {
	Description  *string `form:"Description,omitempty"`
	FriendlyName *string `form:"FriendlyName,omitempty"`
}

// UpdatePluginResponse defines the response fields for the updated plugin
type UpdatePluginResponse struct {
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

// Update modifies a plugin resource
// See https://www.twilio.com/docs/flex/developer/plugins/api/plugin#update-a-plugin-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
// This resource is currently in beta and subject to change. Please use with caution
func (c Client) Update(input *UpdatePluginInput) (*UpdatePluginResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a plugin resource
// See https://www.twilio.com/docs/flex/developer/plugins/api/plugin#update-a-plugin-resource for more details
// This resource is currently in beta and subject to change. Please use with caution
func (c Client) UpdateWithContext(context context.Context, input *UpdatePluginInput) (*UpdatePluginResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/PluginService/Plugins/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	if input == nil {
		input = &UpdatePluginInput{}
	}

	response := &UpdatePluginResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
