// Package roles contains auto-generated files. DO NOT MODIFY
package roles

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreateRoleInput defines the input fields for creating a new role resource
type CreateRoleInput struct {
	FriendlyName string   `validate:"required" form:"FriendlyName"`
	Permissions  []string `validate:"required" form:"Permission"`
	Type         string   `validate:"required" form:"Type"`
}

// CreateRoleResponse defines the response fields for the created role
type CreateRoleResponse struct {
	AccountSid     string     `json:"account_sid"`
	ChatServiceSid string     `json:"chat_service_sid"`
	DateCreated    time.Time  `json:"date_created"`
	DateUpdated    *time.Time `json:"date_updated,omitempty"`
	FriendlyName   string     `json:"friendly_name"`
	Permissions    []string   `json:"permissions"`
	Sid            string     `json:"sid"`
	Type           string     `json:"type"`
	URL            string     `json:"url"`
}

// Create creates a new role
// See https://www.twilio.com/docs/conversations/api/role-resource#create-a-role-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateRoleInput) (*CreateRoleResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new role
// See https://www.twilio.com/docs/conversations/api/role-resource#create-a-role-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateRoleInput) (*CreateRoleResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Roles",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
		},
	}

	if input == nil {
		input = &CreateRoleInput{}
	}

	response := &CreateRoleResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
