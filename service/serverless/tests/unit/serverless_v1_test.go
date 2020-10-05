package tests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service/environment/logs"
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
		Fail(fmt.Sprintf("%s", err.Error()))
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
				ExpectInternalServerError(err)
			})

			It("Then the create service response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of services are successfully retrieved", func() {
			pageOptions := &services.ServicesPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/servicesPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := servicesClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the services page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://serverless.twilio.com/v1/Services?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://serverless.twilio.com/v1/Services?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("services"))

				services := resp.Services
				Expect(services).ToNot(BeNil())
				Expect(len(services)).To(Equal(1))

				Expect(services[0].Sid).To(Equal("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(services[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(services[0].FriendlyName).To(Equal("Test 2"))
				Expect(services[0].UniqueName).To(Equal("Unique Test 2"))
				Expect(services[0].IncludeCredentials).To(Equal(true))
				Expect(services[0].UiEditable).To(Equal(false))
				Expect(services[0].DateCreated.Format(time.RFC3339)).To(Equal("2018-11-10T20:00:00Z"))
				Expect(services[0].DateUpdated).To(BeNil())
				Expect(services[0].URL).To(Equal("https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of services api returns a 500 response", func() {
			pageOptions := &services.ServicesPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := servicesClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the services page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated services are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/servicesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/servicesPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := servicesClient.NewServicesPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated services current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated services results should be returned", func() {
				Expect(len(paginator.Services)).To(Equal(3))
			})
		})

		Describe("When the services api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/servicesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := servicesClient.NewServicesPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated services current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
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

			resp, err := serviceClient.Fetch()
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

			resp, err := serverlessSession.Service("ZS71").Fetch()
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
				ExpectInternalServerError(err)
			})

			It("Then the create environment response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of environments are successfully retrieved", func() {
			pageOptions := &environments.EnvironmentsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/environmentsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := environmentsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the environments page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("environments"))

				environments := resp.Environments
				Expect(environments).ToNot(BeNil())
				Expect(len(environments)).To(Equal(1))

				Expect(environments[0].Sid).To(Equal("ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(environments[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(environments[0].ServiceSid).To(Equal("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(environments[0].UniqueName).To(Equal("test-2"))
				Expect(environments[0].DomainSuffix).To(BeNil())
				Expect(environments[0].DomainName).To(Equal("test-2.twil.io"))
				Expect(environments[0].DateCreated.Format(time.RFC3339)).To(Equal("2018-11-10T20:00:00Z"))
				Expect(environments[0].DateUpdated).To(BeNil())
				Expect(environments[0].URL).To(Equal("https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of environments api returns a 500 response", func() {
			pageOptions := &environments.EnvironmentsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := environmentsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the environments page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated environments are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/environmentsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/environmentsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := environmentsClient.NewEnvironmentsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated environments current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated environments results should be returned", func() {
				Expect(len(paginator.Environments)).To(Equal(3))
			})
		})

		Describe("When the environments api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/environmentsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := environmentsClient.NewEnvironmentsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated environments current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
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

			resp, err := environmentClient.Fetch()
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

			resp, err := serverlessSession.Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Environment("ZE71").Fetch()
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
				ExpectInternalServerError(err)
			})

			It("Then the create variable response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of variables are successfully retrieved", func() {
			pageOptions := &variables.VariablesPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Variables?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/variablesPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := variablesClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the variables page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Variables?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Variables?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("variables"))

				variables := resp.Variables
				Expect(variables).ToNot(BeNil())
				Expect(len(variables)).To(Equal(1))

				Expect(variables[0].Sid).To(Equal("ZVXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(variables[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(variables[0].ServiceSid).To(Equal("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(variables[0].EnvironmentSid).To(Equal("ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(variables[0].Key).To(Equal("key"))
				Expect(variables[0].Value).To(Equal("value"))
				Expect(variables[0].DateCreated.Format(time.RFC3339)).To(Equal("2018-11-10T20:00:00Z"))
				Expect(variables[0].DateUpdated).To(BeNil())
				Expect(variables[0].URL).To(Equal("https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Variables/ZVXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of variables api returns a 500 response", func() {
			pageOptions := &variables.VariablesPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Variables?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := variablesClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the variables page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated variables are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Variables",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/variablesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Variables?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/variablesPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := variablesClient.NewVariablesPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated variables current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated variables results should be returned", func() {
				Expect(len(paginator.Variables)).To(Equal(3))
			})
		})

		Describe("When the variables api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Variables",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/variablesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Variables?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := variablesClient.NewVariablesPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated variables current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
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

			resp, err := variableClient.Fetch()
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

			resp, err := serverlessSession.Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Environment("ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Variable("ZV71").Fetch()
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
				ExpectInternalServerError(err)
			})

			It("Then the create deployment response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of deployments are successfully retrieved", func() {
			pageOptions := &deployments.DeploymentsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Deployments?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/deploymentsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := deploymentsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the deployments page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Deployments?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Deployments?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("deployments"))

				deployments := resp.Deployments
				Expect(deployments).ToNot(BeNil())
				Expect(len(deployments)).To(Equal(1))

				Expect(deployments[0].Sid).To(Equal("ZDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(deployments[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(deployments[0].ServiceSid).To(Equal("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(deployments[0].EnvironmentSid).To(Equal("ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(deployments[0].BuildSid).To(Equal("ZBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(deployments[0].DateCreated.Format(time.RFC3339)).To(Equal("2018-11-10T20:00:00Z"))
				Expect(deployments[0].DateUpdated).To(BeNil())
				Expect(deployments[0].URL).To(Equal("https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Deployments/ZDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of deployments api returns a 500 response", func() {
			pageOptions := &deployments.DeploymentsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Deployments?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := deploymentsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the deployments page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated deployments are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Deployments",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/deploymentsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Deployments?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/deploymentsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := deploymentsClient.NewDeploymentsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated deployments current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated deployments results should be returned", func() {
				Expect(len(paginator.Deployments)).To(Equal(3))
			})
		})

		Describe("When the deployments api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Deployments",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/deploymentsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Deployments?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := deploymentsClient.NewDeploymentsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated deployments current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a deployment sid", func() {
		deploymentClient := serverlessSession.Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Environment("ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Deployment("ZDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the deployment is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Deployments/ZDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/deploymentResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := deploymentClient.Fetch()
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

			resp, err := serverlessSession.Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Environment("ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Deployment("ZD71").Fetch()
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
				ExpectInternalServerError(err)
			})

			It("Then the create function response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of functions are successfully retrieved", func() {
			pageOptions := &functions.FunctionsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Functions?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/functionsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := functionsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the functions page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Functions?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Functions?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("functions"))

				functions := resp.Functions
				Expect(functions).ToNot(BeNil())
				Expect(len(functions)).To(Equal(1))

				Expect(functions[0].Sid).To(Equal("ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(functions[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(functions[0].ServiceSid).To(Equal("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(functions[0].FriendlyName).To(Equal("Test 2"))
				Expect(functions[0].DateCreated.Format(time.RFC3339)).To(Equal("2018-11-10T20:00:00Z"))
				Expect(functions[0].DateUpdated).To(BeNil())
				Expect(functions[0].URL).To(Equal("https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Functions/ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of functions api returns a 500 response", func() {
			pageOptions := &functions.FunctionsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Functions?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := functionsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the functions page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated functions are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Functions",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/functionsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Functions?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/functionsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := functionsClient.NewFunctionsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated functions current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated functions results should be returned", func() {
				Expect(len(paginator.Functions)).To(Equal(3))
			})
		})

		Describe("When the functions api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Functions",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/functionsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Functions?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := functionsClient.NewFunctionsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated functions current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
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

			resp, err := functionClient.Fetch()
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

			resp, err := serverlessSession.Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Function("ZH71").Fetch()
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
				Content: functionVersions.CreateContentDetails{
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
				Content: functionVersions.CreateContentDetails{
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
				Content: functionVersions.CreateContentDetails{
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
				Content: functionVersions.CreateContentDetails{
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
				Content: functionVersions.CreateContentDetails{
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
				Content: functionVersions.CreateContentDetails{
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
				Content: functionVersions.CreateContentDetails{
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
				ExpectInternalServerError(err)
			})

			It("Then the create function version response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of function versions are successfully retrieved", func() {
			pageOptions := &functionVersions.VersionsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Functions/ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Versions?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/functionVersionsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := functionVersionsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the function versions page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Functions/ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Versions?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Functions/ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Versions?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("function_versions"))

				functionVersions := resp.Versions
				Expect(functionVersions).ToNot(BeNil())
				Expect(len(functionVersions)).To(Equal(1))

				Expect(functionVersions[0].Sid).To(Equal("ZNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(functionVersions[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(functionVersions[0].ServiceSid).To(Equal("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(functionVersions[0].FunctionSid).To(Equal("ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(functionVersions[0].Path).To(Equal("/test"))
				Expect(functionVersions[0].Visibility).To(Equal("public"))
				Expect(functionVersions[0].DateCreated.Format(time.RFC3339)).To(Equal("2018-11-10T20:00:00Z"))
				Expect(functionVersions[0].URL).To(Equal("https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Functions/ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Versions/ZNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of function versions api returns a 500 response", func() {
			pageOptions := &functionVersions.VersionsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Functions/ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Versions?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := functionVersionsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the function versions page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated function versions are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Functions/ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Versions",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/functionVersionsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Functions/ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Versions?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/functionVersionsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := functionVersionsClient.NewVersionsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated function versions current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated function versions results should be returned", func() {
				Expect(len(paginator.Versions)).To(Equal(3))
			})
		})

		Describe("When the function versions api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Functions/ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Versions",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/functionVersionsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Functions/ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Versions?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := functionVersionsClient.NewVersionsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated function versions current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
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

			resp, err := functionVersionClient.Fetch()
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

			resp, err := serverlessSession.Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Function("ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Version("ZN71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get function version response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a function version content client", func() {
		functionVersionContentClient := serverlessSession.Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Function("ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Version("ZNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Content()

		Describe("When the function version content is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Functions/ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Versions/ZNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Content",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/functionVersionContentResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := functionVersionContentClient.Fetch()
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

			resp, err := serverlessSession.Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Function("ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Version("ZN71").Content().Fetch()
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
				ExpectInternalServerError(err)
			})

			It("Then the create asset response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of assets are successfully retrieved", func() {
			pageOptions := &assets.AssetsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Assets?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/assetsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := assetsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the assets page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Assets?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Assets?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("assets"))

				assets := resp.Assets
				Expect(assets).ToNot(BeNil())
				Expect(len(assets)).To(Equal(1))

				Expect(assets[0].Sid).To(Equal("ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(assets[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(assets[0].ServiceSid).To(Equal("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(assets[0].FriendlyName).To(Equal("Test Asset 2"))
				Expect(assets[0].DateCreated.Format(time.RFC3339)).To(Equal("2018-11-10T20:00:00Z"))
				Expect(assets[0].DateUpdated).To(BeNil())
				Expect(assets[0].URL).To(Equal("https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Assets/ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of assets api returns a 500 response", func() {
			pageOptions := &assets.AssetsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Assets?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := assetsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the assets page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated assets are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Assets",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/assetsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Assets?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/assetsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := assetsClient.NewAssetsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated assets current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated assets results should be returned", func() {
				Expect(len(paginator.Assets)).To(Equal(3))
			})
		})

		Describe("When the assets api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Assets",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/assetsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Assets?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := assetsClient.NewAssetsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated assets current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
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

			resp, err := assetClient.Fetch()
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

			resp, err := serverlessSession.Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Asset("ZH71").Fetch()
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
				Content: assetVersions.CreateContentDetails{
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
				Content: assetVersions.CreateContentDetails{
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
				Content: assetVersions.CreateContentDetails{
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
				Content: assetVersions.CreateContentDetails{
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
				Content: assetVersions.CreateContentDetails{
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
				Content: assetVersions.CreateContentDetails{
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
				Content: assetVersions.CreateContentDetails{
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
				ExpectInternalServerError(err)
			})

			It("Then the create asset version response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of asset versions are successfully retrieved", func() {
			pageOptions := &assetVersions.VersionsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Assets/ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Versions?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/assetVersionsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := assetVersionsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the asset versions page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Assets/ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Versions?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Assets/ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Versions?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("asset_versions"))

				assetVersions := resp.Versions
				Expect(assetVersions).ToNot(BeNil())
				Expect(len(assetVersions)).To(Equal(1))

				Expect(assetVersions[0].Sid).To(Equal("ZNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(assetVersions[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(assetVersions[0].ServiceSid).To(Equal("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(assetVersions[0].AssetSid).To(Equal("ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(assetVersions[0].Path).To(Equal("/test"))
				Expect(assetVersions[0].Visibility).To(Equal("public"))
				Expect(assetVersions[0].DateCreated.Format(time.RFC3339)).To(Equal("2018-11-10T20:00:00Z"))
				Expect(assetVersions[0].URL).To(Equal("https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Assets/ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Versions/ZNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of asset versions api returns a 500 response", func() {
			pageOptions := &assetVersions.VersionsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Assets/ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Versions?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := assetVersionsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the asset versions page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated asset versions are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Assets/ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Versions",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/assetVersionsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Assets/ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Versions?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/assetVersionsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := assetVersionsClient.NewVersionsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated asset versions current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated asset versions results should be returned", func() {
				Expect(len(paginator.Versions)).To(Equal(3))
			})
		})

		Describe("When the asset versions api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Assets/ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Versions",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/assetVersionsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Assets/ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Versions?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := assetVersionsClient.NewVersionsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated asset versions current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
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

			resp, err := assetVersionClient.Fetch()
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

			resp, err := serverlessSession.Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Asset("ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Version("ZN71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get asset version response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given the builds client", func() {
		buildsClient := serverlessSession.Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Builds

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

			resp, err := buildsClient.Create(createInput)
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

			resp, err := buildsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create build response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of builds are successfully retrieved", func() {
			pageOptions := &builds.BuildsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Builds?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/buildsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := buildsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the builds page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Builds?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Builds?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("builds"))

				builds := resp.Builds
				Expect(builds).ToNot(BeNil())
				Expect(len(builds)).To(Equal(1))

				Expect(builds[0].Sid).To(Equal("ZBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(builds[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(builds[0].ServiceSid).To(Equal("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))

				Expect(builds[0].AssetVersions).ToNot(BeNil())

				assetVersions := *builds[0].AssetVersions
				Expect(len(assetVersions)).To(Equal(1))
				Expect(assetVersions[0].Sid).To(Equal("ZNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(assetVersions[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(assetVersions[0].ServiceSid).To(Equal("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(assetVersions[0].AssetSid).To(Equal("ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(assetVersions[0].DateCreated.Format(time.RFC3339)).To(Equal("2018-11-10T20:00:00Z"))
				Expect(assetVersions[0].Path).To(Equal("/asset-test"))
				Expect(assetVersions[0].Visibility).To(Equal("PUBLIC"))

				Expect(builds[0].FunctionVersions).ToNot(BeNil())

				functionVersions := *builds[0].FunctionVersions
				Expect(len(functionVersions)).To(Equal(1))
				Expect(functionVersions[0].Sid).To(Equal("ZNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX1"))
				Expect(functionVersions[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(functionVersions[0].ServiceSid).To(Equal("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(functionVersions[0].FunctionSid).To(Equal("ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX1"))
				Expect(functionVersions[0].DateCreated.Format(time.RFC3339)).To(Equal("2018-11-10T20:00:00Z"))
				Expect(functionVersions[0].Path).To(Equal("/function-test"))
				Expect(functionVersions[0].Visibility).To(Equal("PUBLIC"))

				Expect(builds[0].Dependencies).ToNot(BeNil())

				dependencies := *builds[0].Dependencies
				Expect(len(dependencies)).To(Equal(1))
				Expect(dependencies[0].Name).To(Equal("twilio"))
				Expect(dependencies[0].Version).To(Equal("3.6.3"))

				Expect(builds[0].Status).To(Equal("building"))
				Expect(builds[0].DateCreated.Format(time.RFC3339)).To(Equal("2018-11-10T20:00:00Z"))
				Expect(builds[0].DateUpdated.Format(time.RFC3339)).To(Equal("2018-11-10T20:00:00Z"))
				Expect(builds[0].URL).To(Equal("https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Builds/ZBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of builds api returns a 500 response", func() {
			pageOptions := &builds.BuildsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Builds?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := buildsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the builds page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated builds are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Builds",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/buildsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Builds?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/buildsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := buildsClient.NewBuildsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated builds current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated builds results should be returned", func() {
				Expect(len(paginator.Builds)).To(Equal(3))
			})
		})

		Describe("When the builds api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Builds",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/buildsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Builds?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := buildsClient.NewBuildsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated builds current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
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

			resp, err := buildClient.Fetch()
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

			resp, err := serverlessSession.Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Build("ZB71").Fetch()
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

		Describe("When the delete build api returns a 404", func() {
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

	Describe("Given I have a status client", func() {
		statusClient := serverlessSession.Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Build("ZBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Status()

		Describe("When the status is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Builds/ZBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Status",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/statusResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := statusClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get status response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("ZBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Status).To(Equal("completed"))
				Expect(resp.URL).To(Equal("https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Builds/ZBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Status"))
			})
		})

		Describe("When the get status api returns a 500 response", func() {
			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Builds/ZBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Status",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := statusClient.Fetch()
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the get status response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given the logs client", func() {
		logsClient := serverlessSession.Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Environment("ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Logs

		Describe("When the page of logs are successfully retrieved", func() {
			pageOptions := &logs.LogsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Logs?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/logsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := logsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the log page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Logs?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Logs?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("logs"))

				logs := resp.Logs
				Expect(logs).ToNot(BeNil())
				Expect(len(logs)).To(Equal(1))

				Expect(logs[0].Sid).To(Equal("NOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(logs[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(logs[0].ServiceSid).To(Equal("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(logs[0].EnvironmentSid).To(Equal("ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(logs[0].BuildSid).To(Equal("ZBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(logs[0].FunctionSid).To(Equal("ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(logs[0].RequestSid).To(Equal("RQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(logs[0].DeploymentSid).To(Equal("ZDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(logs[0].Level).To(Equal("INFO"))
				Expect(logs[0].Message).To(Equal("Test Log"))
				Expect(logs[0].URL).To(Equal("https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Logs/NOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of logs api returns a 500 response", func() {
			pageOptions := &logs.LogsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Logs?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := logsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the logs page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated logs are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Logs",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/logsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Logs?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/logsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := logsClient.NewLogsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated logs current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated logs results should be returned", func() {
				Expect(len(paginator.Logs)).To(Equal(3))
			})
		})

		Describe("When the logs api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Logs",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/logsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Logs?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := logsClient.NewLogsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated log current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a log sid", func() {
		logClient := serverlessSession.Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Environment("ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Log("NOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the log is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Logs/NOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/logResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := logClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get log response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("NOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.EnvironmentSid).To(Equal("ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.BuildSid).To(Equal("ZBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FunctionSid).To(Equal("ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.RequestSid).To(Equal("RQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.DeploymentSid).To(Equal("ZDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Level).To(Equal("INFO"))
				Expect(resp.Message).To(Equal("Test Log"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2018-11-10T20:00:00Z"))
				Expect(resp.URL).To(Equal("https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Logs/NOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the get log response returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://serverless.twilio.com/v1/Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Logs/NO71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := serverlessSession.Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Environment("ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Log("NO71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get log response should be nil", func() {
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
