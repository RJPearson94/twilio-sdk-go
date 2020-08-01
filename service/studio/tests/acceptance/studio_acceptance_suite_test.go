// +build acceptance studio_acceptance

package acceptance

import (
	"fmt"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// Prerequisites

// 1) Twilio Account SID set in the environment variaable - TWILIO_ACCOUNT_SID
// 2) Twilio Auth Token set in the environment variaable - TWILIO_AUTH_TOKEN
// 3) A url for a TwiML bin which has the following TwiML body
// <?xml version="1.0" encoding="UTF-8"?>
// <Response>
// 	<Pause length="10"/>
// </Response>
// The URL for the TwiML bin set in the environment variable - TWILIO_DELAY_TWIML_URL

func TestAcceptance(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Studio Acceptance Test Suite")
}

var _ = BeforeSuite(func() {
	variables := []string{
		"TWILIO_ACCOUNT_SID",
		"TWILIO_AUTH_TOKEN",
		"TWILIO_DELAY_TWIML_URL",
	}

	for _, variable := range variables {
		if value := os.Getenv(variable); value == "" {
			Fail(fmt.Sprintf("`%s` are required for running acceptance tests", variable))
		}
	}
})
