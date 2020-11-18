// Package query contains auto-generated files. DO NOT MODIFY
package query

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// UpdateQueryInput defines the input fields for updating a query
type UpdateQueryInput struct {
	SampleSid *string `form:"SampleSid,omitempty"`
	Status    *string `form:"Status,omitempty"`
}

type UpdateQueryFieldResponse struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type UpdateQueryResultResponse struct {
	Fields []UpdateQueryFieldResponse `json:"fields"`
	Task   string                     `json:"task"`
}

// UpdateQueryResponse defines the response fields for the updated query
type UpdateQueryResponse struct {
	AccountSid    string                    `json:"account_sid"`
	AssistantSid  string                    `json:"assistant_sid"`
	DateCreated   time.Time                 `json:"date_created"`
	DateUpdated   *time.Time                `json:"date_updated,omitempty"`
	DialogueSid   *string                   `json:"dialogue_sid,omitempty"`
	Language      string                    `json:"language"`
	ModelBuildSid string                    `json:"model_build_sid"`
	Query         string                    `json:"query"`
	Results       UpdateQueryResultResponse `json:"results"`
	SampleSid     string                    `json:"sample_sid"`
	Sid           string                    `json:"sid"`
	SourceChannel string                    `json:"source_channel"`
	Status        string                    `json:"status"`
	URL           string                    `json:"url"`
}

// Update modifies a query resource
// See https://www.twilio.com/docs/autopilot/api/query#update-a-query-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateQueryInput) (*UpdateQueryResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a query resource
// See https://www.twilio.com/docs/autopilot/api/query#update-a-query-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateQueryInput) (*UpdateQueryResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Assistants/{assistantSid}/Queries/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
			"sid":          c.sid,
		},
	}

	if input == nil {
		input = &UpdateQueryInput{}
	}

	response := &UpdateQueryResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
