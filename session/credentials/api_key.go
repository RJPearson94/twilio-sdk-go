package credentials

import (
	"fmt"

	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// APIKey represents a Twilio API Key which can be used to authenticate
// against the Twilio APIs
type APIKey struct {
	Sid   string `validate:"required,startswith=SK"`
	Value string `validate:"required"`
}

// Validate ensures the API Keys is valid
func (apiKey APIKey) Validate() error {
	if err := utils.ValidateInput(&apiKey); err != nil {
		return fmt.Errorf("API Key Details Specified are invalid")
	}
	return nil
}

func (apiKey APIKey) username() string {
	return apiKey.Sid
}

func (apiKey APIKey) password() string {
	return apiKey.Value
}
