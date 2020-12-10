// Package recording contains auto-generated files. DO NOT MODIFY
package recording

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// UpdateRecordingInput defines input fields for updating a recording resource
type UpdateRecordingInput struct {
	Mode *string `form:"Mode,omitempty"`
	Trim *string `form:"Trim,omitempty"`
}

// UpdateRecordingResponse defines the response fields for the updated recording
type UpdateRecordingResponse struct {
	Mode string `json:"mode"`
	Trim string `json:"trim"`
}

// Update modifies a recording resource
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateRecordingInput) (*UpdateRecordingResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a recording resource
func (c Client) UpdateWithContext(context context.Context, input *UpdateRecordingInput) (*UpdateRecordingResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Trunks/{trunkSid}/Recording",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"trunkSid": c.trunkSid,
		},
	}

	if input == nil {
		input = &UpdateRecordingInput{}
	}

	response := &UpdateRecordingResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
