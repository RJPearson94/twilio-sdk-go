// Package phone_numbers contains auto-generated files. DO NOT MODIFY
package phone_numbers

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreatePhoneNumberInput defines the input fields for creating a new phone number resource
type CreatePhoneNumberInput struct {
	PhoneNumberSid string `validate:"required" form:"PhoneNumberSid"`
}

// CreatePhoneNumberResponse defines the response fields for the created phone number
type CreatePhoneNumberResponse struct {
	AccountSid   string     `json:"account_sid"`
	Capabilities []string   `json:"capabilities"`
	CountryCode  string     `json:"country_code"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	PhoneNumber  string     `json:"phone_number"`
	ServiceSid   string     `json:"service_sid"`
	Sid          string     `json:"sid"`
	URL          string     `json:"url"`
}

// Create creates a new phone number
// See https://www.twilio.com/docs/sms/services/api/phonenumber-resource#create-a-phonenumber-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreatePhoneNumberInput) (*CreatePhoneNumberResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new phone number
// See https://www.twilio.com/docs/sms/services/api/phonenumber-resource#create-a-phonenumber-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreatePhoneNumberInput) (*CreatePhoneNumberResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/PhoneNumbers",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
		},
	}

	if input == nil {
		input = &CreatePhoneNumberInput{}
	}

	response := &CreatePhoneNumberResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
