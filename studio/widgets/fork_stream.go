// Package widgets contains auto-generated files. DO NOT MODIFY
package widgets

import (
	"fmt"

	"github.com/RJPearson94/twilio-sdk-go/studio/flow"
	"github.com/RJPearson94/twilio-sdk-go/studio/properties"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

type ForkStreamNextTransitions struct {
	Next *string
}

type ForkStreamStreamParameter struct {
	Key   string `validate:"required" json:"key"`
	Value string `validate:"required" json:"value"`
}

type ForkStreamProperties struct {
	Offset              *properties.Offset           `json:"offset,omitempty"`
	StreamAction        string                       `validate:"required" json:"stream_action"`
	StreamConnector     *string                      `json:"stream_connector,omitempty"`
	StreamName          *string                      `json:"stream_name,omitempty"`
	StreamParameters    *[]ForkStreamStreamParameter `json:"stream_parameters,omitempty"`
	StreamTrack         *string                      `json:"stream_track,omitempty"`
	StreamTransportType *string                      `json:"stream_transport_type,omitempty"`
	StreamURL           *string                      `json:"stream_url,omitempty"`
}

type ForkStream struct {
	NextTransitions ForkStreamNextTransitions
	Properties      ForkStreamProperties `validate:"required"`
	Name            string               `validate:"required"`
}

// Validate checks the widget is correctly configured
func (widget ForkStream) Validate() error {
	if err := utils.ValidateInput(widget); err != nil {
		return fmt.Errorf("Invalid input supplied. Errors %s", err.Error())
	}
	return nil
}

// ToState returns a populated Studio Widget State struct
func (widget ForkStream) ToState() (*flow.State, error) {
	transitions := []flow.Transition{
		{
			Event: "next",
			Next:  widget.NextTransitions.Next,
		},
	}

	return &flow.State{
		Name:        widget.Name,
		Type:        "fork-stream",
		Transitions: transitions,
		Properties:  widget.Properties,
	}, nil
}
