// Package composition contains auto-generated files. DO NOT MODIFY
package composition

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchCompositionResponse defines the response fields for the retrieved composition
type FetchCompositionResponse struct {
	AccountSid           string                 `json:"account_sid"`
	AudioSources         []string               `json:"audio_sources"`
	AudioSourcesExcluded []string               `json:"audio_sources_excluded"`
	Bitrate              int                    `json:"bitrate"`
	DateCompleted        *time.Time             `json:"date_completed,omitempty"`
	DateCreated          time.Time              `json:"date_created"`
	DateDeleted          *time.Time             `json:"date_deleted,omitempty"`
	Duration             int                    `json:"duration"`
	Format               string                 `json:"format"`
	Resolution           string                 `json:"resolution"`
	RoomSid              string                 `json:"room_sid"`
	Sid                  string                 `json:"sid"`
	Size                 int                    `json:"size"`
	Status               string                 `json:"status"`
	Trim                 bool                   `json:"trim"`
	URL                  string                 `json:"url"`
	VideoLayout          map[string]interface{} `json:"video_layout"`
}

// Fetch retrieves a composition resource
// See https://www.twilio.com/docs/video/api/compositions-resource#get-composition-http-get for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchCompositionResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a composition resource
// See https://www.twilio.com/docs/video/api/compositions-resource#get-composition-http-get for more details
func (c Client) FetchWithContext(context context.Context) (*FetchCompositionResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Compositions/{sid}",
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	response := &FetchCompositionResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
