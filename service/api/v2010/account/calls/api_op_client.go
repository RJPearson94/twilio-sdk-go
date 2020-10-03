// Package calls contains auto-generated files. DO NOT MODIFY
package calls

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/calls/feedback_summaries"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/calls/feedback_summary"
)

// Client for managing call resources
// See https://www.twilio.com/docs/voice/api/call-resource for more details
type Client struct {
	client *client.Client

	accountSid string

	FeedbackSummaries *feedback_summaries.Client
	FeedbackSummary   func(string) *feedback_summary.Client
}

// ClientProperties are the properties required to manage the calls resources
type ClientProperties struct {
	AccountSid string
}

// New creates a new instance of the calls client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		accountSid: properties.AccountSid,

		FeedbackSummaries: feedback_summaries.New(client, feedback_summaries.ClientProperties{
			AccountSid: properties.AccountSid,
		}),
		FeedbackSummary: func(feedbackSummarySid string) *feedback_summary.Client {
			return feedback_summary.New(client, feedback_summary.ClientProperties{
				AccountSid: properties.AccountSid,
				Sid:        feedbackSummarySid,
			})
		},
	}
}
