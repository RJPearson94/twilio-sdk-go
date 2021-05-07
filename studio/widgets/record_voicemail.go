// Package widgets contains auto-generated files. DO NOT MODIFY
package widgets

import (
	"fmt"

	"github.com/RJPearson94/twilio-sdk-go/studio/flow"
	"github.com/RJPearson94/twilio-sdk-go/studio/properties"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

type RecordVoicemailNextTransitions struct {
	Hangup            *string
	NoAudio           *string
	RecordingComplete *string
}

type RecordVoicemailProperties struct {
	FinishOnKey                *string            `json:"finish_on_key,omitempty"`
	MaxLength                  *int               `json:"max_length,omitempty"`
	Offset                     *properties.Offset `json:"offset,omitempty"`
	PlayBeep                   *string            `json:"play_beep,omitempty"`
	RecordingStatusCallbackURL *string            `json:"recording_status_callback_url,omitempty"`
	Timeout                    *int               `json:"timeout,omitempty"`
	Transcribe                 *bool              `json:"transcribe,omitempty"`
	TranscriptionCallbackURL   *string            `json:"transcription_callback_url,omitempty"`
	Trim                       *string            `json:"trim,omitempty"`
}

type RecordVoicemail struct {
	NextTransitions RecordVoicemailNextTransitions
	Properties      RecordVoicemailProperties `validate:"required"`
	Name            string                    `validate:"required"`
}

// Validate checks the widget is correctly configured
func (widget RecordVoicemail) Validate() error {
	if err := utils.ValidateInput(widget); err != nil {
		return fmt.Errorf("Invalid input supplied. Errors %s", err.Error())
	}
	return nil
}

// ToState returns a populated Studio Widget State struct
func (widget RecordVoicemail) ToState() (*flow.State, error) {
	transitions := []flow.Transition{
		{
			Event: "hangup",
			Next:  widget.NextTransitions.Hangup,
		},
		{
			Event: "noAudio",
			Next:  widget.NextTransitions.NoAudio,
		},
		{
			Event: "recordingComplete",
			Next:  widget.NextTransitions.RecordingComplete,
		},
	}

	return &flow.State{
		Name:        widget.Name,
		Type:        "record-voicemail",
		Transitions: transitions,
		Properties:  widget.Properties,
	}, nil
}
