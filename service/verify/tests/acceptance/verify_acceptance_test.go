package acceptance

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/RJPearson94/twilio-sdk-go"
	messagingServices "github.com/RJPearson94/twilio-sdk-go/service/messaging/v1/services"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/access_tokens"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/entities"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/messaging_configuration"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/messaging_configurations"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/rate_limit"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/rate_limit/bucket"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/rate_limit/buckets"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/rate_limits"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/verification"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/verifications"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/webhook"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/webhooks"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/services"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

var _ = Describe("Verify Acceptance Tests", func() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		Fail(fmt.Sprintf("Failed to create credentials. Error %s", err.Error()))
	}

	messagingSession := twilio.NewWithCredentials(creds).Messaging.V1
	verifySession := twilio.NewWithCredentials(creds).Verify.V2

	Describe("Given the verify service clients", func() {
		It("Then the service is created, fetched, updated and deleted", func() {
			servicesClient := verifySession.Services

			createResp, createErr := servicesClient.Create(&services.CreateServiceInput{
				FriendlyName: "Test Service",
			})
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

			serviceClient := verifySession.Service(createResp.Sid)

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

	Describe("Given the verify rate limit clients", func() {

		var serviceSid string

		BeforeEach(func() {
			resp, err := verifySession.Services.Create(&services.CreateServiceInput{
				FriendlyName: "Test Service",
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create service. Error %s", err.Error()))
			}
			serviceSid = resp.Sid
		})

		AfterEach(func() {
			if err := verifySession.Service(serviceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete service. Error %s", err.Error()))
			}
		})

		It("Then the rate limit is created, fetched, updated and deleted", func() {
			rateLimitsClient := verifySession.Service(serviceSid).RateLimits

			createResp, createErr := rateLimitsClient.Create(&rate_limits.CreateRateLimitInput{
				UniqueName: uuid.New().String(),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := rateLimitsClient.Page(&rate_limits.RateLimitsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.RateLimits)).Should(BeNumerically(">=", 1))

			paginator := rateLimitsClient.NewRateLimitsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.RateLimits)).Should(BeNumerically(">=", 1))

			rateLimitClient := verifySession.Service(serviceSid).RateLimit(createResp.Sid)

			fetchResp, fetchErr := rateLimitClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := rateLimitClient.Update(&rate_limit.UpdateRateLimitInput{})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := rateLimitClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the verify rate limit bucket clients", func() {

		var serviceSid string
		var rateLimitSid string

		BeforeEach(func() {
			resp, err := verifySession.Services.Create(&services.CreateServiceInput{
				FriendlyName: "Test Service",
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create service. Error %s", err.Error()))
			}
			serviceSid = resp.Sid

			rateLimitResp, rateLimitErr := verifySession.Service(serviceSid).RateLimits.Create(&rate_limits.CreateRateLimitInput{
				UniqueName: uuid.New().String(),
			})
			if rateLimitErr != nil {
				Fail(fmt.Sprintf("Failed to create rate limit. Error %s", rateLimitErr.Error()))
			}
			rateLimitSid = rateLimitResp.Sid
		})

		AfterEach(func() {
			if err := verifySession.Service(serviceSid).RateLimit(rateLimitSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete rate limit. Error %s", err.Error()))
			}

			if err := verifySession.Service(serviceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete service. Error %s", err.Error()))
			}
		})

		It("Then the bucket is created, fetched, updated and deleted", func() {
			bucketsClient := verifySession.Service(serviceSid).RateLimit(rateLimitSid).Buckets

			createResp, createErr := bucketsClient.Create(&buckets.CreateBucketInput{
				Max:      4,
				Interval: 10,
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := bucketsClient.Page(&buckets.BucketsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Buckets)).Should(BeNumerically(">=", 1))

			paginator := bucketsClient.NewBucketsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Buckets)).Should(BeNumerically(">=", 1))

			bucketClient := verifySession.Service(serviceSid).RateLimit(rateLimitSid).Bucket(createResp.Sid)

			fetchResp, fetchErr := bucketClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := bucketClient.Update(&bucket.UpdateBucketInput{})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := bucketClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the verify verification clients", func() {

		var serviceSid string

		BeforeEach(func() {
			resp, err := verifySession.Services.Create(&services.CreateServiceInput{
				FriendlyName: "Test Service",
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create service. Error %s", err.Error()))
			}
			serviceSid = resp.Sid
		})

		AfterEach(func() {
			if err := verifySession.Service(serviceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete service. Error %s", err.Error()))
			}
		})

		It("Then the verification is created, fetched and updated", func() {
			verificationsClient := verifySession.Service(serviceSid).Verifications

			createResp, createErr := verificationsClient.Create(&verifications.CreateVerificationInput{
				To:      os.Getenv("DESTINATION_PHONE_NUMBER"),
				Channel: "sms",
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			verificationClient := verifySession.Service(serviceSid).Verification(createResp.Sid)

			fetchResp, fetchErr := verificationClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := verificationClient.Update(&verification.UpdateVerificationInput{
				Status: "canceled",
			})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())
		})
	})

	Describe("Given the verify access token client", func() {

		var serviceSid string

		BeforeEach(func() {
			resp, err := verifySession.Services.Create(&services.CreateServiceInput{
				FriendlyName: "Test Service",
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create service. Error %s", err.Error()))
			}
			serviceSid = resp.Sid
		})

		AfterEach(func() {
			if err := verifySession.Service(serviceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete service. Error %s", err.Error()))
			}
		})

		It("Then the access token is created", func() {
			accessTokensClient := verifySession.Service(serviceSid).AccessTokens

			createResp, createErr := accessTokensClient.Create(&access_tokens.CreateAccessTokenInput{
				Identity:   uuid.New().String(),
				FactorType: "push",
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Token).ToNot(BeNil())
		})
	})

	Describe("Given the verify entities clients", func() {

		var serviceSid string

		BeforeEach(func() {
			resp, err := verifySession.Services.Create(&services.CreateServiceInput{
				FriendlyName: "Test Service",
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create service. Error %s", err.Error()))
			}
			serviceSid = resp.Sid
		})

		AfterEach(func() {
			if err := verifySession.Service(serviceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete service. Error %s", err.Error()))
			}
		})

		It("Then the entity is created, fetched and deleted", func() {
			entitiesClient := verifySession.Service(serviceSid).Entities

			createResp, createErr := entitiesClient.Create(&entities.CreateEntityInput{
				Identity: uuid.New().String(),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := entitiesClient.Page(&entities.EntitiesPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Entities)).Should(BeNumerically(">=", 1))

			paginator := entitiesClient.NewEntitiesPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Entities)).Should(BeNumerically(">=", 1))

			entityClient := verifySession.Service(serviceSid).Entity(createResp.Identity)

			fetchResp, fetchErr := entityClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			deleteErr := entityClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the verify webhook clients", func() {

		var serviceSid string

		BeforeEach(func() {
			resp, err := verifySession.Services.Create(&services.CreateServiceInput{
				FriendlyName: "Test Service",
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create service. Error %s", err.Error()))
			}
			serviceSid = resp.Sid
		})

		AfterEach(func() {
			if err := verifySession.Service(serviceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete service. Error %s", err.Error()))
			}
		})

		It("Then the webhook is created, fetched, updated and deleted", func() {
			webhooksClient := verifySession.Service(serviceSid).Webhooks

			createResp, createErr := webhooksClient.Create(&webhooks.CreateWebhookInput{
				FriendlyName: uuid.New().String(),
				EventTypes:   []string{"*"},
				WebhookURL:   "https://example.com/webhook",
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

			webhookClient := verifySession.Service(serviceSid).Webhook(createResp.Sid)

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

	Describe("Given the verify messaging configuration clients", func() {

		var serviceSid string
		var messagingServiceSid string

		BeforeEach(func() {
			resp, err := verifySession.Services.Create(&services.CreateServiceInput{
				FriendlyName: "Test Service",
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create service. Error %s", err.Error()))
			}
			serviceSid = resp.Sid

			messagingServiceResp, messagingServiceErr := messagingSession.Services.Create(&messagingServices.CreateServiceInput{
				FriendlyName: uuid.New().String(),
			})
			if messagingServiceErr != nil {
				Fail(fmt.Sprintf("Failed to create messaging service. Error %s", messagingServiceErr.Error()))
			}
			messagingServiceSid = messagingServiceResp.Sid
		})

		AfterEach(func() {
			if err := verifySession.Service(serviceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete service. Error %s", err.Error()))
			}
			if err := messagingSession.Service(messagingServiceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete messaging service. Error %s", err.Error()))
			}
		})

		It("Then the messaging configuration is created, fetched, updated and deleted", func() {
			messagingConfigurationsClient := verifySession.Service(serviceSid).MessagingConfigurations

			createResp, createErr := messagingConfigurationsClient.Create(&messaging_configurations.CreateMessagingConfigurationInput{
				Country:             "all",
				MessagingServiceSid: messagingServiceSid,
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Country).ToNot(BeNil())

			pageResp, pageErr := messagingConfigurationsClient.Page(&messaging_configurations.MessagingConfigurationsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.MessagingConfigurations)).Should(BeNumerically(">=", 1))

			paginator := messagingConfigurationsClient.NewMessagingConfigurationsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.MessagingConfigurations)).Should(BeNumerically(">=", 1))

			messagingConfigurationClient := verifySession.Service(serviceSid).MessagingConfiguration(createResp.Country)

			fetchResp, fetchErr := messagingConfigurationClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := messagingConfigurationClient.Update(&messaging_configuration.UpdateMessagingConfigurationInput{
				MessagingServiceSid: messagingServiceSid,
			})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := messagingConfigurationClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	// TODO Add Verification Check tests
	// TODO Add Factor tests
	// TODO Add Challenge tests
})
