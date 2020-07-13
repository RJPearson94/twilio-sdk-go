package tests

import (
	"encoding/json"
	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/RJPearson94/twilio-sdk-go/tools/cli/codegen/client"
)

var _ = Describe("Client CodeGen", func() {
	Describe("Given I need to generate a client", func() {
		Describe("When the client is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/client.golden")
			clientJSON, _ := ioutil.ReadFile("testdata/client.json")
			var clientData interface{}
			_ = json.Unmarshal(clientJSON, &clientData)

			resp, err := client.Generate(clientData, false)

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the response should match the golden data", func() {
				Expect(string(*resp)).To(Equal(string(goldenData)))
			})
		})
	})
})
