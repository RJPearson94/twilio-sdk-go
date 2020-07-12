package tests

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAPIOperation(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "API Operation CodeGen Suite")
}
