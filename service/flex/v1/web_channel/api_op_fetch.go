// Package web_channel contains auto-generated files. DO NOT MODIFY
package web_channel

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchWebChannelResponse defines the response fields for the retrieved web channel
type FetchWebChannelResponse struct {
	AccountSid  string     `json:"account_sid"`
	DateCreated time.Time  `json:"date_created"`
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	FlexFlowSid string     `json:"flex_flow_sid"`
	Sid         string     `json:"sid"`
	URL         string     `json:"url"`
}

// Fetch retrieves a web channel resource
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchWebChannelResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a web channel resource
func (c Client) FetchWithContext(context context.Context) (*FetchWebChannelResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/WebChannels/{sid}",
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	response := &FetchWebChannelResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
