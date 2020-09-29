// Package flow contains auto-generated files. DO NOT MODIFY
package flow

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchFlowResponse defines the response fields for the retrieved flow
type FetchFlowResponse struct {
	AccountSid    string                 `json:"account_sid"`
	CommitMessage *string                `json:"commit_message,omitempty"`
	DateCreated   time.Time              `json:"date_created"`
	DateUpdated   *time.Time             `json:"date_updated,omitempty"`
	Definition    map[string]interface{} `json:"definition"`
	Errors        *[]interface{}         `json:"errors,omitempty"`
	FriendlyName  string                 `json:"friendly_name"`
	Revision      int                    `json:"revision"`
	Sid           string                 `json:"sid"`
	Status        string                 `json:"status"`
	URL           string                 `json:"url"`
	Valid         bool                   `json:"valid"`
	Warnings      *[]interface{}         `json:"warnings,omitempty"`
	WebhookURL    string                 `json:"webhook_url"`
}

// Fetch retrieves a flow resource
// See https://www.twilio.com/docs/studio/rest-api/v2/flow#fetch-a-flow-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchFlowResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a flow resource
// See https://www.twilio.com/docs/studio/rest-api/v2/flow#fetch-a-flow-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchFlowResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Flows/{sid}",
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	response := &FetchFlowResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
