// Package rooms contains auto-generated files. DO NOT MODIFY
package rooms

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreateRoomInput defines the input fields for creating a new room
type CreateRoomInput struct {
	MaxParticipants             *int      `form:"MaxParticipants,omitempty"`
	MediaRegion                 *string   `form:"MediaRegion,omitempty"`
	RecordParticipantsOnConnect *bool     `form:"RecordParticipantsOnConnect,omitempty"`
	StatusCallback              *string   `form:"StatusCallback,omitempty"`
	StatusCallbackMethod        *string   `form:"StatusCallbackMethod,omitempty"`
	Type                        *string   `form:"Type,omitempty"`
	UniqueName                  *string   `form:"UniqueName,omitempty"`
	VideoCodecs                 *[]string `form:"VideoCodecs,omitempty"`
}

// CreateRoomResponse defines the response fields for the created room
type CreateRoomResponse struct {
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
	StatusCallbackMethod         string     `json:"status_callback_method"`
	Type                         string     `json:"type"`
	URL                          string     `json:"url"`
	UniqueName                   string     `json:"unique_name"`
	VideoCodecs                  *[]string  `json:"video_codecs,omitempty"`
}

// Create creates a new room
// See https://www.twilio.com/docs/video/api/rooms-resource#create-room for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateRoomInput) (*CreateRoomResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new room
// See https://www.twilio.com/docs/video/api/rooms-resource#create-room for more details
func (c Client) CreateWithContext(context context.Context, input *CreateRoomInput) (*CreateRoomResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Rooms",
		ContentType: client.URLEncoded,
	}

	if input == nil {
		input = &CreateRoomInput{}
	}

	response := &CreateRoomResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
