// +build acceptance proxy_acceptance

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
// 3) SID for a Twilio Phone Number set in the environment variaable - TWILIO_PHONE_NUMBER_SID
// 4) Phone Number to send messages to - DESTINATION_PHONE_NUMBER

func TestAcceptance(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Proxy Acceptance Test Suite")
}

var _ = BeforeSuite(func() {
	variables := []string{
		"TWILIO_ACCOUNT_SID",
		"TWILIO_AUTH_TOKEN",
		"TWILIO_PHONE_NUMBER_SID",
		"DESTINATION_PHONE_NUMBER",
	}

	for _, variable := range variables {
		if value := os.Getenv(variable); value == "" {
			Fail(fmt.Sprintf("`%s` are required for running acceptance tests", variable))
		}
	}
})
