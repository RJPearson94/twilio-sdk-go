// Package widgets contains auto-generated files. DO NOT MODIFY
package widgets

import (
	"fmt"

	"github.com/RJPearson94/twilio-sdk-go/studio/flow"
	"github.com/RJPearson94/twilio-sdk-go/studio/properties"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

type EnqueueCallNextTransitions struct {
	CallComplete    *string
	CallFailure     *string
	FailedToEnqueue *string
}

type EnqueueCallProperties struct {
	Offset         *properties.Offset `json:"offset,omitempty"`
	Priority       *int               `json:"priority,omitempty"`
	QueueName      *string            `json:"queue_name,omitempty"`
	TaskAttributes *string            `json:"task_attributes,omitempty"`
	Timeout        *int               `json:"timeout,omitempty"`
	WaitURL        *string            `json:"wait_url,omitempty"`
	WaitURLMethod  *string            `json:"wait_url_method,omitempty"`
	WorkflowSid    *string            `json:"workflow_sid,omitempty"`
}

type EnqueueCall struct {
	NextTransitions EnqueueCallNextTransitions
	Properties      EnqueueCallProperties `validate:"required"`
	Name            string                `validate:"required"`
}

// Validate checks the widget is correctly configured
func (widget EnqueueCall) Validate() error {
	if err := utils.ValidateInput(widget); err != nil {
		return fmt.Errorf("Invalid input supplied. Errors %s", err.Error())
	}
	return nil
}

// ToState returns a populated Studio Widget State struct
func (widget EnqueueCall) ToState() (*flow.State, error) {
	transitions := []flow.Transition{
		{
			Event: "callComplete",
			Next:  widget.NextTransitions.CallComplete,
		},
		{
			Event: "callFailure",
			Next:  widget.NextTransitions.CallFailure,
		},
		{
			Event: "failedToEnqueue",
			Next:  widget.NextTransitions.FailedToEnqueue,
		},
	}

	return &flow.State{
		Name:        widget.Name,
		Type:        "enqueue-call",
		Transitions: transitions,
		Properties:  widget.Properties,
	}, nil
}
