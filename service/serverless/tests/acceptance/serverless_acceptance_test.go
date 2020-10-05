package acceptance

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/RJPearson94/twilio-sdk-go"
	v1 "github.com/RJPearson94/twilio-sdk-go/service/serverless/v1"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service/asset"
	assetVersions "github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service/asset/versions"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service/assets"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service/builds"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service/environment/deployments"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service/environment/variable"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service/environment/variables"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service/environments"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service/function"
	functionVersions "github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service/function/versions"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service/functions"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/services"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

var _ = Describe("Serverless Acceptance Tests", func() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		Fail(fmt.Sprintf("Failed to create credentials. Error %s", err.Error()))
	}

	serverlessSession := twilio.NewWithCredentials(creds).Serverless.V1

	Describe("Given the serverless service clients", func() {
		It("Then the service is created, fetched, updated and deleted", func() {
			servicesClient := serverlessSession.Services

			createResp, createErr := servicesClient.Create(&services.CreateServiceInput{
				FriendlyName: uuid.New().String(),
				UniqueName:   uuid.New().String(),
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

			serviceClient := serverlessSession.Service(createResp.Sid)

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

	Describe("Given the serverless environment clients", func() {

		var serviceSid string

		BeforeEach(func() {
			resp, err := serverlessSession.Services.Create(&services.CreateServiceInput{
				FriendlyName: uuid.New().String(),
				UniqueName:   uuid.New().String(),
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create service. Error %s", err.Error()))
			}
			serviceSid = resp.Sid
		})

		AfterEach(func() {
			if err := serverlessSession.Service(serviceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete service. Error %s", err.Error()))
			}
		})

		It("Then the environment is created, fetched and deleted", func() {
			environmentsClient := serverlessSession.Service(serviceSid).Environments

			createResp, createErr := environmentsClient.Create(&environments.CreateEnvironmentInput{
				UniqueName: uuid.New().String(),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := environmentsClient.Page(&environments.EnvironmentsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Environments)).Should(BeNumerically(">=", 1))

			paginator := environmentsClient.NewEnvironmentsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Environments)).Should(BeNumerically(">=", 1))

			environmentClient := serverlessSession.Service(serviceSid).Environment(createResp.Sid)

			fetchResp, fetchErr := environmentClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			deleteErr := environmentClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the serverless environment variable clients", func() {

		var serviceSid string
		var environmentSid string

		BeforeEach(func() {
			resp, err := serverlessSession.Services.Create(&services.CreateServiceInput{
				FriendlyName: uuid.New().String(),
				UniqueName:   uuid.New().String(),
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create service. Error %s", err.Error()))
			}
			serviceSid = resp.Sid

			environmentResp, environmentErr := serverlessSession.Service(serviceSid).Environments.Create(&environments.CreateEnvironmentInput{
				UniqueName: uuid.New().String(),
			})
			if environmentErr != nil {
				Fail(fmt.Sprintf("Failed to create environment. Error %s", environmentErr.Error()))
			}
			environmentSid = environmentResp.Sid
		})

		AfterEach(func() {
			if err := serverlessSession.Service(serviceSid).Environment(environmentSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete environment. Error %s", err.Error()))
			}

			if err := serverlessSession.Service(serviceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete service. Error %s", err.Error()))
			}
		})

		It("Then the environment variable is created, fetched, updated and deleted", func() {
			variablesClient := serverlessSession.Service(serviceSid).Environment(environmentSid).Variables

			createResp, createErr := variablesClient.Create(&variables.CreateVariableInput{
				Key:   "key",
				Value: "value",
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := variablesClient.Page(&variables.VariablesPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Variables)).Should(BeNumerically(">=", 1))

			paginator := variablesClient.NewVariablesPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Variables)).Should(BeNumerically(">=", 1))

			variableClient := serverlessSession.Service(serviceSid).Environment(environmentSid).Variable(createResp.Sid)

			fetchResp, fetchErr := variableClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := variableClient.Update(&variable.UpdateVariableInput{})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := variableClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the serverless function clients", func() {

		var serviceSid string

		BeforeEach(func() {
			resp, err := serverlessSession.Services.Create(&services.CreateServiceInput{
				FriendlyName: uuid.New().String(),
				UniqueName:   uuid.New().String(),
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create service. Error %s", err.Error()))
			}
			serviceSid = resp.Sid
		})

		AfterEach(func() {
			if err := serverlessSession.Service(serviceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete service. Error %s", err.Error()))
			}
		})

		It("Then the function is created, fetched, updated and deleted", func() {
			functionsClient := serverlessSession.Service(serviceSid).Functions

			createResp, createErr := functionsClient.Create(&functions.CreateFunctionInput{
				FriendlyName: uuid.New().String(),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := functionsClient.Page(&functions.FunctionsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Functions)).Should(BeNumerically(">=", 1))

			paginator := functionsClient.NewFunctionsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Functions)).Should(BeNumerically(">=", 1))

			functionClient := serverlessSession.Service(serviceSid).Function(createResp.Sid)

			fetchResp, fetchErr := functionClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := functionClient.Update(&function.UpdateFunctionInput{
				FriendlyName: uuid.New().String(),
			})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := functionClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the serverless function version clients", func() {

		var serviceSid string
		var functionSid string

		BeforeEach(func() {
			resp, err := serverlessSession.Services.Create(&services.CreateServiceInput{
				FriendlyName: uuid.New().String(),
				UniqueName:   uuid.New().String(),
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create service. Error %s", err.Error()))
			}
			serviceSid = resp.Sid

			functionResp, functionErr := serverlessSession.Service(serviceSid).Functions.Create(&functions.CreateFunctionInput{
				FriendlyName: uuid.New().String(),
			})
			if functionErr != nil {
				Fail(fmt.Sprintf("Failed to create function. Error %s", functionErr.Error()))
			}
			functionSid = functionResp.Sid
		})

		AfterEach(func() {
			if err := serverlessSession.Service(serviceSid).Function(functionSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete function. Error %s", err.Error()))
			}

			if err := serverlessSession.Service(serviceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete service. Error %s", err.Error()))
			}
		})

		It("Then the function version is created and fetched", func() {
			functionVersionsClient := serverlessSession.Service(serviceSid).Function(functionSid).Versions

			createResp, createErr := createFunctionVersion(serverlessSession, serviceSid, functionSid)
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := functionVersionsClient.Page(&functionVersions.VersionsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Versions)).Should(BeNumerically(">=", 1))

			paginator := functionVersionsClient.NewVersionsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Versions)).Should(BeNumerically(">=", 1))

			functionVersionClient := serverlessSession.Service(serviceSid).Function(functionSid).Version(createResp.Sid)

			fetchResp, fetchErr := functionVersionClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())
		})
	})

	Describe("Given the serverless function version content client", func() {

		var serviceSid string
		var functionSid string
		var versionSid string

		BeforeEach(func() {
			resp, err := serverlessSession.Services.Create(&services.CreateServiceInput{
				FriendlyName: uuid.New().String(),
				UniqueName:   uuid.New().String(),
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create service. Error %s", err.Error()))
			}
			serviceSid = resp.Sid

			functionResp, functionErr := serverlessSession.Service(serviceSid).Functions.Create(&functions.CreateFunctionInput{
				FriendlyName: uuid.New().String(),
			})
			if functionErr != nil {
				Fail(fmt.Sprintf("Failed to create function. Error %s", functionErr.Error()))
			}
			functionSid = functionResp.Sid

			versionResp, versionErr := createFunctionVersion(serverlessSession, serviceSid, functionSid)
			if versionErr != nil {
				Fail(fmt.Sprintf("Failed to create function version. Error %s", versionErr.Error()))
			}
			versionSid = versionResp.Sid
		})

		AfterEach(func() {
			if err := serverlessSession.Service(serviceSid).Function(functionSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete function. Error %s", err.Error()))
			}

			if err := serverlessSession.Service(serviceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete service. Error %s", err.Error()))
			}
		})

		It("Then the function version content is fetched", func() {
			functionVersionContentClient := serverlessSession.Service(serviceSid).Function(functionSid).Version(versionSid).Content()

			fetchResp, fetchErr := functionVersionContentClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())
		})
	})

	Describe("Given the serverless asset clients", func() {

		var serviceSid string

		BeforeEach(func() {
			resp, err := serverlessSession.Services.Create(&services.CreateServiceInput{
				FriendlyName: uuid.New().String(),
				UniqueName:   uuid.New().String(),
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create service. Error %s", err.Error()))
			}
			serviceSid = resp.Sid
		})

		AfterEach(func() {
			if err := serverlessSession.Service(serviceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete service. Error %s", err.Error()))
			}
		})

		It("Then the asset is created, fetched, updated and deleted", func() {
			assetsClient := serverlessSession.Service(serviceSid).Assets

			createResp, createErr := assetsClient.Create(&assets.CreateAssetInput{
				FriendlyName: uuid.New().String(),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := assetsClient.Page(&assets.AssetsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Assets)).Should(BeNumerically(">=", 1))

			paginator := assetsClient.NewAssetsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Assets)).Should(BeNumerically(">=", 1))

			assetClient := serverlessSession.Service(serviceSid).Asset(createResp.Sid)

			fetchResp, fetchErr := assetClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := assetClient.Update(&asset.UpdateAssetInput{
				FriendlyName: uuid.New().String(),
			})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := assetClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the serverless asset version clients", func() {

		var serviceSid string
		var assetSid string

		BeforeEach(func() {
			resp, err := serverlessSession.Services.Create(&services.CreateServiceInput{
				FriendlyName: uuid.New().String(),
				UniqueName:   uuid.New().String(),
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create service. Error %s", err.Error()))
			}
			serviceSid = resp.Sid

			assetResp, assetErr := serverlessSession.Service(serviceSid).Assets.Create(&assets.CreateAssetInput{
				FriendlyName: uuid.New().String(),
			})
			if assetErr != nil {
				Fail(fmt.Sprintf("Failed to create asset. Error %s", assetErr.Error()))
			}
			assetSid = assetResp.Sid
		})

		AfterEach(func() {
			if err := serverlessSession.Service(serviceSid).Asset(assetSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete asset. Error %s", err.Error()))
			}

			if err := serverlessSession.Service(serviceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete service. Error %s", err.Error()))
			}
		})

		It("Then the asset version is created and fetched", func() {
			assetVersionsClient := serverlessSession.Service(serviceSid).Asset(assetSid).Versions

			createResp, createErr := createAssetVersion(serverlessSession, serviceSid, assetSid)
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := assetVersionsClient.Page(&assetVersions.VersionsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Versions)).Should(BeNumerically(">=", 1))

			paginator := assetVersionsClient.NewVersionsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Versions)).Should(BeNumerically(">=", 1))

			assetVersionClient := serverlessSession.Service(serviceSid).Asset(assetSid).Version(createResp.Sid)

			fetchResp, fetchErr := assetVersionClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())
		})
	})

	Describe("Given the serverless build clients", func() {

		var serviceSid string
		var assetSid string
		var assetVersionSid string

		BeforeEach(func() {
			resp, err := serverlessSession.Services.Create(&services.CreateServiceInput{
				FriendlyName: uuid.New().String(),
				UniqueName:   uuid.New().String(),
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create service. Error %s", err.Error()))
			}
			serviceSid = resp.Sid

			assetResp, assetErr := serverlessSession.Service(serviceSid).Assets.Create(&assets.CreateAssetInput{
				FriendlyName: uuid.New().String(),
			})
			if assetErr != nil {
				Fail(fmt.Sprintf("Failed to create asset. Error %s", assetErr.Error()))
			}
			assetSid = assetResp.Sid

			assetVersionResp, assetVersionErr := createAssetVersion(serverlessSession, serviceSid, assetSid)
			if assetVersionErr != nil {
				Fail(fmt.Sprintf("Failed to create asset version. Error %s", assetVersionErr.Error()))
			}
			assetVersionSid = assetVersionResp.Sid
		})

		AfterEach(func() {
			if err := serverlessSession.Service(serviceSid).Asset(assetSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete asset. Error %s", err.Error()))
			}

			if err := serverlessSession.Service(serviceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete service. Error %s", err.Error()))
			}
		})

		It("Then the build is created, fetched and deleted", func() {
			buildsClient := serverlessSession.Service(serviceSid).Builds

			createResp, createErr := buildsClient.Create(&builds.CreateBuildInput{
				AssetVersions: &[]string{assetVersionSid},
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := buildsClient.Page(&builds.BuildsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Builds)).Should(BeNumerically(">=", 1))

			paginator := buildsClient.NewBuildsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Builds)).Should(BeNumerically(">=", 1))

			buildClient := serverlessSession.Service(serviceSid).Build(createResp.Sid)

			fetchResp, fetchErr := buildClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			deleteErr := buildClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the serverless deployments clients", func() {

		var serviceSid string
		var environmentSid string
		var assetSid string
		var buildSid string

		BeforeEach(func() {
			resp, err := serverlessSession.Services.Create(&services.CreateServiceInput{
				FriendlyName: uuid.New().String(),
				UniqueName:   uuid.New().String(),
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create service. Error %s", err.Error()))
			}
			serviceSid = resp.Sid

			environmentResp, environmentErr := serverlessSession.Service(serviceSid).Environments.Create(&environments.CreateEnvironmentInput{
				UniqueName: uuid.New().String(),
			})
			if environmentErr != nil {
				Fail(fmt.Sprintf("Failed to create environment. Error %s", environmentErr.Error()))
			}
			environmentSid = environmentResp.Sid

			assetResp, assetErr := serverlessSession.Service(serviceSid).Assets.Create(&assets.CreateAssetInput{
				FriendlyName: uuid.New().String(),
			})
			if assetErr != nil {
				Fail(fmt.Sprintf("Failed to create asset. Error %s", assetErr.Error()))
			}
			assetSid = assetResp.Sid

			assetVersionResp, assetVersionErr := createAssetVersion(serverlessSession, serviceSid, assetSid)
			if assetVersionErr != nil {
				Fail(fmt.Sprintf("Failed to create asset version. Error %s", assetVersionErr.Error()))
			}

			buildResp, buildErr := serverlessSession.Service(serviceSid).Builds.Create(&builds.CreateBuildInput{
				AssetVersions: &[]string{assetVersionResp.Sid},
			})
			if buildErr != nil {
				Fail(fmt.Sprintf("Failed to create build. Error %s", buildErr.Error()))
			}
			buildSid = buildResp.Sid

			// The build needs to be complete before it can be used in a deployment
			pollErr := poll(30, 1000, serverlessSession, serviceSid, buildSid)
			if pollErr != nil {
				Fail(pollErr.Error())
			}
		})

		AfterEach(func() {
			if err := serverlessSession.Service(serviceSid).Build(buildSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete build. Error %s", err.Error()))
			}

			if err := serverlessSession.Service(serviceSid).Asset(assetSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete asset. Error %s", err.Error()))
			}

			if err := serverlessSession.Service(serviceSid).Environment(environmentSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete environment. Error %s", err.Error()))
			}

			if err := serverlessSession.Service(serviceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete service. Error %s", err.Error()))
			}
		})

		It("Then the deployment is created, fetched and redeployed", func() {
			deploymentsClient := serverlessSession.Service(serviceSid).Environment(environmentSid).Deployments

			createResp, createErr := deploymentsClient.Create(&deployments.CreateDeploymentInput{
				BuildSid: utils.String(buildSid),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := deploymentsClient.Page(&deployments.DeploymentsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Deployments)).Should(BeNumerically(">=", 1))

			paginator := deploymentsClient.NewDeploymentsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Deployments)).Should(BeNumerically(">=", 1))

			deploymentClient := serverlessSession.Service(serviceSid).Environment(environmentSid).Deployment(createResp.Sid)

			fetchResp, fetchErr := deploymentClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			// Redploy to remove the current build
			redeployResp, redeployErr := serverlessSession.Service(serviceSid).Environment(environmentSid).Deployments.Create(&deployments.CreateDeploymentInput{})
			Expect(redeployErr).To(BeNil())
			Expect(redeployResp).ToNot(BeNil())
			Expect(redeployResp.Sid).ToNot(BeNil())
		})
	})

	// TODO Log tests
})

func createAssetVersion(client *v1.Serverless, serviceSid string, assetSid string) (*assetVersions.CreateVersionResponse, error) {
	return client.Service(serviceSid).Asset(assetSid).Versions.Create(&assetVersions.CreateVersionInput{
		Content: assetVersions.CreateContentDetails{
			Body:        strings.NewReader("{}"),
			ContentType: "application/json",
			FileName:    "test.json",
		},
		Path:       "/test",
		Visibility: "private",
	})
}

func createFunctionVersion(client *v1.Serverless, serviceSid string, functionSid string) (*functionVersions.CreateVersionResponse, error) {
	return client.Service(serviceSid).Function(functionSid).Versions.Create(&functionVersions.CreateVersionInput{
		Content: functionVersions.CreateContentDetails{
			Body:        strings.NewReader(`exports.handler = function (context, event, callback) { callback(null, "Hello World"); };`),
			ContentType: "application/javascript",
			FileName:    "test.js",
		},
		Path:       "/test",
		Visibility: "private",
	})
}

func poll(maxAttempts int, delay int, client *v1.Serverless, serviceSid string, buildSid string) error {
	for i := 0; i < maxAttempts; i++ {
		resp, err := client.Service(serviceSid).Build(buildSid).Status().Fetch()
		if err != nil {
			return fmt.Errorf("Failed to poll serverless build: %s", err)
		}

		if resp.Status == "failed" {
			return fmt.Errorf("Serverless build failed")
		}
		if resp.Status == "completed" {
			return nil
		}
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}
	return fmt.Errorf("Reached max polling attempts without a completed build")
}
