// Package feedbacks contains auto-generated files. DO NOT MODIFY
package feedbacks

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// CreateFeedbackInput defines input fields for creating a new feedback resource
type CreateFeedbackInput struct {
	Issues       *[]string `form:"Issue,omitempty"`
	QualityScore int       `validate:"required" form:"QualityScore"`
}

// CreateFeedbackResponse defines the response fields for creating a new feedback resource
type CreateFeedbackResponse struct {
	AccountSid   string             `json:"account_sid"`
	DateCreated  utils.RFC2822Time  `json:"date_created"`
	DateUpdated  *utils.RFC2822Time `json:"date_updated,omitempty"`
	Issues       []string           `json:"issues"`
	QualityScore int                `json:"quality_score"`
	Sid          string             `json:"sid"`
}

// Create creates a new feedback resource
// See https://www.twilio.com/docs/voice/api/feedback-resource#create-a-call-feedback-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateFeedbackInput) (*CreateFeedbackResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new feedback resource
// See https://www.twilio.com/docs/voice/api/feedback-resource#create-a-call-feedback-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateFeedbackInput) (*CreateFeedbackResponse, error) {
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
		input = &CreateFeedbackInput{}
	}

	response := &CreateFeedbackResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
