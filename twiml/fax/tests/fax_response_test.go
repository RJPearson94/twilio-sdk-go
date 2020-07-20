package tests

import (
	"io/ioutil"
	"regexp"

	"github.com/RJPearson94/twilio-sdk-go/twiml/fax"
	"github.com/RJPearson94/twilio-sdk-go/twiml/fax/verbs"
	"github.com/RJPearson94/twilio-sdk-go/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var ignoreRegex = regexp.MustCompile(`\r?\n|\s`)

var _ = Describe("Fax Response TwiML", func() {
	Describe("Given I need to generate a fax response with an action receive attribute", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/receive.golden.xml")

			response := fax.New()
			response.ReceiveWithAttributes(verbs.ReceiveAttributes{
				Action: utils.String("http://localhost/action"),
			})
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a fax response with receive media attributes", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/receiveMedia.golden.xml")

			response := fax.New()
			response.ReceiveWithAttributes(verbs.ReceiveAttributes{
				MediaType:  utils.String("application/pdf"),
				StoreMedia: utils.Bool(true),
			})
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a fax response with a reject", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/reject.golden.xml")

			response := fax.New()
			response.Reject()
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a fax response with receive and reject verbs", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/receiveAndReject.golden.xml")

			response := fax.New()
			response.ReceiveWithAttributes(verbs.ReceiveAttributes{
				Action: utils.String("http://localhost/action"),
			})
			response.Reject()
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})
})

func CompareXML(actual string, expected string) {
	Expect(ignoreRegex.ReplaceAllString(actual, "")).To(Equal(ignoreRegex.ReplaceAllString(expected, "")))
}
