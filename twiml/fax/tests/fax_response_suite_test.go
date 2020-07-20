package tests

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFaxResponse(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TwiML Fax Response Suite")
}
