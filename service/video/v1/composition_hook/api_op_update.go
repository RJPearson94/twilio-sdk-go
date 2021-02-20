// Package composition_hook contains auto-generated files. DO NOT MODIFY
package composition_hook

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// UpdateCompositionHookInput defines input fields for updating a composition hook resource. NOTE: This API does not support partial updates, please supply all necesary fields
type UpdateCompositionHookInput struct {
	AudioSources         *[]string `form:"AudioSources,omitempty"`
	AudioSourcesExcluded *[]string `form:"AudioSourcesExcluded,omitempty"`
	Enabled              *bool     `form:"Enabled,omitempty"`
	Format               *string   `form:"Format,omitempty"`
	FriendlyName         string    `validate:"required" form:"FriendlyName"`
	Resolution           *string   `form:"Resolution,omitempty"`
	StatusCallback       *string   `form:"StatusCallback,omitempty"`
	StatusCallbackMethod *string   `form:"StatusCallbackMethod,omitempty"`
	Trim                 *bool     `form:"Trim,omitempty"`
	VideoLayout          *string   `form:"VideoLayout,omitempty"`
}

// UpdateCompositionHookResponse defines the response fields for the updated composition hook
type UpdateCompositionHookResponse struct {
	AccountSid           string                 `json:"account_sid"`
	AudioSources         []string               `json:"audio_sources"`
	AudioSourcesExcluded []string               `json:"audio_sources_excluded"`
	DateCreated          time.Time              `json:"date_created"`
	DateUpdated          *time.Time             `json:"date_updated,omitempty"`
	Enabled              bool                   `json:"enabled"`
	Format               string                 `json:"format"`
	FriendlyName         string                 `json:"friendly_name"`
	Resolution           string                 `json:"resolution"`
	Sid                  string                 `json:"sid"`
	StatusCallback       *string                `json:"status_callback,omitempty"`
	StatusCallbackMethod string                 `json:"status_callback_method"`
	Trim                 bool                   `json:"trim"`
	URL                  string                 `json:"url"`
	VideoLayout          map[string]interface{} `json:"video_layout"`
}

// Update modifies a composition hook resource
// See https://www.twilio.com/docs/video/api/composition-hooks#hk-post for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateCompositionHookInput) (*UpdateCompositionHookResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a composition hook resource
// See https://www.twilio.com/docs/video/api/composition-hooks#hk-post for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateCompositionHookInput) (*UpdateCompositionHookResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/CompositionHooks/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	if input == nil {
		input = &UpdateCompositionHookInput{}
	}

	response := &UpdateCompositionHookResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
