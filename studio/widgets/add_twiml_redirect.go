// Package widgets contains auto-generated files. DO NOT MODIFY
package widgets

import (
	"fmt"

	"github.com/RJPearson94/twilio-sdk-go/studio/flow"
	"github.com/RJPearson94/twilio-sdk-go/studio/properties"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

type AddTwimlRedirectNextTransitions struct {
	Fail    *string
	Return  *string
	Timeout *string
}

type AddTwimlRedirectProperties struct {
	Method  *string            `json:"method,omitempty"`
	Offset  *properties.Offset `json:"offset,omitempty"`
	Timeout *string            `json:"timeout,omitempty"`
	URL     string             `validate:"required" json:"url"`
}

type AddTwimlRedirect struct {
	NextTransitions AddTwimlRedirectNextTransitions
	Properties      AddTwimlRedirectProperties `validate:"required"`
	Name            string                     `validate:"required"`
}

// Validate checks the widget is correctly configured
func (widget AddTwimlRedirect) Validate() error {
	if err := utils.ValidateInput(widget); err != nil {
		return fmt.Errorf("Invalid input supplied. Errors %s", err.Error())
	}
	return nil
}

// ToState returns a populated Studio Widget State struct
func (widget AddTwimlRedirect) ToState() (*flow.State, error) {
	transitions := []flow.Transition{
		{
			Event: "fail",
			Next:  widget.NextTransitions.Fail,
		},
		{
			Event: "return",
			Next:  widget.NextTransitions.Return,
		},
		{
			Event: "timeout",
			Next:  widget.NextTransitions.Timeout,
		},
	}

	return &flow.State{
		Name:        widget.Name,
		Type:        "add-twiml-redirect",
		Transitions: transitions,
		Properties:  widget.Properties,
	}, nil
}
