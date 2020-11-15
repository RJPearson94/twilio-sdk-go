// Package factor contains auto-generated files. DO NOT MODIFY
package factor

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type FetchFactorConfigResponse struct {
	AppId                string  `json:"app_id"`
	NotificationPlatform string  `json:"notification_platform"`
	NotificationToken    string  `json:"notification_token"`
	SdkVersion           *string `json:"sdk_version,omitempty"`
}

// FetchFactorResponse defines the response fields for the retrieved factor
type FetchFactorResponse struct {
	AccountSid   string                    `json:"account_sid"`
	Config       FetchFactorConfigResponse `json:"config"`
	DateCreated  time.Time                 `json:"date_created"`
	DateUpdated  *time.Time                `json:"date_updated,omitempty"`
	EntitySid    string                    `json:"entity_sid"`
	FactorType   string                    `json:"factor_type"`
	FriendlyName string                    `json:"friendly_name"`
	Identity     string                    `json:"identity"`
	ServiceSid   string                    `json:"service_sid"`
	Sid          string                    `json:"sid"`
	Status       string                    `json:"status"`
	URL          string                    `json:"url"`
}

// Fetch retrieves a factor resource
// See https://www.twilio.com/docs/verify/api/factor#fetch-a-factor-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchFactorResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a factor resource
// See https://www.twilio.com/docs/verify/api/factor#fetch-a-factor-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchFactorResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Entities/{identity}/Factors/{sid}",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"identity":   c.identity,
			"sid":        c.sid,
		},
	}

	response := &FetchFactorResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
