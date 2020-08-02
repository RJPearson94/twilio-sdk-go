package acceptance

import (
	"fmt"
	"os"

	"github.com/google/uuid"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/RJPearson94/twilio-sdk-go"
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

var _ = Describe("Taskrouter Acceptance Tests", func() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		Fail(fmt.Sprintf("Failed to create credentials. Error %s", err.Error()))
	}

	taskrouterSession := twilio.NewWithCredentials(creds).TaskRouter.V1

	Describe("Given the TaskRouter Workspace clients", func() {
		It("Then the workspace is create, fetched, updated and deleted", func() {
			createResp, createErr := taskrouterSession.Workspaces.Create(&workspaces.CreateWorkspaceInput{
				FriendlyName: uuid.New().String(),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			workspaceClient := taskrouterSession.Workspace(createResp.Sid)

			fetchResp, fetchErr := workspaceClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := workspaceClient.Update(&workspace.UpdateWorkspaceInput{})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := workspaceClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the TaskRouter Task Queue clients", func() {

		var workspaceSid string

		BeforeEach(func() {
			resp, err := taskrouterSession.Workspaces.Create(&workspaces.CreateWorkspaceInput{
				FriendlyName: uuid.New().String(),
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create workspace. Error %s", err.Error()))
			}
			workspaceSid = resp.Sid
		})

		AfterEach(func() {
			if err := taskrouterSession.Workspace(workspaceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete workspace. Error %s", err.Error()))
			}
		})

		It("Then the task queue is create, fetched, updated and deleted", func() {
			createResp, createErr := taskrouterSession.Workspace(workspaceSid).TaskQueues.Create(&task_queues.CreateTaskQueueInput{
				FriendlyName: uuid.New().String(),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			taskQueueClient := taskrouterSession.Workspace(workspaceSid).TaskQueue(createResp.Sid)

			fetchResp, fetchErr := taskQueueClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := taskQueueClient.Update(&task_queue.UpdateTaskQueueInput{})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := taskQueueClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the TaskRouter Workflow clients", func() {

		var workspaceSid string
		var taskQueueSid string

		BeforeEach(func() {
			resp, err := taskrouterSession.Workspaces.Create(&workspaces.CreateWorkspaceInput{
				FriendlyName: uuid.New().String(),
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create workspace. Error %s", err.Error()))
			}
			workspaceSid = resp.Sid

			taskQueueResp, taskQueueErr := taskrouterSession.Workspace(workspaceSid).TaskQueues.Create(&task_queues.CreateTaskQueueInput{
				FriendlyName: uuid.New().String(),
			})
			if taskQueueErr != nil {
				Fail(fmt.Sprintf("Failed to create task queue. Error %s", taskQueueErr.Error()))
			}
			taskQueueSid = taskQueueResp.Sid
		})

		AfterEach(func() {
			if err := taskrouterSession.Workspace(workspaceSid).TaskQueue(taskQueueSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete task queue. Error %s", err.Error()))
			}

			if err := taskrouterSession.Workspace(workspaceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete workspace. Error %s", err.Error()))
			}
		})

		It("Then the workflow is create, fetched, updated and deleted", func() {
			createResp, createErr := taskrouterSession.Workspace(workspaceSid).Workflows.Create(&workflows.CreateWorkflowInput{
				FriendlyName:  uuid.New().String(),
				Configuration: fmt.Sprintf(`{ "task_routing": { "default_filter": { "queue": "%s" } } }`, taskQueueSid),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			workflowClient := taskrouterSession.Workspace(workspaceSid).Workflow(createResp.Sid)

			fetchResp, fetchErr := workflowClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := workflowClient.Update(&workflow.UpdateWorkflowInput{})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := workflowClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the TaskRouter Worker clients", func() {

		var workspaceSid string

		BeforeEach(func() {
			resp, err := taskrouterSession.Workspaces.Create(&workspaces.CreateWorkspaceInput{
				FriendlyName: uuid.New().String(),
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create workspace. Error %s", err.Error()))
			}
			workspaceSid = resp.Sid
		})

		AfterEach(func() {
			if err := taskrouterSession.Workspace(workspaceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete workspace. Error %s", err.Error()))
			}
		})

		It("Then the worker is create, fetched, updated and deleted", func() {
			createResp, createErr := taskrouterSession.Workspace(workspaceSid).Workers.Create(&workers.CreateWorkerInput{
				FriendlyName: uuid.New().String(),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			workerClient := taskrouterSession.Workspace(workspaceSid).Worker(createResp.Sid)

			fetchResp, fetchErr := workerClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := workerClient.Update(&worker.UpdateWorkerInput{})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := workerClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the TaskRouter Activity clients", func() {

		var workspaceSid string

		BeforeEach(func() {
			resp, err := taskrouterSession.Workspaces.Create(&workspaces.CreateWorkspaceInput{
				FriendlyName: uuid.New().String(),
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create workspace. Error %s", err.Error()))
			}
			workspaceSid = resp.Sid
		})

		AfterEach(func() {
			if err := taskrouterSession.Workspace(workspaceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete workspace. Error %s", err.Error()))
			}
		})

		It("Then the activity is create, fetched, updated and deleted", func() {
			createResp, createErr := taskrouterSession.Workspace(workspaceSid).Activities.Create(&activities.CreateActivityInput{
				FriendlyName: uuid.New().String(),
				Available:    utils.Bool(true),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			activityClient := taskrouterSession.Workspace(workspaceSid).Activity(createResp.Sid)

			fetchResp, fetchErr := activityClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := activityClient.Update(&activity.UpdateActivityInput{})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := activityClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the TaskRouter Task Channel clients", func() {

		var workspaceSid string

		BeforeEach(func() {
			resp, err := taskrouterSession.Workspaces.Create(&workspaces.CreateWorkspaceInput{
				FriendlyName: uuid.New().String(),
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create workspace. Error %s", err.Error()))
			}
			workspaceSid = resp.Sid
		})

		AfterEach(func() {
			if err := taskrouterSession.Workspace(workspaceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete workspace. Error %s", err.Error()))
			}
		})

		It("Then the task channel is create, fetched, updated and deleted", func() {
			createResp, createErr := taskrouterSession.Workspace(workspaceSid).TaskChannels.Create(&task_channels.CreateTaskChannelInput{
				FriendlyName: uuid.New().String(),
				UniqueName:   uuid.New().String(),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			taskChannelClient := taskrouterSession.Workspace(workspaceSid).TaskChannel(createResp.Sid)

			fetchResp, fetchErr := taskChannelClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := taskChannelClient.Update(&task_channel.UpdateTaskChannelInput{
				ChannelOptimizedRouting: utils.Bool(true),
			})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := taskChannelClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the TaskRouter Task clients", func() {

		var workspaceSid string
		var taskQueueSid string
		var workflowSid string

		BeforeEach(func() {
			resp, err := taskrouterSession.Workspaces.Create(&workspaces.CreateWorkspaceInput{
				FriendlyName: uuid.New().String(),
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create workspace. Error %s", err.Error()))
			}
			workspaceSid = resp.Sid

			taskQueueResp, taskQueueErr := taskrouterSession.Workspace(workspaceSid).TaskQueues.Create(&task_queues.CreateTaskQueueInput{
				FriendlyName: uuid.New().String(),
			})
			if taskQueueErr != nil {
				Fail(fmt.Sprintf("Failed to create task queue. Error %s", taskQueueErr.Error()))
			}
			taskQueueSid = taskQueueResp.Sid

			workflowResp, workflowErr := taskrouterSession.Workspace(workspaceSid).Workflows.Create(&workflows.CreateWorkflowInput{
				FriendlyName:  uuid.New().String(),
				Configuration: fmt.Sprintf(`{ "task_routing": { "default_filter": { "queue": "%s" } } }`, taskQueueSid),
			})
			if workflowErr != nil {
				Fail(fmt.Sprintf("Failed to create workflow. Error %s", workflowErr.Error()))
			}
			workflowSid = workflowResp.Sid
		})

		AfterEach(func() {
			if err := taskrouterSession.Workspace(workspaceSid).Workflow(workflowSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete workflow. Error %s", err.Error()))
			}

			if err := taskrouterSession.Workspace(workspaceSid).TaskQueue(taskQueueSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete task queue. Error %s", err.Error()))
			}

			if err := taskrouterSession.Workspace(workspaceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete workspace. Error %s", err.Error()))
			}
		})

		It("Then the task is create, fetched, updated and deleted", func() {
			createResp, createErr := taskrouterSession.Workspace(workspaceSid).Tasks.Create(&tasks.CreateTaskInput{
				TaskChannel: utils.String("default"),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			taskClient := taskrouterSession.Workspace(workspaceSid).Task(createResp.Sid)

			fetchResp, fetchErr := taskClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := taskClient.Update(&task.UpdateTaskInput{})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := taskClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})
})
