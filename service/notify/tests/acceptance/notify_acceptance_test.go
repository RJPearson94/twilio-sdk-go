package acceptance

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/RJPearson94/twilio-sdk-go"
	"github.com/RJPearson94/twilio-sdk-go/service/notify/v1/credential"
	"github.com/RJPearson94/twilio-sdk-go/service/notify/v1/credentials"
	"github.com/RJPearson94/twilio-sdk-go/service/notify/v1/service"
	"github.com/RJPearson94/twilio-sdk-go/service/notify/v1/service/bindings"
	"github.com/RJPearson94/twilio-sdk-go/service/notify/v1/service/notifications"
	"github.com/RJPearson94/twilio-sdk-go/service/notify/v1/services"
	sessionCredentials "github.com/RJPearson94/twilio-sdk-go/session/credentials"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

var _ = Describe("Notify Acceptance Tests", func() {
	creds, err := sessionCredentials.New(sessionCredentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		Fail(fmt.Sprintf("Failed to create credentials. Error %s", err.Error()))
	}

	notifySession := twilio.NewWithCredentials(creds).Notify.V1

	Describe("Given the notify service clients", func() {
		It("Then the service is created, fetched and deleted", func() {
			servicesClient := notifySession.Services

			createResp, createErr := servicesClient.Create(&services.CreateServiceInput{})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := servicesClient.Page(&services.ServicesPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Services)).Should(BeNumerically(">=", 1))

			paginator := servicesClient.NewServicesPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Services)).Should(BeNumerically(">=", 1))

			serviceClient := notifySession.Service(createResp.Sid)

			fetchResp, fetchErr := serviceClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := serviceClient.Update(&service.UpdateServiceInput{})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := serviceClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the notify credential clients", func() {
		It("Then the credential is created, fetched, updated and deleted", func() {
			credentialsClient := notifySession.Credentials

			createResp, createErr := credentialsClient.Create(&credentials.CreateCredentialInput{
				Type:   "fcm",
				Secret: utils.String(uuid.New().String()),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := credentialsClient.Page(&credentials.CredentialsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Credentials)).Should(BeNumerically(">=", 1))

			paginator := credentialsClient.NewCredentialsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Credentials)).Should(BeNumerically(">=", 1))

			credentialClient := notifySession.Credential(createResp.Sid)

			fetchResp, fetchErr := credentialClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := credentialClient.Update(&credential.UpdateCredentialInput{
				FriendlyName: utils.String(uuid.New().String()),
			})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := credentialClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the notify binding clients", func() {

		var serviceSid string

		BeforeEach(func() {
			resp, err := notifySession.Services.Create(&services.CreateServiceInput{})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create service. Error %s", err.Error()))
			}
			serviceSid = resp.Sid
		})

		AfterEach(func() {
			if err := notifySession.Service(serviceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete service. Error %s", err.Error()))
			}
		})

		It("Then the binding is created, fetched and deleted", func() {
			bindingsClient := notifySession.Service(serviceSid).Bindings

			createResp, createErr := bindingsClient.Create(&bindings.CreateBindingInput{
				Identity:    uuid.New().String(),
				BindingType: "sms",
				Address:     os.Getenv("DESTINATION_PHONE_NUMBER"),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := bindingsClient.Page(&bindings.BindingsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Bindings)).Should(BeNumerically(">=", 1))

			paginator := bindingsClient.NewBindingsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Bindings)).Should(BeNumerically(">=", 1))

			bindingClient := notifySession.Service(serviceSid).Binding(createResp.Sid)

			fetchResp, fetchErr := bindingClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			deleteErr := bindingClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the notify notification client", func() {

		var serviceSid string

		BeforeEach(func() {
			resp, err := notifySession.Services.Create(&services.CreateServiceInput{})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create service. Error %s", err.Error()))
			}
			serviceSid = resp.Sid
		})

		AfterEach(func() {
			if err := notifySession.Service(serviceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete service. Error %s", err.Error()))
			}
		})

		It("Then the notification is created", func() {
			notificationsClient := notifySession.Service(serviceSid).Notifications

			createResp, createErr := notificationsClient.Create(&notifications.CreateNotificationInput{
				ToBindings: &[]string{
					os.ExpandEnv("{\"binding_type\":\"sms\", \"address\":\"${DESTINATION_PHONE_NUMBER}\"}"),
				},
				Body: utils.String("Hello World"),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())
		})
	})
})
