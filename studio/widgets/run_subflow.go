// Package widgets contains auto-generated files. DO NOT MODIFY
package widgets

import (
	"fmt"

	"github.com/RJPearson94/twilio-sdk-go/studio/flow"
	"github.com/RJPearson94/twilio-sdk-go/studio/properties"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

type RunSubflowNextTransitions struct {
	Completed *string
	Failed    *string
}

type RunSubflowParameter struct {
	Key   string `validate:"required" json:"key"`
	Value string `validate:"required" json:"value"`
}

type RunSubflowProperties struct {
	FlowRevision string                 `validate:"required" json:"flow_revision"`
	FlowSid      string                 `validate:"required" json:"flow_sid"`
	Offset       *properties.Offset     `json:"offset,omitempty"`
	Parameters   *[]RunSubflowParameter `json:"parameters,omitempty"`
}

type RunSubflow struct {
	NextTransitions RunSubflowNextTransitions
	Properties      RunSubflowProperties `validate:"required"`
	Name            string               `validate:"required"`
}

// Validate checks the widget is correctly configured
func (widget RunSubflow) Validate() error {
	if err := utils.ValidateInput(widget); err != nil {
		return fmt.Errorf("Invalid input supplied. Errors %s", err.Error())
	}
	return nil
}

// ToState returns a populated Studio Widget State struct
func (widget RunSubflow) ToState() (*flow.State, error) {
	transitions := []flow.Transition{
		{
			Event: "completed",
			Next:  widget.NextTransitions.Completed,
		},
		{
			Event: "failed",
			Next:  widget.NextTransitions.Failed,
		},
	}

	return &flow.State{
		Name:        widget.Name,
		Type:        "run-subflow",
		Transitions: transitions,
		Properties:  widget.Properties,
	}, nil
}
