package credentials

import (
	"fmt"

	"github.com/RJPearson94/twilio-sdk-go/utils"
)

type APIKey struct {
	Sid   string `validate:"required,startswith=SK"`
	Value string `validate:"required"`
}

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
