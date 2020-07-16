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
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/defaults"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/field_type"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/field_type/field_values"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/field_types"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/model_build"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/model_builds"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/style_sheet"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/task"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/task/actions"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/task/fields"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/task/sample"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/task/samples"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/tasks"
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
				Expect(resp.CallbackURL).To(BeNil())
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
				Expect(resp.CallbackURL).To(BeNil())
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
				Expect(resp.CallbackURL).To(BeNil())
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

	Describe("Given I have a defaults client", func() {
		defaultsClient := autopilotSession.Assistant("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Defaults()

		Describe("When the default is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Defaults",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/defaultsResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := defaultsClient.Get()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get defaults response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AssistantSid).To(Equal("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))

				data := make(map[string]interface{})
				Expect(resp.Data).To(Equal(data))
				Expect(resp.URL).To(Equal("https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Defaults"))
			})
		})

		Describe("When the defaults api returns a 500", func() {
			httpmock.RegisterResponder("GET", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Defaults",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := defaultsClient.Get()
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the get defaults response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the defaults are successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Defaults",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateDefaultsResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &defaults.UpdateDefaultInput{
				Defaults: utils.String(`{ "defaults": { "fallback": "http://localhost/fallback" } }`),
			}

			resp, err := defaultsClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update defaults response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AssistantSid).To(Equal("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))

				defaultsData, _ := ioutil.ReadFile("testdata/defaultsData.json")
				data := make(map[string]interface{})
				json.Unmarshal(defaultsData, &data)
				Expect(resp.Data).To(Equal(data))

				Expect(resp.URL).To(Equal("https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Defaults"))
			})
		})

		Describe("When the defaults api returns a 500", func() {
			httpmock.RegisterResponder("POST", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Defaults",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			updateInput := &defaults.UpdateDefaultInput{
				Defaults: utils.String(`{ "defaults": { "fallback": "http://localhost/fallback" } }`),
			}

			resp, err := defaultsClient.Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the get defaults response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a stylesheet client", func() {
		stylesheetClient := autopilotSession.Assistant("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").StyleSheet()

		Describe("When the stylesheet is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/StyleSheet",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/stylesheetResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := stylesheetClient.Get()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get stylesheet response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AssistantSid).To(Equal("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))

				data := make(map[string]interface{})
				Expect(resp.Data).To(Equal(data))
				Expect(resp.URL).To(Equal("https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/StyleSheet"))
			})
		})

		Describe("When the stylesheet api returns a 500", func() {
			httpmock.RegisterResponder("GET", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/StyleSheet",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := stylesheetClient.Get()
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the get stylesheet response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the stylesheet are successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/StyleSheet",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateStylesheetResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &style_sheet.UpdateStyleSheetInput{
				StyleSheet: utils.String(`{ "voice": { "say_voice": "Polly.Matthew" } }`),
			}

			resp, err := stylesheetClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update stylesheet response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AssistantSid).To(Equal("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))

				stylesheetData, _ := ioutil.ReadFile("testdata/stylesheetData.json")
				data := make(map[string]interface{})
				json.Unmarshal(stylesheetData, &data)
				Expect(resp.Data).To(Equal(data))

				Expect(resp.URL).To(Equal("https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/StyleSheet"))
			})
		})

		Describe("When the stylesheet api returns a 500", func() {
			httpmock.RegisterResponder("POST", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/StyleSheet",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			updateInput := &style_sheet.UpdateStyleSheetInput{
				StyleSheet: utils.String(`{ "voice": { "say_voice": "Polly.Matthew" } }`),
			}

			resp, err := stylesheetClient.Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the get stylesheet response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given the tasks client", func() {
		tasksClient := autopilotSession.Assistant("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Tasks

		Describe("When the task is successfully created", func() {
			createInput := &tasks.CreateTaskInput{
				UniqueName: "test",
			}

			httpmock.RegisterResponder("POST", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/taskResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := tasksClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create task response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AssistantSid).To(Equal("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.UniqueName).To(Equal("test"))
				Expect(resp.FriendlyName).To(BeNil())
				Expect(resp.ActionsURL).To(Equal("https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Actions"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the Task does not contain a unique name", func() {
			createInput := &tasks.CreateTaskInput{}

			resp, err := tasksClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create task response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create task api returns a 500 response", func() {
			createInput := &tasks.CreateTaskInput{
				UniqueName: "test",
			}

			httpmock.RegisterResponder("POST", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := tasksClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create task response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a task sid", func() {
		taskClient := autopilotSession.Assistant("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Task("UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the task is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/taskResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := taskClient.Get()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get task response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AssistantSid).To(Equal("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.UniqueName).To(Equal("test"))
				Expect(resp.FriendlyName).To(BeNil())
				Expect(resp.ActionsURL).To(Equal("https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Actions"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the task api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/UD71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := autopilotSession.Assistant("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Task("UD71").Get()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get task response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the task is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateTaskResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &task.UpdateTaskInput{
				FriendlyName: utils.String("test name"),
			}

			resp, err := taskClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update task response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AssistantSid).To(Equal("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.UniqueName).To(Equal("test"))
				Expect(resp.FriendlyName).To(Equal(utils.String("test name")))
				Expect(resp.ActionsURL).To(Equal("https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Actions"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-27T23:10:00Z"))
				Expect(resp.URL).To(Equal("https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the task api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/UD71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &task.UpdateTaskInput{
				FriendlyName: utils.String("test name"),
			}

			resp, err := autopilotSession.Assistant("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Task("UD71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update task response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the task is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := taskClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the task api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/UD71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := autopilotSession.Assistant("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Task("UD71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given I have a actions client", func() {
		actionsClient := autopilotSession.Assistant("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Task("UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Actions()

		Describe("When the actions is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Actions",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/actionsResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := actionsClient.Get()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get actions response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AssistantSid).To(Equal("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.TaskSid).To(Equal("UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))

				data := make(map[string]interface{})
				Expect(resp.Data).To(Equal(data))
				Expect(resp.URL).To(Equal("https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Actions"))
			})
		})

		Describe("When the actions api returns a 500", func() {
			httpmock.RegisterResponder("GET", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Actions",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := actionsClient.Get()
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the get actions response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the actions are successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Actions",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateActionResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &actions.UpdateActionInput{
				Actions: utils.String(`{"actions": [{ "say": { "speech": "Test Speech" } }] }`),
			}

			resp, err := actionsClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update actions response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AssistantSid).To(Equal("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.TaskSid).To(Equal("UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))

				actionsData, _ := ioutil.ReadFile("testdata/actionsData.json")
				data := make(map[string]interface{})
				json.Unmarshal(actionsData, &data)
				Expect(resp.Data).To(Equal(data))

				Expect(resp.URL).To(Equal("https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Actions"))
			})
		})

		Describe("When the actions api returns a 500", func() {
			httpmock.RegisterResponder("POST", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Actions",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			updateInput := &actions.UpdateActionInput{
				Actions: utils.String(`{"actions": [{ "say": { "speech": "Test Speech" } }] }`),
			}

			resp, err := actionsClient.Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the get actions response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a actions client", func() {
		statisticsClient := autopilotSession.Assistant("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Task("UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Statistics()

		Describe("When the actions is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Statistics",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/statisticsResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := statisticsClient.Get()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get statistics response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AssistantSid).To(Equal("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.TaskSid).To(Equal("UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.SamplesCount).To(Equal(1))
				Expect(resp.FieldsCount).To(Equal(2))
				Expect(resp.URL).To(Equal("https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Statistics"))
			})
		})

		Describe("When the statistics api returns a 500", func() {
			httpmock.RegisterResponder("GET", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Statistics",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := statisticsClient.Get()
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the get statistics response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given the samples client", func() {
		samplesClient := autopilotSession.Assistant("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Task("UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Samples

		Describe("When the sample is successfully created", func() {
			createInput := &samples.CreateSampleInput{
				Language:   "en-US",
				TaggedText: "hello world",
			}

			httpmock.RegisterResponder("POST", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Samples",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/sampleResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := samplesClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create sample response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("UFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AssistantSid).To(Equal("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.TaskSid).To(Equal("UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Language).To(Equal("en-US"))
				Expect(resp.TaggedText).To(Equal("hello world"))
				Expect(resp.SourceChannel).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Samples/UFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the sample does not contain a language", func() {
			createInput := &samples.CreateSampleInput{
				TaggedText: "hello world",
			}

			resp, err := samplesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create sample response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the sample does not contain a tagged text", func() {
			createInput := &samples.CreateSampleInput{
				Language: "en-US",
			}

			resp, err := samplesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create sample response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create sample api returns a 500 response", func() {
			createInput := &samples.CreateSampleInput{
				Language:   "en-US",
				TaggedText: "hello world",
			}

			httpmock.RegisterResponder("POST", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Samples",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := samplesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create sample response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a sample sid", func() {
		sampleClient := autopilotSession.Assistant("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Task("UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Sample("UFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the sample is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Samples/UFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/sampleResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := sampleClient.Get()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get sample response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("UFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AssistantSid).To(Equal("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.TaskSid).To(Equal("UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Language).To(Equal("en-US"))
				Expect(resp.TaggedText).To(Equal("hello world"))
				Expect(resp.SourceChannel).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Samples/UFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the sample api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Samples/UF71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := autopilotSession.Assistant("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Task("UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Sample("UF71").Get()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get sample response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the sample is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Samples/UFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateSampleResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &sample.UpdateSampleInput{
				SourceChannel: utils.String("alexa"),
			}

			resp, err := sampleClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update sample response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("UFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AssistantSid).To(Equal("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.TaskSid).To(Equal("UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Language).To(Equal("en-US"))
				Expect(resp.TaggedText).To(Equal("hello world"))
				Expect(resp.SourceChannel).To(Equal(utils.String("alexa")))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-27T23:10:00Z"))
				Expect(resp.URL).To(Equal("https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Samples/UFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the sample api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Samples/UF71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &sample.UpdateSampleInput{
				SourceChannel: utils.String("alexa"),
			}

			resp, err := autopilotSession.Assistant("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Task("UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Sample("UF71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update sample response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the sample is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Samples/UFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := sampleClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the sample api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Samples/UF71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := autopilotSession.Assistant("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Task("UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Sample("UF71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the fields client", func() {
		fieldsClient := autopilotSession.Assistant("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Task("UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Fields

		Describe("When the field is successfully created", func() {
			createInput := &fields.CreateFieldInput{
				UniqueName: "test",
				FieldType:  "test field type",
			}

			httpmock.RegisterResponder("POST", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Fields",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/fieldResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := fieldsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create field response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("UEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AssistantSid).To(Equal("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.TaskSid).To(Equal("UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.UniqueName).To(Equal("test"))
				Expect(resp.FieldType).To(Equal("test field type"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Fields/UEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the field does not contain a unique name", func() {
			createInput := &fields.CreateFieldInput{
				FieldType: "test field type",
			}

			resp, err := fieldsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create field response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the field does not contain a field name", func() {
			createInput := &fields.CreateFieldInput{
				UniqueName: "test",
			}

			resp, err := fieldsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create field response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create field api returns a 500 response", func() {
			createInput := &fields.CreateFieldInput{
				UniqueName: "test",
				FieldType:  "test field type",
			}

			httpmock.RegisterResponder("POST", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Fields",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := fieldsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create field response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a field sid", func() {
		fieldClient := autopilotSession.Assistant("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Task("UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Field("UEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the field is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Fields/UEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/fieldResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := fieldClient.Get()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get field response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("UEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AssistantSid).To(Equal("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.TaskSid).To(Equal("UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.UniqueName).To(Equal("test"))
				Expect(resp.FieldType).To(Equal("test field type"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Fields/UEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the field api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Fields/UE71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := autopilotSession.Assistant("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Task("UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Field("UE71").Get()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get field response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the field is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Fields/UEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := fieldClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the field api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Fields/UE71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := autopilotSession.Assistant("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Task("UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Field("UE71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the field types client", func() {
		fieldTypesClient := autopilotSession.Assistant("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").FieldTypes

		Describe("When the field types is successfully created", func() {
			createInput := &field_types.CreateFieldTypeInput{
				UniqueName: "test",
			}

			httpmock.RegisterResponder("POST", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/FieldTypes",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/fieldTypeResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := fieldTypesClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create field type response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("UBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AssistantSid).To(Equal("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.UniqueName).To(Equal("test"))
				Expect(resp.FriendlyName).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/FieldTypes/UBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the field type does not contain a unique name", func() {
			createInput := &field_types.CreateFieldTypeInput{}

			resp, err := fieldTypesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create field type response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create field types api returns a 500 response", func() {
			createInput := &field_types.CreateFieldTypeInput{
				UniqueName: "test",
			}

			httpmock.RegisterResponder("POST", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/FieldTypes",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := fieldTypesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create field type response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a field type sid", func() {
		fieldTypeClient := autopilotSession.Assistant("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").FieldType("UBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the field type is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/FieldTypes/UBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/fieldTypeResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := fieldTypeClient.Get()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get field type response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("UBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AssistantSid).To(Equal("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.UniqueName).To(Equal("test"))
				Expect(resp.FriendlyName).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/FieldTypes/UBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the field type api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/FieldTypes/UB71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := autopilotSession.Assistant("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").FieldType("UB71").Get()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get field type response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the field type is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/FieldTypes/UBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateFieldTypeResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &field_type.UpdateFieldTypeInput{
				FriendlyName: utils.String("test name"),
			}

			resp, err := fieldTypeClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update field type response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("UBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AssistantSid).To(Equal("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.UniqueName).To(Equal("test"))
				Expect(resp.FriendlyName).To(Equal(utils.String("test name")))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-27T23:10:00Z"))
				Expect(resp.URL).To(Equal("https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/FieldTypes/UBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the field types api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/FieldTypes/UB71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &field_type.UpdateFieldTypeInput{
				FriendlyName: utils.String("test name"),
			}

			resp, err := autopilotSession.Assistant("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").FieldType("UB71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update field type response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the field type is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/FieldTypes/UBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := fieldTypeClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the field types api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/FieldTypes/UB71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := autopilotSession.Assistant("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").FieldType("UB71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the field values client", func() {
		fieldValuesClient := autopilotSession.Assistant("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").FieldType("UBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").FieldValues

		Describe("When the field values is successfully created", func() {
			createInput := &field_values.CreateFieldValueInput{
				Language: "en-US",
				Value:    "test",
			}

			httpmock.RegisterResponder("POST", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/FieldTypes/UBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/FieldValues",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/fieldValueResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := fieldValuesClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create field value response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("UCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AssistantSid).To(Equal("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FieldTypeSid).To(Equal("UBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Language).To(Equal("en-US"))
				Expect(resp.Value).To(Equal("test"))
				Expect(resp.SynonymOf).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/FieldTypes/UBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/FieldValues/UCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the field value does not contain a language", func() {
			createInput := &field_values.CreateFieldValueInput{
				Value: "test",
			}

			resp, err := fieldValuesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create field value response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the field value does not contain a value", func() {
			createInput := &field_values.CreateFieldValueInput{
				Language: "en-US",
			}

			resp, err := fieldValuesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create field value response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create field values api returns a 500 response", func() {
			createInput := &field_values.CreateFieldValueInput{
				Language: "en-US",
				Value:    "test",
			}

			httpmock.RegisterResponder("POST", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/FieldTypes/UBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/FieldValues",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := fieldValuesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create field value response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a field value sid", func() {
		fieldValueClient := autopilotSession.Assistant("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").FieldType("UBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").FieldValue("UCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the field value is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/FieldTypes/UBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/FieldValues/UCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/fieldValueResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := fieldValueClient.Get()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get field value response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("UCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AssistantSid).To(Equal("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FieldTypeSid).To(Equal("UBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Language).To(Equal("en-US"))
				Expect(resp.Value).To(Equal("test"))
				Expect(resp.SynonymOf).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/FieldTypes/UBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/FieldValues/UCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the field value api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/FieldTypes/UBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/FieldValues/UC71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := autopilotSession.Assistant("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").FieldType("UBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").FieldValue("UC71").Get()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get field value response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the field value is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/FieldTypes/UBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/FieldValues/UCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := fieldValueClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the field values api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/FieldTypes/UBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/FieldValues/UC71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := autopilotSession.Assistant("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").FieldType("UBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").FieldValue("UC71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the model builds client", func() {
		modelBuildsClient := autopilotSession.Assistant("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").ModelBuilds

		Describe("When the model build is successfully created", func() {
			createInput := &model_builds.CreateModelBuildInput{
				UniqueName: utils.String("test"),
			}

			httpmock.RegisterResponder("POST", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/ModelBuilds",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/modelBuildResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := modelBuildsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create field type response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("UGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AssistantSid).To(Equal("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.UniqueName).To(Equal("test"))
				Expect(resp.Status).To(Equal("enqueued"))
				Expect(resp.BuildDuration).To(BeNil())
				Expect(resp.ErrorCode).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/ModelBuilds/UGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the create model build api returns a 500 response", func() {
			createInput := &model_builds.CreateModelBuildInput{
				UniqueName: utils.String("test"),
			}

			httpmock.RegisterResponder("POST", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/ModelBuilds",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := modelBuildsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create model build response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a model build sid", func() {
		modelBuildClient := autopilotSession.Assistant("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").ModelBuild("UGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the model buildis successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/ModelBuilds/UGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/modelBuildResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := modelBuildClient.Get()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get model build response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("UGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AssistantSid).To(Equal("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.UniqueName).To(Equal("test"))
				Expect(resp.Status).To(Equal("enqueued"))
				Expect(resp.BuildDuration).To(BeNil())
				Expect(resp.ErrorCode).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/ModelBuilds/UGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the model build api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/ModelBuilds/UG71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := autopilotSession.Assistant("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").ModelBuild("UG71").Get()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get model build response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the model build is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/ModelBuilds/UGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateModelBuildResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &model_build.UpdateModelBuildInput{
				UniqueName: utils.String("new name"),
			}

			resp, err := modelBuildClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update model build response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("UGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AssistantSid).To(Equal("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.UniqueName).To(Equal("new name"))
				Expect(resp.Status).To(Equal("completed"))
				Expect(resp.BuildDuration).To(Equal(utils.Int(1000)))
				Expect(resp.ErrorCode).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-27T23:10:00Z"))
				Expect(resp.URL).To(Equal("https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/ModelBuilds/UGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the model builds api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/ModelBuilds/UG71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &model_build.UpdateModelBuildInput{
				UniqueName: utils.String("new name"),
			}

			resp, err := autopilotSession.Assistant("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").ModelBuild("UG71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update model build response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the model build is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/ModelBuilds/UGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := modelBuildClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the model buildss api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://autopilot.twilio.com/v1/Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/ModelBuilds/UG71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := autopilotSession.Assistant("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").ModelBuild("UG71").Delete()
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
