// Package channel contains auto-generated files. DO NOT MODIFY
package channel

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchChannelResponse defines the response fields for the retrieved channel
type FetchChannelResponse struct {
	AccountSid    string     `json:"account_sid"`
	Attributes    *string    `json:"attributes,omitempty"`
	CreatedBy     string     `json:"created_by"`
	DateCreated   time.Time  `json:"date_created"`
	DateUpdated   *time.Time `json:"date_updated,omitempty"`
	FriendlyName  *string    `json:"friendly_name,omitempty"`
	MembersCount  int        `json:"members_count"`
	MessagesCount int        `json:"messages_count"`
	ServiceSid    string     `json:"service_sid"`
	Sid           string     `json:"sid"`
	Type          string     `json:"type"`
	URL           string     `json:"url"`
	UniqueName    *string    `json:"unique_name,omitempty"`
}

// Fetch retrieves a channel resource
// See https://www.twilio.com/docs/chat/rest/channel-resource#fetch-a-channel-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchChannelResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a channel resource
// See https://www.twilio.com/docs/chat/rest/channel-resource#fetch-a-channel-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchChannelResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Channels/{sid}",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"sid":        c.sid,
		},
	}

	response := &FetchChannelResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
