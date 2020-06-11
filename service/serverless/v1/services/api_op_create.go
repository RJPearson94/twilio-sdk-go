// This is an autogenerated file. DO NOT MODIFY
package services

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type CreateServiceInput struct {
	UniqueName         string `validate:"required" form:"UniqueName"`
	FriendlyName       string `validate:"required" form:"FriendlyName"`
	IncludeCredentials *bool  `form:"IncludeCredentials,omitempty"`
	UiEditable         *bool  `form:"UiEditable,omitempty"`
}

type CreateServiceOutput struct {
	Sid                string     `json:"sid"`
	AccountSid         string     `json:"account_sid"`
	FriendlyName       string     `json:"friendly_name"`
	UniqueName         string     `json:"unique_name"`
	IncludeCredentials bool       `json:"include_credentials"`
	UiEditable         bool       `json:"ui_editable"`
	DateCreated        time.Time  `json:"date_created"`
	DateUpdated        *time.Time `json:"date_updated,omitempty"`
	URL                string     `json:"url"`
}

func (c Client) Create(input *CreateServiceInput) (*CreateServiceOutput, error) {
	return c.CreateWithContext(context.Background(), input)
}

func (c Client) CreateWithContext(context context.Context, input *CreateServiceInput) (*CreateServiceOutput, error) {
	op := client.Operation{
		HTTPMethod:  http.MethodPost,
		HTTPPath:    "/Services",
		ContentType: client.URLEncoded,
	}

	output := &CreateServiceOutput{}
	if err := c.client.Send(context, op, input, output); err != nil {
		return nil, err
	}
	return output, nil
}
