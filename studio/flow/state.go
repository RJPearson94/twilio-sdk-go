package flow

import (
	"encoding/json"
	"fmt"

	"github.com/RJPearson94/twilio-sdk-go/utils"
)

type Condition struct {
	Arguments    []string `validate:"required" json:"arguments"`
	FriendlyName string   `validate:"required" json:"friendly_name"`
	Type         string   `validate:"required" json:"type"`
	Value        string   `validate:"required" json:"value"`
}

type Transition struct {
	Event      string       `validate:"required" json:"event"`
	Next       *string      `json:"next,omitempty"`
	Conditions *[]Condition `json:"conditions,omitempty"`
}

type State struct {
	Name        string       `validate:"required" json:"name"`
	Properties  interface{}  `validate:"required" json:"properties"`
	Transitions []Transition `validate:"required" json:"transitions"`
	Type        string       `validate:"required" json:"type"`
}

// Validate checks the state is correctly configured
func (state State) Validate() error {
	if err := utils.ValidateInput(state); err != nil {
		return fmt.Errorf("Invalid input supplied. Errors %s", err.Error())
	}
	return nil
}

// ToString marshalls the struct to a json string
func (state State) ToString() (*string, error) {
	bytesArray, err := json.Marshal(state)
	if err != nil {
		return nil, err
	}
	return utils.String(string(bytesArray)), nil
}
