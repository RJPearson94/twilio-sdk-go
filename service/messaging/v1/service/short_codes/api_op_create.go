// Package short_codes contains auto-generated files. DO NOT MODIFY
package short_codes

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreateShortCodeInput defines the input fields for creating a new short code resource
type CreateShortCodeInput struct {
	ShortCodeSid string `validate:"required" form:"ShortCodeSid"`
}

// CreateShortCodeResponse defines the response fields for the created short code
type CreateShortCodeResponse struct {
	AccountSid   string     `json:"account_sid"`
	Capabilities []string   `json:"capabilities"`
	CountryCode  string     `json:"country_code"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	ServiceSid   string     `json:"service_sid"`
	ShortCode    string     `json:"short_code"`
	Sid          string     `json:"sid"`
	URL          string     `json:"url"`
}

// Create creates a new short code
// See https://www.twilio.com/docs/sms/services/api/shortcode-resource#create-a-shortcode-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateShortCodeInput) (*CreateShortCodeResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new short code
// See https://www.twilio.com/docs/sms/services/api/shortcode-resource#create-a-shortcode-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateShortCodeInput) (*CreateShortCodeResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/ShortCodes",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
		},
	}

	if input == nil {
		input = &CreateShortCodeInput{}
	}

	response := &CreateShortCodeResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
