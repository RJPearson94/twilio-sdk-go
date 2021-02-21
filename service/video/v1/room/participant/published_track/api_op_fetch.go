// Package published_track contains auto-generated files. DO NOT MODIFY
package published_track

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchPublishedTrackResponse defines the response fields for the retrieved published track
type FetchPublishedTrackResponse struct {
	DateCreated    time.Time  `json:"date_created"`
	DateUpdated    *time.Time `json:"date_updated,omitempty"`
	Enabled        bool       `json:"enabled"`
	Kind           string     `json:"kind"`
	Name           string     `json:"name"`
	ParticipantSid string     `json:"participant_sid"`
	RoomSid        string     `json:"room_sid"`
	Sid            string     `json:"sid"`
	URL            string     `json:"url"`
}

// Fetch retrieves a published track resource
// See https://www.twilio.com/docs/video/api/track-subscriptions#get-st for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchPublishedTrackResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a published track resource
// See https://www.twilio.com/docs/video/api/track-subscriptions#get-st for more details
func (c Client) FetchWithContext(context context.Context) (*FetchPublishedTrackResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Rooms/{roomSid}/Participants/{participantSid}/PublishedTracks/{sid}",
		PathParams: map[string]string{
			"roomSid":        c.roomSid,
			"participantSid": c.participantSid,
			"sid":            c.sid,
		},
	}

	response := &FetchPublishedTrackResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
