package tests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/Jeffail/gabs/v2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	client "github.com/RJPearson94/twilio-sdk-go/tools/cli/codegen/client"
)

var _ = Describe("Client CodeGen", func() {
	Describe("Given I need to generate a client", func() {
		Describe("When the client is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/client.golden")
			clientJSON, _ := ioutil.ReadFile("testdata/client.json")
			var clientData interface{}
			_ = json.Unmarshal(clientJSON, &clientData)

			resp, err := client.Generate(clientData, true)

			fmt.Println(string(*resp))

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the response should match the golden data", func() {
				Expect(string(*resp)).To(Equal(string(goldenData)))
			})
		})
	})

	Describe("Given the api json", func() {
		Describe("When the json is translated", func() {
			apiJSON, _ := ioutil.ReadFile("testdata/client.json")

			clientJSON, _ := ioutil.ReadFile("testdata/translationOutput.json")
			clientData, _ := gabs.ParseJSON(clientJSON)

			resp, err := client.Translate(apiJSON)

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the response should match the golden data", func() {
				Expect(*resp).To(Equal(clientData.Data()))
			})
		})
	})
})
