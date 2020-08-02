// +build unit

package tests

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestV2010(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "API Test Suite")
}
