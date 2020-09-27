package tests

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jarcoal/httpmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/RJPearson94/twilio-sdk-go/service/lookups"
	"github.com/RJPearson94/twilio-sdk-go/service/lookups/v1/phone_number"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

var _ = Describe("Lookups V1", func() {
	creds, err := credentials.New(credentials.Account{
		Sid:       "ACxxx",
		AuthToken: "Test",
	})
	if err != nil {
		log.Panicf("%s", err)
	}

	lookupsSession := lookups.NewWithCredentials(creds).V1

	httpmock.ActivateNonDefault(lookupsSession.GetClient().GetRestyClient().GetClient())
	defer httpmock.DeactivateAndReset()

	Describe("Given I have a phone number", func() {
		phoneNumberClient := lookupsSession.PhoneNumber("+447111111111")

		Describe("When the phone number is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://lookups.twilio.com/v1/PhoneNumbers/+447111111111",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/phoneNumberResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := phoneNumberClient.Fetch(nil)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the phone number details response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				Expect(resp.CallerName).To(BeNil())
				Expect(resp.CountryCode).To(Equal("GB"))
				Expect(resp.PhoneNumber).To(Equal("+447111111111"))
				Expect(resp.NationalFormat).To(Equal("07111 111111"))
				Expect(resp.Carrier).To(BeNil())
				Expect(resp.AddOns).To(BeNil())
				Expect(resp.URL).To(Equal("https://lookups.twilio.com/v1/PhoneNumbers/+447111111111"))
			})
		})

		Describe("When the phone number with carrier and caller name details is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://lookups.twilio.com/v1/PhoneNumbers/+18551112222?Type=carrier&Type=caller-name",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/phoneNumberResponseWithCarrierAndCallerName.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := lookupsSession.PhoneNumber("+18551112222").Fetch(&phone_number.FetchPhoneNumberOptions{
				Type: &[]string{"carrier", "caller-name"},
			})
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the phone number details response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				Expect(resp.CallerName).To(Equal(&phone_number.FetchCallerNameResponse{
					CallerName: utils.String("TOLL FREE CALL"),
					CallerType: nil,
					ErrorCode:  nil,
				}))
				Expect(resp.CountryCode).To(Equal("US"))
				Expect(resp.PhoneNumber).To(Equal("+18551112222"))
				Expect(resp.NationalFormat).To(Equal("(855) 111-2222"))
				Expect(resp.Carrier).To(Equal(&phone_number.FetchCarrierResponse{
					ErrorCode:         nil,
					MobileCountryCode: nil,
					MobileNetworkCode: nil,
					Name:              utils.String("Twilio - Toll-Free"),
					Type:              utils.String("voip"),
				}))
				Expect(resp.AddOns).To(BeNil())
				Expect(resp.URL).To(Equal("https://lookups.twilio.com/v1/PhoneNumbers/+18551112222?Type=carrier&Type=caller-name"))
			})
		})

		Describe("When the phone number with addon is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://lookups.twilio.com/v1/PhoneNumbers/+447111111111?AddOns=testing&AddOnsData.testing.country_code=GB",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/phoneNumberResponseWithAddon.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := phoneNumberClient.Fetch(&phone_number.FetchPhoneNumberOptions{
				AddOns:     &[]string{"testing"},
				AddOnsData: &map[string]interface{}{"testing.country_code": "GB"},
			})
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the phone number details response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				Expect(resp.CallerName).To(BeNil())
				Expect(resp.CountryCode).To(Equal("GB"))
				Expect(resp.PhoneNumber).To(Equal("+447111111111"))
				Expect(resp.NationalFormat).To(Equal("07111 111111"))
				Expect(resp.Carrier).To(BeNil())

				addonsFixture, _ := ioutil.ReadFile("testdata/addons.json")
				addons := make(map[string]interface{})
				json.Unmarshal(addonsFixture, &addons)

				Expect(resp.AddOns).To(Equal(&addons))
				Expect(resp.URL).To(Equal("https://lookups.twilio.com/v1/PhoneNumbers/+447111111111?AddOns=testing&AddOns.testing.country_code=GB"))
			})
		})

		Describe("When the get phone number response returns a 500", func() {
			httpmock.RegisterResponder("GET", "https://lookups.twilio.com/v1/PhoneNumbers/+447111111111",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := phoneNumberClient.Fetch(nil)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the phone number details response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})
})

func ExpectInternalServerError(err error) {
	Expect(err).ToNot(BeNil())
	twilioErr, ok := err.(*utils.TwilioError)
	Expect(ok).To(Equal(true))
	Expect(twilioErr.Code).To(BeNil())
	Expect(twilioErr.Message).To(Equal("An error occurred"))
	Expect(twilioErr.MoreInfo).To(BeNil())
	Expect(twilioErr.Status).To(Equal(500))
}
