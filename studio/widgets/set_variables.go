// Package widgets contains auto-generated files. DO NOT MODIFY
package widgets

import (
	"fmt"

	"github.com/RJPearson94/twilio-sdk-go/studio/flow"
	"github.com/RJPearson94/twilio-sdk-go/studio/properties"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

type SetVariablesNextTransitions struct {
	Next *string
}

type SetVariablesVariable struct {
	Key   string `validate:"required" json:"key"`
	Value string `validate:"required" json:"value"`
}

type SetVariablesProperties struct {
	Offset    *properties.Offset      `json:"offset,omitempty"`
	Variables *[]SetVariablesVariable `json:"variables,omitempty"`
}

type SetVariables struct {
	NextTransitions SetVariablesNextTransitions
	Properties      SetVariablesProperties `validate:"required"`
	Name            string                 `validate:"required"`
}

// Validate checks the widget is correctly configured
func (widget SetVariables) Validate() error {
	if err := utils.ValidateInput(widget); err != nil {
		return fmt.Errorf("Invalid input supplied. Errors %s", err.Error())
	}
	return nil
}

// ToState returns a populated Studio Widget State struct
func (widget SetVariables) ToState() (*flow.State, error) {
	transitions := []flow.Transition{
		{
			Event: "next",
			Next:  widget.NextTransitions.Next,
		},
	}

	return &flow.State{
		Name:        widget.Name,
		Type:        "set-variables",
		Transitions: transitions,
		Properties:  widget.Properties,
	}, nil
}
