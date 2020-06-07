// This is an autogenerated file. DO NOT MODIFY
package context

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type GetContextOutput struct {
	Sid          string      `json:"sid"`
	AccountSid   string      `json:"account_sid"`
	FlowSid      string      `json:"flow_sid"`
	ExecutionSid string      `json:"execution_sid"`
	StepSid      string      `json:"step_sid"`
	Context      interface{} `json:"context"`
	URL          string      `json:"url"`
}

func (c Client) Get() (*GetContextOutput, error) {
	return c.GetWithContext(context.Background())
}

func (c Client) GetWithContext(context context.Context) (*GetContextOutput, error) {
	op := client.Operation{
		HTTPMethod: http.MethodGet,
		HTTPPath:   "/Flows/{flowSid}/Executions/{executionSid}/Steps/{stepSid}/Context",
		PathParams: map[string]string{
			"flowSid":      c.flowSid,
			"executionSid": c.executionSid,
			"stepSid":      c.stepSid,
		},
	}

	output := &GetContextOutput{}
	if err := c.client.Send(context, op, nil, output); err != nil {
		return nil, err
	}
	return output, nil
}