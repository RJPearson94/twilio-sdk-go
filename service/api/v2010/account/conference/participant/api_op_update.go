// Package participant contains auto-generated files. DO NOT MODIFY
package participant

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// UpdateParticipantInput defines input fields for updating a participant
type UpdateParticipantInput struct {
	AnnounceMethod      *string `form:"AnnounceMethod,omitempty"`
	AnnounceURL         *string `form:"AnnounceUrl,omitempty"`
	BeepOnExit          *bool   `form:"BeepOnExit,omitempty"`
	CallSidToCoach      *string `form:"CallSidToCoach,omitempty"`
	Coaching            *bool   `form:"Coaching,omitempty"`
	EndConferenceOnExit *bool   `form:"EndConferenceOnExit,omitempty"`
	Hold                *bool   `form:"Hold,omitempty"`
	HoldMethod          *string `form:"HoldMethod,omitempty"`
	HoldURL             *string `form:"HoldUrl,omitempty"`
	Muted               *bool   `form:"Muted,omitempty"`
	WaitMethod          *string `form:"WaitMethod,omitempty"`
	WaitURL             *string `form:"WaitUrl,omitempty"`
}

// UpdateParticipantResponse defines the response fields for the updated participant
type UpdateParticipantResponse struct {
	AccountSid             string             `json:"account_sid"`
	CallSid                string             `json:"call_sid"`
	CallSidToCoach         *string            `json:"call_sid_to_coach,omitempty"`
	Coaching               bool               `json:"coaching"`
	ConferenceSid          string             `json:"conference_sid"`
	DateCreated            utils.RFC2822Time  `json:"date_created"`
	DateUpdated            *utils.RFC2822Time `json:"date_updated,omitempty"`
	EndConferenceOnExit    bool               `json:"end_conference_on_exit"`
	Hold                   bool               `json:"hold"`
	Label                  *string            `json:"label,omitempty"`
	Muted                  bool               `json:"muted"`
	StartConferenceOnEnter bool               `json:"start_conference_on_enter"`
	Status                 string             `json:"status"`
}

// Update modifies a participant resource
// See https://www.twilio.com/docs/voice/api/conference-participant-resource#update-a-participant-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateParticipantInput) (*UpdateParticipantResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a participant resource
// See https://www.twilio.com/docs/voice/api/conference-participant-resource#update-a-participant-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateParticipantInput) (*UpdateParticipantResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Accounts/{accountSid}/Conferences/{conferenceSid}/Participants/{sid}.json",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"accountSid":    c.accountSid,
			"conferenceSid": c.conferenceSid,
			"sid":           c.sid,
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
