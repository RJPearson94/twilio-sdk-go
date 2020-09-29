// Package message contains auto-generated files. DO NOT MODIFY
package message

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchChannelMessageResponse defines the response fields for the retrieved message
type FetchChannelMessageResponse struct {
	AccountSid    string                  `json:"account_sid"`
	Attributes    *string                 `json:"attributes,omitempty"`
	Body          *string                 `json:"body,omitempty"`
	ChannelSid    string                  `json:"channel_sid"`
	DateCreated   time.Time               `json:"date_created"`
	DateUpdated   *time.Time              `json:"date_updated,omitempty"`
	From          *string                 `json:"from,omitempty"`
	Index         *int                    `json:"index,omitempty"`
	LastUpdatedBy *string                 `json:"last_updated_by,omitempty"`
	Media         *map[string]interface{} `json:"media,omitempty"`
	ServiceSid    string                  `json:"service_sid"`
	Sid           string                  `json:"sid"`
	To            *string                 `json:"to,omitempty"`
	Type          *string                 `json:"type,omitempty"`
	URL           string                  `json:"url"`
	WasEdited     *bool                   `json:"was_edited,omitempty"`
}

// Fetch retrieves a message resource
// See https://www.twilio.com/docs/chat/rest/message-resource#fetch-a-message-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchChannelMessageResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a message resource
// See https://www.twilio.com/docs/chat/rest/message-resource#fetch-a-message-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchChannelMessageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Channels/{channelSid}/Messages/{sid}",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"channelSid": c.channelSid,
			"sid":        c.sid,
		},
	}

	response := &FetchChannelMessageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
