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

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/sync"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/document"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/document/permission"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/document/permissions"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/documents"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/sync_list"
	syncListItem "github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/sync_list/item"
	syncListItems "github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/sync_list/items"
	syncListPermission "github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/sync_list/permission"
	syncListPermissions "github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/sync_list/permissions"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/sync_lists"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/sync_map"
	syncMapItem "github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/sync_map/item"
	syncMapItems "github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/sync_map/items"
	syncMapPermission "github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/sync_map/permission"
	syncMapPermissions "github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/sync_map/permissions"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/sync_maps"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/sync_stream"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/sync_stream/messages"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/sync_streams"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/services"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

var _ = Describe("Sync V1", func() {
	creds, err := credentials.New(credentials.Account{
		Sid:       "ACxxx",
		AuthToken: "Test",
	})
	if err != nil {
		log.Panicf("%s", err)
	}

	syncSession := sync.New(session.New(creds), &client.Config{
		RetryAttempts: utils.Int(0),
	}).V1

	httpmock.ActivateNonDefault(syncSession.GetClient().GetRestyClient().GetClient())
	defer httpmock.DeactivateAndReset()

	Describe("Given the services client", func() {
		servicesClient := syncSession.Services

		Describe("When the service is successfully created", func() {
			createInput := &services.CreateServiceInput{}

			httpmock.RegisterResponder("POST", "https://sync.twilio.com/v1/Services",
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
				Expect(resp.Sid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.UniqueName).To(BeNil())
				Expect(resp.FriendlyName).To(BeNil())
				Expect(resp.ReachabilityDebouncingWindow).To(Equal(5000))
				Expect(resp.ReachabilityDebouncingEnabled).To(Equal(false))
				Expect(resp.ReachabilityWebhooksEnabled).To(Equal(false))
				Expect(resp.WebhookURL).To(BeNil())
				Expect(resp.WebhooksFromRestEnabled).To(Equal(false))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the create service api returns a 500 response", func() {
			createInput := &services.CreateServiceInput{}

			httpmock.RegisterResponder("POST", "https://sync.twilio.com/v1/Services",
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

			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services?Page=0&PageSize=50",
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
				Expect(meta.FirstPageURL).To(Equal("https://sync.twilio.com/v1/Services?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://sync.twilio.com/v1/Services?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("services"))

				services := resp.Services
				Expect(services).ToNot(BeNil())
				Expect(len(services)).To(Equal(1))

				Expect(services[0].Sid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(services[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(services[0].UniqueName).To(BeNil())
				Expect(services[0].FriendlyName).To(BeNil())
				Expect(services[0].ReachabilityDebouncingWindow).To(Equal(5000))
				Expect(services[0].ReachabilityDebouncingEnabled).To(Equal(false))
				Expect(services[0].ReachabilityWebhooksEnabled).To(Equal(false))
				Expect(services[0].WebhookURL).To(BeNil())
				Expect(services[0].WebhooksFromRestEnabled).To(Equal(false))
				Expect(services[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(services[0].DateUpdated).To(BeNil())
				Expect(services[0].URL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of services api returns a 500 response", func() {
			pageOptions := &services.ServicesPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services?Page=0&PageSize=50",
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
			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/servicesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services?Page=1&PageSize=50&PageToken=abc1234",
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
			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/servicesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services?Page=1&PageSize=50&PageToken=abc1234",
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
		serviceClient := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the service is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
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
				Expect(resp.Sid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.UniqueName).To(BeNil())
				Expect(resp.FriendlyName).To(BeNil())
				Expect(resp.ReachabilityDebouncingWindow).To(Equal(5000))
				Expect(resp.ReachabilityDebouncingEnabled).To(Equal(false))
				Expect(resp.ReachabilityWebhooksEnabled).To(Equal(false))
				Expect(resp.WebhookURL).To(BeNil())
				Expect(resp.WebhooksFromRestEnabled).To(Equal(false))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the service api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/IS71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := syncSession.Service("IS71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get service response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the service is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateServiceResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &service.UpdateServiceInput{
				FriendlyName: utils.String("Test"),
			}

			resp, err := serviceClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update service response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.UniqueName).To(BeNil())
				Expect(resp.FriendlyName).To(Equal(utils.String("test")))
				Expect(resp.ReachabilityDebouncingWindow).To(Equal(5000))
				Expect(resp.ReachabilityDebouncingEnabled).To(Equal(false))
				Expect(resp.ReachabilityWebhooksEnabled).To(Equal(false))
				Expect(resp.WebhookURL).To(BeNil())
				Expect(resp.WebhooksFromRestEnabled).To(Equal(false))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-20T21:00:24Z"))
				Expect(resp.URL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update service response returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://sync.twilio.com/v1/Services/IS71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &service.UpdateServiceInput{
				FriendlyName: utils.String("Test 2"),
			}

			resp, err := syncSession.Service("IS71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update service response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the service is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := serviceClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the service api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://sync.twilio.com/v1/Services/IS71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := syncSession.Service("IS71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the documents client", func() {
		documentsClient := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Documents

		Describe("When the document is successfully created", func() {
			createInput := &documents.CreateDocumentInput{}

			httpmock.RegisterResponder("POST", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Documents",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/documentResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := documentsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create document response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("ETXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.CreatedBy).To(Equal("system"))
				Expect(resp.UniqueName).To(Equal(utils.String("test")))
				Expect(resp.Revision).To(Equal("0"))

				data := make(map[string]interface{}, 0)
				Expect(resp.Data).To(Equal(data))

				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateExpires).To(BeNil())
				Expect(resp.URL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Documents/ETXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the create document api returns a 500 response", func() {
			createInput := &documents.CreateDocumentInput{}

			httpmock.RegisterResponder("POST", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Documents",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := documentsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create document response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of documents are successfully retrieved", func() {
			pageOptions := &documents.DocumentsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Documents?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/documentsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := documentsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the documents page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Documents?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Documents?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("documents"))

				documents := resp.Documents
				Expect(documents).ToNot(BeNil())
				Expect(len(documents)).To(Equal(1))

				Expect(documents[0].Sid).To(Equal("ETXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(documents[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(documents[0].ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(documents[0].CreatedBy).To(Equal("system"))
				Expect(documents[0].UniqueName).To(Equal(utils.String("test")))
				Expect(documents[0].Revision).To(Equal("0"))

				data := make(map[string]interface{}, 0)
				Expect(documents[0].Data).To(Equal(data))

				Expect(documents[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(documents[0].DateUpdated).To(BeNil())
				Expect(documents[0].DateExpires).To(BeNil())
				Expect(documents[0].URL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Documents/ETXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of documents api returns a 500 response", func() {
			pageOptions := &documents.DocumentsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Documents?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := documentsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the documents page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated documents are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Documents",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/documentsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Documents?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/documentsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := documentsClient.NewDocumentsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated documents current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated documents results should be returned", func() {
				Expect(len(paginator.Documents)).To(Equal(3))
			})
		})

		Describe("When the documents api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Documents",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/documentsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Documents?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := documentsClient.NewDocumentsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated documents current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a document sid", func() {
		documentClient := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Document("ETXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the document is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Documents/ETXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/documentResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := documentClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get document response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("ETXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.CreatedBy).To(Equal("system"))
				Expect(resp.UniqueName).To(Equal(utils.String("test")))
				Expect(resp.Revision).To(Equal("0"))

				data := make(map[string]interface{}, 0)
				Expect(resp.Data).To(Equal(data))

				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateExpires).To(BeNil())
				Expect(resp.URL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Documents/ETXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the document api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Documents/ET71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Document("ET71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get document response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the document is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Documents/ETXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateDocumentResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &document.UpdateDocumentInput{
				Data: utils.String("{\"message\":\"Hello World\"}"),
			}

			resp, err := documentClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update document response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("ETXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.CreatedBy).To(Equal("system"))
				Expect(resp.UniqueName).To(Equal(utils.String("test")))
				Expect(resp.Revision).To(Equal("0"))

				data := make(map[string]interface{}, 0)
				data["message"] = "Hello World"
				Expect(resp.Data).To(Equal(data))

				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-20T21:00:24Z"))
				Expect(resp.DateExpires).To(BeNil())
				Expect(resp.URL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Documents/ETXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update document response returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Documents/ET71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &document.UpdateDocumentInput{
				Data: utils.String("{\"message\":\"Hello World\"}"),
			}

			resp, err := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Document("ET71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update document response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the document is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Documents/ETXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := documentClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the document api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Documents/ET71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Document("ET71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the document permissions client", func() {
		documentPermissionsClient := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Document("ETXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Permissions

		Describe("When the page of document permissions are successfully retrieved", func() {
			pageOptions := &permissions.DocumentPermissionsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Documents/ETXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Permissions?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/documentPermissionsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := documentPermissionsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the document permissions page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Documents/ETXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Permissions?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Documents/ETXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Permissions?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("permissions"))

				documentPermissions := resp.Permissions
				Expect(documentPermissions).ToNot(BeNil())
				Expect(len(documentPermissions)).To(Equal(1))

				Expect(documentPermissions[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(documentPermissions[0].ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(documentPermissions[0].DocumentSid).To(Equal("ETXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(documentPermissions[0].Identity).To(Equal("test"))
				Expect(documentPermissions[0].Read).To(Equal(true))
				Expect(documentPermissions[0].Write).To(Equal(true))
				Expect(documentPermissions[0].Manage).To(Equal(true))
				Expect(documentPermissions[0].URL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Documents/ETXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Permissions/test"))
			})
		})

		Describe("When the page of document permissions api returns a 500 response", func() {
			pageOptions := &permissions.DocumentPermissionsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Documents/ETXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Permissions?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := documentPermissionsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the document permissions page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated document permissions are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Documents/ETXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Permissions",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/documentPermissionsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Documents/ETXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Permissions?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/documentPermissionsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := documentPermissionsClient.NewDocumentPermissionsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated document permissions current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated document permissions results should be returned", func() {
				Expect(len(paginator.Permissions)).To(Equal(3))
			})
		})

		Describe("When the document permissions api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Documents/ETXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Permissions",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/documentPermissionsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Documents/ETXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Permissions?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := documentPermissionsClient.NewDocumentPermissionsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated document permissions current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a document permission identity", func() {
		documentPermissionsClient := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Document("ETXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Permission("test")

		Describe("When the document permissions are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Documents/ETXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Permissions/test",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/documentPermissionsResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := documentPermissionsClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get document permissions response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.DocumentSid).To(Equal("ETXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Identity).To(Equal("test"))
				Expect(resp.Read).To(Equal(true))
				Expect(resp.Write).To(Equal(true))
				Expect(resp.Manage).To(Equal(true))
				Expect(resp.URL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Documents/ETXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Permissions/test"))
			})
		})

		Describe("When the document permissions api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Documents/ETXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Permissions/unknownIdentity",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Document("ETXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Permission("unknownIdentity").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get document permissions response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the document permissions are successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Documents/ETXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Permissions/test",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateDocumentPermissionsResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &permission.UpdateDocumentPermissionsInput{
				Read:   true,
				Write:  false,
				Manage: false,
			}

			resp, err := documentPermissionsClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update document permissions response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.DocumentSid).To(Equal("ETXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Identity).To(Equal("test"))
				Expect(resp.Read).To(Equal(true))
				Expect(resp.Write).To(Equal(false))
				Expect(resp.Manage).To(Equal(false))
				Expect(resp.URL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Documents/ETXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Permissions/test"))
			})
		})

		Describe("When the update document permissions api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Documents/ETXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Permissions/unknownIdentity",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &permission.UpdateDocumentPermissionsInput{
				Read:   true,
				Write:  false,
				Manage: false,
			}

			resp, err := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Document("ETXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Permission("unknownIdentity").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update document permissions response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the document permissions are successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Documents/ETXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Permissions/test", httpmock.NewStringResponder(204, ""))

			err := documentPermissionsClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the document permissions api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Documents/ETXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Permissions/unknownIdentity",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Document("ETXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Permission("unknownIdentity").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the sync lists client", func() {
		syncListsClient := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").SyncLists

		Describe("When the sync list is successfully created", func() {
			createInput := &sync_lists.CreateSyncListInput{}

			httpmock.RegisterResponder("POST", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/syncListResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := syncListsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create sync list response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.CreatedBy).To(Equal("system"))
				Expect(resp.UniqueName).To(BeNil())
				Expect(resp.Revision).To(Equal("0"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateExpires).To(BeNil())
				Expect(resp.URL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists/ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the create sync list api returns a 500 response", func() {
			createInput := &sync_lists.CreateSyncListInput{}

			httpmock.RegisterResponder("POST", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := syncListsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create sync list response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of sync lists are successfully retrieved", func() {
			pageOptions := &sync_lists.SyncListsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/syncListsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := syncListsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the sync lists page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("lists"))

				syncLists := resp.SyncLists
				Expect(syncLists).ToNot(BeNil())
				Expect(len(syncLists)).To(Equal(1))

				Expect(syncLists[0].Sid).To(Equal("ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(syncLists[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(syncLists[0].ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(syncLists[0].CreatedBy).To(Equal("system"))
				Expect(syncLists[0].UniqueName).To(BeNil())
				Expect(syncLists[0].Revision).To(Equal("0"))
				Expect(syncLists[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(syncLists[0].DateUpdated).To(BeNil())
				Expect(syncLists[0].DateExpires).To(BeNil())
				Expect(syncLists[0].URL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists/ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of sync lists api returns a 500 response", func() {
			pageOptions := &sync_lists.SyncListsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := syncListsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the sync lists page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated sync lists are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/syncListsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/syncListsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := syncListsClient.NewSyncListsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated sync lists current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated sync lists results should be returned", func() {
				Expect(len(paginator.SyncLists)).To(Equal(3))
			})
		})

		Describe("When the sync lists api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/syncListsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := syncListsClient.NewSyncListsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated sync lists current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a sync list sid", func() {
		syncListClient := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").SyncList("ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the sync list is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists/ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/syncListResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := syncListClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get sync list response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.CreatedBy).To(Equal("system"))
				Expect(resp.UniqueName).To(BeNil())
				Expect(resp.Revision).To(Equal("0"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateExpires).To(BeNil())
				Expect(resp.URL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists/ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the sync list api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists/ES71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").SyncList("ES71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get sync list response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the sync list is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists/ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateSyncListResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &sync_list.UpdateSyncListInput{
				CollectionTtl: utils.Int(31536000),
			}

			resp, err := syncListClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update sync list response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.CreatedBy).To(Equal("system"))
				Expect(resp.UniqueName).To(BeNil())
				Expect(resp.Revision).To(Equal("0"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-20T21:00:24Z"))
				Expect(resp.DateExpires.Format(time.RFC3339)).To(Equal("2021-06-20T21:00:24Z"))
				Expect(resp.URL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists/ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update sync list api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists/ES71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &sync_list.UpdateSyncListInput{
				CollectionTtl: utils.Int(31536000),
			}

			resp, err := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").SyncList("ES71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update sync list response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the sync list is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists/ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := syncListClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the sync list api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists/ES71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").SyncList("ES71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the sync list items client", func() {
		syncListItemsClient := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").SyncList("ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Items

		Describe("When the sync list item is successfully created", func() {
			createInput := &syncListItems.CreateSyncListItemInput{
				Data: "{}",
			}

			httpmock.RegisterResponder("POST", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists/ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Items",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/syncListItemResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := syncListItemsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create sync list item response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Index).To(Equal(0))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ListSid).To(Equal("ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.CreatedBy).To(Equal("system"))
				Expect(resp.Revision).To(Equal("0"))

				data := make(map[string]interface{}, 0)
				Expect(resp.Data).To(Equal(data))

				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateExpires).To(BeNil())
				Expect(resp.URL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists/ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Items/0"))
			})
		})

		Describe("When the sync list item does not contain data", func() {
			createInput := &syncListItems.CreateSyncListItemInput{}

			resp, err := syncListItemsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create sync list item response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create sync list item api returns a 500 response", func() {
			createInput := &syncListItems.CreateSyncListItemInput{
				Data: "{}",
			}

			httpmock.RegisterResponder("POST", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists/ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Items",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := syncListItemsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create sync list item response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of sync list items are successfully retrieved", func() {
			pageOptions := &syncListItems.SyncListItemsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists/ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Items?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/syncListItemsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := syncListItemsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the sync list items page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists/ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Items?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists/ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Items?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("items"))

				items := resp.Items
				Expect(items).ToNot(BeNil())
				Expect(len(items)).To(Equal(1))

				Expect(items[0].Index).To(Equal(0))
				Expect(items[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(items[0].ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(items[0].ListSid).To(Equal("ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(items[0].CreatedBy).To(Equal("system"))
				Expect(items[0].Revision).To(Equal("0"))

				data := make(map[string]interface{}, 0)
				Expect(items[0].Data).To(Equal(data))

				Expect(items[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(items[0].DateUpdated).To(BeNil())
				Expect(items[0].DateExpires).To(BeNil())
				Expect(items[0].URL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists/ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Items/0"))
			})
		})

		Describe("When the page of sync list items api returns a 500 response", func() {
			pageOptions := &syncListItems.SyncListItemsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists/ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Items?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := syncListItemsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the sync list items page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated sync list items are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists/ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Items",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/syncListItemsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists/ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Items?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/syncListItemsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := syncListItemsClient.NewSyncListItemsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated sync list items current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated sync list items results should be returned", func() {
				Expect(len(paginator.Items)).To(Equal(3))
			})
		})

		Describe("When the sync list items api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists/ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Items",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/syncListItemsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists/ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Items?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := syncListItemsClient.NewSyncListItemsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated sync list items current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a sync list item index", func() {
		syncListItemClient := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").SyncList("ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Item(0)

		Describe("When the sync list item is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists/ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Items/0",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/syncListItemResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := syncListItemClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get sync list item response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Index).To(Equal(0))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ListSid).To(Equal("ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.CreatedBy).To(Equal("system"))
				Expect(resp.Revision).To(Equal("0"))

				data := make(map[string]interface{}, 0)
				Expect(resp.Data).To(Equal(data))

				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateExpires).To(BeNil())
				Expect(resp.URL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists/ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Items/0"))
			})
		})

		Describe("When the sync list item api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists/ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Items/1000",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").SyncList("ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Item(1000).Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get sync list item response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the sync list item is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists/ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Items/0",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateSyncListItemResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &syncListItem.UpdateSyncListItemInput{
				Data: utils.String("{\"message\":\"Hello World\"}"),
			}

			resp, err := syncListItemClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update sync list item response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Index).To(Equal(0))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ListSid).To(Equal("ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.CreatedBy).To(Equal("system"))
				Expect(resp.Revision).To(Equal("1"))

				data := make(map[string]interface{}, 0)
				data["message"] = "Hello World"
				Expect(resp.Data).To(Equal(data))

				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-20T21:00:24Z"))
				Expect(resp.DateExpires).To(BeNil())
				Expect(resp.URL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists/ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Items/0"))
			})
		})

		Describe("When the update sync list item api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists/ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Items/1000",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &syncListItem.UpdateSyncListItemInput{
				Data: utils.String("{\"message\":\"Hello World\"}"),
			}

			resp, err := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").SyncList("ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Item(1000).Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update sync list item response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the sync list item is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists/ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Items/0", httpmock.NewStringResponder(204, ""))

			err := syncListItemClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the sync list item api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists/ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Items/1000",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").SyncList("ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Item(1000).Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the sync list permissions client", func() {
		syncListPermissionsClient := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").SyncList("ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Permissions

		Describe("When the page of sync list permissions are successfully retrieved", func() {
			pageOptions := &syncListPermissions.SyncListPermissionsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists/ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Permissions?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/syncListPermissionsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := syncListPermissionsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the sync list permissions page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists/ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Permissions?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists/ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Permissions?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("permissions"))

				syncListPermissions := resp.Permissions
				Expect(syncListPermissions).ToNot(BeNil())
				Expect(len(syncListPermissions)).To(Equal(1))

				Expect(syncListPermissions[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(syncListPermissions[0].ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(syncListPermissions[0].ListSid).To(Equal("ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(syncListPermissions[0].Identity).To(Equal("test"))
				Expect(syncListPermissions[0].Read).To(Equal(true))
				Expect(syncListPermissions[0].Write).To(Equal(true))
				Expect(syncListPermissions[0].Manage).To(Equal(true))
				Expect(syncListPermissions[0].URL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists/ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Permissions/test"))
			})
		})

		Describe("When the page of sync list permissions api returns a 500 response", func() {
			pageOptions := &syncListPermissions.SyncListPermissionsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists/ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Permissions?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := syncListPermissionsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the sync list permissions page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated sync list permissions are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists/ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Permissions",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/syncListPermissionsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists/ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Permissions?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/syncListPermissionsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := syncListPermissionsClient.NewSyncListPermissionsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated sync list permissions current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated sync list permissions results should be returned", func() {
				Expect(len(paginator.Permissions)).To(Equal(3))
			})
		})

		Describe("When the sync list permissions api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists/ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Permissions",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/syncListPermissionsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists/ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Permissions?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := syncListPermissionsClient.NewSyncListPermissionsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated sync list permissions current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a sync list permissions identity", func() {
		syncListPermissionsClient := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").SyncList("ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Permission("test")

		Describe("When the list permissions are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists/ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Permissions/test",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/syncListPermissionsResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := syncListPermissionsClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get sync list permissions response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ListSid).To(Equal("ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Identity).To(Equal("test"))
				Expect(resp.Read).To(Equal(true))
				Expect(resp.Write).To(Equal(true))
				Expect(resp.Manage).To(Equal(true))
				Expect(resp.URL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists/ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Permissions/test"))
			})
		})

		Describe("When the sync list permissions api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists/ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Permissions/unknownIdentity",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").SyncList("ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Permission("unknownIdentity").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get sync list permissions response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the sync list permissions are successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists/ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Permissions/test",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateSyncListPermissionsResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &syncListPermission.UpdateSyncListPermissionsInput{
				Read:   true,
				Write:  false,
				Manage: false,
			}

			resp, err := syncListPermissionsClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update sync list permissions response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ListSid).To(Equal("ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Identity).To(Equal("test"))
				Expect(resp.Read).To(Equal(true))
				Expect(resp.Write).To(Equal(false))
				Expect(resp.Manage).To(Equal(false))
				Expect(resp.URL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists/ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Permissions/test"))
			})
		})

		Describe("When the update sync list permissions api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists/ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Permissions/unknownIdentity",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &syncListPermission.UpdateSyncListPermissionsInput{
				Read:   true,
				Write:  false,
				Manage: false,
			}

			resp, err := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").SyncList("ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Permission("unknownIdentity").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update sync list permissions response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the sync list permissions are successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists/ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Permissions/test", httpmock.NewStringResponder(204, ""))

			err := syncListPermissionsClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the sync list permissions api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists/ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Permissions/unknownIdentity",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").SyncList("ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Permission("unknownIdentity").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the sync maps client", func() {
		syncMapsClient := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").SyncMaps

		Describe("When the sync map is successfully created", func() {
			createInput := &sync_maps.CreateSyncMapInput{}

			httpmock.RegisterResponder("POST", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/syncMapResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := syncMapsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create sync map response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.CreatedBy).To(Equal("system"))
				Expect(resp.UniqueName).To(BeNil())
				Expect(resp.Revision).To(Equal("0"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateExpires).To(BeNil())
				Expect(resp.URL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps/MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the create sync map api returns a 500 response", func() {
			createInput := &sync_maps.CreateSyncMapInput{}

			httpmock.RegisterResponder("POST", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := syncMapsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create sync map response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of sync maps are successfully retrieved", func() {
			pageOptions := &sync_maps.SyncMapsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/syncMapsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := syncMapsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the syncMaps page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("maps"))

				syncMaps := resp.SyncMaps
				Expect(syncMaps).ToNot(BeNil())
				Expect(len(syncMaps)).To(Equal(1))

				Expect(syncMaps[0].Sid).To(Equal("MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(syncMaps[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(syncMaps[0].ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(syncMaps[0].CreatedBy).To(Equal("system"))
				Expect(syncMaps[0].UniqueName).To(BeNil())
				Expect(syncMaps[0].Revision).To(Equal("0"))
				Expect(syncMaps[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(syncMaps[0].DateUpdated).To(BeNil())
				Expect(syncMaps[0].DateExpires).To(BeNil())
				Expect(syncMaps[0].URL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps/MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of sync maps api returns a 500 response", func() {
			pageOptions := &sync_maps.SyncMapsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := syncMapsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the sync maps page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated sync maps are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/syncMapsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/syncMapsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := syncMapsClient.NewSyncMapsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated sync maps current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated sync maps results should be returned", func() {
				Expect(len(paginator.SyncMaps)).To(Equal(3))
			})
		})

		Describe("When the sync maps api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/syncMapsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := syncMapsClient.NewSyncMapsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated sync maps current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a sync map sid", func() {
		syncMapClient := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").SyncMap("MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the sync map is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps/MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/syncMapResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := syncMapClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get sync map response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.CreatedBy).To(Equal("system"))
				Expect(resp.UniqueName).To(BeNil())
				Expect(resp.Revision).To(Equal("0"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateExpires).To(BeNil())
				Expect(resp.URL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps/MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the sync map api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps/MP71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").SyncMap("MP71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get sync map response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the sync map is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps/MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateSyncMapResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &sync_map.UpdateSyncMapInput{
				CollectionTtl: utils.Int(31536000),
			}

			resp, err := syncMapClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update sync map response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.CreatedBy).To(Equal("system"))
				Expect(resp.UniqueName).To(BeNil())
				Expect(resp.Revision).To(Equal("0"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-20T21:00:24Z"))
				Expect(resp.DateExpires.Format(time.RFC3339)).To(Equal("2021-06-20T21:00:24Z"))
				Expect(resp.URL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps/MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update sync map api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps/MP71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &sync_map.UpdateSyncMapInput{
				CollectionTtl: utils.Int(31536000),
			}

			resp, err := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").SyncMap("MP71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update sync map response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the sync map is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps/MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := syncMapClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the sync map api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps/MP71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").SyncMap("MP71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the sync map items client", func() {
		syncMapItemsClient := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").SyncMap("MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Items

		Describe("When the sync map item is successfully created", func() {
			createInput := &syncMapItems.CreateSyncMapItemInput{
				Data: "{}",
				Key:  "test",
			}

			httpmock.RegisterResponder("POST", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps/MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Items",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/syncMapItemResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := syncMapItemsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create sync map item response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Key).To(Equal("key1"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.MapSid).To(Equal("MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.CreatedBy).To(Equal("system"))
				Expect(resp.Revision).To(Equal("0"))

				data := make(map[string]interface{}, 0)
				Expect(resp.Data).To(Equal(data))

				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateExpires).To(BeNil())
				Expect(resp.URL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps/MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Items/key1"))
			})
		})

		Describe("When the sync map item does not contain data", func() {
			createInput := &syncMapItems.CreateSyncMapItemInput{
				Key: "test",
			}

			resp, err := syncMapItemsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create sync map item response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the sync map item does not contain key", func() {
			createInput := &syncMapItems.CreateSyncMapItemInput{
				Data: "{}",
			}

			resp, err := syncMapItemsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create sync map item response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create sync map item api returns a 500 response", func() {
			createInput := &syncMapItems.CreateSyncMapItemInput{
				Data: "{}",
				Key:  "test",
			}

			httpmock.RegisterResponder("POST", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps/MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Items",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := syncMapItemsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create sync map item response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of sync map items are successfully retrieved", func() {
			pageOptions := &syncMapItems.SyncMapItemsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps/MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Items?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/syncMapItemsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := syncMapItemsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the sync map items page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps/MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Items?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps/MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Items?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("items"))

				items := resp.Items
				Expect(items).ToNot(BeNil())
				Expect(len(items)).To(Equal(1))

				Expect(items[0].Key).To(Equal("key1"))
				Expect(items[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(items[0].ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(items[0].MapSid).To(Equal("MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(items[0].CreatedBy).To(Equal("system"))
				Expect(items[0].Revision).To(Equal("0"))

				data := make(map[string]interface{}, 0)
				Expect(items[0].Data).To(Equal(data))

				Expect(items[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(items[0].DateUpdated).To(BeNil())
				Expect(items[0].DateExpires).To(BeNil())
				Expect(items[0].URL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps/MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Items/key1"))
			})
		})

		Describe("When the page of sync map items api returns a 500 response", func() {
			pageOptions := &syncMapItems.SyncMapItemsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps/MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Items?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := syncMapItemsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the syncMapItems page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated sync map items are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps/MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Items",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/syncMapItemsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps/MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Items?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/syncMapItemsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := syncMapItemsClient.NewSyncMapItemsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated sync map items current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated sync map items results should be returned", func() {
				Expect(len(paginator.Items)).To(Equal(3))
			})
		})

		Describe("When the sync map items api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps/MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Items",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/syncMapItemsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps/MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Items?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := syncMapItemsClient.NewSyncMapItemsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated sync map items current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a sync map item key", func() {
		syncMapItemClient := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").SyncMap("MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Item("key1")

		Describe("When the sync map item is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps/MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Items/key1",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/syncMapItemResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := syncMapItemClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get sync map item response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Key).To(Equal("key1"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.MapSid).To(Equal("MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.CreatedBy).To(Equal("system"))
				Expect(resp.Revision).To(Equal("0"))

				data := make(map[string]interface{}, 0)
				Expect(resp.Data).To(Equal(data))

				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateExpires).To(BeNil())
				Expect(resp.URL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps/MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Items/key1"))
			})
		})

		Describe("When the sync map item api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps/MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Items/unknownKey",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").SyncMap("MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Item("unknownKey").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get sync map item response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the sync map item is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps/MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Items/key1",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateSyncMapItemResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &syncMapItem.UpdateSyncMapItemInput{
				Data: utils.String("{\"message\":\"Hello World\"}"),
			}

			resp, err := syncMapItemClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update sync map item response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Key).To(Equal("key1"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.MapSid).To(Equal("MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.CreatedBy).To(Equal("system"))
				Expect(resp.Revision).To(Equal("1"))

				data := make(map[string]interface{}, 0)
				data["message"] = "Hello World"
				Expect(resp.Data).To(Equal(data))

				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-20T21:00:24Z"))
				Expect(resp.DateExpires).To(BeNil())
				Expect(resp.URL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps/MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Items/key1"))
			})
		})

		Describe("When the update sync map item api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps/MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Items/unknownKey",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &syncMapItem.UpdateSyncMapItemInput{
				Data: utils.String("{\"message\":\"Hello World\"}"),
			}

			resp, err := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").SyncMap("MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Item("unknownKey").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update sync map item response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the sync map item is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps/MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Items/key1", httpmock.NewStringResponder(204, ""))

			err := syncMapItemClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the sync map item api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps/MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Items/unknownKey",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").SyncMap("MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Item("unknownKey").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the sync map permissions client", func() {
		syncMapPermissionsClient := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").SyncMap("MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Permissions

		Describe("When the page of sync map permissions are successfully retrieved", func() {
			pageOptions := &syncMapPermissions.SyncMapPermissionsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps/MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Permissions?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/syncMapPermissionsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := syncMapPermissionsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the sync map permissions page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps/MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Permissions?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps/MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Permissions?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("permissions"))

				syncMapPermissions := resp.Permissions
				Expect(syncMapPermissions).ToNot(BeNil())
				Expect(len(syncMapPermissions)).To(Equal(1))

				Expect(syncMapPermissions[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(syncMapPermissions[0].ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(syncMapPermissions[0].MapSid).To(Equal("MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(syncMapPermissions[0].Identity).To(Equal("test"))
				Expect(syncMapPermissions[0].Read).To(Equal(true))
				Expect(syncMapPermissions[0].Write).To(Equal(true))
				Expect(syncMapPermissions[0].Manage).To(Equal(true))
				Expect(syncMapPermissions[0].URL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps/MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Permissions/test"))
			})
		})

		Describe("When the page of sync map permissions api returns a 500 response", func() {
			pageOptions := &syncMapPermissions.SyncMapPermissionsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps/MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Permissions?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := syncMapPermissionsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the sync map permissions page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated sync map permissions are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps/MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Permissions",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/syncMapPermissionsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps/MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Permissions?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/syncMapPermissionsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := syncMapPermissionsClient.NewSyncMapPermissionsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated sync map permissions current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated sync map permissions results should be returned", func() {
				Expect(len(paginator.Permissions)).To(Equal(3))
			})
		})

		Describe("When the sync map permissions api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps/MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Permissions",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/syncMapPermissionsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps/MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Permissions?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := syncMapPermissionsClient.NewSyncMapPermissionsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated sync map permissions current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a sync map permissions identity", func() {
		syncMapPermissionsClient := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").SyncMap("MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Permission("test")

		Describe("When the sync map permissions are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps/MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Permissions/test",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/syncMapPermissionsResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := syncMapPermissionsClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get sync map permissions response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.MapSid).To(Equal("MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Identity).To(Equal("test"))
				Expect(resp.Read).To(Equal(true))
				Expect(resp.Write).To(Equal(true))
				Expect(resp.Manage).To(Equal(true))
				Expect(resp.URL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps/MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Permissions/test"))
			})
		})

		Describe("When the sync map permissions api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps/MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Permissions/unknownIdentity",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").SyncMap("MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Permission("unknownIdentity").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get sync map permissions response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the sync map permissions are successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps/MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Permissions/test",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateSyncMapPermissionsResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &syncMapPermission.UpdateSyncMapPermissionsInput{
				Read:   true,
				Write:  false,
				Manage: false,
			}

			resp, err := syncMapPermissionsClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update sync map permissions response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.MapSid).To(Equal("MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Identity).To(Equal("test"))
				Expect(resp.Read).To(Equal(true))
				Expect(resp.Write).To(Equal(false))
				Expect(resp.Manage).To(Equal(false))
				Expect(resp.URL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps/MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Permissions/test"))
			})
		})

		Describe("When the update sync map permissions api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps/MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Permissions/unknownIdentity",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &syncMapPermission.UpdateSyncMapPermissionsInput{
				Read:   true,
				Write:  false,
				Manage: false,
			}

			resp, err := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").SyncMap("MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Permission("unknownIdentity").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update sync map permissions response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the sync map permissions are successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps/MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Permissions/test", httpmock.NewStringResponder(204, ""))

			err := syncMapPermissionsClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the sync map permissions api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps/MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Permissions/unknownIdentity",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").SyncMap("MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Permission("unknownIdentity").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the sync streams client", func() {
		syncStreamsClient := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").SyncStreams

		Describe("When the sync stream is successfully created", func() {
			createInput := &sync_streams.CreateSyncStreamInput{}

			httpmock.RegisterResponder("POST", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Streams",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/syncStreamResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := syncStreamsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create sync stream response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("TOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.CreatedBy).To(Equal("system"))
				Expect(resp.UniqueName).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateExpires).To(BeNil())
				Expect(resp.URL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Streams/TOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the create sync stream api returns a 500 response", func() {
			createInput := &sync_streams.CreateSyncStreamInput{}

			httpmock.RegisterResponder("POST", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Streams",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := syncStreamsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create sync stream response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of sync streams are successfully retrieved", func() {
			pageOptions := &sync_streams.SyncStreamsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Streams?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/syncStreamsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := syncStreamsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the sync streams page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Streams?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Streams?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("streams"))

				syncStreams := resp.SyncStreams
				Expect(syncStreams).ToNot(BeNil())
				Expect(len(syncStreams)).To(Equal(1))

				Expect(syncStreams[0].Sid).To(Equal("TOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(syncStreams[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(syncStreams[0].ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(syncStreams[0].CreatedBy).To(Equal("system"))
				Expect(syncStreams[0].UniqueName).To(BeNil())
				Expect(syncStreams[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(syncStreams[0].DateUpdated).To(BeNil())
				Expect(syncStreams[0].DateExpires).To(BeNil())
				Expect(syncStreams[0].URL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Streams/TOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of sync streams api returns a 500 response", func() {
			pageOptions := &sync_streams.SyncStreamsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Streams?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := syncStreamsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the sync streams page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated sync streams are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Streams",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/syncStreamsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Streams?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/syncStreamsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := syncStreamsClient.NewSyncStreamsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated sync streams current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated sync streams results should be returned", func() {
				Expect(len(paginator.SyncStreams)).To(Equal(3))
			})
		})

		Describe("When the sync streams api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Streams",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/syncStreamsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Streams?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := syncStreamsClient.NewSyncStreamsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated sync streams current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a sync stream sid", func() {
		syncStreamClient := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").SyncStream("TOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the sync stream is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Streams/TOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/syncStreamResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := syncStreamClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get sync stream response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("TOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.CreatedBy).To(Equal("system"))
				Expect(resp.UniqueName).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateExpires).To(BeNil())
				Expect(resp.URL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Streams/TOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the sync stream api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Streams/TO71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").SyncStream("TO71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get sync stream response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the sync stream is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Streams/TOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateSyncStreamResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &sync_stream.UpdateSyncStreamInput{
				Ttl: utils.Int(31536000),
			}

			resp, err := syncStreamClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update sync stream response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("TOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.CreatedBy).To(Equal("system"))
				Expect(resp.UniqueName).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-20T21:00:24Z"))
				Expect(resp.DateExpires.Format(time.RFC3339)).To(Equal("2021-06-20T21:00:24Z"))
				Expect(resp.URL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Streams/TOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update sync stream api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Streams/TO71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &sync_stream.UpdateSyncStreamInput{
				Ttl: utils.Int(31536000),
			}

			resp, err := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").SyncStream("TO71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update sync stream response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the sync stream is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Streams/TOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := syncStreamClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the sync stream api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Streams/TO71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").SyncStream("TO71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the sync streams messages client", func() {
		syncStreamMessagesClient := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").SyncStream("TOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Messages

		Describe("When the sync stream message is successfully created", func() {
			createInput := &messages.CreateSyncStreamMessageInput{
				Data: "{}",
			}

			httpmock.RegisterResponder("POST", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Streams/TOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/syncStreamMessageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := syncStreamMessagesClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create sync stream message response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("TZXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))

				data := make(map[string]interface{}, 0)
				Expect(resp.Data).To(Equal(data))
			})
		})

		Describe("When the sync stream message does not contain data", func() {
			createInput := &messages.CreateSyncStreamMessageInput{}

			resp, err := syncStreamMessagesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create sync stream message response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create sync stream message api returns a 500 response", func() {
			createInput := &messages.CreateSyncStreamMessageInput{
				Data: "{}",
			}

			httpmock.RegisterResponder("POST", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Streams/TOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := syncStreamMessagesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create sync stream message response should be nil", func() {
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
	Expect(twilioErr.Code).To(Equal(utils.Int(20404)))
	Expect(twilioErr.Message).To(Equal("The requested resource /Services/IS71 was not found"))
	Expect(twilioErr.MoreInfo).To(Equal(utils.String("https://www.twilio.com/docs/errors/20404")))
	Expect(twilioErr.Status).To(Equal(404))
}

func ExpectErrorToNotBeATwilioError(err error) {
	Expect(err).ToNot(BeNil())
	_, ok := err.(*utils.TwilioError)
	Expect(ok).To(Equal(false))
}
