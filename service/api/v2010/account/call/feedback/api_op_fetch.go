// Package feedback contains auto-generated files. DO NOT MODIFY
package feedback

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// FetchFeedbackResponse defines the response fields for retrieving feedback
type FetchFeedbackResponse struct {
	AccountSid   string             `json:"account_sid"`
	DateCreated  utils.RFC2822Time  `json:"date_created"`
	DateUpdated  *utils.RFC2822Time `json:"date_updated,omitempty"`
	Issues       []string           `json:"issues"`
	QualityScore int                `json:"quality_score"`
	Sid          string             `json:"sid"`
}

// Fetch retrieves the feedback resource
// See https://www.twilio.com/docs/voice/api/feedback-resource#fetch-a-call-feedback-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchFeedbackResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves the feedback resource
// See https://www.twilio.com/docs/voice/api/feedback-resource#fetch-a-call-feedback-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchFeedbackResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Accounts/{accountSid}/Calls/{callSid}/Feedback.json",
		PathParams: map[string]string{
			"accountSid": c.accountSid,
			"callSid":    c.callSid,
		},
	}

	response := &FetchFeedbackResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
