// Package users contains auto-generated files. DO NOT MODIFY
package users

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreateUserInput defines the input fields for creating a new user resource
type CreateUserInput struct {
	Attributes   *string `form:"Attributes,omitempty"`
	FriendlyName *string `form:"FriendlyName,omitempty"`
	Identity     string  `validate:"required" form:"Identity"`
	RoleSid      *string `form:"RoleSid,omitempty"`
}

// CreateUserResponse defines the response fields for the created user
type CreateUserResponse struct {
	AccountSid     string     `json:"account_sid"`
	Attributes     string     `json:"attributes"`
	ChatServiceSid string     `json:"chat_service_sid"`
	DateCreated    time.Time  `json:"date_created"`
	DateUpdated    *time.Time `json:"date_updated,omitempty"`
	FriendlyName   *string    `json:"friendly_name,omitempty"`
	Identity       string     `json:"identity"`
	IsNotifiable   *bool      `json:"is_notifiable,omitempty"`
	IsOnline       *bool      `json:"is_online,omitempty"`
	RoleSid        string     `json:"role_sid"`
	Sid            string     `json:"sid"`
	URL            string     `json:"url"`
}

// Create creates a new user
// See https://www.twilio.com/docs/conversations/api/user-resource#create-a-conversations-user for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateUserInput) (*CreateUserResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new user
// See https://www.twilio.com/docs/conversations/api/user-resource#create-a-conversations-user for more details
func (c Client) CreateWithContext(context context.Context, input *CreateUserInput) (*CreateUserResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Users",
		ContentType: client.URLEncoded,
	}

	if input == nil {
		input = &CreateUserInput{}
	}

	response := &CreateUserResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
