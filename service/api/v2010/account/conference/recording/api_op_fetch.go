// Package recording contains auto-generated files. DO NOT MODIFY
package recording

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// FetchRecordingResponse defines the response fields for retrieving a recording
type FetchRecordingResponse struct {
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

// Fetch retrieves a recording resource
// See https://www.twilio.com/docs/voice/api/recording#fetch-a-recording-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchRecordingResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a recording resource
// See https://www.twilio.com/docs/voice/api/recording#fetch-a-recording-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchRecordingResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Accounts/{accountSid}/Conferences/{conferenceSid}/Recordings/{sid}.json",
		PathParams: map[string]string{
			"accountSid":    c.accountSid,
			"conferenceSid": c.conferenceSid,
			"sid":           c.sid,
		},
	}

	response := &FetchRecordingResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
