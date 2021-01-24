// Package service contains auto-generated files. DO NOT MODIFY
package service

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// UpdateServiceInput defines input fields for updating a service resource
type UpdateServiceInput struct {
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

// UpdateServiceResponse defines the response fields for the updated service
type UpdateServiceResponse struct {
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

// Update modifies a service resource
// See https://www.twilio.com/docs/notify/api/service-resource#update-a-service-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateServiceInput) (*UpdateServiceResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a service resource
// See https://www.twilio.com/docs/notify/api/service-resource#update-a-service-resource for more details
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
