// Package widgets contains auto-generated files. DO NOT MODIFY
package widgets

import (
	"fmt"

	"github.com/RJPearson94/twilio-sdk-go/studio/flow"
	"github.com/RJPearson94/twilio-sdk-go/studio/properties"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

type SendMessageNextTransitions struct {
	Failed *string
	Sent   *string
}

type SendMessageProperties struct {
	Attributes *string            `json:"attributes,omitempty"`
	Body       string             `validate:"required" json:"body"`
	Channel    *string            `json:"channel,omitempty"`
	From       string             `validate:"required" json:"from"`
	MediaURL   *string            `json:"media_url,omitempty"`
	Offset     *properties.Offset `json:"offset,omitempty"`
	Service    *string            `json:"service,omitempty"`
	To         string             `validate:"required" json:"to"`
}

type SendMessage struct {
	NextTransitions SendMessageNextTransitions
	Properties      SendMessageProperties `validate:"required"`
	Name            string                `validate:"required"`
}

// Validate checks the widget is correctly configured
func (widget SendMessage) Validate() error {
	if err := utils.ValidateInput(widget); err != nil {
		return fmt.Errorf("Invalid input supplied. Errors %s", err.Error())
	}
	return nil
}

// ToState returns a populated Studio Widget State struct
func (widget SendMessage) ToState() (*flow.State, error) {
	transitions := []flow.Transition{
		{
			Event: "failed",
			Next:  widget.NextTransitions.Failed,
		},
		{
			Event: "sent",
			Next:  widget.NextTransitions.Sent,
		},
	}

	return &flow.State{
		Name:        widget.Name,
		Type:        "send-message",
		Transitions: transitions,
		Properties:  widget.Properties,
	}, nil
}
