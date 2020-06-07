// This is an autogenerated file. DO NOT MODIFY
package flows

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type CreateFlowInput struct {
	FriendlyName  string `validate:"required" form:"FriendlyName"`
	Status        string `validate:"required" form:"Status"`
	Definition    string `validate:"required" form:"Definition"`
	CommitMessage string `form:"CommitMessage,omitempty"`
}

type CreateFlowOutput struct {
	Sid           string         `json:"sid"`
	AccountSid    string         `json:"account_sid"`
	FriendlyName  string         `json:"friendly_name"`
	Definition    interface{}    `json:"definition"`
	Status        string         `json:"status"`
	Revision      int            `json:"revision"`
	CommitMessage *string        `json:"commit_message,omitempty"`
	Valid         bool           `json:"valid"`
	Errors        *[]interface{} `json:"errors,omitempty"`
	Warnings      *[]interface{} `json:"warnings,omitempty"`
	DateCreated   time.Time      `json:"date_created"`
	DateUpdated   *time.Time     `json:"date_updated,omitempty"`
	WebhookURL    string         `json:"webhook_url"`
	URL           string         `json:"url"`
}

func (c Client) Create(input *CreateFlowInput) (*CreateFlowOutput, error) {
	return c.CreateWithContext(context.Background(), input)
}

func (c Client) CreateWithContext(context context.Context, input *CreateFlowInput) (*CreateFlowOutput, error) {
	op := client.Operation{
		HTTPMethod:  http.MethodPost,
		HTTPPath:    "/Flows",
		ContentType: client.URLEncoded,
	}

	output := &CreateFlowOutput{}
	if err := c.client.Send(context, op, input, output); err != nil {
		return nil, err
	}
	return output, nil
}
