// Package widgets contains auto-generated files. DO NOT MODIFY
package widgets

import (
	"fmt"

	"github.com/RJPearson94/twilio-sdk-go/studio/flow"
	"github.com/RJPearson94/twilio-sdk-go/studio/properties"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

type SendToAutopilotNextTransitions struct {
	Failure      *string
	SessionEnded *string
	Timeout      *string
}

type SendToAutopilotMemoryParameter struct {
	Key   string `validate:"required" json:"key"`
	Value string `validate:"required" json:"value"`
}

type SendToAutopilotProperties struct {
	AutopilotAssistantSid string                            `validate:"required" json:"autopilot_assistant_sid"`
	Body                  string                            `validate:"required" json:"body"`
	ChatAttributes        *string                           `json:"chat_attributes,omitempty"`
	ChatChannel           *string                           `json:"chat_channel,omitempty"`
	ChatService           *string                           `json:"chat_service,omitempty"`
	From                  string                            `validate:"required" json:"from"`
	MemoryParameters      *[]SendToAutopilotMemoryParameter `json:"memory_parameters,omitempty"`
	Offset                *properties.Offset                `json:"offset,omitempty"`
	TargetTask            *string                           `json:"target_task,omitempty"`
	Timeout               int                               `validate:"required" json:"timeout"`
}

type SendToAutopilot struct {
	NextTransitions SendToAutopilotNextTransitions
	Properties      SendToAutopilotProperties `validate:"required"`
	Name            string                    `validate:"required"`
}

// Validate checks the widget is correctly configured
func (widget SendToAutopilot) Validate() error {
	if err := utils.ValidateInput(widget); err != nil {
		return fmt.Errorf("Invalid input supplied. Errors %s", err.Error())
	}
	return nil
}

// ToState returns a populated Studio Widget State struct
func (widget SendToAutopilot) ToState() (*flow.State, error) {
	transitions := []flow.Transition{
		{
			Event: "failure",
			Next:  widget.NextTransitions.Failure,
		},
		{
			Event: "sessionEnded",
			Next:  widget.NextTransitions.SessionEnded,
		},
		{
			Event: "timeout",
			Next:  widget.NextTransitions.Timeout,
		},
	}

	return &flow.State{
		Name:        widget.Name,
		Type:        "send-to-auto-pilot",
		Transitions: transitions,
		Properties:  widget.Properties,
	}, nil
}
