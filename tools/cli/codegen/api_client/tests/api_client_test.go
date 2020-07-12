package tests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	apiclient "github.com/RJPearson94/twilio-sdk-go/tools/cli/codegen/api_client"
)

var _ = Describe("API Client CodeGen", func() {
	Describe("Given I need to generate a api client", func() {
		Describe("When the api client is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/apiClient.golden")
			apiClientJSON, _ := ioutil.ReadFile("testdata/apiClient.json")
			var apiClientData interface{}
			_ = json.Unmarshal(apiClientJSON, &apiClientData)

			resp, err := apiclient.Generate(apiClientData)

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the response should match the golden data", func() {
				fmt.Println(string(*resp))
				fmt.Println(string(goldenData))
				Expect(string(*resp)).To(Equal(string(goldenData)))
			})
		})
	})
})
