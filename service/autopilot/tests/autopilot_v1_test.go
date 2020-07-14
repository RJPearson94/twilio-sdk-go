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
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/style_sheet"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/task"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/task/actions"
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
