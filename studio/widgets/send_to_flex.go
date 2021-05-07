// Package widgets contains auto-generated files. DO NOT MODIFY
package widgets

import (
	"fmt"

	"github.com/RJPearson94/twilio-sdk-go/studio/flow"
	"github.com/RJPearson94/twilio-sdk-go/studio/properties"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

type SendToFlexNextTransitions struct {
	CallComplete    *string
	CallFailure     *string
	FailedToEnqueue *string
}

type SendToFlexProperties struct {
	Attributes    *string            `json:"attributes,omitempty"`
	Channel       string             `validate:"required" json:"channel"`
	Offset        *properties.Offset `json:"offset,omitempty"`
	Priority      *string            `json:"priority,omitempty"`
	Timeout       *string            `json:"timeout,omitempty"`
	WaitURL       *string            `json:"waitUrl,omitempty"`
	WaitURLMethod *string            `json:"waitUrlMethod,omitempty"`
	Workflow      string             `validate:"required" json:"workflow"`
}

type SendToFlex struct {
	NextTransitions SendToFlexNextTransitions
	Properties      SendToFlexProperties `validate:"required"`
	Name            string               `validate:"required"`
}

// Validate checks the widget is correctly configured
func (widget SendToFlex) Validate() error {
	if err := utils.ValidateInput(widget); err != nil {
		return fmt.Errorf("Invalid input supplied. Errors %s", err.Error())
	}
	return nil
}

// ToState returns a populated Studio Widget State struct
func (widget SendToFlex) ToState() (*flow.State, error) {
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
		Type:        "send-to-flex",
		Transitions: transitions,
		Properties:  widget.Properties,
	}, nil
}
