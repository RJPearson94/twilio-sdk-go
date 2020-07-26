// This is an autogenerated file. DO NOT MODIFY
package message

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// Resource/ response properties for the retrieved message
type GetMessageResponse struct {
	// The api version which was responsible for dealing with the message
	APIVersion string `json:"api_version"`
	// The SID of the account which sent the message
	AccountSid string `json:"account_sid"`
	// The message contents. Can be plain text or twiML
	Body string `json:"body"`
	// The date and time (in RFC2822 format) when the resource was created
	DateCreated utils.RFC2822Time `json:"date_created"`
	// The date and time (in RFC2822 format) the message was sent
	DateSent utils.RFC2822Time `json:"date_sent"`
	// The date and time (in RFC2822 format) when the resource was last updated
	DateUpdated *utils.RFC2822Time `json:"date_updated,omitempty"`
	// The message direction
	Direction string `json:"direction"`
	// The Twilio error code for the issue that occurred
	ErrorCode *int `json:"error_code,omitempty"`
	// The description of the error that occurred
	ErrorMessage *string `json:"error_message,omitempty"`
	// The sender of the message
	From *string `json:"from,omitempty"`
	// The messaging service that is associate with the message
	MessagingServiceSid *string `json:"messaging_service_sid,omitempty"`
	// The number of media attachments
	NumMedia string `json:"num_media"`
	// The number of segments the message was split into
	NumSegments string `json:"num_segments"`
	// The cost of the message
	Price *string `json:"price,omitempty"`
	// The currency format for the message price
	PriceUnit string `json:"price_unit"`
	// The unique alphanumeric string for the resource
	Sid string `json:"sid"`
	// The current status of the message
	Status string `json:"status"`
	// The intended recipient of the message
	To string `json:"to"`
}

// Retrieve a message resource
// See https://www.twilio.com/docs/sms/api/message-resource#fetch-a-message-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Get() (*GetMessageResponse, error) {
	return c.GetWithContext(context.Background())
}

// Retrieve a message resource
// See https://www.twilio.com/docs/sms/api/message-resource#fetch-a-message-resource for more details
func (c Client) GetWithContext(context context.Context) (*GetMessageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Accounts/{accountSid}/Messages/{sid}.json",
		PathParams: map[string]string{
			"accountSid": c.accountSid,
			"sid":        c.sid,
		},
	}

	response := &GetMessageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
