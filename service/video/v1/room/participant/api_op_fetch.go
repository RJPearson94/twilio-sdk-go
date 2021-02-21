// Package participant contains auto-generated files. DO NOT MODIFY
package participant

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchParticipantResponse defines the response fields for the retrieved participant
type FetchParticipantResponse struct {
	AccountSid  string     `json:"account_sid"`
	DateCreated time.Time  `json:"date_created"`
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	Duration    *int       `json:"duration,omitempty"`
	EndTime     *time.Time `json:"end_time,omitempty"`
	Identity    string     `json:"identity"`
	RoomSid     string     `json:"room_sid"`
	Sid         string     `json:"sid"`
	StartTime   time.Time  `json:"start_time"`
	Status      string     `json:"status"`
	URL         string     `json:"url"`
}

// Fetch retrieves a participant resource
// See https://www.twilio.com/docs/video/api/participants#http-get-1 for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchParticipantResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a participant resource
// See https://www.twilio.com/docs/video/api/participants#http-get-1 for more details
func (c Client) FetchWithContext(context context.Context) (*FetchParticipantResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Rooms/{roomSid}/Participants/{sid}",
		PathParams: map[string]string{
			"roomSid": c.roomSid,
			"sid":     c.sid,
		},
	}

	response := &FetchParticipantResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
