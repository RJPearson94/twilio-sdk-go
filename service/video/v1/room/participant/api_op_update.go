// Package participant contains auto-generated files. DO NOT MODIFY
package participant

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// UpdateParticipantInput defines input fields for updating a participant resource
type UpdateParticipantInput struct {
	Status *string `form:"Status,omitempty"`
}

// UpdateParticipantResponse defines the response fields for the updated participant
type UpdateParticipantResponse struct {
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

// Update modifies a participant resource
// See https://www.twilio.com/docs/video/api/participants#http-post for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateParticipantInput) (*UpdateParticipantResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a participant resource
// See https://www.twilio.com/docs/video/api/participants#http-post for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateParticipantInput) (*UpdateParticipantResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Rooms/{roomSid}/Participants/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"roomSid": c.roomSid,
			"sid":     c.sid,
		},
	}

	if input == nil {
		input = &UpdateParticipantInput{}
	}

	response := &UpdateParticipantResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
