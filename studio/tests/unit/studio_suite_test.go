// +build unit

package tests

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestStudio(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Studio Suite")
}
