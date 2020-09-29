// Package task_channel contains auto-generated files. DO NOT MODIFY
package task_channel

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchTaskChannelResponse defines the response fields for the retrieved task channel
type FetchTaskChannelResponse struct {
	AccountSid              string     `json:"account_sid"`
	ChannelOptimizedRouting *bool      `json:"channel_optimized_routing,omitempty"`
	DateCreated             time.Time  `json:"date_created"`
	DateUpdated             *time.Time `json:"date_updated,omitempty"`
	FriendlyName            string     `json:"friendly_name"`
	Sid                     string     `json:"sid"`
	URL                     string     `json:"url"`
	UniqueName              string     `json:"unique_name"`
	WorkspaceSid            string     `json:"workspace_sid"`
}

// Fetch retrieves an task channel resource
// See twilio.com/docs/taskrouter/api/task-channel#fetch-a-taskchannel-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchTaskChannelResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves an task channel resource
// See twilio.com/docs/taskrouter/api/task-channel#fetch-a-taskchannel-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchTaskChannelResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Workspaces/{workspaceSid}/TaskChannels/{sid}",
		PathParams: map[string]string{
			"workspaceSid": c.workspaceSid,
			"sid":          c.sid,
		},
	}

	response := &FetchTaskChannelResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
