package tests

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMessagingResponse(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TwiML Messaging Response Suite")
}
