// Package widgets contains auto-generated files. DO NOT MODIFY
package widgets

import (
	"fmt"

	"github.com/RJPearson94/twilio-sdk-go/studio/flow"
	"github.com/RJPearson94/twilio-sdk-go/studio/properties"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

type CapturePaymentsNextTransitions struct {
	Hangup            *string
	MaxFailedAttempts *string
	PayInterrupted    *string
	ProviderError     *string
	Success           *string
	ValidationError   *string
}

type CapturePaymentsParameter struct {
	Key   string `validate:"required" json:"key"`
	Value string `validate:"required" json:"value"`
}

type CapturePaymentsProperties struct {
	Currency            *string                     `json:"currency,omitempty"`
	Description         *string                     `json:"description,omitempty"`
	Language            *string                     `json:"language,omitempty"`
	MaxAttempts         *int                        `json:"max_attempts,omitempty"`
	MinPostalCodeLength *int                        `json:"min_postal_code_length,omitempty"`
	Offset              *properties.Offset          `json:"offset,omitempty"`
	Parameters          *[]CapturePaymentsParameter `json:"parameters,omitempty"`
	PaymentAmount       *string                     `json:"payment_amount,omitempty"`
	PaymentConnector    *string                     `json:"payment_connector,omitempty"`
	PaymentMethod       *string                     `json:"payment_method,omitempty"`
	PaymentTokenType    *string                     `json:"payment_token_type,omitempty"`
	PostalCode          *string                     `json:"postal_code,omitempty"`
	SecurityCode        *bool                       `json:"security_code,omitempty"`
	Timeout             *int                        `json:"timeout,omitempty"`
	ValidCardTypes      *[]string                   `json:"valid_card_types,omitempty"`
}

type CapturePayments struct {
	NextTransitions CapturePaymentsNextTransitions
	Properties      CapturePaymentsProperties `validate:"required"`
	Name            string                    `validate:"required"`
}

// Validate checks the widget is correctly configured
func (widget CapturePayments) Validate() error {
	if err := utils.ValidateInput(widget); err != nil {
		return fmt.Errorf("Invalid input supplied. Errors %s", err.Error())
	}
	return nil
}

// ToState returns a populated Studio Widget State struct
func (widget CapturePayments) ToState() (*flow.State, error) {
	transitions := []flow.Transition{
		{
			Event: "hangup",
			Next:  widget.NextTransitions.Hangup,
		},
		{
			Event: "maxFailedAttempts",
			Next:  widget.NextTransitions.MaxFailedAttempts,
		},
		{
			Event: "payInterrupted",
			Next:  widget.NextTransitions.PayInterrupted,
		},
		{
			Event: "providerError",
			Next:  widget.NextTransitions.ProviderError,
		},
		{
			Event: "success",
			Next:  widget.NextTransitions.Success,
		},
		{
			Event: "validationError",
			Next:  widget.NextTransitions.ValidationError,
		},
	}

	return &flow.State{
		Name:        widget.Name,
		Type:        "capture-payments",
		Transitions: transitions,
		Properties:  widget.Properties,
	}, nil
}
