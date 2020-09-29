// Package member contains auto-generated files. DO NOT MODIFY
package member

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// FetchMemberResponse defines the response fields for retrieving a member
type FetchMemberResponse struct {
	CallSid      string            `json:"call_sid"`
	DateEnqueued utils.RFC2822Time `json:"date_enqueued"`
	Position     int               `json:"position"`
	QueueSid     string            `json:"queue_sid"`
	WaitTime     int               `json:"wait_time"`
}

// Fetch retrieves the member resource
// See https://www.twilio.com/docs/voice/api/member-resource#fetch-a-member-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchMemberResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves the member resource
// See https://www.twilio.com/docs/voice/api/member-resource#fetch-a-member-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchMemberResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Accounts/{accountSid}/Queues/{queueSid}/Members/{sid}.json",
		PathParams: map[string]string{
			"accountSid": c.accountSid,
			"queueSid":   c.queueSid,
			"sid":        c.sid,
		},
	}

	response := &FetchMemberResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
