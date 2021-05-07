// Package widgets contains auto-generated files. DO NOT MODIFY
package widgets

import (
	"fmt"

	"github.com/RJPearson94/twilio-sdk-go/studio/flow"
	"github.com/RJPearson94/twilio-sdk-go/studio/properties"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

type TriggerNextTransitions struct {
	IncomingCall    *string
	IncomingMessage *string
	IncomingRequest *string
}

type TriggerProperties struct {
	Offset *properties.Offset `json:"offset,omitempty"`
}

type Trigger struct {
	NextTransitions TriggerNextTransitions
	Properties      TriggerProperties `validate:"required"`
	Name            string            `validate:"required"`
}

// Validate checks the widget is correctly configured
func (widget Trigger) Validate() error {
	if err := utils.ValidateInput(widget); err != nil {
		return fmt.Errorf("Invalid input supplied. Errors %s", err.Error())
	}
	return nil
}

// ToState returns a populated Studio Widget State struct
func (widget Trigger) ToState() (*flow.State, error) {
	transitions := []flow.Transition{
		{
			Event: "incomingCall",
			Next:  widget.NextTransitions.IncomingCall,
		},
		{
			Event: "incomingMessage",
			Next:  widget.NextTransitions.IncomingMessage,
		},
		{
			Event: "incomingRequest",
			Next:  widget.NextTransitions.IncomingRequest,
		},
	}

	return &flow.State{
		Name:        widget.Name,
		Type:        "trigger",
		Transitions: transitions,
		Properties:  widget.Properties,
	}, nil
}
