package tests

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/worker"

	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/workers"

	"github.com/jarcoal/httpmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/task_queue"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/task_queues"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/workflow"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/workflows"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspaces"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

var _ = Describe("Taskrouter V1", func() {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	creds, err := credentials.New(credentials.Account{
		Sid:       "ACxxx",
		AuthToken: "Test",
	})
	if err != nil {
		log.Panicf("%s", err)
	}

	taskrouterSession := taskrouter.NewWithCredentials(creds).V1

	Describe("Given the Workspace Client", func() {
		workspacesClient := taskrouterSession.Workspaces

		Describe("When the Task Queue is successfully created", func() {
			createInput := &workspaces.CreateWorkspaceInput{
				FriendlyName: "Test 2",
			}

			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/workspaceResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := workspacesClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create workspace response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.DefaultActivityName).To(Equal("Offline"))
				Expect(resp.DefaultActivitySid).To(Equal("WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2016-08-01T22:10:40Z"))
				eventCallbackURL := "https://ngrok.com"
				Expect(resp.EventCallbackURL).To(Equal(&eventCallbackURL))
				Expect(resp.EventsFilter).To(BeNil())
				Expect(resp.FriendlyName).To(Equal("NewWorkspace"))
				Expect(resp.MultiTaskEnabled).To(Equal(false))
				Expect(resp.PrioritizeQueueOrder).To(Equal("FIFO"))
				Expect(resp.TimeoutActivityName).To(Equal("Offline"))
				Expect(resp.TimeoutActivitySid).To(Equal("WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX1"))
				Expect(resp.Sid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the Workspace request does not contain a friendly name", func() {
			createInput := &workspaces.CreateWorkspaceInput{}

			resp, err := workspacesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create workspace response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the Workspaces API returns a 500 response", func() {
			createInput := &workspaces.CreateWorkspaceInput{
				FriendlyName: "Test 2",
			}

			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := workspacesClient.Create(createInput)
			It("Then an error should be returned", func() {
				Expect(err).ToNot(BeNil())
				twilioErr, ok := err.(*utils.TwilioError)
				Expect(ok).To(Equal(true))
				Expect(twilioErr.Code).To(BeNil())
				Expect(twilioErr.Message).To(Equal("An error occurred"))
				Expect(twilioErr.MoreInfo).To(BeNil())
				Expect(twilioErr.Status).To(Equal(500))
			})

			It("Then the create workspace response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a Workspace SID", func() {
		workspaceClient := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the Workspace is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/workspaceResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := workspaceClient.Get()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get workspace response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.DefaultActivityName).To(Equal("Offline"))
				Expect(resp.DefaultActivitySid).To(Equal("WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2016-08-01T22:10:40Z"))
				eventCallbackURL := "https://ngrok.com"
				Expect(resp.EventCallbackURL).To(Equal(&eventCallbackURL))
				Expect(resp.EventsFilter).To(BeNil())
				Expect(resp.FriendlyName).To(Equal("NewWorkspace"))
				Expect(resp.MultiTaskEnabled).To(Equal(false))
				Expect(resp.PrioritizeQueueOrder).To(Equal("FIFO"))
				Expect(resp.TimeoutActivityName).To(Equal("Offline"))
				Expect(resp.TimeoutActivitySid).To(Equal("WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX1"))
				Expect(resp.Sid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the get workspace response returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WS71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := taskrouterSession.Workspace("WS71").Get()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get workspace response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the Workspace is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updatedWorkspaceResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &workspace.UpdateWorkspaceInput{}

			resp, err := workspaceClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update workspace response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.DefaultActivityName).To(Equal("Offline"))
				Expect(resp.DefaultActivitySid).To(Equal("WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2016-08-01T23:10:40Z"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2016-08-01T22:10:40Z"))
				eventCallbackURL := "https://ngrok.com"
				Expect(resp.EventCallbackURL).To(Equal(&eventCallbackURL))
				Expect(resp.EventsFilter).To(BeNil())
				Expect(resp.FriendlyName).To(Equal("NewWorkspace"))
				Expect(resp.MultiTaskEnabled).To(Equal(false))
				Expect(resp.PrioritizeQueueOrder).To(Equal("FIFO"))
				Expect(resp.TimeoutActivityName).To(Equal("Offline"))
				Expect(resp.TimeoutActivitySid).To(Equal("WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX1"))
				Expect(resp.Sid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update flow response returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WS71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &workspace.UpdateWorkspaceInput{
				FriendlyName: "Test Workspace",
			}

			resp, err := taskrouterSession.Workspace("WS71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update workspace response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the workspace is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := workspaceClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the delete workspace response returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://taskrouter.twilio.com/v1/Workspaces/WS71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := taskrouterSession.Workspace("WS71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the Task Queues Client", func() {
		taskQueuesClient := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").TaskQueues

		Describe("When the Task Queue is successfully created", func() {
			createInput := &task_queues.CreateTaskQueueInput{
				FriendlyName: "Test 2",
			}

			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskQueues",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/taskQueueResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := taskQueuesClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create task queue response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AssignmentActivityName).To(Equal("817ca1c5-3a05-11e5-9292-98e0d9a1eb73"))
				Expect(resp.AssignmentActivitySid).To(Equal("WA21d51f4c72583766988f9860de3e130a"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2015-08-04T01:31:41Z"))
				Expect(resp.FriendlyName).To(Equal("English"))
				Expect(resp.MaxReservedWorkers).To(Equal(1))
				Expect(resp.ReservationActivityName).To(Equal("80fa2beb-3a05-11e5-8fc8-98e0d9a1eb74"))
				Expect(resp.ReservationActivitySid).To(Equal("WAea296a56ebce4bfbff0e99abadf16934"))
				Expect(resp.Sid).To(Equal("WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.TargetWorkers).To(Equal("languages HAS \"english\""))
				Expect(resp.TaskOrder).To(Equal("FIFO"))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskQueues/WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.WorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})
	})

	Describe("Given I have a Task Queue SID", func() {
		taskQueueClient := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").TaskQueue("WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the Task Queue is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskQueues/WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/taskQueueResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := taskQueueClient.Get()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get task queue response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AssignmentActivityName).To(Equal("817ca1c5-3a05-11e5-9292-98e0d9a1eb73"))
				Expect(resp.AssignmentActivitySid).To(Equal("WA21d51f4c72583766988f9860de3e130a"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2015-08-04T01:31:41Z"))
				Expect(resp.FriendlyName).To(Equal("English"))
				Expect(resp.MaxReservedWorkers).To(Equal(1))
				Expect(resp.ReservationActivityName).To(Equal("80fa2beb-3a05-11e5-8fc8-98e0d9a1eb74"))
				Expect(resp.ReservationActivitySid).To(Equal("WAea296a56ebce4bfbff0e99abadf16934"))
				Expect(resp.Sid).To(Equal("WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.TargetWorkers).To(Equal("languages HAS \"english\""))
				Expect(resp.TaskOrder).To(Equal("FIFO"))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskQueues/WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.WorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the get workspace response returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskQueues/WQ71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").TaskQueue("WQ71").Get()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get task queue response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the Workspace is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskQueues/WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updatedTaskQueueResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &task_queue.UpdateTaskQueueInput{}

			resp, err := taskQueueClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update workspace response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AssignmentActivityName).To(Equal("817ca1c5-3a05-11e5-9292-98e0d9a1eb73"))
				Expect(resp.AssignmentActivitySid).To(Equal("WA21d51f4c72583766988f9860de3e130a"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2015-08-04T02:31:41Z"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2015-08-04T01:31:41Z"))
				Expect(resp.FriendlyName).To(Equal("English"))
				Expect(resp.MaxReservedWorkers).To(Equal(1))
				Expect(resp.ReservationActivityName).To(Equal("80fa2beb-3a05-11e5-8fc8-98e0d9a1eb74"))
				Expect(resp.ReservationActivitySid).To(Equal("WAea296a56ebce4bfbff0e99abadf16934"))
				Expect(resp.Sid).To(Equal("WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.TargetWorkers).To(Equal("languages HAS \"english\""))
				Expect(resp.TaskOrder).To(Equal("FIFO"))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskQueues/WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.WorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update flow response returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskQueues/WQ71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &task_queue.UpdateTaskQueueInput{
				FriendlyName: "Test Queue",
			}

			resp, err := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").TaskQueue("WQ71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update task queue response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the task is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskQueues/WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := taskQueueClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the delete workspace response returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskQueues/WQ71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").TaskQueue("WQ71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the Workflow Client", func() {
		workflowsClient := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Workflows

		Describe("When the Workflow is successfully created", func() {
			taskRoutingConfiguration, _ := ioutil.ReadFile("testdata/taskRoutingConfiguration.json")

			createInput := &workflows.CreateWorkflowInput{
				FriendlyName:  "Test 2",
				Configuration: string(taskRoutingConfiguration),
			}

			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workflows",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/workflowsResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := workflowsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create workflow response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))

				assignmentCallbackURL := "https://example.com/"
				Expect(resp.AssignmentCallbackURL).To(Equal(&assignmentCallbackURL))

				configuration := make(map[string]interface{})
				json.Unmarshal(taskRoutingConfiguration, &configuration)
				Expect(resp.Configuration).To(Equal(configuration))

				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2014-05-14T10:50:02Z"))
				Expect(resp.DocumentContentType).To(Equal("application/json"))

				fallbackAssignmentCallbackURL := "https://example2.com/"
				Expect(resp.FallbackAssignmentCallbackURL).To(Equal(&fallbackAssignmentCallbackURL))

				Expect(resp.FriendlyName).To(Equal("Sales, Marketing, Support Workflow"))
				Expect(resp.Sid).To(Equal("WFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.TaskReservationTimeout).To(Equal(120))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workflows/WFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.WorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})
	})

	Describe("Given I have a Workflow SID", func() {
		workflowClient := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Workflow("WFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the Workflow is successfully retrieved", func() {
			taskRoutingConfiguration, _ := ioutil.ReadFile("testdata/taskRoutingConfiguration.json")

			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workflows/WFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/workflowsResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := workflowClient.Get()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get workflow response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				assignmentCallbackURL := "https://example.com/"
				Expect(resp.AssignmentCallbackURL).To(Equal(&assignmentCallbackURL))

				configuration := make(map[string]interface{})
				json.Unmarshal(taskRoutingConfiguration, &configuration)
				Expect(resp.Configuration).To(Equal(configuration))

				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2014-05-14T10:50:02Z"))
				Expect(resp.DocumentContentType).To(Equal("application/json"))

				fallbackAssignmentCallbackURL := "https://example2.com/"
				Expect(resp.FallbackAssignmentCallbackURL).To(Equal(&fallbackAssignmentCallbackURL))

				Expect(resp.FriendlyName).To(Equal("Sales, Marketing, Support Workflow"))
				Expect(resp.Sid).To(Equal("WFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.TaskReservationTimeout).To(Equal(120))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workflows/WFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.WorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the get workflow response returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workflows/WF71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Workflow("WF71").Get()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get workflow response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the Workflow is successfully updated", func() {
			taskRoutingConfiguration, _ := ioutil.ReadFile("testdata/taskRoutingConfiguration.json")

			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workflows/WFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updatedWorkflowsResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &workflow.UpdateWorkflowInput{}

			resp, err := workflowClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update workflow response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))

				assignmentCallbackURL := "https://example.com/"
				Expect(resp.AssignmentCallbackURL).To(Equal(&assignmentCallbackURL))

				configuration := make(map[string]interface{})
				json.Unmarshal(taskRoutingConfiguration, &configuration)
				Expect(resp.Configuration).To(Equal(configuration))

				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2014-05-14T11:50:02Z"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2014-05-14T10:50:02Z"))
				Expect(resp.DocumentContentType).To(Equal("application/json"))

				fallbackAssignmentCallbackURL := "https://example2.com/"
				Expect(resp.FallbackAssignmentCallbackURL).To(Equal(&fallbackAssignmentCallbackURL))

				Expect(resp.FriendlyName).To(Equal("Sales, Marketing, Support Workflow"))
				Expect(resp.Sid).To(Equal("WFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.TaskReservationTimeout).To(Equal(120))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workflows/WFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.WorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update workflow response returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workflows/WF71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &workflow.UpdateWorkflowInput{
				FriendlyName: "Test Queue",
			}

			resp, err := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Workflow("WF71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update task queue response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the workflow is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workflows/WFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := workflowClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the delete workflow response returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workflows/WF71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").TaskQueue("WQ71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the Worker Client", func() {
		workersClient := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Workers

		Describe("When the Worker is successfully created", func() {
			createInput := &workers.CreateWorkerInput{
				FriendlyName: "NewWorker",
			}

			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/workersResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := workersClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create workers response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("NewWorker"))
				Expect(resp.ActivitySid).To(Equal("WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ActivityName).To(Equal("Offline"))
				Expect(resp.WorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))

				attributes := make(map[string]interface{})
				Expect(resp.Attributes).To(Equal(attributes))
				Expect(resp.Available).To(Equal(false))

				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2017-05-30T23:19:38Z"))
				Expect(resp.DateStatusChange.Format(time.RFC3339)).To(Equal("2017-05-30T23:19:38Z"))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the Worker request does not contain a friendly name", func() {
			createInput := &workers.CreateWorkerInput{}

			resp, err := workersClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create worker response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the Worker API returns a 500 response", func() {
			createInput := &workers.CreateWorkerInput{
				FriendlyName: "NewWorker",
			}

			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := workersClient.Create(createInput)
			It("Then an error should be returned", func() {
				Expect(err).ToNot(BeNil())
				twilioErr, ok := err.(*utils.TwilioError)
				Expect(ok).To(Equal(true))
				Expect(twilioErr.Code).To(BeNil())
				Expect(twilioErr.Message).To(Equal("An error occurred"))
				Expect(twilioErr.MoreInfo).To(BeNil())
				Expect(twilioErr.Status).To(Equal(500))
			})

			It("Then the create worker response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a Worker SID", func() {
		workersClient := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Worker("WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the Worker is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/workersResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := workersClient.Get()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get worker response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("NewWorker"))
				Expect(resp.ActivitySid).To(Equal("WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ActivityName).To(Equal("Offline"))
				Expect(resp.WorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))

				attributes := make(map[string]interface{})
				Expect(resp.Attributes).To(Equal(attributes))
				Expect(resp.Available).To(Equal(false))

				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2017-05-30T23:19:38Z"))
				Expect(resp.DateStatusChange.Format(time.RFC3339)).To(Equal("2017-05-30T23:19:38Z"))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the get worker response returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WK71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Worker("WK71").Get()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get worker response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the Worker is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updatedWorkersResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &worker.UpdateWorkerInput{}

			resp, err := workersClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update worker response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("NewWorker"))
				Expect(resp.ActivitySid).To(Equal("WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ActivityName).To(Equal("Offline"))
				Expect(resp.WorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))

				attributes := make(map[string]interface{})
				Expect(resp.Attributes).To(Equal(attributes))
				Expect(resp.Available).To(Equal(false))

				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2017-05-31T23:19:38Z"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2017-05-30T23:19:38Z"))
				Expect(resp.DateStatusChange.Format(time.RFC3339)).To(Equal("2017-05-30T23:19:38Z"))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update worker response returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WK71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &worker.UpdateWorkerInput{
				FriendlyName: "Test Worker",
			}

			resp, err := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Worker("WK71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update worker response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the worker is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := workersClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the delete worker response returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WK71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Worker("WK71").Delete()
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
