// +build acceptance flex_acceptance

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
// 3) SID for a Twilio Chat Service which is configured for Flex set as an environment variable - TWILIO_PHONE_NUMBER_SID
// 3) SID for a default plugin configuration to superseed the test deployment with set as an environment variable - TWILIO_FLEX_DEFAULT_CONFIGURATION

func TestAcceptance(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Flex Acceptance Test Suite")
}

var _ = BeforeSuite(func() {
	variables := []string{
		"TWILIO_ACCOUNT_SID",
		"TWILIO_AUTH_TOKEN",
		"TWILIO_FLEX_CHANNEL_SERVICE_SID",
		"TWILIO_FLEX_DEFAULT_CONFIGURATION",
	}

	for _, variable := range variables {
		if value := os.Getenv(variable); value == "" {
			Fail(fmt.Sprintf("`%s` are required for running acceptance tests", variable))
		}
	}
})
