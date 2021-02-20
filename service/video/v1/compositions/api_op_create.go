// Package compositions contains auto-generated files. DO NOT MODIFY
package compositions

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreateCompositionInput defines the input fields for creating a new composition
type CreateCompositionInput struct {
	AudioSources         *[]string `form:"AudioSources,omitempty"`
	AudioSourcesExcluded *[]string `form:"AudioSourcesExcluded,omitempty"`
	Format               *string   `form:"Format,omitempty"`
	Resolution           *string   `form:"Resolution,omitempty"`
	RoomSid              string    `validate:"required" form:"RoomSid"`
	StatusCallback       *string   `form:"StatusCallback,omitempty"`
	StatusCallbackMethod *string   `form:"StatusCallbackMethod,omitempty"`
	Trim                 *bool     `form:"Trim,omitempty"`
	VideoLayout          *string   `form:"VideoLayout,omitempty"`
}

// CreateCompositionResponse defines the response fields for the created composition
type CreateCompositionResponse struct {
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

// Create creates a new composition
// See https://www.twilio.com/docs/video/api/compositions-resource#create-composition-http-post for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateCompositionInput) (*CreateCompositionResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new composition
// See https://www.twilio.com/docs/video/api/compositions-resource#create-composition-http-post for more details
func (c Client) CreateWithContext(context context.Context, input *CreateCompositionInput) (*CreateCompositionResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Compositions",
		ContentType: client.URLEncoded,
	}

	if input == nil {
		input = &CreateCompositionInput{}
	}

	response := &CreateCompositionResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
