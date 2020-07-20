package tests

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFaxV1(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Fax V1 Test Suite")
}
