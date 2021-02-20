// Package room contains auto-generated files. DO NOT MODIFY
package room

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchRoomResponse defines the response fields for the retrieved room
type FetchRoomResponse struct {
	AccountSid                   string     `json:"account_sid"`
	DateCreated                  time.Time  `json:"date_created"`
	DateUpdated                  *time.Time `json:"date_updated,omitempty"`
	Duration                     *int       `json:"duration,omitempty"`
	EndTime                      *time.Time `json:"end_time,omitempty"`
	MaxConcurrentPublishedTracks *int       `json:"max_concurrent_published_tracks,omitempty"`
	MaxParticipants              int        `json:"max_participants"`
	MediaRegion                  *string    `json:"media_region,omitempty"`
	RecordParticipantsOnConnect  bool       `json:"record_participants_on_connect"`
	Sid                          string     `json:"sid"`
	Status                       string     `json:"status"`
	StatusCallback               *string    `json:"status_callback,omitempty"`
	StatusCallbackMethod         *string    `json:"status_callback_method,omitempty"`
	Type                         string     `json:"type"`
	URL                          string     `json:"url"`
	UniqueName                   string     `json:"unique_name"`
	VideoCodecs                  *[]string  `json:"video_codecs,omitempty"`
}

// Fetch retrieves a room resource
// See https://www.twilio.com/docs/video/api/rooms-resource#get-instance for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchRoomResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a room resource
// See https://www.twilio.com/docs/video/api/rooms-resource#get-instance for more details
func (c Client) FetchWithContext(context context.Context) (*FetchRoomResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Rooms/{sid}",
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	response := &FetchRoomResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
