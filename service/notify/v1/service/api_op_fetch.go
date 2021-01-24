// Package service contains auto-generated files. DO NOT MODIFY
package service

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchServiceResponse defines the response fields for the retrieved service
type FetchServiceResponse struct {
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

// Fetch retrieves a service resource
// See https://www.twilio.com/docs/notify/api/service-resource#fetch-a-service-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchServiceResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a service resource
// See https://www.twilio.com/docs/notify/api/service-resource#fetch-a-service-resource for more details
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
