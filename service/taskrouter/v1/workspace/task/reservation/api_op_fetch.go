// Package reservation contains auto-generated files. DO NOT MODIFY
package reservation

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchReservationResponse defines the response fields for the retrieved task reservation
type FetchReservationResponse struct {
	AccountSid        string     `json:"account_sid"`
	DateCreated       time.Time  `json:"date_created"`
	DateUpdated       *time.Time `json:"date_updated,omitempty"`
	ReservationStatus string     `json:"reservation_status"`
	Sid               string     `json:"sid"`
	TaskSid           string     `json:"task_sid"`
	URL               string     `json:"url"`
	WorkerName        string     `json:"worker_name"`
	WorkerSid         string     `json:"worker_sid"`
	WorkspaceSid      string     `json:"workspace_sid"`
}

// Fetch retrieves an task reservation resource
// See https://www.twilio.com/docs/taskrouter/api/reservations#action-get for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchReservationResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves an task reservation resource
// See https://www.twilio.com/docs/taskrouter/api/reservations#action-get for more details
func (c Client) FetchWithContext(context context.Context) (*FetchReservationResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Workspaces/{workspaceSid}/Tasks/{taskSid}/Reservations/{sid}",
		PathParams: map[string]string{
			"workspaceSid": c.workspaceSid,
			"taskSid":      c.taskSid,
			"sid":          c.sid,
		},
	}

	response := &FetchReservationResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
