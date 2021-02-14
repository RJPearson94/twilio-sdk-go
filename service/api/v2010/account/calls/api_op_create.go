// Package calls contains auto-generated files. DO NOT MODIFY
package calls

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// CreateCallInput defines input fields for making a new call
type CreateCallInput struct {
	ApplicationSid                     *string   `form:"ApplicationSid,omitempty"`
	AsyncAMD                           *bool     `form:"AsyncAmd,omitempty"`
	AsyncAMDStatusCallback             *string   `form:"AsyncAmdStatusCallback,omitempty"`
	AsyncAmdStatusCallbackMethod       *string   `form:"AsyncAmdStatusCallbackMethod,omitempty"`
	Byoc                               *string   `form:"Byoc,omitempty"`
	CallReason                         *string   `form:"CallReason,omitempty"`
	CallerID                           *string   `form:"CallerId,omitempty"`
	FallbackMethod                     *string   `form:"FallbackMethod,omitempty"`
	FallbackURL                        *string   `form:"FallbackUrl,omitempty"`
	From                               string    `validate:"required" form:"From"`
	MachineDetection                   *string   `form:"MachineDetection,omitempty"`
	MachineDetectionSilenceTimeout     *int      `form:"MachineDetectionSilenceTimeout,omitempty"`
	MachineDetectionSpeechEndThreshold *int      `form:"MachineDetectionSpeechEndThreshold,omitempty"`
	MachineDetectionSpeechThreshold    *int      `form:"MachineDetectionSpeechThreshold,omitempty"`
	MachineDetectionTimeout            *int      `form:"MachineDetectionTimeout,omitempty"`
	Method                             *string   `form:"Method,omitempty"`
	Record                             *bool     `form:"Record,omitempty"`
	RecordingChannels                  *string   `form:"RecordingChannels,omitempty"`
	RecordingStatusCallback            *string   `form:"RecordingStatusCallback,omitempty"`
	RecordingStatusCallbackEvents      *[]string `form:"RecordingStatusCallbackEvent,omitempty"`
	RecordingStatusCallbackMethod      *string   `form:"RecordingStatusCallbackMethod,omitempty"`
	RecordingTrack                     *string   `form:"RecordingTrack,omitempty"`
	SendDigits                         *string   `form:"SendDigits,omitempty"`
	SipAuthPassword                    *string   `form:"SipAuthPassword,omitempty"`
	SipAuthUsername                    *string   `form:"SipAuthUsername,omitempty"`
	StatusCallback                     *string   `form:"StatusCallback,omitempty"`
	StatusCallbackEvents               *[]string `form:"StatusCallbacks,omitempty"`
	StatusCallbackMethod               *string   `form:"StatusCallbackMethod,omitempty"`
	Timeout                            *int      `form:"Timeout,omitempty"`
	To                                 string    `validate:"required" form:"To"`
	Trim                               *string   `form:"Trim,omitempty"`
	TwiML                              *string   `form:"Twiml,omitempty"`
	URL                                *string   `form:"Url,omitempty"`
}

// CreateCallResponse defines the response fields for making a new call
type CreateCallResponse struct {
	APIVersion     string             `json:"api_version"`
	AccountSid     string             `json:"account_sid"`
	AnsweredBy     *string            `json:"answered_by,omitempty"`
	CallerName     *string            `json:"caller_name,omitempty"`
	DateCreated    utils.RFC2822Time  `json:"date_created"`
	DateUpdated    *utils.RFC2822Time `json:"date_updated,omitempty"`
	Direction      string             `json:"direction"`
	Duration       string             `json:"duration"`
	EndTime        *utils.RFC2822Time `json:"end_time,omitempty"`
	ForwardedFrom  *string            `json:"forwarded_from,omitempty"`
	From           string             `json:"from"`
	FromFormatted  string             `json:"from_formatted"`
	GroupSid       *string            `json:"group_sid,omitempty"`
	ParentCallSid  *string            `json:"parent_call_sid,omitempty"`
	PhoneNumberSid string             `json:"phone_number_sid"`
	Price          *string            `json:"price,omitempty"`
	PriceUnit      *string            `json:"price_unit,omitempty"`
	QueueTime      string             `json:"queue_time"`
	Sid            string             `json:"sid"`
	StartTime      *utils.RFC2822Time `json:"start_time,omitempty"`
	Status         string             `json:"status"`
	To             string             `json:"to"`
	ToFormatted    string             `json:"to_formatted"`
	TrunkSid       *string            `json:"trunk_sid,omitempty"`
}

// Create creates a new call resource
// See https://www.twilio.com/docs/voice/api/call-resource#create-a-call-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateCallInput) (*CreateCallResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new call resource
// See https://www.twilio.com/docs/voice/api/call-resource#create-a-call-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateCallInput) (*CreateCallResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Accounts/{accountSid}/Calls.json",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"accountSid": c.accountSid,
		},
	}

	if input == nil {
		input = &CreateCallInput{}
	}

	response := &CreateCallResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
