// Package message contains auto-generated files. DO NOT MODIFY
package message

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type FetchMessageDeliveryResponse struct {
	Delivered   string `json:"delivered"`
	Failed      string `json:"failed"`
	Read        string `json:"read"`
	Sent        string `json:"sent"`
	Total       int    `json:"total"`
	Undelivered string `json:"undelivered"`
}

type FetchMessageMediaResponse struct {
	ContentType string `json:"content_type"`
	Filename    string `json:"filename"`
	Sid         string `json:"sid"`
	Size        int    `json:"size"`
}

// FetchMessageResponse defines the response fields for the retrieved message
type FetchMessageResponse struct {
	AccountSid      string                        `json:"account_sid"`
	Attributes      string                        `json:"attributes"`
	Author          string                        `json:"author"`
	Body            *string                       `json:"body,omitempty"`
	ConversationSid string                        `json:"conversation_sid"`
	DateCreated     time.Time                     `json:"date_created"`
	DateUpdated     *time.Time                    `json:"date_updated,omitempty"`
	Delivery        *FetchMessageDeliveryResponse `json:"delivery,omitempty"`
	Index           int                           `json:"index"`
	Media           *[]FetchMessageMediaResponse  `json:"media,omitempty"`
	ParticipantSid  *string                       `json:"participant_sid,omitempty"`
	Sid             string                        `json:"sid"`
	URL             string                        `json:"url"`
}

// Fetch retrieves a message resource
// See https://www.twilio.com/docs/conversations/api/conversation-message-resource#fetch-a-conversationmessage-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchMessageResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a message resource
// See https://www.twilio.com/docs/conversations/api/conversation-message-resource#fetch-a-conversationmessage-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchMessageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Conversations/{conversationSid}/Messages/{sid}",
		PathParams: map[string]string{
			"conversationSid": c.conversationSid,
			"sid":             c.sid,
		},
	}

	response := &FetchMessageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
