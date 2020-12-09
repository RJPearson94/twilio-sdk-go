// +build acceptance autopilot_acceptance

package acceptance

import (
	"fmt"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// Prerequisites

// 1) Twilio Account SID set as an environment variable - TWILIO_ACCOUNT_SID
// 2) Twilio Auth Token set as an environment variable - TWILIO_AUTH_TOKEN
// 3) Twilio Phone Number set as an environment variable - TWILIO_PHONE_NUMBER
// 4) Phone Number to send messages to set as an environment variable - DESTINATION_PHONE_NUMBER
// 5) Your name as an environment variable - TWILIO_CUSTOMER_NAME
// 6) The street name of your address as an environment variable - TWILIO_ADDRESS_STREET
// 7) The city for your address as an environment variable - TWILIO_ADDRESS_CITY
// 8) The region for your address as an environment variable - TWILIO_ADDRESS_REGION
// 9) The postal code for your address as an environment variable - TWILIO_ADDRESS_POSTAL_CODE
// 10) The iso country code for your address as an environment variable - TWILIO_ADDRESS_ISO_COUNTRY

func TestAcceptance(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "API Acceptance Test Suite")
}

var _ = BeforeSuite(func() {
	variables := []string{
		"TWILIO_ACCOUNT_SID",
		"TWILIO_AUTH_TOKEN",
		"TWILIO_PHONE_NUMBER",
		"DESTINATION_PHONE_NUMBER",
		"TWILIO_CUSTOMER_NAME",
		"TWILIO_ADDRESS_STREET",
		"TWILIO_ADDRESS_CITY",
		"TWILIO_ADDRESS_REGION",
		"TWILIO_ADDRESS_POSTAL_CODE",
		"TWILIO_ADDRESS_ISO_COUNTRY",
	}

	for _, variable := range variables {
		if value := os.Getenv(variable); value == "" {
			Fail(fmt.Sprintf("`%s` are required for running acceptance tests", variable))
		}
	}
})
