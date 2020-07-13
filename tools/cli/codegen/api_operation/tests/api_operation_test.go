package tests

import (
	"encoding/json"
	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	apioperation "github.com/RJPearson94/twilio-sdk-go/tools/cli/codegen/api_operation"
)

var _ = Describe("API Operation CodeGen", func() {
	Describe("Given I need to generate a api operation", func() {
		Describe("When the api operation is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/apiOperation.golden")
			apiOperationJSON, _ := ioutil.ReadFile("testdata/apiOperation.json")
			var apiOperationData interface{}
			_ = json.Unmarshal(apiOperationJSON, &apiOperationData)

			resp, err := apioperation.Generate(apiOperationData, true)

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the response should match the golden data", func() {
				Expect(string(*resp)).To(Equal(string(goldenData)))
			})
		})
	})
})
