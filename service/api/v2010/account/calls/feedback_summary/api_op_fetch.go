// Package feedback_summary contains auto-generated files. DO NOT MODIFY
package feedback_summary

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// FetchFeedbackSummaryResponse defines the response fields for retrieving a feedback summary
type FetchFeedbackSummaryResponse struct {
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

// Fetch retrieves a feedback summary resource
// See https://www.twilio.com/docs/voice/api/feedbacksummary-resource#fetch-a-callfeedbacksummary-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchFeedbackSummaryResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a feedback summary resource
// See https://www.twilio.com/docs/voice/api/feedbacksummary-resource#fetch-a-callfeedbacksummary-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchFeedbackSummaryResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Accounts/{accountSid}/Calls/FeedbackSummary/{sid}.json",
		PathParams: map[string]string{
			"accountSid": c.accountSid,
			"sid":        c.sid,
		},
	}

	response := &FetchFeedbackSummaryResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
