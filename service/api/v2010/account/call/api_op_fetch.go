// Package call contains auto-generated files. DO NOT MODIFY
package call

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// FetchCallResponse defines the response fields for retrieving a call
type FetchCallResponse struct {
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

// Fetch retrieves a call resource
// See https://www.twilio.com/docs/voice/api/call-resource#fetch-a-call-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchCallResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a call resource
// See https://www.twilio.com/docs/voice/api/call-resource#fetch-a-call-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchCallResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Accounts/{accountSid}/Calls/{sid}.json",
		PathParams: map[string]string{
			"accountSid": c.accountSid,
			"sid":        c.sid,
		},
	}

	response := &FetchCallResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
