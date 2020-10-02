// Package recordings contains auto-generated files. DO NOT MODIFY
package recordings

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// CreateRecordingInput defines input fields for creating a new recording
type CreateRecordingInput struct {
	RecordingChannels             *string   `form:"RecordingChannels,omitempty"`
	RecordingStatusCallback       *string   `form:"RecordingStatusCallback,omitempty"`
	RecordingStatusCallbackEvents *[]string `form:"RecordingStatusCallbackEvent,omitempty"`
	RecordingStatusCallbackMethod *string   `form:"RecordingStatusCallbackMethod,omitempty"`
	Trim                          *string   `form:"Trim,omitempty"`
}

// CreateRecordingResponse defines the response fields for creating a new recording resource
type CreateRecordingResponse struct {
	APIVersion        string                  `json:"api_version"`
	CallSid           string                  `json:"call_sid"`
	Channels          int                     `json:"channels"`
	ConferenceSid     *string                 `json:"conference_sid,omitempty"`
	DateCreated       utils.RFC2822Time       `json:"date_created"`
	DateUpdated       *utils.RFC2822Time      `json:"date_updated,omitempty"`
	Duration          *string                 `json:"duration,omitempty"`
	EncryptionDetails *map[string]interface{} `json:"encryption_details,omitempty"`
	ErrorCode         *string                 `json:"error_code,omitempty"`
	Price             *string                 `json:"price,omitempty"`
	PriceUnit         *string                 `json:"price_unit,omitempty"`
	Sid               string                  `json:"sid"`
	Source            string                  `json:"source"`
	StartTime         utils.RFC2822Time       `json:"start_time"`
	Status            string                  `json:"status"`
}

// Create creates a new recording resource
// See https://www.twilio.com/docs/voice/api/recording#create-a-recording-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateRecordingInput) (*CreateRecordingResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new recording resource
// See https://www.twilio.com/docs/voice/api/recording#create-a-recording-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateRecordingInput) (*CreateRecordingResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Accounts/{accountSid}/Calls/{callSid}/Recordings.json",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"accountSid": c.accountSid,
			"callSid":    c.callSid,
		},
	}

	if input == nil {
		input = &CreateRecordingInput{}
	}

	response := &CreateRecordingResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
