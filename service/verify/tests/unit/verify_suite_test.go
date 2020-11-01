// +build unit

package tests

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestVerifyV2(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Verify V2 Test Suite")
}
