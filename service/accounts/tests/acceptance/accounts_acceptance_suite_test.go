// +build acceptance accounts_acceptance

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
// 3) A public key which can be uploaded to Twilio set as an environment variable - TWILIO_PUBLIC_KEY
// 4) An AWS access key ID which can be uploaded to Twilio set as an environment variable - TWILIO_AWS_ACCESS_KEY_ID
// 5) An AWS secret access key which can be uploaded to Twilio set as an environment variable - TWILIO_AWS_SECRET_ACCESS_KEY

func TestAcceptance(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "API Acceptance Test Suite")
}

var _ = BeforeSuite(func() {
	variables := []string{
		"TWILIO_ACCOUNT_SID",
		"TWILIO_AUTH_TOKEN",
		"TWILIO_PUBLIC_KEY",
		"TWILIO_AWS_ACCESS_KEY_ID",
		"TWILIO_AWS_SECRET_ACCESS_KEY",
	}

	for _, variable := range variables {
		if value := os.Getenv(variable); value == "" {
			Fail(fmt.Sprintf("`%s` are required for running acceptance tests", variable))
		}
	}
})
