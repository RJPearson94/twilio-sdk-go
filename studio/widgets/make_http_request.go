// Package widgets contains auto-generated files. DO NOT MODIFY
package widgets

import (
	"fmt"

	"github.com/RJPearson94/twilio-sdk-go/studio/flow"
	"github.com/RJPearson94/twilio-sdk-go/studio/properties"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

type MakeHTTPRequestNextTransitions struct {
	Failed  *string
	Success *string
}

type MakeHTTPRequestParameter struct {
	Key   string `validate:"required" json:"key"`
	Value string `validate:"required" json:"value"`
}

type MakeHTTPRequestProperties struct {
	Body        *string                     `json:"body,omitempty"`
	ContentType string                      `validate:"required" json:"content_type"`
	Method      string                      `validate:"required" json:"method"`
	Offset      *properties.Offset          `json:"offset,omitempty"`
	Parameters  *[]MakeHTTPRequestParameter `json:"parameters,omitempty"`
	URL         string                      `validate:"required" json:"url"`
}

type MakeHTTPRequest struct {
	NextTransitions MakeHTTPRequestNextTransitions
	Properties      MakeHTTPRequestProperties `validate:"required"`
	Name            string                    `validate:"required"`
}

// Validate checks the widget is correctly configured
func (widget MakeHTTPRequest) Validate() error {
	if err := utils.ValidateInput(widget); err != nil {
		return fmt.Errorf("Invalid input supplied. Errors %s", err.Error())
	}
	return nil
}

// ToState returns a populated Studio Widget State struct
func (widget MakeHTTPRequest) ToState() (*flow.State, error) {
	transitions := []flow.Transition{
		{
			Event: "failed",
			Next:  widget.NextTransitions.Failed,
		},
		{
			Event: "success",
			Next:  widget.NextTransitions.Success,
		},
	}

	return &flow.State{
		Name:        widget.Name,
		Type:        "make-http-request",
		Transitions: transitions,
		Properties:  widget.Properties,
	}, nil
}
