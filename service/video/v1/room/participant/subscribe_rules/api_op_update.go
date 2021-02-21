// Package subscribe_rules contains auto-generated files. DO NOT MODIFY
package subscribe_rules

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// UpdateSubscribeRulesInput defines input fields for updating a subscribe rules resource
type UpdateSubscribeRulesInput struct {
	Rules string `validate:"required" form:"Rules"`
}

type UpdateSubscribeRuleResponse struct {
	All       *bool   `json:"all,omitempty"`
	Kind      *string `json:"kind,omitempty"`
	Priority  *string `json:"priority,omitempty"`
	Publisher *string `json:"publisher,omitempty"`
	Track     *string `json:"track,omitempty"`
	Type      string  `json:"type"`
}

// UpdateSubscribeRulesResponse defines the response fields for the updated subscribe rules
type UpdateSubscribeRulesResponse struct {
	DateCreated    time.Time                     `json:"date_created"`
	DateUpdated    *time.Time                    `json:"date_updated,omitempty"`
	ParticipantSid string                        `json:"participant_sid"`
	RoomSid        string                        `json:"room_sid"`
	Rules          []UpdateSubscribeRuleResponse `json:"rules"`
}

// Update modifies a subscribe rules resource
// See https://www.twilio.com/docs/video/api/track-subscriptions#sr-post for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateSubscribeRulesInput) (*UpdateSubscribeRulesResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a subscribe rules resource
// See https://www.twilio.com/docs/video/api/track-subscriptions#sr-post for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateSubscribeRulesInput) (*UpdateSubscribeRulesResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Rooms/{roomSid}/Participants/{participantSid}/SubscribeRules",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"roomSid":        c.roomSid,
			"participantSid": c.participantSid,
		},
	}

	if input == nil {
		input = &UpdateSubscribeRulesInput{}
	}

	response := &UpdateSubscribeRulesResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
