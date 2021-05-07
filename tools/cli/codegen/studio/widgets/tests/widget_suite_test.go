// +build unit

package tests

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestWidget(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Widget CodeGen Suite")
}
