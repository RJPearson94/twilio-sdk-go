// Package notifications contains auto-generated files. DO NOT MODIFY
package notifications

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreateNotificationInput defines the input fields for creating a new service notification
type CreateNotificationInput struct {
	APN                 *string   `form:"Apn,omitempty"`
	Action              *string   `form:"Action,omitempty"`
	Body                *string   `form:"Body,omitempty"`
	Data                *string   `form:"Data,omitempty"`
	DeliveryCallbackURL *string   `form:"DeliveryCallbackUrl,omitempty"`
	FCM                 *string   `form:"Fcm,omitempty"`
	Identities          *[]string `form:"Identity,omitempty"`
	Priority            *string   `form:"Priority,omitempty"`
	SMS                 *string   `form:"Sms,omitempty"`
	Sound               *string   `form:"Sound,omitempty"`
	Tags                *[]string `form:"Tag,omitempty"`
	Title               *string   `form:"Title,omitempty"`
	ToBindings          *[]string `form:"ToBinding,omitempty"`
	Ttl                 *int      `form:"Ttl,omitempty"`
}

// CreateNotificationResponse defines the response fields for the retrieved service notification
type CreateNotificationResponse struct {
	APN         *map[string]interface{} `json:"apn,omitempty"`
	AccountSid  string                  `json:"account_sid"`
	Action      *string                 `json:"action,omitempty"`
	Body        *string                 `json:"body,omitempty"`
	Data        *map[string]interface{} `json:"data,omitempty"`
	DateCreated time.Time               `json:"date_created"`
	FCM         *map[string]interface{} `json:"fcm,omitempty"`
	Identities  []string                `json:"identities"`
	Priority    string                  `json:"priority"`
	SMS         *map[string]interface{} `json:"sms,omitempty"`
	ServiceSid  string                  `json:"service_sid"`
	Sid         string                  `json:"sid"`
	Sound       *string                 `json:"sound,omitempty"`
	TTL         int                     `json:"ttl"`
	Tags        []string                `json:"tags"`
}

// Create creates a service notification resource
// See https://www.twilio.com/docs/notify/api/notification-resource#create-a-notification-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateNotificationInput) (*CreateNotificationResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a service notification resource
// See https://www.twilio.com/docs/notify/api/notification-resource#create-a-notification-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateNotificationInput) (*CreateNotificationResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Notifications",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
		},
	}

	if input == nil {
		input = &CreateNotificationInput{}
	}

	response := &CreateNotificationResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
