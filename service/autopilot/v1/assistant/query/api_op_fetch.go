// Package query contains auto-generated files. DO NOT MODIFY
package query

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type FetchQueryResponseField struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type FetchQueryResponseResult struct {
	Fields []FetchQueryResponseField `json:"fields"`
	Task   string                    `json:"task"`
}

// FetchQueryResponse defines the response fields for the retrieved query
type FetchQueryResponse struct {
	AccountSid    string                   `json:"account_sid"`
	AssistantSid  string                   `json:"assistant_sid"`
	DateCreated   time.Time                `json:"date_created"`
	DateUpdated   *time.Time               `json:"date_updated,omitempty"`
	DialogueSid   *string                  `json:"dialogue_sid,omitempty"`
	Language      string                   `json:"language"`
	ModelBuildSid string                   `json:"model_build_sid"`
	Query         string                   `json:"query"`
	Results       FetchQueryResponseResult `json:"results"`
	SampleSid     string                   `json:"sample_sid"`
	Sid           string                   `json:"sid"`
	SourceChannel string                   `json:"source_channel"`
	Status        string                   `json:"status"`
	URL           string                   `json:"url"`
}

// Fetch retrieves a query resource
// See https://www.twilio.com/docs/autopilot/api/query#fetch-a-query-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchQueryResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a query resource
// See https://www.twilio.com/docs/autopilot/api/query#fetch-a-query-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchQueryResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Assistants/{assistantSid}/Queries/{sid}",
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
			"sid":          c.sid,
		},
	}

	response := &FetchQueryResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
