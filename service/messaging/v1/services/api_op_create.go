// Package services contains auto-generated files. DO NOT MODIFY
package services

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreateServiceInput defines the input fields for creating a new messaging resource
type CreateServiceInput struct {
	AreaCodeGeomatch      *bool   `form:"AreaCodeGeomatch,omitempty"`
	FallbackMethod        *string `form:"FallbackMethod,omitempty"`
	FallbackToLongCode    *bool   `form:"FallbackToLongCode,omitempty"`
	FallbackURL           *string `form:"FallbackUrl,omitempty"`
	FriendlyName          string  `validate:"required" form:"FriendlyName"`
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

// CreateServiceResponse defines the response fields for the created messaging
type CreateServiceResponse struct {
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

// Create creates a new messaging
// See https://www.twilio.com/docs/sms/services/api#create-a-service-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateServiceInput) (*CreateServiceResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new messaging
// See https://www.twilio.com/docs/sms/services/api#create-a-service-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateServiceInput) (*CreateServiceResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services",
		ContentType: client.URLEncoded,
	}

	if input == nil {
		input = &CreateServiceInput{}
	}

	response := &CreateServiceResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
