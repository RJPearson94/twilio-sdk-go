// Package message contains auto-generated files. DO NOT MODIFY
package message

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// FetchMessageResponse defines the response fields for the retrieved message
type FetchMessageResponse struct {
	APIVersion          string             `json:"api_version"`
	AccountSid          string             `json:"account_sid"`
	Body                string             `json:"body"`
	DateCreated         utils.RFC2822Time  `json:"date_created"`
	DateSent            utils.RFC2822Time  `json:"date_sent"`
	DateUpdated         *utils.RFC2822Time `json:"date_updated,omitempty"`
	Direction           string             `json:"direction"`
	ErrorCode           *int               `json:"error_code,omitempty"`
	ErrorMessage        *string            `json:"error_message,omitempty"`
	From                *string            `json:"from,omitempty"`
	MessagingServiceSid *string            `json:"messaging_service_sid,omitempty"`
	NumMedia            string             `json:"num_media"`
	NumSegments         string             `json:"num_segments"`
	Price               *string            `json:"price,omitempty"`
	PriceUnit           string             `json:"price_unit"`
	Sid                 string             `json:"sid"`
	Status              string             `json:"status"`
	To                  string             `json:"to"`
}

// Fetch retrieves a message resource
// See https://www.twilio.com/docs/sms/api/message-resource#fetch-a-message-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchMessageResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a message resource
// See https://www.twilio.com/docs/sms/api/message-resource#fetch-a-message-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchMessageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Accounts/{accountSid}/Messages/{sid}.json",
		PathParams: map[string]string{
			"accountSid": c.accountSid,
			"sid":        c.sid,
		},
	}

	response := &FetchMessageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
