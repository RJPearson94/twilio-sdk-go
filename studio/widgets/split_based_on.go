// Package widgets contains auto-generated files. DO NOT MODIFY
package widgets

import (
	"fmt"

	"github.com/RJPearson94/twilio-sdk-go/studio/flow"
	"github.com/RJPearson94/twilio-sdk-go/studio/properties"
	"github.com/RJPearson94/twilio-sdk-go/studio/transition"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

type SplitBasedOnNextTransitions struct {
	NoMatch *string

	Matches *[]transition.Conditional
}

type SplitBasedOnProperties struct {
	Input  string             `validate:"required" json:"input"`
	Offset *properties.Offset `json:"offset,omitempty"`
}

type SplitBasedOn struct {
	NextTransitions SplitBasedOnNextTransitions
	Properties      SplitBasedOnProperties `validate:"required"`
	Name            string                 `validate:"required"`
}

// Validate checks the widget is correctly configured
func (widget SplitBasedOn) Validate() error {
	if err := utils.ValidateInput(widget); err != nil {
		return fmt.Errorf("Invalid input supplied. Errors %s", err.Error())
	}
	return nil
}

// ToState returns a populated Studio Widget State struct
func (widget SplitBasedOn) ToState() (*flow.State, error) {
	transitions := []flow.Transition{
		{
			Event: "noMatch",
			Next:  widget.NextTransitions.NoMatch,
		},
	}

	if widget.NextTransitions.Matches != nil {
		for _, value := range *widget.NextTransitions.Matches {
			if err := value.Validate(); err != nil {
				return nil, err
			}

			transitions = append(transitions, flow.Transition{
				Event:      "match",
				Next:       utils.String(value.Next),
				Conditions: value.Conditions,
			})
		}
	}

	return &flow.State{
		Name:        widget.Name,
		Type:        "split-based-on",
		Transitions: transitions,
		Properties:  widget.Properties,
	}, nil
}
