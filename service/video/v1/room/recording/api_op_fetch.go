// Package recording contains auto-generated files. DO NOT MODIFY
package recording

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type FetchRecordingGroupingSidsResponse struct {
	ParticipantSid string `json:"participant_sid"`
	RoomSid        string `json:"room_sid"`
}

// FetchRecordingResponse defines the response fields for the retrieved recording
type FetchRecordingResponse struct {
	AccountSid      string                             `json:"account_sid"`
	Codec           string                             `json:"codec"`
	ContainerFormat string                             `json:"container_format"`
	DateCreated     time.Time                          `json:"date_created"`
	Duration        int                                `json:"duration"`
	GroupingSids    FetchRecordingGroupingSidsResponse `json:"grouping_sids"`
	Offset          int                                `json:"offset"`
	RoomSid         string                             `json:"room_sid"`
	Sid             string                             `json:"sid"`
	Size            int                                `json:"size"`
	SourceSid       string                             `json:"source_sid"`
	Status          string                             `json:"status"`
	TrackName       string                             `json:"track_name"`
	Type            string                             `json:"type"`
	URL             string                             `json:"url"`
}

// Fetch retrieves a recording resource
// See https://www.twilio.com/docs/video/api/recordings-resource#get-instance for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchRecordingResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a recording resource
// See https://www.twilio.com/docs/video/api/recordings-resource#get-instance for more details
func (c Client) FetchWithContext(context context.Context) (*FetchRecordingResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Rooms/{roomSid}/Recordings/{sid}",
		PathParams: map[string]string{
			"roomSid": c.roomSid,
			"sid":     c.sid,
		},
	}

	response := &FetchRecordingResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
