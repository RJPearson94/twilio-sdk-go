package tests

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestVoiceResponse(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TwiML Voice Response Suite")
}
