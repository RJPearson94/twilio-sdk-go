// Package recording contains auto-generated files. DO NOT MODIFY
package recording

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchRecordingResponse defines the response fields for the retrieved recording resource
type FetchRecordingResponse struct {
	Mode string `json:"mode"`
	Trim string `json:"trim"`
}

// Fetch retrieves a recording resource
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchRecordingResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a recording resource
func (c Client) FetchWithContext(context context.Context) (*FetchRecordingResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Trunks/{trunkSid}/Recording",
		PathParams: map[string]string{
			"trunkSid": c.trunkSid,
		},
	}

	response := &FetchRecordingResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
