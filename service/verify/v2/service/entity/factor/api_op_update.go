// Package factor contains auto-generated files. DO NOT MODIFY
package factor

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type UpdateFactorConfigInput struct {
	NotificationToken *string `form:"NotificationToken,omitempty"`
	SdkVersion        *string `form:"SdkVersion,omitempty"`
}

// UpdateFactorInput defines input fields for updating a factor resource
type UpdateFactorInput struct {
	AuthPayload  *string                  `form:"AuthPayload,omitempty"`
	Config       *UpdateFactorConfigInput `form:"Config,omitempty"`
	FriendlyName *string                  `form:"FriendlyName,omitempty"`
}

type UpdateFactorConfigResponse struct {
	AppId                string  `json:"app_id"`
	NotificationPlatform string  `json:"notification_platform"`
	NotificationToken    string  `json:"notification_token"`
	SdkVersion           *string `json:"sdk_version,omitempty"`
}

// UpdateFactorResponse defines the response fields for the updated factor
type UpdateFactorResponse struct {
	AccountSid   string                     `json:"account_sid"`
	Config       UpdateFactorConfigResponse `json:"config"`
	DateCreated  time.Time                  `json:"date_created"`
	DateUpdated  *time.Time                 `json:"date_updated,omitempty"`
	EntitySid    string                     `json:"entity_sid"`
	FactorType   string                     `json:"factor_type"`
	FriendlyName string                     `json:"friendly_name"`
	Identity     string                     `json:"identity"`
	ServiceSid   string                     `json:"service_sid"`
	Sid          string                     `json:"sid"`
	Status       string                     `json:"status"`
	URL          string                     `json:"url"`
}

// Update modifies a factor resource
// See https://www.twilio.com/docs/verify/api/factor#update-a-factor-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
// This resource is currently in beta and subject to change. Please use with caution
func (c Client) Update(input *UpdateFactorInput) (*UpdateFactorResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a factor resource
// See https://www.twilio.com/docs/verify/api/factor#update-a-factor-resource for more details
// This resource is currently in beta and subject to change. Please use with caution
func (c Client) UpdateWithContext(context context.Context, input *UpdateFactorInput) (*UpdateFactorResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Entities/{identity}/Factors/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"identity":   c.identity,
			"sid":        c.sid,
		},
	}

	if input == nil {
		input = &UpdateFactorInput{}
	}

	response := &UpdateFactorResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
