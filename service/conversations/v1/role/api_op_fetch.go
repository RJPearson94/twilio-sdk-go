// Package role contains auto-generated files. DO NOT MODIFY
package role

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchRoleResponse defines the response fields for the retrieved role
type FetchRoleResponse struct {
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

// Fetch retrieves a role resource
// See https://www.twilio.com/docs/conversations/api/role-resource#fetch-a-role-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchRoleResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a role resource
// See https://www.twilio.com/docs/conversations/api/role-resource#fetch-a-role-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchRoleResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Roles/{sid}",
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	response := &FetchRoleResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
