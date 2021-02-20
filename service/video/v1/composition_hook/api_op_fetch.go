// Package composition_hook contains auto-generated files. DO NOT MODIFY
package composition_hook

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchCompositionHookResponse defines the response fields for the retrieved composition hook
type FetchCompositionHookResponse struct {
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

// Fetch retrieves a composition hook resource
// See https://www.twilio.com/docs/video/api/composition-hooks#hk-get for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchCompositionHookResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a composition hook resource
// See https://www.twilio.com/docs/video/api/composition-hooks#hk-get for more details
func (c Client) FetchWithContext(context context.Context) (*FetchCompositionHookResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/CompositionHooks/{sid}",
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	response := &FetchCompositionHookResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
