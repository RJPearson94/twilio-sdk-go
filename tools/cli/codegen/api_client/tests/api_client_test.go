package tests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/Jeffail/gabs/v2"
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

			resp, err := apiclient.Generate(apiClientData, true)

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the response should match the golden data", func() {
				fmt.Println(string(*resp))
				Expect(string(*resp)).To(Equal(string(goldenData)))
			})
		})
	})

	Describe("Given the api json", func() {
		Describe("When the json is translated", func() {
			apiJSON, _ := ioutil.ReadFile("testdata/subClient.json")

			apiClientJSON, _ := ioutil.ReadFile("testdata/translationOutput.json")
			apiClientData, _ := gabs.ParseJSON(apiClientJSON)

			resp, err := apiclient.Translate(apiJSON)

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the response should match the golden data", func() {
				Expect(*resp).To(Equal(apiClientData.Data()))
			})
		})
	})
})
