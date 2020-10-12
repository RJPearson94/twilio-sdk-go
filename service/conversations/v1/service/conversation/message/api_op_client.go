// Package message contains auto-generated files. DO NOT MODIFY
package message

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/service/conversation/message/delivery_receipt"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/service/conversation/message/delivery_receipts"
)

// Client for managing a specific message resource
// See https://www.twilio.com/docs/conversations/api/conversation-message-resource for more details
type Client struct {
	client *client.Client

	conversationSid string
	serviceSid      string
	sid             string

	DeliveryReceipt  func(string) *delivery_receipt.Client
	DeliveryReceipts *delivery_receipts.Client
}

// ClientProperties are the properties required to manage the message resources
type ClientProperties struct {
	ConversationSid string
	ServiceSid      string
	Sid             string
}

// New creates a new instance of the message client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		conversationSid: properties.ConversationSid,
		serviceSid:      properties.ServiceSid,
		sid:             properties.Sid,

		DeliveryReceipt: func(deliveryReceiptSid string) *delivery_receipt.Client {
			return delivery_receipt.New(client, delivery_receipt.ClientProperties{
				ConversationSid: properties.ConversationSid,
				MessageSid:      properties.Sid,
				ServiceSid:      properties.ServiceSid,
				Sid:             deliveryReceiptSid,
			})
		},
		DeliveryReceipts: delivery_receipts.New(client, delivery_receipts.ClientProperties{
			ConversationSid: properties.ConversationSid,
			MessageSid:      properties.Sid,
			ServiceSid:      properties.ServiceSid,
		}),
	}
}
