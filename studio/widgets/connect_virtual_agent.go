// Package widgets contains auto-generated files. DO NOT MODIFY
package widgets

import (
	"fmt"

	"github.com/RJPearson94/twilio-sdk-go/studio/flow"
	"github.com/RJPearson94/twilio-sdk-go/studio/properties"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

type ConnectVirtualAgentNextTransitions struct {
	Hangup *string
	Return *string
}

type ConnectVirtualAgentProperties struct {
	Connector         string             `validate:"required" json:"connector"`
	Language          *string            `json:"language,omitempty"`
	Offset            *properties.Offset `json:"offset,omitempty"`
	SentimentAnalysis *string            `json:"sentiment_analysis,omitempty"`
	StatusCallbackURL *string            `json:"status_callback,omitempty"`
}

type ConnectVirtualAgent struct {
	NextTransitions ConnectVirtualAgentNextTransitions
	Properties      ConnectVirtualAgentProperties `validate:"required"`
	Name            string                        `validate:"required"`
}

// Validate checks the widget is correctly configured
func (widget ConnectVirtualAgent) Validate() error {
	if err := utils.ValidateInput(widget); err != nil {
		return fmt.Errorf("Invalid input supplied. Errors %s", err.Error())
	}
	return nil
}

// ToState returns a populated Studio Widget State struct
func (widget ConnectVirtualAgent) ToState() (*flow.State, error) {
	transitions := []flow.Transition{
		{
			Event: "hangup",
			Next:  widget.NextTransitions.Hangup,
		},
		{
			Event: "return",
			Next:  widget.NextTransitions.Return,
		},
	}

	return &flow.State{
		Name:        widget.Name,
		Type:        "connect-virtual-agent",
		Transitions: transitions,
		Properties:  widget.Properties,
	}, nil
}
