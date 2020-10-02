// Package recording contains auto-generated files. DO NOT MODIFY
package recording

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// UpdateRecordingInput defines input fields for updating a recording
type UpdateRecordingInput struct {
	PauseBehavior *string `form:"PauseBehavior,omitempty"`
	Status        string  `validate:"required" form:"Status"`
}

// UpdateRecordingResponse defines the response fields for the updated recording
type UpdateRecordingResponse struct {
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

// Update modifies a recording resource
// See https://www.twilio.com/docs/voice/api/recording#update-a-recording-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateRecordingInput) (*UpdateRecordingResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a recording resource
// See https://www.twilio.com/docs/voice/api/recording#update-a-recording-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateRecordingInput) (*UpdateRecordingResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Accounts/{accountSid}/Recordings/{sid}.json",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"accountSid": c.accountSid,
			"sid":        c.sid,
		},
	}

	if input == nil {
		input = &UpdateRecordingInput{}
	}

	response := &UpdateRecordingResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
