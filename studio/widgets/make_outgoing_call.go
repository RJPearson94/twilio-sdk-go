// Package widgets contains auto-generated files. DO NOT MODIFY
package widgets

import (
	"fmt"

	"github.com/RJPearson94/twilio-sdk-go/studio/flow"
	"github.com/RJPearson94/twilio-sdk-go/studio/properties"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

type MakeOutgoingCallNextTransitions struct {
	Answered *string
	Busy     *string
	Failed   *string
	NoAnswer *string
}

type MakeOutgoingCallProperties struct {
	DetectAnsweringMachine             *bool              `json:"detect_answering_machine,omitempty"`
	From                               string             `validate:"required" json:"from"`
	MachineDetection                   *string            `json:"machine_detection,omitempty"`
	MachineDetectionSilenceTimeout     *string            `json:"machine_detection_silence_timeout,omitempty"`
	MachineDetectionSpeechEndThreshold *string            `json:"machine_detection_speech_end_threshold,omitempty"`
	MachineDetectionSpeechThreshold    *string            `json:"machine_detection_speech_threshold,omitempty"`
	MachineDetectionTimeout            *string            `json:"machine_detection_timeout,omitempty"`
	Offset                             *properties.Offset `json:"offset,omitempty"`
	Record                             *bool              `json:"record,omitempty"`
	RecordingChannels                  *string            `json:"recording_channels,omitempty"`
	RecordingStatusCallbackURL         *string            `json:"recording_status_callback,omitempty"`
	SendDigits                         *string            `json:"send_digits,omitempty"`
	SipAuthPassword                    *string            `json:"sip_auth_password,omitempty"`
	SipAuthUsername                    *string            `json:"sip_auth_username,omitempty"`
	Timeout                            *int               `json:"timeout,omitempty"`
	To                                 string             `validate:"required" json:"to"`
	Trim                               *string            `json:"trim,omitempty"`
}

type MakeOutgoingCall struct {
	NextTransitions MakeOutgoingCallNextTransitions
	Properties      MakeOutgoingCallProperties `validate:"required"`
	Name            string                     `validate:"required"`
}

// Validate checks the widget is correctly configured
func (widget MakeOutgoingCall) Validate() error {
	if err := utils.ValidateInput(widget); err != nil {
		return fmt.Errorf("Invalid input supplied. Errors %s", err.Error())
	}
	return nil
}

// ToState returns a populated Studio Widget State struct
func (widget MakeOutgoingCall) ToState() (*flow.State, error) {
	transitions := []flow.Transition{
		{
			Event: "answered",
			Next:  widget.NextTransitions.Answered,
		},
		{
			Event: "busy",
			Next:  widget.NextTransitions.Busy,
		},
		{
			Event: "failed",
			Next:  widget.NextTransitions.Failed,
		},
		{
			Event: "noAnswer",
			Next:  widget.NextTransitions.NoAnswer,
		},
	}

	return &flow.State{
		Name:        widget.Name,
		Type:        "make-outgoing-call-v2",
		Transitions: transitions,
		Properties:  widget.Properties,
	}, nil
}
