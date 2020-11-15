// Package factors contains auto-generated files. DO NOT MODIFY
package factors

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreateFactorInput defines the input fields for creating a new factor
type CreateFactorInput struct {
	BindingAlg                 string  `validate:"required" form:"Binding.Alg"`
	BindingPublicKey           string  `validate:"required" form:"Binding.PublicKey"`
	ConfigAppId                string  `validate:"required" form:"Config.AppId"`
	ConfigNotificationPlatform string  `validate:"required" form:"Config.NotificationPlatform"`
	ConfigNotificationToken    string  `validate:"required" form:"Config.NotificationToken"`
	ConfigSdkVersion           *string `form:"Config.SdkVersion,omitempty"`
	FactorType                 string  `validate:"required" form:"FactorType"`
	FriendlyName               string  `validate:"required" form:"FriendlyName"`
}

type CreateFactorConfigResponse struct {
	AppId                string  `json:"app_id"`
	NotificationPlatform string  `json:"notification_platform"`
	NotificationToken    string  `json:"notification_token"`
	SdkVersion           *string `json:"sdk_version,omitempty"`
}

// CreateFactorResponse defines the response fields for the created factor
type CreateFactorResponse struct {
	AccountSid   string                     `json:"account_sid"`
	Config       CreateFactorConfigResponse `json:"config"`
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

// Create creates a new factor
// See https://www.twilio.com/docs/verify/api/factor#create-a-factor-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateFactorInput) (*CreateFactorResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new factor
// See https://www.twilio.com/docs/verify/api/factor#create-a-factor-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateFactorInput) (*CreateFactorResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Entities/{identity}/Factors",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"identity":   c.identity,
		},
	}

	if input == nil {
		input = &CreateFactorInput{}
	}

	response := &CreateFactorResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
