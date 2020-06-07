// This is an autogenerated file. DO NOT MODIFY
package step

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type GetStepOutput struct {
	Sid              string      `json:"sid"`
	AccountSid       string      `json:"account_sid"`
	FlowSid          string      `json:"flow_sid"`
	ExecutionSid     string      `json:"execution_sid"`
	Name             string      `json:"name"`
	Context          interface{} `json:"context"`
	TransitionedFrom string      `json:"transitioned_from"`
	TransitionedTo   string      `json:"transitioned_to"`
	DateCreated      time.Time   `json:"date_created"`
	DateUpdated      *time.Time  `json:"date_updated,omitempty"`
	URL              string      `json:"url"`
}

func (c Client) Get() (*GetStepOutput, error) {
	return c.GetWithContext(context.Background())
}

func (c Client) GetWithContext(context context.Context) (*GetStepOutput, error) {
	op := client.Operation{
		HTTPMethod: http.MethodGet,
		HTTPPath:   "/Flows/{flowSid}/Executions/{executionSid}/Steps/{sid}",
		PathParams: map[string]string{
			"flowSid":      c.flowSid,
			"executionSid": c.executionSid,
			"sid":          c.sid,
		},
	}

	output := &GetStepOutput{}
	if err := c.client.Send(context, op, nil, output); err != nil {
		return nil, err
	}
	return output, nil
}