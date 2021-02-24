// Package messaging_configuration contains auto-generated files. DO NOT MODIFY
package messaging_configuration

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchMessagingConfigurationResponse defines the response fields for the retrieved messaging configuration
type FetchMessagingConfigurationResponse struct {
	AccountSid          string     `json:"account_sid"`
	Country             string     `json:"country"`
	DateCreated         time.Time  `json:"date_created"`
	DateUpdated         *time.Time `json:"date_updated,omitempty"`
	MessagingServiceSid string     `json:"messaging_service_sid"`
	ServiceSid          string     `json:"service_sid"`
	URL                 string     `json:"url"`
}

// Fetch retrieves a messaging configuration resource
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchMessagingConfigurationResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a messaging configuration resource
func (c Client) FetchWithContext(context context.Context) (*FetchMessagingConfigurationResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/MessagingConfigurations/{countryCode}",
		PathParams: map[string]string{
			"serviceSid":  c.serviceSid,
			"countryCode": c.countryCode,
		},
	}

	response := &FetchMessagingConfigurationResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
