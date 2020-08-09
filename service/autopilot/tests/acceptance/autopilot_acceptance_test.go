package acceptance

import (
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/RJPearson94/twilio-sdk-go"
	v1 "github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/defaults"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/field_type"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/field_type/field_values"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/field_types"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/model_build"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/model_builds"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/queries"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/query"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/style_sheet"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/task"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/task/actions"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/task/fields"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/task/sample"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/task/samples"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/tasks"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/webhook"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/webhooks"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistants"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

var _ = Describe("Autopilot Acceptance Tests", func() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		Fail(fmt.Sprintf("Failed to create credentials. Error %s", err.Error()))
	}

	autopilotSession := twilio.NewWithCredentials(creds).Autopilot.V1

	Describe("Given the autopilot assistant clients", func() {
		It("Then the assistant is created, fetched, updated and deleted", func() {
			createResp, createErr := autopilotSession.Assistants.Create(&assistants.CreateAssistantInput{})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := autopilotSession.Assistants.Page(&assistants.AssistantsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Assistants)).Should(BeNumerically(">=", 1))

			paginator := autopilotSession.Assistants.NewAssistantsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Assistants)).Should(BeNumerically(">=", 1))

			assistantClient := autopilotSession.Assistant(createResp.Sid)

			fetchResp, fetchErr := assistantClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := assistantClient.Update(&assistant.UpdateAssistantInput{})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := assistantClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the autopilot defaults client", func() {

		var assistantSid string

		BeforeEach(func() {
			resp, err := autopilotSession.Assistants.Create(&assistants.CreateAssistantInput{})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create assistant. Error %s", err.Error()))
			}
			assistantSid = resp.Sid
		})

		AfterEach(func() {
			if err := autopilotSession.Assistant(assistantSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete assistant. Error %s", err.Error()))
			}
		})

		It("Then the defaults are fetched and updated", func() {
			defaultsClient := autopilotSession.Assistant(assistantSid).Defaults()

			fetchResp, fetchErr := defaultsClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := defaultsClient.Update(&defaults.UpdateDefaultInput{})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())
		})
	})

	Describe("Given the autopilot style sheet client", func() {

		var assistantSid string

		BeforeEach(func() {
			resp, err := autopilotSession.Assistants.Create(&assistants.CreateAssistantInput{})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create assistant. Error %s", err.Error()))
			}
			assistantSid = resp.Sid
		})

		AfterEach(func() {
			if err := autopilotSession.Assistant(assistantSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete assistant. Error %s", err.Error()))
			}
		})

		It("Then the style sheet is fetched and updated", func() {
			stylesheetClient := autopilotSession.Assistant(assistantSid).StyleSheet()

			fetchResp, fetchErr := stylesheetClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := stylesheetClient.Update(&style_sheet.UpdateStyleSheetInput{})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())
		})
	})

	Describe("Given the autopilot task clients", func() {

		var assistantSid string

		BeforeEach(func() {
			resp, err := autopilotSession.Assistants.Create(&assistants.CreateAssistantInput{})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create assistant. Error %s", err.Error()))
			}
			assistantSid = resp.Sid
		})

		AfterEach(func() {
			if err := autopilotSession.Assistant(assistantSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete assistant. Error %s", err.Error()))
			}
		})

		It("Then the task is created, fetched, updated and deleted", func() {
			tasksClient := autopilotSession.Assistant(assistantSid).Tasks

			createResp, createErr := tasksClient.Create(&tasks.CreateTaskInput{
				UniqueName: uuid.New().String(),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := tasksClient.Page(&tasks.TasksPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Tasks)).Should(BeNumerically(">=", 1))

			paginator := tasksClient.NewTasksPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Tasks)).Should(BeNumerically(">=", 1))

			taskClient := autopilotSession.Assistant(assistantSid).Task(createResp.Sid)

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

	Describe("Given the autopilot task actions client", func() {

		var assistantSid string
		var taskSid string

		BeforeEach(func() {
			resp, err := autopilotSession.Assistants.Create(&assistants.CreateAssistantInput{})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create assistant. Error %s", err.Error()))
			}
			assistantSid = resp.Sid

			taskResp, taskErr := autopilotSession.Assistant(assistantSid).Tasks.Create(&tasks.CreateTaskInput{
				UniqueName: uuid.New().String(),
			})
			if taskErr != nil {
				Fail(fmt.Sprintf("Failed to create task. Error %s", taskErr.Error()))
			}
			taskSid = taskResp.Sid
		})

		AfterEach(func() {
			if err := autopilotSession.Assistant(assistantSid).Task(taskSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete task. Error %s", err.Error()))
			}

			if err := autopilotSession.Assistant(assistantSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete assistant. Error %s", err.Error()))
			}
		})

		It("Then the task actions are fetched and updated", func() {
			actionsClient := autopilotSession.Assistant(assistantSid).Task(taskSid).Actions()

			fetchResp, fetchErr := actionsClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := actionsClient.Update(&actions.UpdateActionInput{})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())
		})
	})

	Describe("Given the autopilot task statistics client", func() {

		var assistantSid string
		var taskSid string

		BeforeEach(func() {
			resp, err := autopilotSession.Assistants.Create(&assistants.CreateAssistantInput{})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create assistant. Error %s", err.Error()))
			}
			assistantSid = resp.Sid

			taskResp, taskErr := autopilotSession.Assistant(assistantSid).Tasks.Create(&tasks.CreateTaskInput{
				UniqueName: uuid.New().String(),
			})
			if taskErr != nil {
				Fail(fmt.Sprintf("Failed to create task. Error %s", taskErr.Error()))
			}
			taskSid = taskResp.Sid
		})

		AfterEach(func() {
			if err := autopilotSession.Assistant(assistantSid).Task(taskSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete task. Error %s", err.Error()))
			}

			if err := autopilotSession.Assistant(assistantSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete assistant. Error %s", err.Error()))
			}
		})

		It("Then the task statistics are fetched", func() {
			actionsClient := autopilotSession.Assistant(assistantSid).Task(taskSid).Statistics()

			fetchResp, fetchErr := actionsClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())
		})
	})

	Describe("Given the autopilot task sample clients", func() {

		var assistantSid string
		var taskSid string

		BeforeEach(func() {
			resp, err := autopilotSession.Assistants.Create(&assistants.CreateAssistantInput{})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create assistant. Error %s", err.Error()))
			}
			assistantSid = resp.Sid

			taskResp, taskErr := autopilotSession.Assistant(assistantSid).Tasks.Create(&tasks.CreateTaskInput{
				UniqueName: uuid.New().String(),
			})
			if taskErr != nil {
				Fail(fmt.Sprintf("Failed to create task. Error %s", taskErr.Error()))
			}
			taskSid = taskResp.Sid
		})

		AfterEach(func() {
			if err := autopilotSession.Assistant(assistantSid).Task(taskSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete task. Error %s", err.Error()))
			}

			if err := autopilotSession.Assistant(assistantSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete assistant. Error %s", err.Error()))
			}
		})

		It("Then the task sample is created, fetched, updated and deleted", func() {
			samplesClient := autopilotSession.Assistant(assistantSid).Task(taskSid).Samples

			createResp, createErr := samplesClient.Create(&samples.CreateSampleInput{
				Language:   "en-US",
				TaggedText: "hello world",
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := samplesClient.Page(&samples.SamplesPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Samples)).Should(BeNumerically(">=", 1))

			paginator := samplesClient.NewSamplesPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Samples)).Should(BeNumerically(">=", 1))

			sampleClient := autopilotSession.Assistant(assistantSid).Task(taskSid).Sample(createResp.Sid)

			fetchResp, fetchErr := sampleClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := sampleClient.Update(&sample.UpdateSampleInput{})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := sampleClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the autopilot task field clients", func() {

		var assistantSid string
		var taskSid string

		BeforeEach(func() {
			resp, err := autopilotSession.Assistants.Create(&assistants.CreateAssistantInput{})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create assistant. Error %s", err.Error()))
			}
			assistantSid = resp.Sid

			taskResp, taskErr := autopilotSession.Assistant(assistantSid).Tasks.Create(&tasks.CreateTaskInput{
				UniqueName: uuid.New().String(),
			})
			if taskErr != nil {
				Fail(fmt.Sprintf("Failed to create task. Error %s", taskErr.Error()))
			}
			taskSid = taskResp.Sid
		})

		AfterEach(func() {
			if err := autopilotSession.Assistant(assistantSid).Task(taskSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete task. Error %s", err.Error()))
			}

			if err := autopilotSession.Assistant(assistantSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete assistant. Error %s", err.Error()))
			}
		})

		It("Then the task field is created, fetched and deleted", func() {
			fieldsClient := autopilotSession.Assistant(assistantSid).Task(taskSid).Fields

			createResp, createErr := fieldsClient.Create(&fields.CreateFieldInput{
				UniqueName: uuid.New().String(),
				FieldType:  "Twilio.YES_NO",
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := fieldsClient.Page(&fields.FieldsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Fields)).Should(BeNumerically(">=", 1))

			paginator := fieldsClient.NewFieldsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Fields)).Should(BeNumerically(">=", 1))

			fieldClient := autopilotSession.Assistant(assistantSid).Task(taskSid).Field(createResp.Sid)

			fetchResp, fetchErr := fieldClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			deleteErr := fieldClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the autopilot field type clients", func() {

		var assistantSid string

		BeforeEach(func() {
			resp, err := autopilotSession.Assistants.Create(&assistants.CreateAssistantInput{})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create assistant. Error %s", err.Error()))
			}
			assistantSid = resp.Sid
		})

		AfterEach(func() {
			if err := autopilotSession.Assistant(assistantSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete assistant. Error %s", err.Error()))
			}
		})

		It("Then the field type is created, fetched, updated and deleted", func() {
			fieldTypesClient := autopilotSession.Assistant(assistantSid).FieldTypes

			createResp, createErr := fieldTypesClient.Create(&field_types.CreateFieldTypeInput{
				UniqueName: uuid.New().String(),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := fieldTypesClient.Page(&field_types.FieldTypesPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.FieldTypes)).Should(BeNumerically(">=", 1))

			paginator := fieldTypesClient.NewFieldTypesPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.FieldTypes)).Should(BeNumerically(">=", 1))

			fieldTypeClient := autopilotSession.Assistant(assistantSid).FieldType(createResp.Sid)

			fetchResp, fetchErr := fieldTypeClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := fieldTypeClient.Update(&field_type.UpdateFieldTypeInput{})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := fieldTypeClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the autopilot field value clients", func() {

		var assistantSid string
		var fieldTypeSid string

		BeforeEach(func() {
			resp, err := autopilotSession.Assistants.Create(&assistants.CreateAssistantInput{})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create assistant. Error %s", err.Error()))
			}
			assistantSid = resp.Sid

			fieldTypeResp, fieldTypeErr := autopilotSession.Assistant(assistantSid).FieldTypes.Create(&field_types.CreateFieldTypeInput{
				UniqueName: uuid.New().String(),
			})
			if fieldTypeErr != nil {
				Fail(fmt.Sprintf("Failed to create field type. Error %s", fieldTypeErr.Error()))
			}
			fieldTypeSid = fieldTypeResp.Sid
		})

		AfterEach(func() {
			if err := autopilotSession.Assistant(assistantSid).FieldType(fieldTypeSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete field type. Error %s", err.Error()))
			}

			if err := autopilotSession.Assistant(assistantSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete assistant. Error %s", err.Error()))
			}
		})

		It("Then the field value is created, fetched and deleted", func() {
			fieldValuesClient := autopilotSession.Assistant(assistantSid).FieldType(fieldTypeSid).FieldValues

			createResp, createErr := fieldValuesClient.Create(&field_values.CreateFieldValueInput{
				Language: "en-US",
				Value:    "test",
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := fieldValuesClient.Page(&field_values.FieldValuesPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.FieldValues)).Should(BeNumerically(">=", 1))

			paginator := fieldValuesClient.NewFieldValuesPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.FieldValues)).Should(BeNumerically(">=", 1))

			fieldValueClient := autopilotSession.Assistant(assistantSid).FieldType(fieldTypeSid).FieldValue(createResp.Sid)

			fetchResp, fetchErr := fieldValueClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			deleteErr := fieldValueClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the autopilot model build clients", func() {

		var assistantSid string
		var taskSid string
		var taskSampleSid string

		BeforeEach(func() {
			resp, err := autopilotSession.Assistants.Create(&assistants.CreateAssistantInput{})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create assistant. Error %s", err.Error()))
			}
			assistantSid = resp.Sid

			taskResp, taskErr := autopilotSession.Assistant(assistantSid).Tasks.Create(&tasks.CreateTaskInput{
				UniqueName: uuid.New().String(),
			})
			if taskErr != nil {
				Fail(fmt.Sprintf("Failed to create task. Error %s", taskErr.Error()))
			}
			taskSid = taskResp.Sid

			taskSampleResp, taskSampleErr := autopilotSession.Assistant(assistantSid).Task(taskSid).Samples.Create(&samples.CreateSampleInput{
				Language:   "en-US",
				TaggedText: "hello world",
			})
			if taskSampleErr != nil {
				Fail(fmt.Sprintf("Failed to create task sample. Error %s", taskSampleErr.Error()))
			}
			taskSampleSid = taskSampleResp.Sid
		})

		AfterEach(func() {
			if err := autopilotSession.Assistant(assistantSid).Task(taskSid).Sample(taskSampleSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete task sample. Error %s", err.Error()))
			}

			if err := autopilotSession.Assistant(assistantSid).Task(taskSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete task. Error %s", err.Error()))
			}

			if err := autopilotSession.Assistant(assistantSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete assistant. Error %s", err.Error()))
			}
		})

		It("Then the model build is created, fetched, updated and deleted", func() {
			modelBuildsClient := autopilotSession.Assistant(assistantSid).ModelBuilds

			createResp, createErr := modelBuildsClient.Create(&model_builds.CreateModelBuildInput{})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := modelBuildsClient.Page(&model_builds.ModelBuildsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.ModelBuilds)).Should(BeNumerically(">=", 1))

			paginator := modelBuildsClient.NewModelBuildsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.ModelBuilds)).Should(BeNumerically(">=", 1))

			modelBuildClient := autopilotSession.Assistant(assistantSid).ModelBuild(createResp.Sid)

			fetchResp, fetchErr := modelBuildClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := modelBuildClient.Update(&model_build.UpdateModelBuildInput{})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := modelBuildClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the autopilot model build clients", func() {

		var assistantSid string
		var taskSid string
		var taskSampleSid string
		var modelBuildSid string

		BeforeEach(func() {
			resp, err := autopilotSession.Assistants.Create(&assistants.CreateAssistantInput{})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create assistant. Error %s", err.Error()))
			}
			assistantSid = resp.Sid

			taskResp, taskErr := autopilotSession.Assistant(assistantSid).Tasks.Create(&tasks.CreateTaskInput{
				UniqueName: uuid.New().String(),
			})
			if taskErr != nil {
				Fail(fmt.Sprintf("Failed to create task. Error %s", taskErr.Error()))
			}
			taskSid = taskResp.Sid

			taskSampleResp, taskSampleErr := autopilotSession.Assistant(assistantSid).Task(taskSid).Samples.Create(&samples.CreateSampleInput{
				Language:   "en-US",
				TaggedText: "hello world",
			})
			if taskSampleErr != nil {
				Fail(fmt.Sprintf("Failed to create task sample. Error %s", taskSampleErr.Error()))
			}
			taskSampleSid = taskSampleResp.Sid

			modelBuildResp, modelBuildErr := autopilotSession.Assistant(assistantSid).ModelBuilds.Create(&model_builds.CreateModelBuildInput{})
			if modelBuildErr != nil {
				Fail(fmt.Sprintf("Failed to create model build. Error %s", modelBuildErr.Error()))
			}
			modelBuildSid = modelBuildResp.Sid

			// The model build needs to be complete before it can be queried
			pollErr := poll(30, 1000, autopilotSession, assistantSid, modelBuildSid)
			if pollErr != nil {
				Fail(pollErr.Error())
			}
		})

		AfterEach(func() {
			if err := autopilotSession.Assistant(assistantSid).ModelBuild(modelBuildSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete model build. Error %s", err.Error()))
			}

			if err := autopilotSession.Assistant(assistantSid).Task(taskSid).Sample(taskSampleSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete task sample. Error %s", err.Error()))
			}

			if err := autopilotSession.Assistant(assistantSid).Task(taskSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete task. Error %s", err.Error()))
			}

			if err := autopilotSession.Assistant(assistantSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete assistant. Error %s", err.Error()))
			}
		})

		It("Then the query is created, fetched, updated and deleted", func() {
			queriesClient := autopilotSession.Assistant(assistantSid).Queries

			createResp, createErr := queriesClient.Create(&queries.CreateQueryInput{
				Language: "en-US",
				Query:    "hello world",
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := queriesClient.Page(&queries.QueriesPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Queries)).Should(BeNumerically(">=", 1))

			paginator := queriesClient.NewQueriesPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Queries)).Should(BeNumerically(">=", 1))

			queryClient := autopilotSession.Assistant(assistantSid).Query(createResp.Sid)

			fetchResp, fetchErr := queryClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := queryClient.Update(&query.UpdateQueryInput{})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := queryClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the autopilot webhook clients", func() {

		var assistantSid string

		BeforeEach(func() {
			resp, err := autopilotSession.Assistants.Create(&assistants.CreateAssistantInput{})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create assistant. Error %s", err.Error()))
			}
			assistantSid = resp.Sid
		})

		AfterEach(func() {
			if err := autopilotSession.Assistant(assistantSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete assistant. Error %s", err.Error()))
			}
		})

		It("Then the webhook is created, fetched, updated and deleted", func() {
			webhooksClient := autopilotSession.Assistant(assistantSid).Webhooks

			createResp, createErr := webhooksClient.Create(&webhooks.CreateWebhookInput{
				UniqueName: uuid.New().String(),
				Events:     "onDialogueEnd",
				WebhookURL: "https://localhost/webhook",
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := webhooksClient.Page(&webhooks.WebhooksPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Webhooks)).Should(BeNumerically(">=", 1))

			paginator := webhooksClient.NewWebhooksPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Webhooks)).Should(BeNumerically(">=", 1))

			webhookClient := autopilotSession.Assistant(assistantSid).Webhook(createResp.Sid)

			fetchResp, fetchErr := webhookClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := webhookClient.Update(&webhook.UpdateWebhookInput{})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := webhookClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})
})

func poll(maxAttempts int, delay int, client *v1.Autopilot, assistantSid string, modelBuildSid string) error {
	for i := 0; i < maxAttempts; i++ {
		resp, err := client.Assistant(assistantSid).ModelBuild(modelBuildSid).Fetch()
		if err != nil {
			return fmt.Errorf("Failed to poll autopilot model build: %s", err)
		}

		if resp.Status == "failed" {
			return fmt.Errorf("Autopilot model build failed")
		}
		if resp.Status == "completed" {
			return nil
		}
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}
	return fmt.Errorf("Reached max polling attempts without a completed model build")
}
