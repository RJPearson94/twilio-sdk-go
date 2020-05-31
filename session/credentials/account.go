package credentials

import (
	"fmt"

	"github.com/RJPearson94/twilio-sdk-go/utils"
)

type Account struct {
	Sid       string `validate:"required,startswith=AC"`
	AuthToken string `validate:"required"`
}

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
