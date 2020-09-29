// Package sessions contains auto-generated files. DO NOT MODIFY
package sessions

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreateSessionInput defines the input fields for creating a new session resource
type CreateSessionInput struct {
	DateExpiry   *time.Time     `form:"DateExpiry,omitempty"`
	Mode         *string        `form:"Mode,omitempty"`
	Participants *[]interface{} `form:"Participants,omitempty"`
	Status       *string        `form:"Status,omitempty"`
	Ttl          *int           `form:"Ttl,omitempty"`
	UniqueName   *string        `form:"UniqueName,omitempty"`
}

// CreateSessionResponse defines the response fields for the created session
type CreateSessionResponse struct {
	AccountSid          string     `json:"account_sid"`
	ClosedReason        *string    `json:"closed_reason,omitempty"`
	DateCreated         time.Time  `json:"date_created"`
	DateEnded           *time.Time `json:"date_ended,omitempty"`
	DateExpiry          *time.Time `json:"date_expiry,omitempty"`
	DateLastInteraction *time.Time `json:"date_last_interaction,omitempty"`
	DateStarted         *time.Time `json:"date_started,omitempty"`
	DateUpdated         *time.Time `json:"date_updated,omitempty"`
	Mode                *string    `json:"mode,omitempty"`
	ServiceSid          string     `json:"service_sid"`
	Sid                 string     `json:"sid"`
	Status              *string    `json:"status,omitempty"`
	Ttl                 *int       `json:"ttl,omitempty"`
	URL                 string     `json:"url"`
	UniqueName          string     `json:"unique_name"`
}

// Create creates a new session
// See https://www.twilio.com/docs/proxy/api/session#create-a-session-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateSessionInput) (*CreateSessionResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new session
// See https://www.twilio.com/docs/proxy/api/session#create-a-session-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateSessionInput) (*CreateSessionResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Sessions",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
		},
	}

	if input == nil {
		input = &CreateSessionInput{}
	}

	response := &CreateSessionResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
