// Package field_type contains auto-generated files. DO NOT MODIFY
package field_type

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchFieldTypeResponse defines the response fields for the retrieved field type
type FetchFieldTypeResponse struct {
	AccountSid   string     `json:"account_sid"`
	AssistantSid string     `json:"assistant_sid"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	FriendlyName *string    `json:"friendly_name,omitempty"`
	Sid          string     `json:"sid"`
	URL          string     `json:"url"`
	UniqueName   string     `json:"unique_name"`
}

// Fetch retrieves a field type resource
// See https://www.twilio.com/docs/autopilot/api/field-type#fetch-a-fieldtype-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchFieldTypeResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a field type resource
// See https://www.twilio.com/docs/autopilot/api/field-type#fetch-a-fieldtype-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchFieldTypeResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Assistants/{assistantSid}/FieldTypes/{sid}",
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
			"sid":          c.sid,
		},
	}

	response := &FetchFieldTypeResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
