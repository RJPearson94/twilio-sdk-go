// Package queues contains auto-generated files. DO NOT MODIFY
package queues

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// CreateQueueInput defines input fields for creating a new queue
type CreateQueueInput struct {
	FriendlyName string `validate:"required" form:"FriendlyName"`
	MaxSize      *int   `form:"MaxSize,omitempty"`
}

// CreateQueueResponse defines the response fields for creating a new queue
type CreateQueueResponse struct {
	AccountSid      string             `json:"account_sid"`
	AverageWaitTime int                `json:"average_wait_time"`
	CurrentSize     int                `json:"current_size"`
	DateCreated     utils.RFC2822Time  `json:"date_created"`
	DateUpdated     *utils.RFC2822Time `json:"date_updated,omitempty"`
	FriendlyName    string             `json:"friendly_name"`
	MaxSize         int                `json:"max_size"`
	Sid             string             `json:"sid"`
}

// Create creates a new queue resource
// See https://www.twilio.com/docs/voice/api/queue-resource#create-a-queue-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateQueueInput) (*CreateQueueResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new queue resource
// See https://www.twilio.com/docs/voice/api/queue-resource#create-a-queue-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateQueueInput) (*CreateQueueResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Accounts/{accountSid}/Queues.json",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"accountSid": c.accountSid,
		},
	}

	if input == nil {
		input = &CreateQueueInput{}
	}

	response := &CreateQueueResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
