// Package call contains auto-generated files. DO NOT MODIFY
package call

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// UpdateCallInput defines input fields for updating a call
type UpdateCallInput struct {
	FallbackMethod       *string `form:"FallbackMethod,omitempty"`
	FallbackURL          *string `form:"FallbackUrl,omitempty"`
	Method               *string `form:"Method,omitempty"`
	Status               *string `form:"Status,omitempty"`
	StatusCallback       *string `form:"StatusCallback,omitempty"`
	StatusCallbackMethod *string `form:"StatusCallbackMethod,omitempty"`
	TwiML                *string `form:"Twiml,omitempty"`
	URL                  *string `form:"Url,omitempty"`
}

// UpdateCallResponse defines the response fields for the updated call
type UpdateCallResponse struct {
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

// Update modifies a call resource
// See https://www.twilio.com/docs/voice/api/call-resource#update-a-call-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateCallInput) (*UpdateCallResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a call resource
// See https://www.twilio.com/docs/voice/api/call-resource#update-a-call-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateCallInput) (*UpdateCallResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Accounts/{accountSid}/Calls/{sid}.json",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"accountSid": c.accountSid,
			"sid":        c.sid,
		},
	}

	if input == nil {
		input = &UpdateCallInput{}
	}

	response := &UpdateCallResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
