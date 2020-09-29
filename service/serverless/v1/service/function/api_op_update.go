// Package function contains auto-generated files. DO NOT MODIFY
package function

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// UpdateFunctionInput defines input fields for updating a function resource
type UpdateFunctionInput struct {
	FriendlyName string `validate:"required" form:"FriendlyName"`
}

// UpdateFunctionResponse defines the response fields for the updated function
type UpdateFunctionResponse struct {
	AccountSid   string     `json:"account_sid"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	FriendlyName string     `json:"friendly_name"`
	ServiceSid   string     `json:"service_sid"`
	Sid          string     `json:"sid"`
	URL          string     `json:"url"`
}

// Update modifies a function resource
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/function#update-a-function-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateFunctionInput) (*UpdateFunctionResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a function resource
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/function#update-a-function-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateFunctionInput) (*UpdateFunctionResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Functions/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"sid":        c.sid,
		},
	}

	if input == nil {
		input = &UpdateFunctionInput{}
	}

	response := &UpdateFunctionResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
