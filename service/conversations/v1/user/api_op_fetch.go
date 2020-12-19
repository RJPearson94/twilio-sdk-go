// Package user contains auto-generated files. DO NOT MODIFY
package user

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchUserResponse defines the response fields for the retrieved user
type FetchUserResponse struct {
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

// Fetch retrieves a user resource
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchUserResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a user resource
func (c Client) FetchWithContext(context context.Context) (*FetchUserResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Users/{sid}",
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	response := &FetchUserResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
