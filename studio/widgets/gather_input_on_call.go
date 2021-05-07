// Package widgets contains auto-generated files. DO NOT MODIFY
package widgets

import (
	"fmt"

	"github.com/RJPearson94/twilio-sdk-go/studio/flow"
	"github.com/RJPearson94/twilio-sdk-go/studio/properties"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

type GatherInputOnCallNextTransitions struct {
	Keypress *string
	Speech   *string
	Timeout  *string
}

type GatherInputOnCallProperties struct {
	FinishOnKey     *string            `json:"finish_on_key,omitempty"`
	GatherLanguage  *string            `json:"gather_language,omitempty"`
	Hints           *string            `json:"hints,omitempty"`
	Language        *string            `json:"language,omitempty"`
	Loop            *int               `json:"loop,omitempty"`
	NumberOfDigits  *int               `json:"number_of_digits,omitempty"`
	Offset          *properties.Offset `json:"offset,omitempty"`
	Play            *string            `json:"play,omitempty"`
	ProfanityFilter *string            `json:"profanity_filter,omitempty"`
	Say             *string            `json:"say,omitempty"`
	SpeechModel     *string            `json:"speech_model,omitempty"`
	SpeechTimeout   *string            `json:"speech_timeout,omitempty"`
	StopGather      *bool              `json:"stop_gather,omitempty"`
	Timeout         *int               `json:"timeout,omitempty"`
	Voice           *string            `json:"voice,omitempty"`
}

type GatherInputOnCall struct {
	NextTransitions GatherInputOnCallNextTransitions
	Properties      GatherInputOnCallProperties `validate:"required"`
	Name            string                      `validate:"required"`
}

// Validate checks the widget is correctly configured
func (widget GatherInputOnCall) Validate() error {
	if err := utils.ValidateInput(widget); err != nil {
		return fmt.Errorf("Invalid input supplied. Errors %s", err.Error())
	}
	return nil
}

// ToState returns a populated Studio Widget State struct
func (widget GatherInputOnCall) ToState() (*flow.State, error) {
	transitions := []flow.Transition{
		{
			Event: "keypress",
			Next:  widget.NextTransitions.Keypress,
		},
		{
			Event: "speech",
			Next:  widget.NextTransitions.Speech,
		},
		{
			Event: "timeout",
			Next:  widget.NextTransitions.Timeout,
		},
	}

	return &flow.State{
		Name:        widget.Name,
		Type:        "gather-input-on-call",
		Transitions: transitions,
		Properties:  widget.Properties,
	}, nil
}
