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

	"github.com/RJPearson94/twilio-sdk-go/service/studio"
	"github.com/RJPearson94/twilio-sdk-go/service/studio/v2/flow"
	"github.com/RJPearson94/twilio-sdk-go/service/studio/v2/flow/execution"
	"github.com/RJPearson94/twilio-sdk-go/service/studio/v2/flow/execution/steps"
	"github.com/RJPearson94/twilio-sdk-go/service/studio/v2/flow/executions"
	"github.com/RJPearson94/twilio-sdk-go/service/studio/v2/flow/revisions"
	"github.com/RJPearson94/twilio-sdk-go/service/studio/v2/flow/test_users"
	"github.com/RJPearson94/twilio-sdk-go/service/studio/v2/flow_validation"
	"github.com/RJPearson94/twilio-sdk-go/service/studio/v2/flows"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

var _ = Describe("Studio V2", func() {
	creds, err := credentials.New(credentials.Account{
		Sid:       "ACxxx",
		AuthToken: "Test",
	})
	if err != nil {
		log.Panicf("%s", err)
	}

	studioSession := studio.NewWithCredentials(creds).V2

	httpmock.ActivateNonDefault(studioSession.GetClient().GetRestyClient().GetClient())
	defer httpmock.DeactivateAndReset()

	Describe("Given the Flows Client", func() {
		flowsClient := studioSession.Flows

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

			resp, err := flowsClient.Create(createInput)
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

			resp, err := flowsClient.Create(createInput)
			It("Then an error flowsClient be returned", func() {
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

			resp, err := flowsClient.Create(createInput)
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

			resp, err := flowsClient.Create(createInput)
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

			resp, err := flowsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create flow response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of flows are successfully retrieved", func() {
			pageOptions := &flows.FlowsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://studio.twilio.com/v2/Flows?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/flowsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := flowsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the flows page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://studio.twilio.com/v2/Flows?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://studio.twilio.com/v2/Flows?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("flows"))

				flows := resp.Flows
				Expect(flows).ToNot(BeNil())
				Expect(len(flows)).To(Equal(1))

				Expect(flows[0].Status).To(Equal("draft"))

				flowDefinition, _ := ioutil.ReadFile("testdata/flowDefinition.json")
				definition := make(map[string]interface{})
				json.Unmarshal(flowDefinition, &definition)
				Expect(flows[0].Definition).To(Equal(definition))

				Expect(flows[0].Errors).To(BeNil())
				Expect(flows[0].Warnings).To(BeNil())
				Expect(flows[0].DateUpdated).To(BeNil())
				Expect(flows[0].FriendlyName).To(Equal("Test 2"))
				Expect(flows[0].AccountSid).To(Equal("ACxxxxxxxxxxxx"))
				Expect(flows[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-05-30T22:28:43Z"))
				Expect(flows[0].URL).To(Equal("https://studio.twilio.com/v2/Flows/FWxxxxxxxxxxxx"))
				Expect(flows[0].Valid).To(Equal(true))
				Expect(flows[0].Sid).To(Equal("FWxxxxxxxxxxxx"))
				Expect(flows[0].CommitMessage).To(BeNil())
				Expect(flows[0].WebhookURL).To(Equal("https://webhooks.twilio.com/v1/Accounts/ACxxxxxxxxxxxx/Flows/FWxxxxxxxxxxxx"))
				Expect(flows[0].Revision).To(Equal(1))

			})
		})

		Describe("When the page of flows api returns a 500 response", func() {
			pageOptions := &flows.FlowsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://studio.twilio.com/v2/Flows?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := flowsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the flows page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated flows are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://studio.twilio.com/v2/Flows",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/flowsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://studio.twilio.com/v2/Flows?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/flowsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := flowsClient.NewFlowsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated flows current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated flows results should be returned", func() {
				Expect(len(paginator.Flows)).To(Equal(3))
			})
		})

		Describe("When the flows api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://studio.twilio.com/v2/Flows",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/flowsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://studio.twilio.com/v2/Flows?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := flowsClient.NewFlowsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated flows current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
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

			resp, err := flowClient.Fetch()
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

			resp, err := studioSession.Flow("FW71").Fetch()
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

			updateInput := &flow.UpdateFlowInput{
				FriendlyName: utils.String("Test 2"),
				Definition:   utils.String(string(flowDefinition)),
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

	Describe("Given the Executions Client", func() {
		executionsClient := studioSession.Flow("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Executions

		Describe("When the execution is successfully created", func() {
			createInput := &executions.CreateExecutionInput{
				To:         "+18001234567",
				From:       "+15017122661",
				Parameters: utils.String("{\"name\": \"RJPearson94\"}"),
			}

			httpmock.RegisterResponder("POST", "https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Executions",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/executionResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := executionsClient.Create(createInput)
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
			createInput := &executions.CreateExecutionInput{
				From:       "+15017122661",
				Parameters: utils.String("{\"name\": \"RJPearson94\"}"),
			}

			resp, err := executionsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create flow response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the execution does not contain a from", func() {
			createInput := &executions.CreateExecutionInput{
				To:         "+15017122661",
				Parameters: utils.String("{\"name\": \"RJPearson94\"}"),
			}

			resp, err := executionsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create flow response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the Create Flow API returns a 500 response", func() {
			createInput := &executions.CreateExecutionInput{
				To:         "+18001234567",
				From:       "+15017122661",
				Parameters: utils.String("{\"name\": \"RJPearson94\"}"),
			}

			httpmock.RegisterResponder("POST", "https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Executions",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := executionsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create flow response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of executions are successfully retrieved", func() {
			pageOptions := &executions.ExecutionsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Executions?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/executionsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := executionsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the executions page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Executions?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Executions?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("executions"))

				executions := resp.Executions
				Expect(executions).ToNot(BeNil())
				Expect(len(executions)).To(Equal(1))

				Expect(executions[0].Sid).To(Equal("FNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(executions[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(executions[0].FlowSid).To(Equal("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(executions[0].Context).To(BeEmpty())
				Expect(executions[0].ContactChannelAddress).To(Equal("+18001234567"))
				Expect(executions[0].Status).To(Equal("active"))
				Expect(executions[0].DateUpdated).To(BeNil())
				Expect(executions[0].DateCreated.Format(time.RFC3339)).To(Equal("2015-07-30T20:00:00Z"))
				Expect(executions[0].URL).To(Equal("https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Executions/FNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of executions api returns a 500 response", func() {
			pageOptions := &executions.ExecutionsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Executions?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := executionsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the executions page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated executions are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Executions",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/executionsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Executions?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/executionsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := executionsClient.NewExecutionsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated executions current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated executions results should be returned", func() {
				Expect(len(paginator.Executions)).To(Equal(3))
			})
		})

		Describe("When the executions api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Executions",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/executionsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Executions?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := executionsClient.NewExecutionsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated executions current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
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

			resp, err := executionClient.Fetch()
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

			resp, err := studioSession.Flow("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Execution("FN71").Fetch()
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

			resp, err := executionContextClient.Fetch()
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

			resp, err := studioSession.Flow("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Execution("FN71").Context().Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get execution context response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a Steps Client", func() {
		stepsClient := studioSession.Flow("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Execution("FNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Steps

		Describe("When the page of steps are successfully retrieved", func() {
			pageOptions := &steps.StepsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Executions/FNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Steps?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/stepsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := stepsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the steps page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Executions/FNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Steps?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Executions/FNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Steps?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("steps"))

				steps := resp.Steps
				Expect(steps).ToNot(BeNil())
				Expect(len(steps)).To(Equal(1))

				Expect(steps[0].Sid).To(Equal("FTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(steps[0].ExecutionSid).To(Equal("FNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(steps[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(steps[0].FlowSid).To(Equal("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(steps[0].Name).To(Equal("incomingRequest"))
				Expect(steps[0].Context).To(BeEmpty())
				Expect(steps[0].TransitionedFrom).To(Equal("Trigger"))
				Expect(steps[0].TransitionedTo).To(Equal("Ended"))
				Expect(steps[0].DateUpdated).To(BeNil())
				Expect(steps[0].DateCreated.Format(time.RFC3339)).To(Equal("2017-11-06T12:00:00Z"))
				Expect(steps[0].URL).To(Equal("https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Executions/FNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Steps/FTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of steps api returns a 500 response", func() {
			pageOptions := &steps.StepsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Executions/FNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Steps?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := stepsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the steps page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated steps are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Executions/FNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Steps",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/stepsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Executions/FNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Steps?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/stepsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := stepsClient.NewStepsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated steps current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated steps results should be returned", func() {
				Expect(len(paginator.Steps)).To(Equal(3))
			})
		})

		Describe("When the steps api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Executions/FNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Steps",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/stepsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Executions/FNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Steps?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := stepsClient.NewStepsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated steps current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
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

			resp, err := stepClient.Fetch()
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

			resp, err := studioSession.Flow("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Execution("FNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Step("FT71").Fetch()
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

			resp, err := stepContextClient.Fetch()
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

			resp, err := studioSession.Flow("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Execution("FNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Step("FT71").Context().Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get step context response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a Revisions Client", func() {
		revisionsClient := studioSession.Flow("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Revisions

		Describe("When the page of revisions are successfully retrieved", func() {
			pageOptions := &revisions.RevisionsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Revisions?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/flowRevisionsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := revisionsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the revisions page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Revisions?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Revisions?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("revisions"))

				revisions := resp.Revisions
				Expect(revisions).ToNot(BeNil())
				Expect(len(revisions)).To(Equal(1))

				Expect(revisions[0].Sid).To(Equal("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(revisions[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))

				definition := make(map[string]interface{})
				definition["initial_state"] = "Trigger"

				Expect(revisions[0].Definition).To(Equal(definition))
				Expect(revisions[0].FriendlyName).To(Equal("Test Flow"))
				Expect(revisions[0].Status).To(Equal("published"))
				Expect(revisions[0].CommitMessage).To(BeNil())
				Expect(revisions[0].Valid).To(Equal(true))

				errors := make([]interface{}, 0)
				Expect(revisions[0].Errors).To(Equal(&errors))
				Expect(revisions[0].DateUpdated).To(BeNil())
				Expect(revisions[0].DateCreated.Format(time.RFC3339)).To(Equal("2017-11-06T12:00:00Z"))
				Expect(revisions[0].URL).To(Equal("https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Revisions/1"))
			})
		})

		Describe("When the page of revisions api returns a 500 response", func() {
			pageOptions := &revisions.RevisionsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Revisions?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := revisionsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the revisions page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated revisions are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Revisions",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/flowRevisionsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Revisions?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/flowRevisionsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := revisionsClient.NewRevisionsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated revisions current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated revisions results should be returned", func() {
				Expect(len(paginator.Revisions)).To(Equal(3))
			})
		})

		Describe("When the revisions api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Revisions",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/flowRevisionsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Revisions?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := revisionsClient.NewRevisionsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated revisions current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
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

			resp, err := flowRevision.Fetch()
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

		Describe("When the get flow revision response returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Revisions/100",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := studioSession.Flow("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Revision(100).Fetch()
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

		Describe("When the test users are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TestUsers",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/flowTestUsersResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := testUsersClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get test users response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.TestUsers).ToNot(BeEmpty())
				Expect(resp.TestUsers[0]).To(Equal("+14155551212"))
				Expect(resp.TestUsers[1]).To(Equal("+14155551213"))
				Expect(resp.URL).To(Equal("https://studio.twilio.com/v2/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TestUsers"))
			})
		})

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

		Describe("When the Update Test Users API returns a 500 response", func() {
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
				ExpectInternalServerError(err)
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
