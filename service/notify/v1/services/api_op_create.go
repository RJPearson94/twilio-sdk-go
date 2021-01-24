// Package services contains auto-generated files. DO NOT MODIFY
package services

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreateServiceInput defines the input fields for creating a new service resource
type CreateServiceInput struct {
	APNCredentialSid                      *string `form:"ApnCredentialSid,omitempty"`
	DefaultAPNNotificationProtocolVersion *string `form:"DefaultApnNotificationProtocolVersion,omitempty"`
	DefaultFCMNotificationProtocolVersion *string `form:"DefaultFcmNotificationProtocolVersion,omitempty"`
	DeliveryCallbackEnabled               *bool   `form:"DeliveryCallbackEnabled,omitempty"`
	DeliveryCallbackURL                   *string `form:"DeliveryCallbackUrl,omitempty"`
	FCMCredentialSid                      *string `form:"FcmCredentialSid,omitempty"`
	FriendlyName                          *string `form:"FriendlyName,omitempty"`
	LogEnabled                            *bool   `form:"LogEnabled,omitempty"`
	MessagingServiceSid                   *string `form:"MessagingServiceSid,omitempty"`
}

// CreateServiceResponse defines the response fields for the created service
type CreateServiceResponse struct {
	APNCredentialSid                      *string    `json:"apn_credential_sid,omitempty"`
	AccountSid                            string     `json:"account_sid"`
	DateCreated                           time.Time  `json:"date_created"`
	DateUpdated                           *time.Time `json:"date_updated,omitempty"`
	DefaultAPNNotificationProtocolVersion string     `json:"default_apn_notification_protocol_version"`
	DefaultFCMNotificationProtocolVersion string     `json:"default_fcm_notification_protocol_version"`
	DeliveryCallbackEnabled               bool       `json:"delivery_callback_enabled"`
	DeliveryCallbackURL                   *string    `json:"delivery_callback_url,omitempty"`
	FCMCredentialSid                      *string    `json:"fcm_credential_sid,omitempty"`
	FriendlyName                          *string    `json:"friendly_name,omitempty"`
	LogEnabled                            bool       `json:"log_enabled"`
	MessagingServiceSid                   *string    `json:"messaging_service_sid,omitempty"`
	Sid                                   string     `json:"sid"`
	URL                                   string     `json:"url"`
}

// Create creates a new service
// See https://www.twilio.com/docs/notify/api/service-resource#create-a-service-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateServiceInput) (*CreateServiceResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new service
// See https://www.twilio.com/docs/notify/api/service-resource#create-a-service-resource for more details
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
