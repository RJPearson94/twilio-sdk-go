package credentials

import (
	"fmt"
	"strings"
)

// APIKey represents a Twilio API Key which can be used to authenticate
// against the Twilio APIs
type APIKey struct {
	Account string
	Sid     string
	Value   string
}

// Validate ensures the API Keys is valid
func (apiKey APIKey) Validate() error {
	validationErrors := make([]string, 0)
	if apiKey.Account == "" {
		validationErrors = append(validationErrors, "Account SID is required")
	} else if !strings.HasPrefix(apiKey.Account, "AC") {
		validationErrors = append(validationErrors, fmt.Sprintf("Account SID (%s) must start with AC", apiKey.Account))
	}

	if apiKey.Sid == "" {
		validationErrors = append(validationErrors, "SID is required")
	} else if !strings.HasPrefix(apiKey.Sid, "SK") {
		validationErrors = append(validationErrors, fmt.Sprintf("SID (%s) must start with SK", apiKey.Sid))
	}

	if apiKey.Value == "" {
		validationErrors = append(validationErrors, "Value is required")
	}

	if len(validationErrors) > 0 {
		return fmt.Errorf("API Key details specified are invalid. Validation errors: [%s]", strings.Join(validationErrors, ", "))
	}
	return nil
}

func (apiKey APIKey) AccountSid() string {
	return apiKey.Account
}

func (apiKey APIKey) username() string {
	return apiKey.Sid
}

func (apiKey APIKey) password() string {
	return apiKey.Value
}
