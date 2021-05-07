// Package widgets contains auto-generated files. DO NOT MODIFY
package widgets

import (
	"fmt"

	"github.com/RJPearson94/twilio-sdk-go/studio/flow"
	"github.com/RJPearson94/twilio-sdk-go/studio/properties"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

type SendAndWaitForReplyNextTransitions struct {
	DeliveryFailure *string
	IncomingMessage string `validate:"required"`
	Timeout         *string
}

type SendAndWaitForReplyProperties struct {
	Attributes *string            `json:"attributes,omitempty"`
	Body       string             `validate:"required" json:"body"`
	Channel    *string            `json:"channel,omitempty"`
	From       string             `validate:"required" json:"from"`
	MediaURL   *string            `json:"media_url,omitempty"`
	Offset     *properties.Offset `json:"offset,omitempty"`
	Service    *string            `json:"service,omitempty"`
	Timeout    string             `validate:"required" json:"timeout"`
}

// SendAndWaitForReply widget allows you to send a SMS and wait for a reply
type SendAndWaitForReply struct {
	NextTransitions SendAndWaitForReplyNextTransitions
	Properties      SendAndWaitForReplyProperties `validate:"required"`
	Name            string                        `validate:"required"`
}

// Validate checks the widget is correctly configured
func (widget SendAndWaitForReply) Validate() error {
	if err := utils.ValidateInput(widget); err != nil {
		return fmt.Errorf("Invalid input supplied. Errors %s", err.Error())
	}
	return nil
}

// ToState returns a populated Studio Widget State struct
func (widget SendAndWaitForReply) ToState() (*flow.State, error) {
	transitions := []flow.Transition{
		{
			Event: "deliveryFailure",
			Next:  widget.NextTransitions.DeliveryFailure,
		},
		{
			Event: "incomingMessage",
			Next:  utils.String(widget.NextTransitions.IncomingMessage),
		},
		{
			Event: "timeout",
			Next:  widget.NextTransitions.Timeout,
		},
	}

	return &flow.State{
		Name:        widget.Name,
		Type:        "send-and-wait-for-reply",
		Transitions: transitions,
		Properties:  widget.Properties,
	}, nil
}
