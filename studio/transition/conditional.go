package transition

import (
	"fmt"

	"github.com/RJPearson94/twilio-sdk-go/studio/flow"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

type Conditional struct {
	Next       string            `validate:"required"`
	Conditions *[]flow.Condition `validate:"required"`
}

// Validate checks the conditional is correctly configured
func (conditional Conditional) Validate() error {
	if err := utils.ValidateInput(conditional); err != nil {
		return fmt.Errorf("Invalid input supplied. Errors %s", err.Error())
	}
	return nil
}
