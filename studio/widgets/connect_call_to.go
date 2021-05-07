// Package widgets contains auto-generated files. DO NOT MODIFY
package widgets

import (
	"fmt"

	"github.com/RJPearson94/twilio-sdk-go/studio/flow"
	"github.com/RJPearson94/twilio-sdk-go/studio/properties"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

type ConnectCallToNextTransitions struct {
	CallCompleted *string
	Hangup        *string
}

type ConnectCallToProperties struct {
	CallerID    string             `validate:"required" json:"caller_id"`
	Noun        string             `validate:"required" json:"noun"`
	Offset      *properties.Offset `json:"offset,omitempty"`
	Record      *bool              `json:"record,omitempty"`
	SipEndpoint *string            `json:"sip_endpoint,omitempty"`
	SipPassword *string            `json:"sip_password,omitempty"`
	SipUsername *string            `json:"sip_username,omitempty"`
	Timeout     *int               `json:"timeout,omitempty"`
	To          *string            `json:"to,omitempty"`
}

type ConnectCallTo struct {
	NextTransitions ConnectCallToNextTransitions
	Properties      ConnectCallToProperties `validate:"required"`
	Name            string                  `validate:"required"`
}

// Validate checks the widget is correctly configured
func (widget ConnectCallTo) Validate() error {
	if err := utils.ValidateInput(widget); err != nil {
		return fmt.Errorf("Invalid input supplied. Errors %s", err.Error())
	}
	return nil
}

// ToState returns a populated Studio Widget State struct
func (widget ConnectCallTo) ToState() (*flow.State, error) {
	transitions := []flow.Transition{
		{
			Event: "callCompleted",
			Next:  widget.NextTransitions.CallCompleted,
		},
		{
			Event: "hangup",
			Next:  widget.NextTransitions.Hangup,
		},
	}

	return &flow.State{
		Name:        widget.Name,
		Type:        "connect-call-to",
		Transitions: transitions,
		Properties:  widget.Properties,
	}, nil
}
