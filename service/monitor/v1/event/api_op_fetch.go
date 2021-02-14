// Package event contains auto-generated files. DO NOT MODIFY
package event

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchEventResponse defines the response fields for the retrieved event
type FetchEventResponse struct {
	AccountSid      string                 `json:"account_sid"`
	ActorSid        string                 `json:"actor_sid"`
	ActorType       string                 `json:"actor_type"`
	Description     *string                `json:"description,omitempty"`
	EventData       map[string]interface{} `json:"event_data"`
	EventDate       time.Time              `json:"event_date"`
	EventType       string                 `json:"event_type"`
	ResourceSid     string                 `json:"resource_sid"`
	ResourceType    string                 `json:"resource_type"`
	Sid             string                 `json:"sid"`
	Source          string                 `json:"source"`
	SourceIpAddress string                 `json:"source_ip_address"`
	URL             string                 `json:"url"`
}

// Fetch retrieves an event resource
// See https://www.twilio.com/docs/usage/monitor-events#fetch-an-event-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchEventResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves an event resource
// See https://www.twilio.com/docs/usage/monitor-events#fetch-an-event-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchEventResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Events/{sid}",
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	response := &FetchEventResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
