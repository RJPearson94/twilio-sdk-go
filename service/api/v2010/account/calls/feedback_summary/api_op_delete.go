// Package feedback_summary contains auto-generated files. DO NOT MODIFY
package feedback_summary

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// Delete removes a feedback summary resource from the account
// See https://www.twilio.com/docs/voice/api/feedbacksummary-resource#delete-a-callfeedbacksummary-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Delete() error {
	return c.DeleteWithContext(context.Background())
}

// DeleteWithContext removes a feedback summary resource from the account
// See https://www.twilio.com/docs/voice/api/feedbacksummary-resource#delete-a-callfeedbacksummary-resource for more details
func (c Client) DeleteWithContext(context context.Context) error {
	op := client.Operation{
		Method: http.MethodDelete,
		URI:    "/Accounts/{accountSid}/Calls/FeedbackSummary/{sid}.json",
		PathParams: map[string]string{
			"accountSid": c.accountSid,
			"sid":        c.sid,
		},
	}

	if err := c.client.Send(context, op, nil, nil); err != nil {
		return err
	}
	return nil
}
