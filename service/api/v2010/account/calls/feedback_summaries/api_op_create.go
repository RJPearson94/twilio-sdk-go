// Package feedback_summaries contains auto-generated files. DO NOT MODIFY
package feedback_summaries

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// CreateFeedbackSummaryInput defines input fields for creating a new feedback summary
type CreateFeedbackSummaryInput struct {
	EndDate              string  `validate:"required" form:"EndDate"`
	IncludeSubaccounts   *bool   `form:"IncludeSubaccounts,omitempty"`
	StartDate            string  `validate:"required" form:"StartDate"`
	StatusCallback       *string `form:"StatusCallback,omitempty"`
	StatusCallbackMethod *string `form:"StatusCallbackMethod,omitempty"`
}

// CreateFeedbackSummaryResponse defines the response fields for creating a new feedback summary
type CreateFeedbackSummaryResponse struct {
	AccountSid                    string             `json:"account_sid"`
	CallCount                     int                `json:"call_count"`
	CallFeedbackCount             int                `json:"call_feedback_count"`
	DateCreated                   utils.RFC2822Time  `json:"date_created"`
	DateUpdated                   *utils.RFC2822Time `json:"date_updated,omitempty"`
	EndDate                       string             `json:"end_date"`
	IncludeSubaccounts            bool               `json:"include_subaccounts"`
	Issues                        *[]string          `json:"issues,omitempty"`
	QualityScoreAverage           *float64           `json:"quality_score_average,omitempty"`
	QualityScoreMedian            *float64           `json:"quality_score_median,omitempty"`
	QualityScoreStandardDeviation *float64           `json:"quality_score_standard_deviation,omitempty"`
	Sid                           string             `json:"sid"`
	StartDate                     string             `json:"start_date"`
	Status                        string             `json:"status"`
}

// Create creates a new feedback summary resource
// See https://www.twilio.com/docs/voice/api/feedbacksummary-resource#create-a-callfeedbacksummary-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateFeedbackSummaryInput) (*CreateFeedbackSummaryResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new feedback summary resource
// See https://www.twilio.com/docs/voice/api/feedbacksummary-resource#create-a-callfeedbacksummary-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateFeedbackSummaryInput) (*CreateFeedbackSummaryResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Accounts/{accountSid}/Calls/FeedbackSummary.json",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"accountSid": c.accountSid,
		},
	}

	if input == nil {
		input = &CreateFeedbackSummaryInput{}
	}

	response := &CreateFeedbackSummaryResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
