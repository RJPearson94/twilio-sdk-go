// Package service contains auto-generated files. DO NOT MODIFY
package service

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// UpdateServiceInput defines input fields for updating a messaging resource
type UpdateServiceInput struct {
	AreaCodeGeomatch      *bool   `form:"AreaCodeGeomatch,omitempty"`
	FallbackMethod        *string `form:"FallbackMethod,omitempty"`
	FallbackToLongCode    *bool   `form:"FallbackToLongCode,omitempty"`
	FallbackURL           *string `form:"FallbackUrl,omitempty"`
	FriendlyName          *string `form:"FriendlyName,omitempty"`
	InboundMethod         *string `form:"InboundMethod,omitempty"`
	InboundRequestURL     *string `form:"InboundRequestUrl,omitempty"`
	MmsConverter          *bool   `form:"MmsConverter,omitempty"`
	ScanMessageContent    *string `form:"ScanMessageContent,omitempty"`
	SmartEncoding         *bool   `form:"SmartEncoding,omitempty"`
	StatusCallback        *string `form:"StatusCallback,omitempty"`
	StickySender          *bool   `form:"StickySender,omitempty"`
	SynchronousValidation *bool   `form:"SynchronousValidation,omitempty"`
	ValidityPeriod        *int    `form:"ValidityPeriod,omitempty"`
}

// UpdateServiceResponse defines the response fields for the updated messaging
type UpdateServiceResponse struct {
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

// Update modifies a messaging resource
// See https://www.twilio.com/docs/sms/services/api#update-a-service-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateServiceInput) (*UpdateServiceResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a messaging resource
// See https://www.twilio.com/docs/sms/services/api#update-a-service-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateServiceInput) (*UpdateServiceResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	if input == nil {
		input = &UpdateServiceInput{}
	}

	response := &UpdateServiceResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
