package studio

import (
	"encoding/json"
	"fmt"

	"github.com/RJPearson94/twilio-sdk-go/studio/flow"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// FlowFlags are the advanced features configuration for the Studio Flow
type FlowFlags struct {
	AllowConcurrentCalls bool `json:"allow_concurrent_calls"`
}

// Flow is the struct to create a Studio Flow
type Flow struct {
	Description  string       `json:"description"`
	Flags        *FlowFlags   `json:"flags,omitempty"`
	InitialState string       `validate:"required" json:"initial_state"`
	States       []flow.State `validate:"required" json:"states"`
}

// Validate checks the flow is correctly configured
func (flow Flow) Validate() error {
	if err := utils.ValidateInput(flow); err != nil {
		return fmt.Errorf("Invalid input supplied. Errors %s", err.Error())
	}
	return nil
}

// ToString marshalls the struct to a json string
func (flow Flow) ToString() (*string, error) {
	bytesArray, err := json.Marshal(flow)
	if err != nil {
		return nil, err
	}
	return utils.String(string(bytesArray)), nil
}
