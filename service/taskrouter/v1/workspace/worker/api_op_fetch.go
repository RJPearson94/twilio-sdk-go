// Package worker contains auto-generated files. DO NOT MODIFY
package worker

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchWorkerResponse defines the response fields for the retrieved worker
type FetchWorkerResponse struct {
	AccountSid        string     `json:"account_sid"`
	ActivityName      string     `json:"activity_name"`
	ActivitySid       string     `json:"activity_sid"`
	Attributes        string     `json:"attributes"`
	Available         bool       `json:"available"`
	DateCreated       time.Time  `json:"date_created"`
	DateStatusChanged *time.Time `json:"date_status_changed,omitempty"`
	DateUpdated       *time.Time `json:"date_updated,omitempty"`
	FriendlyName      string     `json:"friendly_name"`
	Sid               string     `json:"sid"`
	URL               string     `json:"url"`
	WorkspaceSid      string     `json:"workspace_sid"`
}

// Fetch retrieves an worker resource
// See https://www.twilio.com/docs/taskrouter/api/worker#fetch-a-worker-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchWorkerResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves an worker resource
// See https://www.twilio.com/docs/taskrouter/api/worker#fetch-a-worker-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchWorkerResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Workspaces/{workspaceSid}/Workers/{sid}",
		PathParams: map[string]string{
			"workspaceSid": c.workspaceSid,
			"sid":          c.sid,
		},
	}

	response := &FetchWorkerResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
