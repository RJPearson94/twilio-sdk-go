// Package queries contains auto-generated files. DO NOT MODIFY
package queries

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreateQueryInput defines the input fields for creating a new query resource
type CreateQueryInput struct {
	Language   string  `validate:"required" form:"Language"`
	ModelBuild *string `form:"ModelBuild,omitempty"`
	Query      string  `validate:"required" form:"Query"`
	Tasks      *string `form:"Tasks,omitempty"`
}

type CreateQueryFieldResponse struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type CreateQueryResultResponse struct {
	Fields []CreateQueryFieldResponse `json:"fields"`
	Task   string                     `json:"task"`
}

// CreateQueryResponse defines the response fields for the created query
type CreateQueryResponse struct {
	AccountSid    string                    `json:"account_sid"`
	AssistantSid  string                    `json:"assistant_sid"`
	DateCreated   time.Time                 `json:"date_created"`
	DateUpdated   *time.Time                `json:"date_updated,omitempty"`
	DialogueSid   *string                   `json:"dialogue_sid,omitempty"`
	Language      string                    `json:"language"`
	ModelBuildSid string                    `json:"model_build_sid"`
	Query         string                    `json:"query"`
	Results       CreateQueryResultResponse `json:"results"`
	SampleSid     string                    `json:"sample_sid"`
	Sid           string                    `json:"sid"`
	SourceChannel string                    `json:"source_channel"`
	Status        string                    `json:"status"`
	URL           string                    `json:"url"`
}

// Create creates a new query resource
// See https://www.twilio.com/docs/autopilot/api/query#create-a-query-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateQueryInput) (*CreateQueryResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new query resource
// See https://www.twilio.com/docs/autopilot/api/query#create-a-query-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateQueryInput) (*CreateQueryResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Assistants/{assistantSid}/Queries",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
		},
	}

	if input == nil {
		input = &CreateQueryInput{}
	}

	response := &CreateQueryResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
