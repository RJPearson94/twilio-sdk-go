// Package role contains auto-generated files. DO NOT MODIFY
package role

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// UpdateRoleInput defines input fields for updating a role resource
type UpdateRoleInput struct {
	Permissions []string `validate:"required" form:"Permission"`
}

// UpdateRoleResponse defines the response fields for the updated role
type UpdateRoleResponse struct {
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

// Update modifies a role resource
// See https://www.twilio.com/docs/conversations/api/role-resource#update-a-role-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateRoleInput) (*UpdateRoleResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a role resource
// See https://www.twilio.com/docs/conversations/api/role-resource#update-a-role-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateRoleInput) (*UpdateRoleResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Roles/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	if input == nil {
		input = &UpdateRoleInput{}
	}

	response := &UpdateRoleResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
