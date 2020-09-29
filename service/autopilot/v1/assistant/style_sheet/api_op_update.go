// Package style_sheet contains auto-generated files. DO NOT MODIFY
package style_sheet

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// UpdateStyleSheetInput defines the input fields for updating a stylesheet
type UpdateStyleSheetInput struct {
	StyleSheet *string `form:"StyleSheet,omitempty"`
}

// UpdateStyleSheetResponse defines the response fields for the updated style sheet
type UpdateStyleSheetResponse struct {
	AccountSid   string                 `json:"account_sid"`
	AssistantSid string                 `json:"assistant_sid"`
	Data         map[string]interface{} `json:"data"`
	URL          string                 `json:"url"`
}

// Update modifies a style sheet resource
// See https://www.twilio.com/docs/autopilot/api/assistant/stylesheet#update-a-stylesheet-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateStyleSheetInput) (*UpdateStyleSheetResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a style sheet resource
// See https://www.twilio.com/docs/autopilot/api/assistant/stylesheet#update-a-stylesheet-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateStyleSheetInput) (*UpdateStyleSheetResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Assistants/{assistantSid}/StyleSheet",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
		},
	}

	if input == nil {
		input = &UpdateStyleSheetInput{}
	}

	response := &UpdateStyleSheetResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
