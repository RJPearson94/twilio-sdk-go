// Package model_build contains auto-generated files. DO NOT MODIFY
package model_build

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// UpdateModelBuildInput defines the input fields for updating a model build
type UpdateModelBuildInput struct {
	UniqueName *string `form:"UniqueName,omitempty"`
}

// UpdateModelBuildResponse defines the response fields for the updated model build
type UpdateModelBuildResponse struct {
	AccountSid    string     `json:"account_sid"`
	AssistantSid  string     `json:"assistant_sid"`
	BuildDuration *int       `json:"build_duration,omitempty"`
	DateCreated   time.Time  `json:"date_created"`
	DateUpdated   *time.Time `json:"date_updated,omitempty"`
	ErrorCode     *int       `json:"error_code,omitempty"`
	Sid           string     `json:"sid"`
	Status        string     `json:"status"`
	URL           string     `json:"url"`
	UniqueName    string     `json:"unique_name"`
}

// Update modifies an model build resource
// See https://www.twilio.com/docs/autopilot/api/model-build#update-a-modelbuild-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateModelBuildInput) (*UpdateModelBuildResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies an model build resource
// See https://www.twilio.com/docs/autopilot/api/model-build#update-a-modelbuild-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateModelBuildInput) (*UpdateModelBuildResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Assistants/{assistantSid}/ModelBuilds/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
			"sid":          c.sid,
		},
	}

	if input == nil {
		input = &UpdateModelBuildInput{}
	}

	response := &UpdateModelBuildResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
