package tests

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/jarcoal/httpmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/RJPearson94/twilio-sdk-go/service/autopilot"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistants"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

var _ = Describe("Autopilot V1", func() {
	creds, err := credentials.New(credentials.Account{
		Sid:       "ACxxx",
		AuthToken: "Test",
	})
	if err != nil {
		log.Panicf("%s", err)
	}

	autopilotSession := autopilot.NewWithCredentials(creds).V1

	httpmock.ActivateNonDefault(autopilotSession.GetClient().GetRestyClient().GetClient())
	defer httpmock.DeactivateAndReset()

	Describe("Given the assistants client", func() {
		assistantsClient := autopilotSession.Assistants

		Describe("When the assistant is successfully created", func() {
			createInput := &assistants.CreateAssistantInput{}

			httpmock.RegisterResponder("POST", "https://autopilot.twilio.com/v1/Assistants",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/assistantResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := assistantsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create assistant response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.UniqueName).To(Equal("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX-1234"))
				Expect(resp.LatestModelBuildSid).To(BeNil())
				Expect(resp.CallbackEvents).To(BeNil())
				Expect(resp.FriendlyName).To(BeNil())
				Expect(resp.LogQueries).To(Equal(true))
				Expect(resp.CallbackUrl).To(BeNil())
				Expect(resp.DevelopmentStage).To(Equal("in-development"))
				Expect(resp.NeedsModelBuild).To(Equal(utils.Bool(false)))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the create assistant api returns a 500 response", func() {
			createInput := &assistants.CreateAssistantInput{}

			httpmock.RegisterResponder("POST", "https://autopilot.twilio.com/v1/Assistants",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := assistantsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create assistant response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a assistant sid", func() {
		assistantClient := autopilotSession.Assistant("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the assistant is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/assistantResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := assistantClient.Get()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get assistant response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.UniqueName).To(Equal("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX-1234"))
				Expect(resp.LatestModelBuildSid).To(BeNil())
				Expect(resp.CallbackEvents).To(BeNil())
				Expect(resp.FriendlyName).To(BeNil())
				Expect(resp.LogQueries).To(Equal(true))
				Expect(resp.CallbackUrl).To(BeNil())
				Expect(resp.DevelopmentStage).To(Equal("in-development"))
				Expect(resp.NeedsModelBuild).To(Equal(utils.Bool(false)))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the assistant api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://autopilot.twilio.com/v1/Assistants/UA71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := autopilotSession.Assistant("UA71").Get()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get assistant response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the assistant is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateAssistantResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &assistant.UpdateAssistantInput{
				FriendlyName: utils.String("Test"),
			}

			resp, err := assistantClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update assistant response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.UniqueName).To(Equal("Test"))
				Expect(resp.LatestModelBuildSid).To(BeNil())
				Expect(resp.CallbackEvents).To(BeNil())
				Expect(resp.FriendlyName).To(BeNil())
				Expect(resp.LogQueries).To(Equal(true))
				Expect(resp.CallbackUrl).To(BeNil())
				Expect(resp.DevelopmentStage).To(Equal("in-development"))
				Expect(resp.NeedsModelBuild).To(Equal(utils.Bool(false)))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-27T23:10:00Z"))
				Expect(resp.URL).To(Equal("https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the assistant api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://autopilot.twilio.com/v1/Assistants/UA71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &assistant.UpdateAssistantInput{
				FriendlyName: utils.String("Test"),
			}

			resp, err := autopilotSession.Assistant("UA71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update assistant response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the assistant is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := assistantClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the assistant api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://autopilot.twilio.com/v1/Assistants/UA71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := autopilotSession.Assistant("UA71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
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

func ExpectInvalidInputError(err error) {
	ExpectErrorToNotBeATwilioError(err)
	Expect(err.Error()).To(Equal("Invalid input supplied"))
}

func ExpectNotFoundError(err error) {
	Expect(err).ToNot(BeNil())
	twilioErr, ok := err.(*utils.TwilioError)
	Expect(ok).To(Equal(true))

	code := 20404
	Expect(twilioErr.Code).To(Equal(&code))
	Expect(twilioErr.Message).To(Equal("The requested resource /Assistants/UA71 was not found"))

	moreInfo := "https://www.twilio.com/docs/errors/20404"
	Expect(twilioErr.MoreInfo).To(Equal(&moreInfo))
	Expect(twilioErr.Status).To(Equal(404))
}

func ExpectErrorToNotBeATwilioError(err error) {
	Expect(err).ToNot(BeNil())
	_, ok := err.(*utils.TwilioError)
	Expect(ok).To(Equal(false))
}
