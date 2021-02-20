// Package composition_hooks contains auto-generated files. DO NOT MODIFY
package composition_hooks

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreateCompositionHookInput defines the input fields for creating a new composition hook
type CreateCompositionHookInput struct {
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

// CreateCompositionHookResponse defines the response fields for the created composition hook
type CreateCompositionHookResponse struct {
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

// Create creates a new composition hook
// See https://www.twilio.com/docs/video/api/composition-hooks#hks-post for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateCompositionHookInput) (*CreateCompositionHookResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new composition hook
// See https://www.twilio.com/docs/video/api/composition-hooks#hks-post for more details
func (c Client) CreateWithContext(context context.Context, input *CreateCompositionHookInput) (*CreateCompositionHookResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/CompositionHooks",
		ContentType: client.URLEncoded,
	}

	if input == nil {
		input = &CreateCompositionHookInput{}
	}

	response := &CreateCompositionHookResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
