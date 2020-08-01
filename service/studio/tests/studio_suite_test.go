// +build unit

package tests

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestV2(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Studio Test Suite")
}
