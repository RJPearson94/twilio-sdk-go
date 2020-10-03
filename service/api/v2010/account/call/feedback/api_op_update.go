// Package feedback contains auto-generated files. DO NOT MODIFY
package feedback

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// UpdateFeedbackInput defines input fields for updating feedback
type UpdateFeedbackInput struct {
	Issues       *[]string `form:"Issue,omitempty"`
	QualityScore int       `validate:"required" form:"QualityScore"`
}

// UpdateFeedbackResponse defines the response fields for the updated feedback
type UpdateFeedbackResponse struct {
	AccountSid   string             `json:"account_sid"`
	DateCreated  utils.RFC2822Time  `json:"date_created"`
	DateUpdated  *utils.RFC2822Time `json:"date_updated,omitempty"`
	Issues       []string           `json:"issues"`
	QualityScore int                `json:"quality_score"`
	Sid          string             `json:"sid"`
}

// Update modifies a feedback resource
// See https://www.twilio.com/docs/voice/api/feedback-resource#update-a-call-feedback-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateFeedbackInput) (*UpdateFeedbackResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a feedback resource
// See https://www.twilio.com/docs/voice/api/feedback-resource#update-a-call-feedback-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateFeedbackInput) (*UpdateFeedbackResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Accounts/{accountSid}/Calls/{callSid}/Feedback.json",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"accountSid": c.accountSid,
			"callSid":    c.callSid,
		},
	}

	if input == nil {
		input = &UpdateFeedbackInput{}
	}

	response := &UpdateFeedbackResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
