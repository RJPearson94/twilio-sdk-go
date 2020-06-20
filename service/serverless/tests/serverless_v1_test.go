package tests

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/jarcoal/httpmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/RJPearson94/twilio-sdk-go/service/serverless"
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

var _ = Describe("Serverless V1", func() {
	creds, err := credentials.New(credentials.Account{
		Sid:       "ACxxx",
		AuthToken: "Test",
	})
	if err != nil {
		log.Panicf("%s", err)
	}

	serverlessSession := serverless.NewWithCredentials(creds).V1

	httpmock.ActivateNonDefault(serverlessSession.GetClient().GetRestyClient().GetClient())
	defer httpmock.DeactivateAndReset()

	Describe("Given the services client", func() {
		servicesClient := serverlessSession.Services

		Describe("When the service is successfully created", func() {
			createInput := &services.CreateServiceInput{
				FriendlyName: "Test 2",
				UniqueName:   "Unique Test 2",
			}

			httpmock.RegisterResponder("POST", "https://serverless.twilio.com/v1/Services",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/serviceResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := servicesClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create service response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test 2"))
				Expect(resp.UniqueName).To(Equal("Unique Test 2"))
				Expect(resp.IncludeCredentials).To(Equal(true))
				Expect(resp.UiEditable).To(Equal(false))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2018-11-10T20:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the service does not contain a friendly name", func() {
			createInput := &services.CreateServiceInput{
				UniqueName: "Unique Test 2",
			}

			resp, err := servicesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create service response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the service does not contain a unique name", func() {
			createInput := &services.CreateServiceInput{
				FriendlyName: "Test 2",
			}

			resp, err := servicesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create service response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create service api returns a 500 response", func() {
			createInput := &services.CreateServiceInput{
				FriendlyName: "Test 2",
				UniqueName:   "Unique Test 2",
			}

			httpmock.RegisterResponder("POST", "https://serverless.twilio.com/v1/Services",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := servicesClient.Create(createInput)
			It("Then an error should be returned", func() {
				Expect(err).ToNot(BeNil())
				twilioErr, ok := err.(*utils.TwilioError)
				Expect(ok).To(Equal(true))
				Expect(twilioErr.Code).To(BeNil())
				Expect(twilioErr.Message).To(Equal("An error occurred"))
				Expect(twilioErr.MoreInfo).To(BeNil())
				Expect(twilioErr.Status).To(Equal(500))
			})

			It("Then the create service response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a service sid", func() {
		serviceClient := serverlessSession.Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the service is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/serviceResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := serviceClient.Get()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get service response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test 2"))
				Expect(resp.UniqueName).To(Equal("Unique Test 2"))
				Expect(resp.IncludeCredentials).To(Equal(true))
				Expect(resp.UiEditable).To(Equal(false))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2018-11-10T20:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the get service response returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZS71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := serverlessSession.Service("ZS71").Get()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get service response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the service is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateServiceResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &service.UpdateServiceInput{}

			resp, err := serviceClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update service response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test 2"))
				Expect(resp.UniqueName).To(Equal("Unique Test 2"))
				Expect(resp.IncludeCredentials).To(Equal(true))
				Expect(resp.UiEditable).To(Equal(false))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2018-11-10T20:00:00Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2018-11-11T20:00:00Z"))
				Expect(resp.URL).To(Equal("https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update service response returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://serverless.twilio.com/v1/Services/ZS71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &service.UpdateServiceInput{
				FriendlyName: utils.String("New Friendly Name"),
			}

			resp, err := serverlessSession.Service("ZS71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update service response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the service is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := serviceClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the delete service response returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://serverless.twilio.com/v1/Services/ZS71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := serverlessSession.Service("ZS71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the environments client", func() {
		environmentsClient := serverlessSession.Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Environments

		Describe("When the environment is successfully created", func() {
			createInput := &environments.CreateEnvironmentInput{
				UniqueName: "test-2",
			}

			httpmock.RegisterResponder("POST", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/environmentResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := environmentsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create environment response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.UniqueName).To(Equal("test-2"))
				Expect(resp.DomainSuffix).To(BeNil())
				Expect(resp.DomainName).To(Equal("test-2.twil.io"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2018-11-10T20:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the environment does not contain a unique name", func() {
			createInput := &environments.CreateEnvironmentInput{}

			resp, err := environmentsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create environment response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create environment api returns a 500 response", func() {
			createInput := &environments.CreateEnvironmentInput{
				UniqueName: "Test 2",
			}

			httpmock.RegisterResponder("POST", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := environmentsClient.Create(createInput)
			It("Then an error should be returned", func() {
				Expect(err).ToNot(BeNil())
				twilioErr, ok := err.(*utils.TwilioError)
				Expect(ok).To(Equal(true))
				Expect(twilioErr.Code).To(BeNil())
				Expect(twilioErr.Message).To(Equal("An error occurred"))
				Expect(twilioErr.MoreInfo).To(BeNil())
				Expect(twilioErr.Status).To(Equal(500))
			})

			It("Then the create environment response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a environment sid", func() {
		environmentClient := serverlessSession.Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Environment("ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the service is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/environmentResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := environmentClient.Get()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get environment response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.UniqueName).To(Equal("test-2"))
				Expect(resp.DomainSuffix).To(BeNil())
				Expect(resp.DomainName).To(Equal("test-2.twil.io"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2018-11-10T20:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the get environment response returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZE71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := serverlessSession.Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Environment("ZE71").Get()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get environment response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the environment is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := environmentClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the delete environment response returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZE71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := serverlessSession.Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Environment("ZE71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the variable client", func() {
		variablesClient := serverlessSession.Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Environment("ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Variables

		Describe("When the variable is successfully created", func() {
			createInput := &variables.CreateVariableInput{
				Key:   "key",
				Value: "value",
			}

			httpmock.RegisterResponder("POST", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Variables",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/variableResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := variablesClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create variable response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("ZVXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.EnvironmentSid).To(Equal("ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Key).To(Equal("key"))
				Expect(resp.Value).To(Equal("value"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2018-11-10T20:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Variables/ZVXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the variable does not contain a value", func() {
			createInput := &variables.CreateVariableInput{
				Key: "key",
			}

			resp, err := variablesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create variable response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the variable does not contain a key", func() {
			createInput := &variables.CreateVariableInput{
				Value: "value",
			}

			resp, err := variablesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create variable response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create variable api returns a 500 response", func() {
			createInput := &variables.CreateVariableInput{
				Key:   "key",
				Value: "value",
			}

			httpmock.RegisterResponder("POST", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Variables",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := variablesClient.Create(createInput)
			It("Then an error should be returned", func() {
				Expect(err).ToNot(BeNil())
				twilioErr, ok := err.(*utils.TwilioError)
				Expect(ok).To(Equal(true))
				Expect(twilioErr.Code).To(BeNil())
				Expect(twilioErr.Message).To(Equal("An error occurred"))
				Expect(twilioErr.MoreInfo).To(BeNil())
				Expect(twilioErr.Status).To(Equal(500))
			})

			It("Then the create variable response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a variable sid", func() {
		variableClient := serverlessSession.Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Environment("ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Variable("ZVXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the variable is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Variables/ZVXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/variableResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := variableClient.Get()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get variable response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("ZVXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.EnvironmentSid).To(Equal("ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Key).To(Equal("key"))
				Expect(resp.Value).To(Equal("value"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2018-11-10T20:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Variables/ZVXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the get variable response returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Variables/ZV71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := serverlessSession.Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Environment("ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Variable("ZV71").Get()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get variable response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the variable is successfully updated", func() {
			updateInput := &variable.UpdateVariableInput{
				Key:   utils.String("key"),
				Value: utils.String("value"),
			}

			httpmock.RegisterResponder("POST", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Variables/ZVXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateVariableResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := variableClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get variable response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("ZVXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.EnvironmentSid).To(Equal("ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Key).To(Equal("key"))
				Expect(resp.Value).To(Equal("value"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2018-11-10T20:00:00Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2018-11-11T20:00:00Z"))
				Expect(resp.URL).To(Equal("https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Variables/ZVXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update variable response returns a 404", func() {
			updateInput := &variable.UpdateVariableInput{
				Key:   utils.String("Key"),
				Value: utils.String("Value"),
			}

			httpmock.RegisterResponder("POST", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Variables/ZV71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := serverlessSession.Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Environment("ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Variable("ZV71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get variable response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the variable is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Variables/ZVXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := variableClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the delete variable response returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Variables/ZV71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := serverlessSession.Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Environment("ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Variable("ZV71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the deployments client", func() {
		deploymentsClient := serverlessSession.Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Environment("ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Deployments

		Describe("When the deployment is successfully created", func() {
			createInput := &deployments.CreateDeploymentInput{}

			httpmock.RegisterResponder("POST", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Deployments",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/deploymentResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := deploymentsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create deployment response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("ZDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.EnvironmentSid).To(Equal("ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.BuildSid).To(Equal("ZBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2018-11-10T20:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Deployments/ZDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the create deployment api returns a 500 response", func() {
			createInput := &deployments.CreateDeploymentInput{}

			httpmock.RegisterResponder("POST", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Deployments",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := deploymentsClient.Create(createInput)
			It("Then an error should be returned", func() {
				Expect(err).ToNot(BeNil())
				twilioErr, ok := err.(*utils.TwilioError)
				Expect(ok).To(Equal(true))
				Expect(twilioErr.Code).To(BeNil())
				Expect(twilioErr.Message).To(Equal("An error occurred"))
				Expect(twilioErr.MoreInfo).To(BeNil())
				Expect(twilioErr.Status).To(Equal(500))
			})

			It("Then the create deployment response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a deployment sid", func() {
		deploymentClient := serverlessSession.Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Environment("ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Deployment("ZDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the variable is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Deployments/ZDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/deploymentResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := deploymentClient.Get()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get deployment response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("ZDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.EnvironmentSid).To(Equal("ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.BuildSid).To(Equal("ZBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2018-11-10T20:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Deployments/ZDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the get deployment response returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Deployments/ZD71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := serverlessSession.Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Environment("ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Deployment("ZD71").Get()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get deployment response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given the functions client", func() {
		functionsClient := serverlessSession.Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Functions

		Describe("When the function is successfully created", func() {
			createInput := &functions.CreateFunctionInput{
				FriendlyName: "Test 2",
			}

			httpmock.RegisterResponder("POST", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Functions",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/functionResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := functionsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create function response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test 2"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2018-11-10T20:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Functions/ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the function does not contain a friendly name", func() {
			createInput := &functions.CreateFunctionInput{}

			resp, err := functionsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create function response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create function api returns a 500 response", func() {
			createInput := &functions.CreateFunctionInput{
				FriendlyName: "Test 2",
			}

			httpmock.RegisterResponder("POST", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Functions",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := functionsClient.Create(createInput)
			It("Then an error should be returned", func() {
				Expect(err).ToNot(BeNil())
				twilioErr, ok := err.(*utils.TwilioError)
				Expect(ok).To(Equal(true))
				Expect(twilioErr.Code).To(BeNil())
				Expect(twilioErr.Message).To(Equal("An error occurred"))
				Expect(twilioErr.MoreInfo).To(BeNil())
				Expect(twilioErr.Status).To(Equal(500))
			})

			It("Then the create function response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a function sid", func() {
		functionClient := serverlessSession.Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Function("ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the function is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Functions/ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/functionResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := functionClient.Get()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get function response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test 2"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2018-11-10T20:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Functions/ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the get function response returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Functions/ZH71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := serverlessSession.Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Function("ZH71").Get()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get function response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the function is successfully updated", func() {
			updateInput := &function.UpdateFunctionInput{
				FriendlyName: "Test 3",
			}

			httpmock.RegisterResponder("POST", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Functions/ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateFunctionResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := functionClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get function response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test 3"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2018-11-10T20:00:00Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2018-11-11T20:00:00Z"))
				Expect(resp.URL).To(Equal("https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Functions/ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update function response returns a 404", func() {
			updateInput := &function.UpdateFunctionInput{
				FriendlyName: "Test 2",
			}

			httpmock.RegisterResponder("POST", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Functions/ZH71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := serverlessSession.Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Function("ZH71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get function response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the function is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Functions/ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := functionClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the delete function response returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Functions/ZH71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := serverlessSession.Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Function("ZH71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the function versions client", func() {
		functionVersionsClient := serverlessSession.Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Function("ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Versions

		Describe("When the function versions is successfully created", func() {
			createInput := &functionVersions.CreateVersionInput{
				Content: functionVersions.ContentDetails{
					Body:        strings.NewReader("Test Content"),
					ContentType: "application/javascript",
					FileName:    "test.js",
				},
				Path:       "/test",
				Visibility: "public",
			}

			httpmock.RegisterResponder("POST", "https://serverless-upload.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Functions/ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Versions",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/functionVersionResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := functionVersionsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create function response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("ZNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FunctionSid).To(Equal("ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Path).To(Equal("/test"))
				Expect(resp.Visibility).To(Equal("public"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2018-11-10T20:00:00Z"))
				Expect(resp.URL).To(Equal("https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Functions/ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Versions/ZNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the function version request does not contain content", func() {
			createInput := &functionVersions.CreateVersionInput{
				Path:       "/test",
				Visibility: "public",
			}

			resp, err := functionVersionsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create function version response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the asset version request does not contain filename", func() {
			createInput := &functionVersions.CreateVersionInput{
				Content: functionVersions.ContentDetails{
					Body:        strings.NewReader("Test Content"),
					ContentType: "application/javascript",
				},
				Path:       "/test",
				Visibility: "public",
			}

			resp, err := functionVersionsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create asset version response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the asset version request does not contain content type", func() {
			createInput := &functionVersions.CreateVersionInput{
				Content: functionVersions.ContentDetails{
					Body:     strings.NewReader("Test Content"),
					FileName: "test.js",
				},
				Path:       "/test",
				Visibility: "public",
			}

			resp, err := functionVersionsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create asset version response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the asset version request does not contain content body", func() {
			createInput := &functionVersions.CreateVersionInput{
				Content: functionVersions.ContentDetails{
					ContentType: "application/javascript",
					FileName:    "test.js",
				},
				Path:       "/test",
				Visibility: "public",
			}

			resp, err := functionVersionsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create asset version response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the function version request does not contain a path", func() {
			createInput := &functionVersions.CreateVersionInput{
				Content: functionVersions.ContentDetails{
					Body:        strings.NewReader("Test Content"),
					ContentType: "application/javascript",
					FileName:    "test.js",
				},
				Visibility: "public",
			}

			resp, err := functionVersionsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create function version response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the function version request does not contain visibility", func() {
			createInput := &functionVersions.CreateVersionInput{
				Content: functionVersions.ContentDetails{
					Body:        strings.NewReader("Test Content"),
					ContentType: "application/javascript",
					FileName:    "test.js",
				},
				Path: "/test",
			}

			resp, err := functionVersionsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create function version response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create function version api returns a 500 response", func() {
			createInput := &functionVersions.CreateVersionInput{
				Content: functionVersions.ContentDetails{
					Body:        strings.NewReader("Test Content"),
					ContentType: "application/javascript",
					FileName:    "test.js",
				},
				Path:       "/test",
				Visibility: "public",
			}

			httpmock.RegisterResponder("POST", "https://serverless-upload.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Functions/ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Versions",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := functionVersionsClient.Create(createInput)
			It("Then an error should be returned", func() {
				Expect(err).ToNot(BeNil())
				twilioErr, ok := err.(*utils.TwilioError)
				Expect(ok).To(Equal(true))
				Expect(twilioErr.Code).To(BeNil())
				Expect(twilioErr.Message).To(Equal("An error occurred"))
				Expect(twilioErr.MoreInfo).To(BeNil())
				Expect(twilioErr.Status).To(Equal(500))
			})

			It("Then the create function version response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a function version sid", func() {
		functionVersionClient := serverlessSession.Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Function("ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Version("ZNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the function version is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Functions/ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Versions/ZNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/functionVersionResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := functionVersionClient.Get()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get function version response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("ZNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FunctionSid).To(Equal("ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Path).To(Equal("/test"))
				Expect(resp.Visibility).To(Equal("public"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2018-11-10T20:00:00Z"))
				Expect(resp.URL).To(Equal("https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Functions/ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Versions/ZNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the get function version response returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Functions/ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Versions/ZN71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := serverlessSession.Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Function("ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Version("ZN71").Get()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get function version response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a function version content client", func() {
		functionVersionContentClient := serverlessSession.Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Function("ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Version("ZNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Content

		Describe("When the function version content is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Functions/ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Versions/ZNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Content",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/functionVersionContentResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := functionVersionContentClient.Get()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get function version response content should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("ZNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FunctionSid).To(Equal("ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Content).To(Equal("exports.handler = function (context, event, callback) { callback(null, \"Hello World\") }"))
				Expect(resp.URL).To(Equal("https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Functions/ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Versions/ZNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Content"))
			})
		})

		Describe("When the get function version content response returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Functions/ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Versions/ZN71/Content",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := serverlessSession.Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Function("ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Version("ZN71").Content.Get()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get function version context response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given the assets client", func() {
		assetsClient := serverlessSession.Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Assets

		Describe("When the asset is successfully created", func() {
			createInput := &assets.CreateAssetInput{
				FriendlyName: "Test Asset 2",
			}

			httpmock.RegisterResponder("POST", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Assets",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/assetResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := assetsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create asset response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test Asset 2"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2018-11-10T20:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Assets/ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the asset does not contain a friendly name", func() {
			createInput := &assets.CreateAssetInput{}

			resp, err := assetsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create asset response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create asset api returns a 500 response", func() {
			createInput := &assets.CreateAssetInput{
				FriendlyName: "Test Asset 2",
			}

			httpmock.RegisterResponder("POST", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Assets",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := assetsClient.Create(createInput)
			It("Then an error should be returned", func() {
				Expect(err).ToNot(BeNil())
				twilioErr, ok := err.(*utils.TwilioError)
				Expect(ok).To(Equal(true))
				Expect(twilioErr.Code).To(BeNil())
				Expect(twilioErr.Message).To(Equal("An error occurred"))
				Expect(twilioErr.MoreInfo).To(BeNil())
				Expect(twilioErr.Status).To(Equal(500))
			})

			It("Then the create asset response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a asset sid", func() {
		assetClient := serverlessSession.Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Asset("ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the asset is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Assets/ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/assetResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := assetClient.Get()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get asset response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test Asset 2"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2018-11-10T20:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Assets/ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the get asset response returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Assets/ZH71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := serverlessSession.Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Asset("ZH71").Get()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get asset response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the asset is successfully retrieved", func() {
			updateInput := &asset.UpdateAssetInput{
				FriendlyName: "Test Asset 3",
			}

			httpmock.RegisterResponder("POST", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Assets/ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateAssetResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := assetClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get asset response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test Asset 3"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2018-11-10T20:00:00Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2018-11-11T20:00:00Z"))
				Expect(resp.URL).To(Equal("https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Assets/ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update asset response returns a 404", func() {
			updateInput := &asset.UpdateAssetInput{
				FriendlyName: "Test 2",
			}

			httpmock.RegisterResponder("POST", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Assets/ZH71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := serverlessSession.Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Asset("ZH71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get asset response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the asset is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Assets/ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := assetClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the delete asset response returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Assets/ZH71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := serverlessSession.Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Asset("ZH71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the asset versions client", func() {
		assetVersionsClient := serverlessSession.Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Asset("ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Versions

		Describe("When the asset versions is successfully created", func() {
			createInput := &assetVersions.CreateVersionInput{
				Content: assetVersions.ContentDetails{
					Body:        strings.NewReader("Test Content"),
					ContentType: "image/png",
					FileName:    "test.png",
				},
				Path:       "/test",
				Visibility: "public",
			}

			httpmock.RegisterResponder("POST", "https://serverless-upload.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Assets/ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Versions",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/assetVersionResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := assetVersionsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create asset response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("ZNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AssetSid).To(Equal("ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Path).To(Equal("/test"))
				Expect(resp.Visibility).To(Equal("public"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2018-11-10T20:00:00Z"))
				Expect(resp.URL).To(Equal("https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Assets/ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Versions/ZNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the asset version request does not contain content", func() {
			createInput := &assetVersions.CreateVersionInput{
				Path:       "/test",
				Visibility: "public",
			}

			resp, err := assetVersionsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create asset version response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the asset version request does not contain filename", func() {
			createInput := &assetVersions.CreateVersionInput{
				Content: assetVersions.ContentDetails{
					Body:        strings.NewReader("Test Content"),
					ContentType: "image/png",
				},
				Path:       "/test",
				Visibility: "public",
			}

			resp, err := assetVersionsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create asset version response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the asset version request does not contain content type", func() {
			createInput := &assetVersions.CreateVersionInput{
				Content: assetVersions.ContentDetails{
					Body:     strings.NewReader("Test Content"),
					FileName: "test.png",
				},
				Path:       "/test",
				Visibility: "public",
			}

			resp, err := assetVersionsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create asset version response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the asset version request does not contain content body", func() {
			createInput := &assetVersions.CreateVersionInput{
				Content: assetVersions.ContentDetails{
					ContentType: "image/png",
					FileName:    "test.png",
				},
				Path:       "/test",
				Visibility: "public",
			}

			resp, err := assetVersionsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create asset version response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the asset version request does not contain a path", func() {
			createInput := &assetVersions.CreateVersionInput{
				Content: assetVersions.ContentDetails{
					Body:        strings.NewReader("Test Content"),
					ContentType: "image/png",
					FileName:    "test.png",
				},
				Visibility: "public",
			}

			resp, err := assetVersionsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create asset version response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the asset version request does not contain visibility", func() {
			createInput := &assetVersions.CreateVersionInput{
				Content: assetVersions.ContentDetails{
					Body:        strings.NewReader("Test Content"),
					ContentType: "image/png",
					FileName:    "test.png",
				},
				Path: "/test",
			}

			resp, err := assetVersionsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create asset version response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create asset version api returns a 500 response", func() {
			createInput := &assetVersions.CreateVersionInput{
				Content: assetVersions.ContentDetails{
					Body:        strings.NewReader("Test Content"),
					ContentType: "image/png",
					FileName:    "test.png",
				},
				Path:       "/test",
				Visibility: "public",
			}

			httpmock.RegisterResponder("POST", "https://serverless-upload.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Assets/ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Versions",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := assetVersionsClient.Create(createInput)
			It("Then an error should be returned", func() {
				Expect(err).ToNot(BeNil())
				twilioErr, ok := err.(*utils.TwilioError)
				Expect(ok).To(Equal(true))
				Expect(twilioErr.Code).To(BeNil())
				Expect(twilioErr.Message).To(Equal("An error occurred"))
				Expect(twilioErr.MoreInfo).To(BeNil())
				Expect(twilioErr.Status).To(Equal(500))
			})

			It("Then the create asset version response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a asset version sid", func() {
		assetVersionClient := serverlessSession.Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Asset("ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Version("ZNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the asset version is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Assets/ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Versions/ZNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/assetVersionResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := assetVersionClient.Get()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get asset version response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("ZNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AssetSid).To(Equal("ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Path).To(Equal("/test"))
				Expect(resp.Visibility).To(Equal("public"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2018-11-10T20:00:00Z"))
				Expect(resp.URL).To(Equal("https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Assets/ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Versions/ZNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the get asset version response returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Assets/ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Versions/ZN71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := serverlessSession.Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Asset("ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Version("ZN71").Get()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get asset version response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given the build client", func() {
		buildClient := serverlessSession.Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Builds

		Describe("When the build is successfully created", func() {
			createInput := &builds.CreateBuildInput{}

			httpmock.RegisterResponder("POST", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Builds",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/buildResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := buildClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create build response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("ZBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))

				Expect(resp.AssetVersions).ToNot(BeNil())

				assetVersions := *resp.AssetVersions
				Expect(len(assetVersions)).To(Equal(1))
				Expect(assetVersions[0].Sid).To(Equal("ZNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(assetVersions[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(assetVersions[0].ServiceSid).To(Equal("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(assetVersions[0].AssetSid).To(Equal("ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(assetVersions[0].DateCreated.Format(time.RFC3339)).To(Equal("2018-11-10T20:00:00Z"))
				Expect(assetVersions[0].Path).To(Equal("/asset-test"))
				Expect(assetVersions[0].Visibility).To(Equal("PUBLIC"))

				Expect(resp.FunctionVersions).ToNot(BeNil())

				functionVersions := *resp.FunctionVersions
				Expect(len(functionVersions)).To(Equal(1))
				Expect(functionVersions[0].Sid).To(Equal("ZNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX1"))
				Expect(functionVersions[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(functionVersions[0].ServiceSid).To(Equal("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(functionVersions[0].FunctionSid).To(Equal("ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX1"))
				Expect(functionVersions[0].DateCreated.Format(time.RFC3339)).To(Equal("2018-11-10T20:00:00Z"))
				Expect(functionVersions[0].Path).To(Equal("/function-test"))
				Expect(functionVersions[0].Visibility).To(Equal("PUBLIC"))

				Expect(resp.Dependencies).ToNot(BeNil())

				dependencies := *resp.Dependencies
				Expect(len(dependencies)).To(Equal(1))
				Expect(dependencies[0].Name).To(Equal("twilio"))
				Expect(dependencies[0].Version).To(Equal("3.6.3"))

				Expect(resp.Status).To(Equal("building"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2018-11-10T20:00:00Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2018-11-10T20:00:00Z"))
				Expect(resp.URL).To(Equal("https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Builds/ZBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the create build api returns a 500 response", func() {
			createInput := &builds.CreateBuildInput{}

			httpmock.RegisterResponder("POST", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Builds",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := buildClient.Create(createInput)
			It("Then an error should be returned", func() {
				Expect(err).ToNot(BeNil())
				twilioErr, ok := err.(*utils.TwilioError)
				Expect(ok).To(Equal(true))
				Expect(twilioErr.Code).To(BeNil())
				Expect(twilioErr.Message).To(Equal("An error occurred"))
				Expect(twilioErr.MoreInfo).To(BeNil())
				Expect(twilioErr.Status).To(Equal(500))
			})

			It("Then the create build response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a build sid", func() {
		buildClient := serverlessSession.Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Build("ZBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the build is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Builds/ZBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/buildResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := buildClient.Get()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get build response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("ZBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))

				Expect(resp.AssetVersions).ToNot(BeNil())

				assetVersions := *resp.AssetVersions
				Expect(len(assetVersions)).To(Equal(1))
				Expect(assetVersions[0].Sid).To(Equal("ZNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(assetVersions[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(assetVersions[0].ServiceSid).To(Equal("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(assetVersions[0].AssetSid).To(Equal("ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(assetVersions[0].DateCreated.Format(time.RFC3339)).To(Equal("2018-11-10T20:00:00Z"))
				Expect(assetVersions[0].Path).To(Equal("/asset-test"))
				Expect(assetVersions[0].Visibility).To(Equal("PUBLIC"))

				Expect(resp.FunctionVersions).ToNot(BeNil())

				functionVersions := *resp.FunctionVersions
				Expect(len(functionVersions)).To(Equal(1))
				Expect(functionVersions[0].Sid).To(Equal("ZNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX1"))
				Expect(functionVersions[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(functionVersions[0].ServiceSid).To(Equal("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(functionVersions[0].FunctionSid).To(Equal("ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX1"))
				Expect(functionVersions[0].DateCreated.Format(time.RFC3339)).To(Equal("2018-11-10T20:00:00Z"))
				Expect(functionVersions[0].Path).To(Equal("/function-test"))
				Expect(functionVersions[0].Visibility).To(Equal("PUBLIC"))

				Expect(resp.Dependencies).ToNot(BeNil())

				dependencies := *resp.Dependencies
				Expect(len(dependencies)).To(Equal(1))
				Expect(dependencies[0].Name).To(Equal("twilio"))
				Expect(dependencies[0].Version).To(Equal("3.6.3"))

				Expect(resp.Status).To(Equal("building"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2018-11-10T20:00:00Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2018-11-10T20:00:00Z"))
				Expect(resp.URL).To(Equal("https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Builds/ZBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the get build response returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Builds/ZB71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := serverlessSession.Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Build("ZB71").Get()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get build response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the build is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Builds/ZBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := buildClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the delete variable response returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Builds/ZB71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := serverlessSession.Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Build("ZB71").Delete()
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
	Expect(twilioErr.Message).To(Equal("The requested resource /Services/ZS71 was not found"))

	moreInfo := "https://www.twilio.com/docs/errors/20404"
	Expect(twilioErr.MoreInfo).To(Equal(&moreInfo))
	Expect(twilioErr.Status).To(Equal(404))
}

func ExpectErrorToNotBeATwilioError(err error) {
	Expect(err).ToNot(BeNil())
	_, ok := err.(*utils.TwilioError)
	Expect(ok).To(Equal(false))
}
