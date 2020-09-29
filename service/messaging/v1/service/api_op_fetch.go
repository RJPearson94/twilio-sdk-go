// Package service contains auto-generated files. DO NOT MODIFY
package service

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchServiceResponse defines the response fields for the retrieved messaging
type FetchServiceResponse struct {
	AccountSid            string     `json:"account_sid"`
	AreaCodeGeomatch      bool       `json:"area_code_geomatch"`
	DateCreated           time.Time  `json:"date_created"`
	DateUpdated           *time.Time `json:"date_updated,omitempty"`
	FallbackMethod        string     `json:"fallback_method"`
	FallbackToLongCode    bool       `json:"fallback_to_long_code"`
	FallbackURL           *string    `json:"fallback_url,omitempty"`
	FriendlyName          string     `json:"friendly_name"`
	InboundMethod         string     `json:"inbound_method"`
	InboundRequestURL     *string    `json:"inbound_request_url,omitempty"`
	MmsConverter          bool       `json:"mms_converter"`
	ScanMessageContent    string     `json:"scan_message_content"`
	Sid                   string     `json:"sid"`
	SmartEncoding         bool       `json:"smart_encoding"`
	StatusCallback        *string    `json:"status_callback,omitempty"`
	StickySender          bool       `json:"sticky_sender"`
	SynchronousValidation bool       `json:"synchronous_validation"`
	URL                   string     `json:"url"`
	ValidityPeriod        int        `json:"validity_period"`
}

// Fetch retrieves a messaging resource
// See https://www.twilio.com/docs/sms/services/api#fetch-a-service-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchServiceResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a messaging resource
// See https://www.twilio.com/docs/sms/services/api#fetch-a-service-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchServiceResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{sid}",
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	response := &FetchServiceResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
