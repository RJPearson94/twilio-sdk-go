// Package room contains auto-generated files. DO NOT MODIFY
package room

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// UpdateRoomInput defines input fields for updating a room resource
type UpdateRoomInput struct {
	Status string `validate:"required" form:"Status"`
}

// UpdateRoomResponse defines the response fields for the updated room
type UpdateRoomResponse struct {
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

// Update modifies a room resource
// See https://www.twilio.com/docs/video/api/rooms-resource#post-instance for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateRoomInput) (*UpdateRoomResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a room resource
// See https://www.twilio.com/docs/video/api/rooms-resource#post-instance for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateRoomInput) (*UpdateRoomResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Rooms/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	if input == nil {
		input = &UpdateRoomInput{}
	}

	response := &UpdateRoomResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
