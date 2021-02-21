// Package subscribe_rules contains auto-generated files. DO NOT MODIFY
package subscribe_rules

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type FetchSubscribeRuleResponse struct {
	All       *bool   `json:"all,omitempty"`
	Kind      *string `json:"kind,omitempty"`
	Priority  *string `json:"priority,omitempty"`
	Publisher *string `json:"publisher,omitempty"`
	Track     *string `json:"track,omitempty"`
	Type      string  `json:"type"`
}

// FetchSubscribeRulesResponse defines the response fields for the retrieved subscribe rules
type FetchSubscribeRulesResponse struct {
	DateCreated    time.Time                    `json:"date_created"`
	DateUpdated    *time.Time                   `json:"date_updated,omitempty"`
	ParticipantSid string                       `json:"participant_sid"`
	RoomSid        string                       `json:"room_sid"`
	Rules          []FetchSubscribeRuleResponse `json:"rules"`
}

// Fetch retrieves a subscribe rules resource
// See https://www.twilio.com/docs/video/api/track-subscriptions#sr-get for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchSubscribeRulesResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a subscribe rules resource
// See https://www.twilio.com/docs/video/api/track-subscriptions#sr-get for more details
func (c Client) FetchWithContext(context context.Context) (*FetchSubscribeRulesResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Rooms/{roomSid}/Participants/{participantSid}/SubscribeRules",
		PathParams: map[string]string{
			"roomSid":        c.roomSid,
			"participantSid": c.participantSid,
		},
	}

	response := &FetchSubscribeRulesResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
