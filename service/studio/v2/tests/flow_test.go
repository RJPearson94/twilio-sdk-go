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

	v2 "github.com/RJPearson94/twilio-sdk-go/service/studio/v2"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

var _ = Describe("Flow", func() {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	creds, err := credentials.New(credentials.Account{
		Sid:       "ACxxx",
		AuthToken: "Test",
	})
	if err != nil {
		log.Panicf("%s", err)
	}

	studioSession := v2.NewWithCredentials(creds)

	Describe("Given the Flow Service", func() {
		flowService := studioSession.Flows

		Describe("When page of Flows is returned with default settings", func() {
			getFlowPageRequest := &v2.GetFlowPageRequest{}

			httpmock.RegisterResponder("GET", "https://studio.twilio.com/v2/Flows",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/defaultFlowPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := flowService.GetPage(getFlowPageRequest)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get flow page response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Meta).ToNot(BeNil())
				Expect(resp.Flows).ToNot(BeNil())
				Expect(len(resp.Flows)).To(Equal(3))
			})
		})

		Describe("When page of Flows is returned with PageSize and Page Specified", func() {
			getFlowPageRequest := &v2.GetFlowPageRequest{
				Page:     "1",
				PageSize: "20",
			}

			httpmock.RegisterResponder("GET", "https://studio.twilio.com/v2/Flows?Page=1&PageSize=20",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/pageSizeAndPageFlowPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := flowService.GetPage(getFlowPageRequest)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get flow page response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Meta).ToNot(BeNil())
				Expect(resp.Flows).ToNot(BeNil())
				Expect(len(resp.Flows)).To(Equal(2))
			})
		})

		Describe("When the Flow is successfully created", func() {
			flowDefinition, _ := ioutil.ReadFile("testdata/flowDefinition.json")

			createInput := &v2.CreateFlowInput{
				FriendlyName: "Test 2",
				Status:       "draft",
				Definition:   string(flowDefinition),
			}

			httpmock.RegisterResponder("POST", "https://studio.twilio.com/v2/Flows",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/flowResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := flowService.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create flow response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Status).To(Equal("draft"))

				definition := make(map[string]interface{})
				json.Unmarshal(flowDefinition, &definition)
				Expect(resp.Definition).To(Equal(definition))
				Expect(resp.Errors).To(BeNil())
				Expect(resp.Warnings).To(BeNil())
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.FriendlyName).To(Equal("Test 2"))
				Expect(resp.AccountSid).To(Equal("ACxxxxxxxxxxxx"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-05-30T22:28:43Z"))
				Expect(resp.URL).To(Equal("https://studio.twilio.com/v2/Flows/FWxxxxxxxxxxxx"))
				Expect(resp.Valid).To(Equal(true))
				Expect(resp.Sid).To(Equal("FWxxxxxxxxxxxx"))
				Expect(resp.CommitMessage).To(BeNil())
				Expect(resp.WebhookURL).To(Equal("https://webhooks.twilio.com/v1/Accounts/ACxxxxxxxxxxxx/Flows/FWxxxxxxxxxxxx"))
				Expect(resp.Revision).To(Equal(1))
			})
		})

		Describe("When the Flow does not contain a friendly name", func() {
			flowDefinition, _ := ioutil.ReadFile("testdata/flowDefinition.json")

			createInput := &v2.CreateFlowInput{
				Status:     "draft",
				Definition: string(flowDefinition),
			}

			resp, err := flowService.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create flow response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the Flow does not contain a status", func() {
			flowDefinition, _ := ioutil.ReadFile("testdata/flowDefinition.json")

			createInput := &v2.CreateFlowInput{
				FriendlyName: "Test 2",
				Definition:   string(flowDefinition),
			}

			resp, err := flowService.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create flow response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the Flow does not contain a definition", func() {
			createInput := &v2.CreateFlowInput{
				FriendlyName: "Test 2",
				Status:       "draft",
			}

			resp, err := flowService.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create flow response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the Create Flow API returns a 500 response", func() {
			flowDefinition, _ := ioutil.ReadFile("testdata/flowDefinition.json")

			createInput := &v2.CreateFlowInput{
				FriendlyName: "Test 2",
				Status:       "draft",
				Definition:   string(flowDefinition),
			}

			httpmock.RegisterResponder("POST", "https://studio.twilio.com/v2/Flows",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := flowService.Create(createInput)
			It("Then an error should be returned", func() {
				Expect(err).ToNot(BeNil())
				twilioErr, ok := err.(*utils.TwilioError)
				Expect(ok).To(Equal(true))
				Expect(twilioErr.Code).To(BeNil())
				Expect(twilioErr.Message).To(Equal("An error occurred"))
				Expect(twilioErr.MoreInfo).To(BeNil())
				Expect(twilioErr.Status).To(Equal(500))
			})

			It("Then the create flow response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a Flow SID", func() {
		flow := studioSession.Flow("FWxxxxxxxxxxxx")

		Describe("When the Flow is successfully retrieved", func() {
			flowDefinition, _ := ioutil.ReadFile("testdata/flowDefinition.json")

			httpmock.RegisterResponder("GET", "https://studio.twilio.com/v2/Flows/FWxxxxxxxxxxxx",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/flowResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := flow.Get()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get flow response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Status).To(Equal("draft"))

				definition := make(map[string]interface{})
				json.Unmarshal(flowDefinition, &definition)
				Expect(resp.Definition).To(Equal(definition))
				Expect(resp.Errors).To(BeNil())
				Expect(resp.Warnings).To(BeNil())
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.FriendlyName).To(Equal("Test 2"))
				Expect(resp.AccountSid).To(Equal("ACxxxxxxxxxxxx"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-05-30T22:28:43Z"))
				Expect(resp.URL).To(Equal("https://studio.twilio.com/v2/Flows/FWxxxxxxxxxxxx"))
				Expect(resp.Valid).To(Equal(true))
				Expect(resp.Sid).To(Equal("FWxxxxxxxxxxxx"))
				Expect(resp.CommitMessage).To(BeNil())
				Expect(resp.WebhookURL).To(Equal("https://webhooks.twilio.com/v1/Accounts/ACxxxxxxxxxxxx/Flows/FWxxxxxxxxxxxx"))
				Expect(resp.Revision).To(Equal(1))
			})
		})

		Describe("When the get flow response returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://studio.twilio.com/v2/Flows/FW71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := studioSession.Flow("FW71").Get()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the create flow response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the Flow is successfully updated", func() {
			flowDefinition, _ := ioutil.ReadFile("testdata/flowDefinition.json")

			httpmock.RegisterResponder("POST", "https://studio.twilio.com/v2/Flows/FWxxxxxxxxxxxx",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateFlowResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &v2.UpdateFlowInput{
				Status: "published",
			}

			resp, err := flow.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update flow response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Status).To(Equal("published"))

				definition := make(map[string]interface{})
				json.Unmarshal(flowDefinition, &definition)
				Expect(resp.Definition).To(Equal(definition))
				Expect(resp.Errors).To(BeNil())
				Expect(resp.Warnings).To(BeNil())
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-01T10:00:00Z"))
				Expect(resp.FriendlyName).To(Equal("Test 2"))
				Expect(resp.AccountSid).To(Equal("ACxxxxxxxxxxxx"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-05-30T22:28:43Z"))
				Expect(resp.URL).To(Equal("https://studio.twilio.com/v2/Flows/FWxxxxxxxxxxxx"))
				Expect(resp.Valid).To(Equal(true))
				Expect(resp.Sid).To(Equal("FWxxxxxxxxxxxx"))
				Expect(resp.CommitMessage).To(BeNil())
				Expect(resp.WebhookURL).To(Equal("https://webhooks.twilio.com/v1/Accounts/ACxxxxxxxxxxxx/Flows/FWxxxxxxxxxxxx"))
				Expect(resp.Revision).To(Equal(2))
			})
		})

		Describe("When the update flow response returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://studio.twilio.com/v2/Flows/FW71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &v2.UpdateFlowInput{
				Status: "published",
			}

			resp, err := studioSession.Flow("FW71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the create flow response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the Update Flow Request does not contain a status", func() {
			flowDefinition, _ := ioutil.ReadFile("testdata/flowDefinition.json")

			updateInput := &v2.UpdateFlowInput{
				FriendlyName: "Test 2",
				Definition:   string(flowDefinition),
			}

			resp, err := flow.Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create flow response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the Flow is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://studio.twilio.com/v2/Flows/FWxxxxxxxxxxxx", httpmock.NewStringResponder(204, ""))

			err := flow.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the delete flow response returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://studio.twilio.com/v2/Flows/FW71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := studioSession.Flow("FW71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})
})

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
	Expect(twilioErr.Message).To(Equal("The requested resource /Flows/FW71 was not found"))

	moreInfo := "https://www.twilio.com/docs/errors/20404"
	Expect(twilioErr.MoreInfo).To(Equal(&moreInfo))
	Expect(twilioErr.Status).To(Equal(404))
}

func ExpectErrorToNotBeATwilioError(err error) {
	Expect(err).ToNot(BeNil())
	_, ok := err.(*utils.TwilioError)
	Expect(ok).To(Equal(false))
}
