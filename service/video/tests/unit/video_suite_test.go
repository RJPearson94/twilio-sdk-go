// +build unit

package tests

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestVideoV1(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Video V1 Test Suite")
}
