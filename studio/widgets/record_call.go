// Package widgets contains auto-generated files. DO NOT MODIFY
package widgets

import (
	"fmt"

	"github.com/RJPearson94/twilio-sdk-go/studio/flow"
	"github.com/RJPearson94/twilio-sdk-go/studio/properties"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

type RecordCallNextTransitions struct {
	Failed  *string
	Success *string
}

type RecordCallProperties struct {
	Offset                        *properties.Offset `json:"offset,omitempty"`
	RecordCall                    bool               `json:"record_call"`
	RecordingChannels             *string            `json:"recording_channels,omitempty"`
	RecordingStatusCallbackEvents *string            `json:"recording_status_callback_events,omitempty"`
	RecordingStatusCallbackMethod *string            `json:"recording_status_callback_method,omitempty"`
	RecordingStatusCallbackURL    *string            `json:"recording_status_callback,omitempty"`
	Trim                          *string            `json:"trim,omitempty"`
}

type RecordCall struct {
	NextTransitions RecordCallNextTransitions
	Properties      RecordCallProperties `validate:"required"`
	Name            string               `validate:"required"`
}

// Validate checks the widget is correctly configured
func (widget RecordCall) Validate() error {
	if err := utils.ValidateInput(widget); err != nil {
		return fmt.Errorf("Invalid input supplied. Errors %s", err.Error())
	}
	return nil
}

// ToState returns a populated Studio Widget State struct
func (widget RecordCall) ToState() (*flow.State, error) {
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
		Type:        "record-call",
		Transitions: transitions,
		Properties:  widget.Properties,
	}, nil
}
