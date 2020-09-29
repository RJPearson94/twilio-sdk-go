// Package assistant contains auto-generated files. DO NOT MODIFY
package assistant

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchAssistantResponse defines the response fields for the retrieved assistant
type FetchAssistantResponse struct {
	AccountSid          string     `json:"account_sid"`
	CallbackEvents      *string    `json:"callback_events,omitempty"`
	CallbackURL         *string    `json:"callback_url,omitempty"`
	DateCreated         time.Time  `json:"date_created"`
	DateUpdated         *time.Time `json:"date_updated,omitempty"`
	DevelopmentStage    string     `json:"development_stage"`
	FriendlyName        *string    `json:"friendly_name,omitempty"`
	LatestModelBuildSid *string    `json:"latest_model_build_sid,omitempty"`
	LogQueries          bool       `json:"log_queries"`
	NeedsModelBuild     *bool      `json:"needs_model_build,omitempty"`
	Sid                 string     `json:"sid"`
	URL                 string     `json:"url"`
	UniqueName          string     `json:"unique_name"`
}

// Fetch retrieves an assistant resource
// See https://www.twilio.com/docs/autopilot/api/assistant#fetch-an-assistant-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchAssistantResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves an assistant resource
// See https://www.twilio.com/docs/autopilot/api/assistant#fetch-an-assistant-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchAssistantResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Assistants/{sid}",
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	response := &FetchAssistantResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
