// +build acceptance sync_acceptance

package acceptance

import (
	"fmt"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// Prerequisites

// 1) Twilio Account SID set as an environment variaable - TWILIO_ACCOUNT_SID
// 2) Twilio Auth Token set as an environment variaable - TWILIO_AUTH_TOKEN

func TestAcceptance(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sync Acceptance Test Suite")
}

var _ = BeforeSuite(func() {
	variables := []string{
		"TWILIO_ACCOUNT_SID",
		"TWILIO_AUTH_TOKEN",
	}

	for _, variable := range variables {
		if value := os.Getenv(variable); value == "" {
			Fail(fmt.Sprintf("`%s` are required for running acceptance tests", variable))
		}
	}
})
