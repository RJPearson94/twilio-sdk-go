// Package member contains auto-generated files. DO NOT MODIFY
package member

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// UpdateMemberInput defines input fields for updating a member
type UpdateMemberInput struct {
	Method *string `form:"Method,omitempty"`
	URL    string  `validate:"required" form:"Url"`
}

// UpdateMemberResponse defines the response fields for the updated member
type UpdateMemberResponse struct {
	CallSid      string            `json:"call_sid"`
	DateEnqueued utils.RFC2822Time `json:"date_enqueued"`
	Position     int               `json:"position"`
	QueueSid     string            `json:"queue_sid"`
	WaitTime     int               `json:"wait_time"`
}

// Update modifies a member resource
// See https://www.twilio.com/docs/voice/api/member-resource#update-a-member-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateMemberInput) (*UpdateMemberResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a member resource
// See https://www.twilio.com/docs/voice/api/member-resource#update-a-member-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateMemberInput) (*UpdateMemberResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Accounts/{accountSid}/Queues/{queueSid}/Members/{sid}.json",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"accountSid": c.accountSid,
			"queueSid":   c.queueSid,
			"sid":        c.sid,
		},
	}

	if input == nil {
		input = &UpdateMemberInput{}
	}

	response := &UpdateMemberResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
