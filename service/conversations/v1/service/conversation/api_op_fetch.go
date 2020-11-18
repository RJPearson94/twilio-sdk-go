// Package conversation contains auto-generated files. DO NOT MODIFY
package conversation

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type FetchConversationTimersResponse struct {
	DateClosed   *time.Time `json:"date_closed,omitempty"`
	DateInactive *time.Time `json:"date_inactive,omitempty"`
}

// FetchConversationResponse defines the response fields for the retrieved conversation
type FetchConversationResponse struct {
	AccountSid          string                          `json:"account_sid"`
	Attributes          string                          `json:"attributes"`
	ChatServiceSid      *string                         `json:"chat_service_sid,omitempty"`
	DateCreated         time.Time                       `json:"date_created"`
	DateUpdated         *time.Time                      `json:"date_updated,omitempty"`
	FriendlyName        *string                         `json:"friendly_name,omitempty"`
	MessagingServiceSid *string                         `json:"messaging_service_sid,omitempty"`
	Sid                 string                          `json:"sid"`
	State               string                          `json:"state"`
	Timers              FetchConversationTimersResponse `json:"timers"`
	URL                 string                          `json:"url"`
}

// Fetch retrieves a conversation resource
// See https://www.twilio.com/docs/conversations/api/conversation-resource#fetch-a-conversation-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchConversationResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a conversation resource
// See https://www.twilio.com/docs/conversations/api/conversation-resource#fetch-a-conversation-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchConversationResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Conversations/{sid}",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"sid":        c.sid,
		},
	}

	response := &FetchConversationResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
