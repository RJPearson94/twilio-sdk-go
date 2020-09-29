// Package member contains auto-generated files. DO NOT MODIFY
package member

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// UpdateChannelMemberInput defines input fields for updating a member resource
type UpdateChannelMemberInput struct {
	Attributes               *string    `form:"Attributes,omitempty"`
	DateCreated              *time.Time `form:"DateCreated,omitempty"`
	DateUpdated              *time.Time `form:"DateUpdated,omitempty"`
	LastConsumedMessageIndex *int       `form:"LastConsumedMessageIndex,omitempty"`
	LastConsumptionTimestamp *time.Time `form:"LastConsumptionTimestamp,omitempty"`
	RoleSid                  *string    `form:"RoleSid,omitempty"`
}

// UpdateChannelMemberResponse defines the response fields for the updated member
type UpdateChannelMemberResponse struct {
	AccountSid               string     `json:"account_sid"`
	Attributes               *string    `json:"attributes,omitempty"`
	ChannelSid               string     `json:"channel_sid"`
	DateCreated              time.Time  `json:"date_created"`
	DateUpdated              *time.Time `json:"date_updated,omitempty"`
	Identity                 string     `json:"identity"`
	LastConsumedMessageIndex *int       `json:"last_consumed_message_index,omitempty"`
	LastConsumedTimestamp    *time.Time `json:"last_consumption_timestamp,omitempty"`
	RoleSid                  *string    `json:"role_sid,omitempty"`
	ServiceSid               string     `json:"service_sid"`
	Sid                      string     `json:"sid"`
	URL                      string     `json:"url"`
}

// Update modifies a member resource
// See https://www.twilio.com/docs/chat/rest/member-resource#update-a-member-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateChannelMemberInput) (*UpdateChannelMemberResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a member resource
// See https://www.twilio.com/docs/chat/rest/member-resource#update-a-member-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateChannelMemberInput) (*UpdateChannelMemberResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Channels/{channelSid}/Members/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"channelSid": c.channelSid,
			"sid":        c.sid,
		},
	}

	if input == nil {
		input = &UpdateChannelMemberInput{}
	}

	response := &UpdateChannelMemberResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
