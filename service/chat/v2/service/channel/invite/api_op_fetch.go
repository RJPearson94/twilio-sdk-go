// Package invite contains auto-generated files. DO NOT MODIFY
package invite

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchChannelInviteResponse defines the response fields for the retrieved invite
type FetchChannelInviteResponse struct {
	AccountSid  string     `json:"account_sid"`
	ChannelSid  string     `json:"channel_sid"`
	CreatedBy   *string    `json:"created_by,omitempty"`
	DateCreated time.Time  `json:"date_created"`
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	Identity    string     `json:"identity"`
	RoleSid     *string    `json:"role_sid,omitempty"`
	ServiceSid  string     `json:"service_sid"`
	Sid         string     `json:"sid"`
	URL         string     `json:"url"`
}

// Fetch retrieves a invite resource
// See https://www.twilio.com/docs/chat/rest/invite-resource#fetch-an-invite-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchChannelInviteResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a invite resource
// See https://www.twilio.com/docs/chat/rest/invite-resource#fetch-an-invite-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchChannelInviteResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Channels/{channelSid}/Invites/{sid}",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"channelSid": c.channelSid,
			"sid":        c.sid,
		},
	}

	response := &FetchChannelInviteResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
