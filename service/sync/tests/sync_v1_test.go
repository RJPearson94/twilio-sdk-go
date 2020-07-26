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

	"github.com/RJPearson94/twilio-sdk-go/service/sync"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/document"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/document/permissions"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/documents"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/sync_list"
	syncListItem "github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/sync_list/item"
	syncListItems "github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/sync_list/items"
	syncListPermissions "github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/sync_list/permissions"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/sync_lists"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/sync_map"
	syncMapItem "github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/sync_map/item"
	syncMapItems "github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/sync_map/items"
	syncMapPermissions "github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/sync_map/permissions"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/sync_maps"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/sync_stream"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/sync_stream/messages"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/sync_streams"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/services"
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

	syncSession := sync.NewWithCredentials(creds).V1

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

			resp, err := serviceClient.Get()
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

			resp, err := syncSession.Service("IS71").Get()
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

			resp, err := documentClient.Get()
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

			resp, err := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Document("ET71").Get()
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

	Describe("Given I have a document permissions identity", func() {
		documentPermissionsClient := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Document("ETXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Permissions("test")

		Describe("When the document permissions are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Documents/ETXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Permissions/test",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/documentPermissionsResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := documentPermissionsClient.Get()
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

			resp, err := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Document("ETXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Permissions("unknownIdentity").Get()
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

			updateInput := &permissions.UpdateDocumentPermissionsInput{
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

			updateInput := &permissions.UpdateDocumentPermissionsInput{
				Read:   true,
				Write:  false,
				Manage: false,
			}

			resp, err := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Document("ETXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Permissions("unknownIdentity").Update(updateInput)
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

			err := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Document("ETXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Permissions("unknownIdentity").Delete()
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

			resp, err := syncListClient.Get()
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

			resp, err := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").SyncList("ES71").Get()
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
	})

	Describe("Given I have a sync list item index", func() {
		syncListItemClient := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").SyncList("ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Item(0)

		Describe("When the sync list item is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists/ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Items/0",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/synclistItemResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := syncListItemClient.Get()
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

			resp, err := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").SyncList("ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Item(1000).Get()
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

	Describe("Given I have a sync list permissions identity", func() {
		syncListPermissionsClient := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").SyncList("ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Permissions("test")

		Describe("When the list permissions are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Lists/ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Permissions/test",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/syncListPermissionsResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := syncListPermissionsClient.Get()
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

			resp, err := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").SyncList("ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Permissions("unknownIdentity").Get()
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

			updateInput := &syncListPermissions.UpdateSyncListPermissionsInput{
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

			updateInput := &syncListPermissions.UpdateSyncListPermissionsInput{
				Read:   true,
				Write:  false,
				Manage: false,
			}

			resp, err := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").SyncList("ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Permissions("unknownIdentity").Update(updateInput)
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

			err := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").SyncList("ESXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Permissions("unknownIdentity").Delete()
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

			resp, err := syncMapClient.Get()
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

			resp, err := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").SyncMap("MP71").Get()
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
			createInput := &syncMapItems.CreateSyncMapItemInput{}

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

			resp, err := syncMapItemClient.Get()
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

			resp, err := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").SyncMap("MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Item("unknownKey").Get()
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

	Describe("Given I have a sync map permissions identity", func() {
		syncMapPermissionsClient := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").SyncMap("MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Permissions("test")

		Describe("When the sync map permissions are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Maps/MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Permissions/test",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/syncMapPermissionsResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := syncMapPermissionsClient.Get()
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

			resp, err := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").SyncMap("MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Permissions("unknownIdentity").Get()
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

			updateInput := &syncMapPermissions.UpdateSyncMapPermissionsInput{
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

			updateInput := &syncMapPermissions.UpdateSyncMapPermissionsInput{
				Read:   true,
				Write:  false,
				Manage: false,
			}

			resp, err := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").SyncMap("MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Permissions("unknownIdentity").Update(updateInput)
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

			err := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").SyncMap("MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Permissions("unknownIdentity").Delete()
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

			resp, err := syncStreamClient.Get()
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

			resp, err := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").SyncStream("TO71").Get()
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
