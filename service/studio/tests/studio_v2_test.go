package tests

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/service/studio/v2/flow/test_users"

	"github.com/jarcoal/httpmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/RJPearson94/twilio-sdk-go/service/studio"
	"github.com/RJPearson94/twilio-sdk-go/service/studio/v2/flow"
	"github.com/RJPearson94/twilio-sdk-go/service/studio/v2/flow/execution"
	"github.com/RJPearson94/twilio-sdk-go/service/studio/v2/flow/executions"
	"github.com/RJPearson94/twilio-sdk-go/service/studio/v2/flow_validation"
	"github.com/RJPearson94/twilio-sdk-go/service/studio/v2/flows"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

var _ = Describe("Studio V2", func() {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	creds, err := credentials.New(credentials.Account{
		Sid:       "ACxxx",
		AuthToken: "Test",
	})
	if err != nil {
		log.Panicf("%s", err)
	}

	studioSession := studio.NewWithCredentials(creds).V2

	Describe("Given the Flows Client", func() {
		flowClient := studioSession.Flows

		Describe("When the Flow is successfully created", func() {
			flowDefinition, _ := ioutil.ReadFile("testdata/flowDefinition.json")

			createInput := &flows.CreateFlowInput{
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

			resp, err := flowClient.Create(createInput)
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

			createInput := &flows.CreateFlowInput{
				Status:     "draft",
				Definition: string(flowDefinition),
			}

			resp, err := flowClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create flow response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the Flow does not contain a status", func() {
			flowDefinition, _ := ioutil.ReadFile("testdata/flowDefinition.json")

			createInput := &flows.CreateFlowInput{
				FriendlyName: "Test 2",
				Definition:   string(flowDefinition),
			}

			resp, err := flowClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create flow response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the Flow does not contain a definition", func() {
			createInput := &flows.CreateFlowInput{
				FriendlyName: "Test 2",
				Status:       "draft",
			}

			resp, err := flowClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create flow response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the Create Flow API returns a 500 response", func() {
			flowDefinition, _ := ioutil.ReadFile("testdata/flowDefinition.json")

			createInput := &flows.CreateFlowInput{
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

			resp, err := flowClient.Create(createInput)
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
		flowClient := studioSession.Flow("FWxxxxxxxxxxxx")

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

			resp, err := flowClient.Get()
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

			It("Then the get flow response should be nil", func() {
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

			updateInput := &flow.UpdateFlowInput{
				Status: "published",
			}

			resp, err := flowClient.Update(updateInput)
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

			updateInput := &flow.UpdateFlowInput{
				Status: "published",
			}

			resp, err := studioSession.Flow("FW71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update flow response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the Update Flow Request does not contain a status", func() {
			flowDefinition, _ := ioutil.ReadFile("testdata/flowDefinition.json")

			friendlyName := "Test 2"
			definition := string(flowDefinition)
			updateInput := &flow.UpdateFlowInput{
				FriendlyName: &friendlyName,
				Definition:   &definition,
			}

			resp, err := flowClient.Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the update flow response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the Flow is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://studio.twilio.com/v2/Flows/FWxxxxxxxxxxxx", httpmock.NewStringResponder(204, ""))

			err := flowClient.Delete()
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

	Describe("Given the Execution Client", func() {
		executionClient := studioSession.Flow("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Executions

		Describe("When the execution is successfully created", func() {
			parameters := "{\"name\": \"RJPearson94\"}"
			createInput := &executions.CreateExecutionInput{
				To:         "+18001234567",
				From:       "+15017122661",
				Parameters: &parameters,
			}

			httpmock.RegisterResponder("POST", "https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Executions",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/executionResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := executionClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create execution response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("FNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FlowSid).To(Equal("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Context).To(BeEmpty())
				Expect(resp.ContactChannelAddress).To(Equal("+18001234567"))
				Expect(resp.Status).To(Equal("active"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2015-07-30T20:00:00Z"))
				Expect(resp.URL).To(Equal("https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Executions/FNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the execution does not contain a to", func() {
			parameters := "{\"name\": \"RJPearson94\"}"
			createInput := &executions.CreateExecutionInput{
				From:       "+15017122661",
				Parameters: &parameters,
			}

			resp, err := executionClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create flow response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the execution does not contain a from", func() {
			parameters := "{\"name\": \"RJPearson94\"}"
			createInput := &executions.CreateExecutionInput{
				To:         "+15017122661",
				Parameters: &parameters,
			}

			resp, err := executionClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create flow response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the Create Flow API returns a 500 response", func() {
			parameters := "{\"name\": \"RJPearson94\"}"
			createInput := &executions.CreateExecutionInput{
				To:         "+18001234567",
				From:       "+15017122661",
				Parameters: &parameters,
			}

			httpmock.RegisterResponder("POST", "https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Executions",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := executionClient.Create(createInput)
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

	Describe("Given I have a Execution SID", func() {
		executionClient := studioSession.Flow("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Execution("FNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the Execution is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Executions/FNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/executionResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := executionClient.Get()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get execution response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("FNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FlowSid).To(Equal("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Context).To(BeEmpty())
				Expect(resp.ContactChannelAddress).To(Equal("+18001234567"))
				Expect(resp.Status).To(Equal("active"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2015-07-30T20:00:00Z"))
				Expect(resp.URL).To(Equal("https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Executions/FNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the get execution response returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Executions/FN71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := studioSession.Flow("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Execution("FN71").Get()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get flow response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the Flow is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Executions/FNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateExecutionResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &execution.UpdateExecutionInput{
				Status: "ended",
			}

			resp, err := executionClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update execution response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("FNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FlowSid).To(Equal("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Context).To(BeEmpty())
				Expect(resp.ContactChannelAddress).To(Equal("+18001234567"))
				Expect(resp.Status).To(Equal("Ended"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2015-07-30T20:00:00Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2015-07-30T21:00:00Z"))
				Expect(resp.URL).To(Equal("https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Executions/FNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update execution response returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Executions/FN71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &execution.UpdateExecutionInput{
				Status: "Ended",
			}

			resp, err := studioSession.Flow("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Execution("FN71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update execution response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the Update Execution Request does not contain a status", func() {
			updateInput := &execution.UpdateExecutionInput{}

			resp, err := executionClient.Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the update execution response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the execution is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Executions/FNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := executionClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the delete execution response returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Executions/FN71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := studioSession.Flow("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Execution("FN71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given I have a Execution Context Client", func() {
		executionContextClient := studioSession.Flow("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Execution("FNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Context()

		Describe("When the Execution Context is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Executions/FNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Context",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/executionContextResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := executionContextClient.Get()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get execution context response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FlowSid).To(Equal("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ExecutionSid).To(Equal("FNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Context).ToNot(BeEmpty())
				Expect(resp.URL).To(Equal("https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Executions/FNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Context"))
			})
		})

		Describe("When the get execution context response returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Executions/FN71/Context",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := studioSession.Flow("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Execution("FN71").Context().Get()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get execution context response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a Step SID", func() {
		stepClient := studioSession.Flow("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Execution("FNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Step("FTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the Step is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Executions/FNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Steps/FTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/stepResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := stepClient.Get()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get step response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("FTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ExecutionSid).To(Equal("FNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FlowSid).To(Equal("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Name).To(Equal("incomingRequest"))
				Expect(resp.Context).To(BeEmpty())
				Expect(resp.TransitionedFrom).To(Equal("Trigger"))
				Expect(resp.TransitionedTo).To(Equal("Ended"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2017-11-06T12:00:00Z"))
				Expect(resp.URL).To(Equal("https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Executions/FNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Steps/FTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the get execution response returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Executions/FNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Steps/FT71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := studioSession.Flow("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Execution("FNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Step("FT71").Get()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get step response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a Step Context Client", func() {
		stepContextClient := studioSession.Flow("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Execution("FNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Step("FTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Context()

		Describe("When the Step Context is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Executions/FNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Steps/FTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Context",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/stepContextResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := stepContextClient.Get()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get step context response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FlowSid).To(Equal("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ExecutionSid).To(Equal("FNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.StepSid).To(Equal("FTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Context).ToNot(BeEmpty())
				Expect(resp.URL).To(Equal("https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Executions/FNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Steps/FTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Context"))
			})
		})

		Describe("When the get execution response returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Executions/FNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Steps/FT71/Context",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := studioSession.Flow("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Execution("FNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Step("FT71").Context().Get()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get step context response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a Revision Number", func() {
		flowRevision := studioSession.Flow("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Revision(1)

		Describe("When the revision is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Revisions/1",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/flowRevisionResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := flowRevision.Get()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get flow revision response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))

				definition := make(map[string]interface{})
				definition["initial_state"] = "Trigger"

				Expect(resp.Definition).To(Equal(definition))
				Expect(resp.FriendlyName).To(Equal("Test Flow"))
				Expect(resp.Status).To(Equal("published"))
				Expect(resp.CommitMessage).To(BeNil())
				Expect(resp.Valid).To(Equal(true))

				errors := make([]interface{}, 0)
				Expect(resp.Errors).To(Equal(&errors))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2017-11-06T12:00:00Z"))
				Expect(resp.URL).To(Equal("https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Revisions/1"))
			})
		})

		Describe("When the get execution response returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Revisions/100",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := studioSession.Flow("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Revision(100).Get()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get flow revision response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given the Test User Client", func() {
		testUsersClient := studioSession.Flow("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").TestUsers()

		Describe("When the test users are successfully updated", func() {
			testUsers := make([]string, 2)
			testUsers[0] = "+14155551212"
			testUsers[1] = "*14155551213"
			updateInput := &test_users.UpdateTestUsersInput{
				TestUsers: testUsers,
			}

			httpmock.RegisterResponder("POST", "https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TestUsers",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/flowTestUsersResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := testUsersClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update test users response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.TestUsers).ToNot(BeEmpty())
				Expect(resp.TestUsers[0]).To(Equal("+14155551212"))
				Expect(resp.TestUsers[1]).To(Equal("+14155551213"))
				Expect(resp.URL).To(Equal("https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TestUsers"))
			})
		})

		Describe("When the test users does not contain test users", func() {
			updateInput := &test_users.UpdateTestUsersInput{}

			resp, err := testUsersClient.Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create flow response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the Create Flow API returns a 500 response", func() {
			testUsers := make([]string, 2)
			testUsers[0] = "+14155551212"
			testUsers[1] = "*14155551213"
			updateInput := &test_users.UpdateTestUsersInput{
				TestUsers: testUsers,
			}

			httpmock.RegisterResponder("POST", "https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TestUsers",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := testUsersClient.Update(updateInput)
			It("Then an error should be returned", func() {
				Expect(err).ToNot(BeNil())
				twilioErr, ok := err.(*utils.TwilioError)
				Expect(ok).To(Equal(true))
				Expect(twilioErr.Code).To(BeNil())
				Expect(twilioErr.Message).To(Equal("An error occurred"))
				Expect(twilioErr.MoreInfo).To(BeNil())
				Expect(twilioErr.Status).To(Equal(500))
			})

			It("Then the update test users response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given the Flow Validation Client", func() {
		flowValidationClient := studioSession.FlowValidation

		Describe("When the Flow is successfully validated", func() {
			flowDefinition, _ := ioutil.ReadFile("testdata/flowDefinition.json")

			validateInput := &flow_validation.ValidateFlowInput{
				FriendlyName: "Test 2",
				Status:       "draft",
				Definition:   string(flowDefinition),
			}

			httpmock.RegisterResponder("POST", "https://studio.twilio.com/v2/Flows/Validate",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/validateFlowResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := flowValidationClient.Validate(validateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the validate flow response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Valid).To(Equal(true))
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
