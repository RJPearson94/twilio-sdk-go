// Package test_users contains auto-generated files. DO NOT MODIFY
package test_users

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchTestUsersResponse defines the response fields for the retrieved test users
type FetchTestUsersResponse struct {
	Sid       string   `json:"sid"`
	TestUsers []string `json:"test_users"`
	URL       string   `json:"url"`
}

// Fetch retrieves a test users resource
// See https://www.twilio.com/docs/studio/rest-api/v2/test-user#fetch-a-testuser-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchTestUsersResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a test users resource
// See https://www.twilio.com/docs/studio/rest-api/v2/test-user#fetch-a-testuser-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchTestUsersResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Flows/{flowSid}/TestUsers",
		PathParams: map[string]string{
			"flowSid": c.flowSid,
		},
	}

	response := &FetchTestUsersResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
