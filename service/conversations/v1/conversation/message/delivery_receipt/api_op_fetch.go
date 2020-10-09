// Package delivery_receipt contains auto-generated files. DO NOT MODIFY
package delivery_receipt

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchDeliveryReceiptResponse defines the response fields for the retrieved delivery receipt
type FetchDeliveryReceiptResponse struct {
	AccountSid        *string    `json:"account_sid,omitempty"`
	ChannelMessageSid string     `json:"channel_message_sid"`
	ConversationSid   string     `json:"conversation_sid"`
	DateCreated       time.Time  `json:"date_created"`
	DateUpdated       *time.Time `json:"date_updated,omitempty"`
	ErrorCode         *int       `json:"error_code,omitempty"`
	MessageSid        string     `json:"message_sid"`
	ParticipantSid    string     `json:"participant_sid"`
	Sid               string     `json:"sid"`
	Status            string     `json:"status"`
	URL               string     `json:"url"`
}

// Fetch retrieves a delivery receipt resource
// See https://www.twilio.com/docs/conversations/api/receipt-resource#fetch-a-conversationmessagereceipt-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchDeliveryReceiptResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a delivery receipt resource
// See https://www.twilio.com/docs/conversations/api/receipt-resource#fetch-a-conversationmessagereceipt-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchDeliveryReceiptResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Conversations/{conversationSid}/Messages/{messageSid}/Receipts/{sid}",
		PathParams: map[string]string{
			"conversationSid": c.conversationSid,
			"messageSid":      c.messageSid,
			"sid":             c.sid,
		},
	}

	response := &FetchDeliveryReceiptResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
