// Package user contains auto-generated files. DO NOT MODIFY
package user

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// UpdateUserInput defines input fields for updating a user resource
type UpdateUserInput struct {
	Attributes   *string `form:"Attributes,omitempty"`
	FriendlyName *string `form:"FriendlyName,omitempty"`
	RoleSid      *string `form:"RoleSid,omitempty"`
}

// UpdateUserResponse defines the response fields for the updated user
type UpdateUserResponse struct {
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

// Update modifies a user resource
// See https://www.twilio.com/docs/conversations/api/user-resource#update-a-conversationuser-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateUserInput) (*UpdateUserResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a user resource
// See https://www.twilio.com/docs/conversations/api/user-resource#update-a-conversationuser-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateUserInput) (*UpdateUserResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Users/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	if input == nil {
		input = &UpdateUserInput{}
	}

	response := &UpdateUserResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
