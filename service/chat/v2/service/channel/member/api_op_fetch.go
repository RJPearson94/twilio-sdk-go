// Package member contains auto-generated files. DO NOT MODIFY
package member

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchChannelMemberResponse defines the response fields for the retrieved member
type FetchChannelMemberResponse struct {
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

// Fetch retrieves a member resource
// See https://www.twilio.com/docs/chat/rest/member-resource#fetch-a-member-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchChannelMemberResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a member resource
// See https://www.twilio.com/docs/chat/rest/member-resource#fetch-a-member-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchChannelMemberResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Channels/{channelSid}/Members/{sid}",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"channelSid": c.channelSid,
			"sid":        c.sid,
		},
	}

	response := &FetchChannelMemberResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
