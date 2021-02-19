package credentials

import (
	"fmt"
	"strings"
)

// Account represents Twilio Account Credentials which can be used to authenticate
// against the Twilio APIs
type Account struct {
	Sid       string
	AuthToken string
}

// Validate ensures the account credentials are valid
func (account Account) Validate() error {
	validationErrors := make([]string, 0)
	if account.Sid == "" {
		validationErrors = append(validationErrors, "SID is required")
	} else if !strings.HasPrefix(account.Sid, "AC") {
		validationErrors = append(validationErrors, fmt.Sprintf("SID (%s) must start with AC", account.Sid))
	}

	if account.AuthToken == "" {
		validationErrors = append(validationErrors, "Auth token is required")
	}

	if len(validationErrors) > 0 {
		return fmt.Errorf("Account details specified are invalid. Validation errors: [%s]", strings.Join(validationErrors, ", "))
	}
	return nil
}

func (account Account) AccountSid() string {
	return account.Sid
}

func (account Account) username() string {
	return account.Sid
}

func (account Account) password() string {
	return account.AuthToken
}
