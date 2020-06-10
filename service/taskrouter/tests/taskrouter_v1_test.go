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

	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/activities"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/activity"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/task"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/task_channel"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/task_channels"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/task_queue"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/task_queues"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/tasks"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/worker"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/workers"
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

	Describe("Given the workspace client", func() {
		workspacesClient := taskrouterSession.Workspaces

		Describe("When the workspace is successfully created", func() {
			createInput := &workspaces.CreateWorkspaceInput{
				FriendlyName:     "Test 2",
				MultiTaskEnabled: utils.Bool(false),
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
				Expect(resp.EventCallbackURL).To(Equal(utils.String("https://ngrok.com")))
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

		Describe("When the workspace request does not contain a friendly name", func() {
			createInput := &workspaces.CreateWorkspaceInput{}

			resp, err := workspacesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create workspace response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the workspaces api returns a 500 response", func() {
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

	Describe("Given I have a workspace sid", func() {
		workspaceClient := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the workspace is successfully retrieved", func() {
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
				Expect(resp.EventCallbackURL).To(Equal(utils.String("https://ngrok.com")))
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
				Expect(resp.EventCallbackURL).To(Equal(utils.String("https://ngrok.com")))
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

		Describe("When the update workspace api returns a 404", func() {
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

	Describe("Given the task queues client", func() {
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
				Expect(resp.AssignmentActivityName).To(Equal(utils.String("817ca1c5-3a05-11e5-9292-98e0d9a1eb73")))
				Expect(resp.AssignmentActivitySid).To(Equal(utils.String("WA21d51f4c72583766988f9860de3e130a")))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2015-08-04T01:31:41Z"))
				Expect(resp.FriendlyName).To(Equal("English"))
				Expect(resp.MaxReservedWorkers).To(Equal(1))
				Expect(resp.ReservationActivityName).To(Equal(utils.String("80fa2beb-3a05-11e5-8fc8-98e0d9a1eb74")))
				Expect(resp.ReservationActivitySid).To(Equal(utils.String("WAea296a56ebce4bfbff0e99abadf16934")))
				Expect(resp.Sid).To(Equal("WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.TargetWorkers).To(Equal(utils.String("languages HAS \"english\"")))
				Expect(resp.TaskOrder).To(Equal("FIFO"))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskQueues/WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.WorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the task queue request does not contain a friendly name", func() {
			createInput := &task_queues.CreateTaskQueueInput{}

			resp, err := taskQueuesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create task queue response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the task queue api returns a 500 response", func() {
			createInput := &task_queues.CreateTaskQueueInput{
				FriendlyName: "Test 2",
			}

			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskQueues",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := taskQueuesClient.Create(createInput)
			It("Then an error should be returned", func() {
				Expect(err).ToNot(BeNil())
				twilioErr, ok := err.(*utils.TwilioError)
				Expect(ok).To(Equal(true))
				Expect(twilioErr.Code).To(BeNil())
				Expect(twilioErr.Message).To(Equal("An error occurred"))
				Expect(twilioErr.MoreInfo).To(BeNil())
				Expect(twilioErr.Status).To(Equal(500))
			})

			It("Then the create task queue response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a task queue sid", func() {
		taskQueueClient := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").TaskQueue("WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the task queue is successfully retrieved", func() {
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
				Expect(resp.AssignmentActivityName).To(Equal(utils.String("817ca1c5-3a05-11e5-9292-98e0d9a1eb73")))
				Expect(resp.AssignmentActivitySid).To(Equal(utils.String("WA21d51f4c72583766988f9860de3e130a")))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2015-08-04T01:31:41Z"))
				Expect(resp.FriendlyName).To(Equal("English"))
				Expect(resp.MaxReservedWorkers).To(Equal(1))
				Expect(resp.ReservationActivityName).To(Equal(utils.String("80fa2beb-3a05-11e5-8fc8-98e0d9a1eb74")))
				Expect(resp.ReservationActivitySid).To(Equal(utils.String("WAea296a56ebce4bfbff0e99abadf16934")))
				Expect(resp.Sid).To(Equal("WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.TargetWorkers).To(Equal(utils.String("languages HAS \"english\"")))
				Expect(resp.TaskOrder).To(Equal("FIFO"))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskQueues/WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.WorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the get task queue api returns a 404", func() {
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

		Describe("When the task queue is successfully updated", func() {
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

			It("Then the update task queue response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AssignmentActivityName).To(Equal(utils.String("817ca1c5-3a05-11e5-9292-98e0d9a1eb73")))
				Expect(resp.AssignmentActivitySid).To(Equal(utils.String("WA21d51f4c72583766988f9860de3e130a")))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2015-08-04T02:31:41Z"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2015-08-04T01:31:41Z"))
				Expect(resp.FriendlyName).To(Equal("English"))
				Expect(resp.MaxReservedWorkers).To(Equal(1))
				Expect(resp.ReservationActivityName).To(Equal(utils.String("80fa2beb-3a05-11e5-8fc8-98e0d9a1eb74")))
				Expect(resp.ReservationActivitySid).To(Equal(utils.String("WAea296a56ebce4bfbff0e99abadf16934")))
				Expect(resp.Sid).To(Equal("WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.TargetWorkers).To(Equal(utils.String("languages HAS \"english\"")))
				Expect(resp.TaskOrder).To(Equal("FIFO"))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskQueues/WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.WorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update task queue api returns a 404", func() {
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

		Describe("When the task queue is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskQueues/WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := taskQueueClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the delete task queue api returns a 404", func() {
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

	Describe("Given the workflows client", func() {
		workflowsClient := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Workflows

		Describe("When the workflow is successfully created", func() {
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
				Expect(resp.AssignmentCallbackURL).To(Equal(utils.String("https://example.com/")))

				configuration := make(map[string]interface{})
				json.Unmarshal(taskRoutingConfiguration, &configuration)
				Expect(resp.Configuration).To(Equal(configuration))

				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2014-05-14T10:50:02Z"))
				Expect(resp.DocumentContentType).To(Equal("application/json"))
				Expect(resp.FallbackAssignmentCallbackURL).To(Equal(utils.String("https://example2.com/")))
				Expect(resp.FriendlyName).To(Equal("Sales, Marketing, Support Workflow"))
				Expect(resp.Sid).To(Equal("WFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.TaskReservationTimeout).To(Equal(120))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workflows/WFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.WorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the workflow request does not contain a friendly name", func() {
			taskRoutingConfiguration, _ := ioutil.ReadFile("testdata/taskRoutingConfiguration.json")

			createInput := &workflows.CreateWorkflowInput{
				Configuration: string(taskRoutingConfiguration),
			}

			resp, err := workflowsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create workflow response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the workflow request does not contain a configuration", func() {
			createInput := &workflows.CreateWorkflowInput{
				FriendlyName: "Test 2",
			}

			resp, err := workflowsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create workflow response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the workflow api returns a 500 response", func() {
			taskRoutingConfiguration, _ := ioutil.ReadFile("testdata/taskRoutingConfiguration.json")

			createInput := &workflows.CreateWorkflowInput{
				FriendlyName:  "Test 2",
				Configuration: string(taskRoutingConfiguration),
			}

			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workflows",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := workflowsClient.Create(createInput)
			It("Then an error should be returned", func() {
				Expect(err).ToNot(BeNil())
				twilioErr, ok := err.(*utils.TwilioError)
				Expect(ok).To(Equal(true))
				Expect(twilioErr.Code).To(BeNil())
				Expect(twilioErr.Message).To(Equal("An error occurred"))
				Expect(twilioErr.MoreInfo).To(BeNil())
				Expect(twilioErr.Status).To(Equal(500))
			})

			It("Then the create workflow response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a workflow sid", func() {
		workflowClient := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Workflow("WFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the workflow is successfully retrieved", func() {
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
				Expect(resp.AssignmentCallbackURL).To(Equal(utils.String("https://example.com/")))

				configuration := make(map[string]interface{})
				json.Unmarshal(taskRoutingConfiguration, &configuration)
				Expect(resp.Configuration).To(Equal(configuration))

				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2014-05-14T10:50:02Z"))
				Expect(resp.DocumentContentType).To(Equal("application/json"))
				Expect(resp.FallbackAssignmentCallbackURL).To(Equal(utils.String("https://example2.com/")))
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

		Describe("When the workflow is successfully updated", func() {
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
				Expect(resp.AssignmentCallbackURL).To(Equal(utils.String("https://example.com/")))

				configuration := make(map[string]interface{})
				json.Unmarshal(taskRoutingConfiguration, &configuration)
				Expect(resp.Configuration).To(Equal(configuration))

				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2014-05-14T11:50:02Z"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2014-05-14T10:50:02Z"))
				Expect(resp.DocumentContentType).To(Equal("application/json"))
				Expect(resp.FallbackAssignmentCallbackURL).To(Equal(utils.String("https://example2.com/")))
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

			It("Then the update workflow response should be nil", func() {
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

		Describe("When the delete workflow API returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workflows/WF71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Workflow("WF71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the workers client", func() {
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
				Expect(resp.DateStatusChanged.Format(time.RFC3339)).To(Equal("2017-05-30T23:19:38Z"))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the worker request does not contain a friendly name", func() {
			createInput := &workers.CreateWorkerInput{}

			resp, err := workersClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create worker response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the worker api returns a 500 response", func() {
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

	Describe("Given I have a worker sid", func() {
		workerClient := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Worker("WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the worker is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/workersResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := workerClient.Get()
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
				Expect(resp.DateStatusChanged.Format(time.RFC3339)).To(Equal("2017-05-30T23:19:38Z"))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the get worker api returns a 404", func() {
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

		Describe("When the worker is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updatedWorkersResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &worker.UpdateWorkerInput{}

			resp, err := workerClient.Update(updateInput)
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
				Expect(resp.DateStatusChanged.Format(time.RFC3339)).To(Equal("2017-05-30T23:19:38Z"))
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

			err := workerClient.Delete()
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

	Describe("Given the activities client", func() {
		activitiesClient := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Activities

		Describe("When the activity is successfully created", func() {
			createInput := &activities.CreateActivityInput{
				FriendlyName: "NewAvailableActivity",
				Available:    utils.Bool(true),
			}

			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Activities",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/activityResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := activitiesClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create activity response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.WorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("NewAvailableActivity"))
				Expect(resp.Available).To(Equal(true))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2014-05-14T10:50:02Z"))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Activities/WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the activity request does not contain a friendly name", func() {
			createInput := &activities.CreateActivityInput{}

			resp, err := activitiesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create activity response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the activities api returns a 500 response", func() {
			createInput := &activities.CreateActivityInput{
				FriendlyName: "NewAvailableActivity",
				Available:    utils.Bool(true),
			}

			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Activities",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := activitiesClient.Create(createInput)
			It("Then an error should be returned", func() {
				Expect(err).ToNot(BeNil())
				twilioErr, ok := err.(*utils.TwilioError)
				Expect(ok).To(Equal(true))
				Expect(twilioErr.Code).To(BeNil())
				Expect(twilioErr.Message).To(Equal("An error occurred"))
				Expect(twilioErr.MoreInfo).To(BeNil())
				Expect(twilioErr.Status).To(Equal(500))
			})

			It("Then the create activity response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a activity sid", func() {
		activityClient := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Activity("WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the activity is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Activities/WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/activityResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := activityClient.Get()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get activity response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.WorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("NewAvailableActivity"))
				Expect(resp.Available).To(Equal(true))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2014-05-14T10:50:02Z"))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Activities/WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the get activity response returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Activities/WA71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Activity("WA71").Get()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get activity response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the activity is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Activities/WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updatedActivityResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &activity.UpdateActivityInput{}

			resp, err := activityClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update activity response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.WorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("NewAvailableActivity"))
				Expect(resp.Available).To(Equal(true))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2014-05-14T23:26:06Z"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2014-05-14T10:50:02Z"))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Activities/WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update activity response returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Activities/WA71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &activity.UpdateActivityInput{
				FriendlyName: "Test Activity",
			}

			resp, err := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Activity("WA71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update activity response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the activity is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Activities/WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := activityClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the delete activity api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Activities/WA71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Activity("WA71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the task channels client", func() {
		taskChannelsClient := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").TaskChannels

		Describe("When the task channel is successfully created", func() {
			createInput := &task_channels.CreateTaskChannelInput{
				FriendlyName: "Test 2",
				UniqueName:   "Unique Test 2",
			}

			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskChannels",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/taskChannelResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := taskChannelsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create task channel response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.WorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Sid).To(Equal("TCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test 2"))
				Expect(resp.UniqueName).To(Equal("Unique Test 2"))
				Expect(resp.ChannelOptimizedRouting).To(Equal(utils.Bool(true)))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2016-04-14T17:35:54Z"))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskChannels/TCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the task channel request does not contain a friendly name", func() {
			createInput := &task_channels.CreateTaskChannelInput{
				UniqueName: "Unique Test 2",
			}

			resp, err := taskChannelsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create task channel response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the task channel request does not contain a unique name", func() {
			createInput := &task_channels.CreateTaskChannelInput{
				FriendlyName: "Test 2",
			}

			resp, err := taskChannelsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create task channel response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the task channel api returns a 500 response", func() {
			createInput := &task_channels.CreateTaskChannelInput{
				FriendlyName: "Test 2",
				UniqueName:   "Unique Test 2",
			}

			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskChannels",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := taskChannelsClient.Create(createInput)
			It("Then an error should be returned", func() {
				Expect(err).ToNot(BeNil())
				twilioErr, ok := err.(*utils.TwilioError)
				Expect(ok).To(Equal(true))
				Expect(twilioErr.Code).To(BeNil())
				Expect(twilioErr.Message).To(Equal("An error occurred"))
				Expect(twilioErr.MoreInfo).To(BeNil())
				Expect(twilioErr.Status).To(Equal(500))
			})

			It("Then the create task channel response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a task channel sid", func() {
		taskChannelClient := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").TaskChannel("TCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the task channel is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskChannels/TCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/taskChannelResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := taskChannelClient.Get()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get task channel response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.WorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Sid).To(Equal("TCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test 2"))
				Expect(resp.UniqueName).To(Equal("Unique Test 2"))
				Expect(resp.ChannelOptimizedRouting).To(Equal(utils.Bool(true)))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2016-04-14T17:35:54Z"))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskChannels/TCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the get task channel response returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskChannels/TC71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").TaskChannel("TC71").Get()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get task channel response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the task channel is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskChannels/TCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updatedTaskChannelResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &task_channel.UpdateTaskChannelInput{}

			resp, err := taskChannelClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update task channel response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.WorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Sid).To(Equal("TCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test 2"))
				Expect(resp.UniqueName).To(Equal("Unique Test 2"))
				Expect(resp.ChannelOptimizedRouting).To(Equal(utils.Bool(true)))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2016-04-15T17:35:54Z"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2016-04-14T17:35:54Z"))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskChannels/TCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update task channel response returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskChannels/TC71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &task_channel.UpdateTaskChannelInput{
				FriendlyName: "Test Task Channel",
			}

			resp, err := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").TaskChannel("TC71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update task channel response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the task channel is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskChannels/TCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := taskChannelClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the delete task channel response returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskChannels/TC71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").TaskChannel("TC71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the tasks client", func() {
		tasksClient := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Tasks

		Describe("When the task is successfully created", func() {
			createInput := &tasks.CreateTaskInput{
				TaskChannel: "default",
			}

			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks",
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
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.WorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Sid).To(Equal("WTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Age).To(Equal(25200))
				Expect(resp.AssignmentStatus).To(Equal("pending"))

				attributes := make(map[string]interface{})
				attributes["type"] = "support"
				Expect(resp.Attributes).To(Equal(attributes))

				Expect(resp.Priority).To(Equal(utils.Int(1)))
				Expect(resp.Reason).To(Equal(utils.String("Test Reason")))
				Expect(resp.TaskQueueSid).To(Equal(utils.String("WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.TaskChannelSid).To(Equal(utils.String("TCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.TaskChannelUniqueName).To(Equal(utils.String("unique")))
				Expect(resp.WorkflowSid).To(Equal(utils.String("WWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.WorkflowFriendlyName).To(Equal(utils.String("Example Workflow")))
				Expect(resp.TaskQueueEnteredDate).To(BeNil())
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2014-05-14T18:50:02Z"))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/WTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the task api returns a 500 response", func() {
			createInput := &tasks.CreateTaskInput{
				TaskChannel: "default",
			}

			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := tasksClient.Create(createInput)
			It("Then an error should be returned", func() {
				Expect(err).ToNot(BeNil())
				twilioErr, ok := err.(*utils.TwilioError)
				Expect(ok).To(Equal(true))
				Expect(twilioErr.Code).To(BeNil())
				Expect(twilioErr.Message).To(Equal("An error occurred"))
				Expect(twilioErr.MoreInfo).To(BeNil())
				Expect(twilioErr.Status).To(Equal(500))
			})

			It("Then the create task response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a task sid", func() {
		taskClient := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Task("WTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the task is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/WTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
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
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.WorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Sid).To(Equal("WTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Age).To(Equal(25200))
				Expect(resp.AssignmentStatus).To(Equal("pending"))

				attributes := make(map[string]interface{})
				attributes["type"] = "support"
				Expect(resp.Attributes).To(Equal(attributes))

				Expect(resp.Priority).To(Equal(utils.Int(1)))
				Expect(resp.Reason).To(Equal(utils.String("Test Reason")))
				Expect(resp.TaskQueueSid).To(Equal(utils.String("WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.TaskChannelSid).To(Equal(utils.String("TCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.TaskChannelUniqueName).To(Equal(utils.String("unique")))
				Expect(resp.WorkflowSid).To(Equal(utils.String("WWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.WorkflowFriendlyName).To(Equal(utils.String("Example Workflow")))
				Expect(resp.TaskQueueEnteredDate).To(BeNil())
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2014-05-14T18:50:02Z"))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/WTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the get task response returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/WT71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Task("WT71").Get()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get task response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the task is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/WTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updatedTaskResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &task.UpdateTaskInput{}

			resp, err := taskClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update task response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.WorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Sid).To(Equal("WTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Age).To(Equal(25200))
				Expect(resp.AssignmentStatus).To(Equal("pending"))

				attributes := make(map[string]interface{})
				attributes["type"] = "support"
				Expect(resp.Attributes).To(Equal(attributes))

				Expect(resp.Priority).To(Equal(utils.Int(1)))
				Expect(resp.Reason).To(Equal(utils.String("Test Reason")))
				Expect(resp.TaskQueueSid).To(Equal(utils.String("WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.TaskChannelSid).To(Equal(utils.String("TCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.TaskChannelUniqueName).To(Equal(utils.String("unique")))
				Expect(resp.WorkflowSid).To(Equal(utils.String("WWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.WorkflowFriendlyName).To(Equal(utils.String("Example Workflow")))
				Expect(resp.TaskQueueEnteredDate).To(BeNil())
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2014-05-15T18:50:02Z"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2014-05-14T18:50:02Z"))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/WTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update task response returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/WT71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &task.UpdateTaskInput{}

			resp, err := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Task("WT71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update task response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the task is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/WTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := taskClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the delete task response returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/WT71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Task("WT71").Delete()
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
	Expect(twilioErr.Code).To(Equal(utils.Int(20404)))
	Expect(twilioErr.Message).To(Equal("The requested resource /Flows/FW71 was not found"))
	Expect(twilioErr.MoreInfo).To(Equal(utils.String("https://www.twilio.com/docs/errors/20404")))
	Expect(twilioErr.Status).To(Equal(404))
}

func ExpectErrorToNotBeATwilioError(err error) {
	Expect(err).ToNot(BeNil())
	_, ok := err.(*utils.TwilioError)
	Expect(ok).To(Equal(false))
}
