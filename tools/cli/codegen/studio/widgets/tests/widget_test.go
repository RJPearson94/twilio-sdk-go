package tests

import (
	"encoding/json"
	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	widget "github.com/RJPearson94/twilio-sdk-go-tools/cli/codegen/studio/widgets"
)

var _ = Describe("Widget CodeGen", func() {
	Describe("Given I need to generate a widget", func() {
		Describe("When the widget is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/widget.golden")
			widgetJSON, _ := ioutil.ReadFile("testdata/widget.json")
			var widgetData interface{}
			_ = json.Unmarshal(widgetJSON, &widgetData)

			resp, err := widget.Generate(widgetData, true)

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the response should match the golden data", func() {
				Expect(string(*resp)).To(Equal(string(goldenData)))
			})
		})
	})
})
