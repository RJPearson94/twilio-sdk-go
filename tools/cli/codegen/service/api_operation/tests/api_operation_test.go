package tests

import (
	"encoding/json"
	"io/ioutil"

	"github.com/Jeffail/gabs/v2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	apioperation "github.com/RJPearson94/twilio-sdk-go-tools/cli/codegen/service/api_operation"
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

		Describe("When the api operation is generated with query params", func() {
			goldenData, _ := ioutil.ReadFile("testdata/apiOperationWithQueryParams.golden")
			apiOperationJSON, _ := ioutil.ReadFile("testdata/apiOperationWithQueryParams.json")
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

		Describe("When the api operation is generated with beta config", func() {
			goldenData, _ := ioutil.ReadFile("testdata/betaApiOperation.golden")
			apiOperationJSON, _ := ioutil.ReadFile("testdata/betaApiOperation.json")
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

	Describe("Given a snippet of api json", func() {
		Describe("When the json is translated", func() {
			apiOperationStructureJSON, _ := ioutil.ReadFile("testdata/apiOperationStructure.json")

			apiOperationJSON, _ := ioutil.ReadFile("testdata/apiOperation.json")
			apiOperationData, _ := gabs.ParseJSON(apiOperationJSON)

			resp, err := apioperation.Translate(apiOperationStructureJSON)

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the response should match the golden data", func() {
				Expect(*resp).To(Equal(apiOperationData.Data()))
			})
		})

		Describe("When the json with arrays is translated", func() {
			apiOperationStructureArraysJSON, _ := ioutil.ReadFile("testdata/apiOperationStructureArrays.json")

			apiOperationArraysJSON, _ := ioutil.ReadFile("testdata/apiOperationArrays.json")
			apiOperationArraysData, _ := gabs.ParseJSON(apiOperationArraysJSON)

			resp, err := apioperation.Translate(apiOperationStructureArraysJSON)

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the response should match the golden data", func() {
				Expect(*resp).To(Equal(apiOperationArraysData.Data()))
			})
		})

		Describe("When the json with a map with nested structure array is translated", func() {
			apiOperationStructureMapWithNestedStructureArrayJSON, _ := ioutil.ReadFile("testdata/apiOperationStructureMapWithNestedStructureArray.json")

			apiOperationMapWithNestedStructureArrayJSON, _ := ioutil.ReadFile("testdata/apiOperationMapWithNestedStructureArray.json")
			apiOperationMapWithNestedStructureArrayData, _ := gabs.ParseJSON(apiOperationMapWithNestedStructureArrayJSON)

			resp, err := apioperation.Translate(apiOperationStructureMapWithNestedStructureArrayJSON)

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the response should match the golden data", func() {
				Expect(*resp).To(Equal(apiOperationMapWithNestedStructureArrayData.Data()))
			})
		})

		Describe("When the json with query params is translated", func() {
			apiOperationStructureMapWithQueryParamsJSON, _ := ioutil.ReadFile("testdata/apiOperationStructureWithQueryParams.json")

			apiOperationMapWithQueryParamsJSON, _ := ioutil.ReadFile("testdata/apiOperationMapWithQueryParams.json")
			apiOperationMapWithQueryParamsData, _ := gabs.ParseJSON(apiOperationMapWithQueryParamsJSON)

			resp, err := apioperation.Translate(apiOperationStructureMapWithQueryParamsJSON)

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the response should match the golden data", func() {
				Expect(*resp).To(Equal(apiOperationMapWithQueryParamsData.Data()))
			})
		})
	})
})
