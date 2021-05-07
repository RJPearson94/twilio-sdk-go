// Package widgets contains auto-generated files. DO NOT MODIFY
package widgets

import (
	"fmt"

	"github.com/RJPearson94/twilio-sdk-go/studio/flow"
	"github.com/RJPearson94/twilio-sdk-go/studio/properties"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

type RunFunctionNextTransitions struct {
	Fail    *string
	Success *string
}

type RunFunctionParameter struct {
	Key   string `validate:"required" json:"key"`
	Value string `validate:"required" json:"value"`
}

type RunFunctionProperties struct {
	EnvironmentSid *string                 `json:"environment_sid,omitempty"`
	FunctionSid    *string                 `json:"function_sid,omitempty"`
	Offset         *properties.Offset      `json:"offset,omitempty"`
	Parameters     *[]RunFunctionParameter `json:"parameters,omitempty"`
	ServiceSid     *string                 `json:"service_sid,omitempty"`
	URL            string                  `validate:"required" json:"url"`
}

type RunFunction struct {
	NextTransitions RunFunctionNextTransitions
	Properties      RunFunctionProperties `validate:"required"`
	Name            string                `validate:"required"`
}

// Validate checks the widget is correctly configured
func (widget RunFunction) Validate() error {
	if err := utils.ValidateInput(widget); err != nil {
		return fmt.Errorf("Invalid input supplied. Errors %s", err.Error())
	}
	return nil
}

// ToState returns a populated Studio Widget State struct
func (widget RunFunction) ToState() (*flow.State, error) {
	transitions := []flow.Transition{
		{
			Event: "fail",
			Next:  widget.NextTransitions.Fail,
		},
		{
			Event: "success",
			Next:  widget.NextTransitions.Success,
		},
	}

	return &flow.State{
		Name:        widget.Name,
		Type:        "run-function",
		Transitions: transitions,
		Properties:  widget.Properties,
	}, nil
}
