package credentials

import (
	"fmt"

	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// Account represents Twilio Account Credentials which can be used to authenticate
// against the Twilio APIs
type Account struct {
	Sid       string `validate:"required,startswith=AC"`
	AuthToken string `validate:"required"`
}

// Validate ensures the account credentials are valid
func (account Account) Validate() error {
	if err := utils.ValidateInput(&account); err != nil {
		return fmt.Errorf("Account Details Specified are invalid")
	}
	return nil
}

func (account Account) username() string {
	return account.Sid
}

func (account Account) password() string {
	return account.AuthToken
}
