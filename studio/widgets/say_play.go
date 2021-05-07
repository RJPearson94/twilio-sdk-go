// Package widgets contains auto-generated files. DO NOT MODIFY
package widgets

import (
	"fmt"

	"github.com/RJPearson94/twilio-sdk-go/studio/flow"
	"github.com/RJPearson94/twilio-sdk-go/studio/properties"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

type SayPlayNextTransitions struct {
	AudioComplete *string
}

type SayPlayProperties struct {
	Digits   *string            `json:"digits,omitempty"`
	Language *string            `json:"language,omitempty"`
	Loop     *int               `json:"loop,omitempty"`
	Offset   *properties.Offset `json:"offset,omitempty"`
	Play     *string            `json:"play,omitempty"`
	Say      *string            `json:"say,omitempty"`
	Voice    *string            `json:"voice,omitempty"`
}

type SayPlay struct {
	NextTransitions SayPlayNextTransitions
	Properties      SayPlayProperties `validate:"required"`
	Name            string            `validate:"required"`
}

// Validate checks the widget is correctly configured
func (widget SayPlay) Validate() error {
	if err := utils.ValidateInput(widget); err != nil {
		return fmt.Errorf("Invalid input supplied. Errors %s", err.Error())
	}
	return nil
}

// ToState returns a populated Studio Widget State struct
func (widget SayPlay) ToState() (*flow.State, error) {
	transitions := []flow.Transition{
		{
			Event: "audioComplete",
			Next:  widget.NextTransitions.AudioComplete,
		},
	}

	return &flow.State{
		Name:        widget.Name,
		Type:        "say-play",
		Transitions: transitions,
		Properties:  widget.Properties,
	}, nil
}
