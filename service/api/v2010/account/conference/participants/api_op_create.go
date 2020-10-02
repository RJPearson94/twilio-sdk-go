// Package participants contains auto-generated files. DO NOT MODIFY
package participants

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// CreateParticipantInput defines input fields for adding a new participant
type CreateParticipantInput struct {
	BYOC                                    *string   `form:"Byoc,omitempty"`
	Beep                                    *string   `form:"Beep,omitempty"`
	CallReason                              *string   `form:"CallReason,omitempty"`
	CallSidToCoach                          *string   `form:"CallSidToCoach,omitempty"`
	CallerID                                *string   `form:"CallerId,omitempty"`
	Coaching                                *bool     `form:"Coaching,omitempty"`
	ConferenceRecord                        *string   `form:"ConferenceRecord,omitempty"`
	ConferenceRecordingStatusCallback       *string   `form:"ConferenceRecordingStatusCallback,omitempty"`
	ConferenceRecordingStatusCallbackEvents *[]string `form:"ConferenceRecordingStatusCallbackEvent,omitempty"`
	ConferenceRecordingStatusCallbackMethod *string   `form:"ConferenceRecordingStatusCallbackMethod,omitempty"`
	ConferenceStatusCallback                *string   `form:"ConferenceStatusCallback,omitempty"`
	ConferenceStatusCallbackEvents          *[]string `form:"ConferenceStatusCallbackEvent,omitempty"`
	ConferenceStatusCallbackMethod          *string   `form:"ConferenceStatusCallbackMethod,omitempty"`
	ConferenceTrim                          *string   `form:"ConferenceTrim,omitempty"`
	EarlyMedia                              *string   `form:"EarlyMedia,omitempty"`
	EndConferenceOnExit                     *bool     `form:"EndConferenceOnExit,omitempty"`
	From                                    string    `validate:"required" form:"From"`
	JitterBufferSize                        *string   `form:"JitterBufferSize,omitempty"`
	Label                                   *string   `form:"Label,omitempty"`
	MaxParticipants                         *int      `form:"MaxParticipants,omitempty"`
	Muted                                   *bool     `form:"Muted,omitempty"`
	Record                                  *bool     `form:"Record,omitempty"`
	RecordingChannels                       *string   `form:"RecordingChannels,omitempty"`
	RecordingStatusCallback                 *string   `form:"RecordingStatusCallback,omitempty"`
	RecordingStatusCallbackEvents           *[]string `form:"RecordingStatusCallbackEvent,omitempty"`
	RecordingStatusCallbackMethod           *string   `form:"RecordingStatusCallbackMethod,omitempty"`
	Region                                  *string   `form:"Region,omitempty"`
	SipAuthPassword                         *string   `form:"SipAuthPassword,omitempty"`
	SipAuthUsername                         *string   `form:"SipAuthUsername,omitempty"`
	StartConferenceOnEnter                  *bool     `form:"StartConferenceOnEnter,omitempty"`
	StatusCallback                          *string   `form:"StatusCallback,omitempty"`
	StatusCallbackEvents                    *[]string `form:"StatusCallbackEvent,omitempty"`
	StatusCallbackMethod                    *string   `form:"StatusCallbackMethod,omitempty"`
	Timeout                                 *int      `form:"Timeout,omitempty"`
	To                                      string    `validate:"required" form:"To"`
	WaitMethod                              *string   `form:"WaitMethod,omitempty"`
	WaitURL                                 *string   `form:"WaitUrl,omitempty"`
}

// CreateParticipantResponse defines the response fields for adding a new participant
type CreateParticipantResponse struct {
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

// Create creates a new participant resource
// See https://www.twilio.com/docs/voice/api/conference-participant-resource#create-a-participant for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateParticipantInput) (*CreateParticipantResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new participant resource
// See https://www.twilio.com/docs/voice/api/conference-participant-resource#create-a-participant for more details
func (c Client) CreateWithContext(context context.Context, input *CreateParticipantInput) (*CreateParticipantResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Accounts/{accountSid}/Conferences/{conferencesSid}/Participants.json",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"accountSid":     c.accountSid,
			"conferencesSid": c.conferencesSid,
		},
	}

	if input == nil {
		input = &CreateParticipantInput{}
	}

	response := &CreateParticipantResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
