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
	"github.com/RJPearson94/twilio-sdk-go/service/trunking"
	"github.com/RJPearson94/twilio-sdk-go/service/trunking/v1/trunk"
	"github.com/RJPearson94/twilio-sdk-go/service/trunking/v1/trunk/credential_lists"
	"github.com/RJPearson94/twilio-sdk-go/service/trunking/v1/trunk/ip_access_control_lists"
	"github.com/RJPearson94/twilio-sdk-go/service/trunking/v1/trunk/origination_url"
	"github.com/RJPearson94/twilio-sdk-go/service/trunking/v1/trunk/origination_urls"
	"github.com/RJPearson94/twilio-sdk-go/service/trunking/v1/trunk/phone_number"
	"github.com/RJPearson94/twilio-sdk-go/service/trunking/v1/trunk/phone_numbers"
	"github.com/RJPearson94/twilio-sdk-go/service/trunking/v1/trunk/recording"
	"github.com/RJPearson94/twilio-sdk-go/service/trunking/v1/trunks"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

var _ = Describe("Trunking V1", func() {
	creds, err := credentials.New(credentials.Account{
		Sid:       "ACxxx",
		AuthToken: "Test",
	})
	if err != nil {
		log.Panicf("%s", err)
	}

	trunkingClient := trunking.New(session.New(creds), &client.Config{
		RetryAttempts: utils.Int(0),
	}).V1

	httpmock.ActivateNonDefault(trunkingClient.GetClient().GetRestyClient().GetClient())
	defer httpmock.DeactivateAndReset()

	Describe("Given the Elastic SIP Trunk Client", func() {
		trunksClient := trunkingClient.Trunks

		Describe("When the trunk is successfully created", func() {
			createInput := &trunks.CreateTrunkInput{}

			httpmock.RegisterResponder("POST", "https://trunking.twilio.com/v1/Trunks",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/trunkResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := trunksClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create trunk response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AuthType).To(Equal(utils.String("")))
				Expect(resp.AuthTypeSet).To(Equal(&[]string{}))
				Expect(resp.TransferMode).To(Equal("disable-all"))
				Expect(resp.Secure).To(Equal(false))
				Expect(resp.CnamLookupEnabled).To(Equal(false))
				Expect(resp.FriendlyName).To(BeNil())
				Expect(resp.DomainName).To(BeNil())
				Expect(resp.Recording).To(Equal(trunks.CreateTrunkRecordingResponse{
					Trim: "do-not-trim",
					Mode: "do-not-record",
				}))
				Expect(resp.DisasterRecoveryURL).To(BeNil())
				Expect(resp.DisasterRecoveryMethod).To(BeNil())
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.URL).To(Equal("https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the Create Trunk API returns a 500 response", func() {
			createInput := &trunks.CreateTrunkInput{}

			httpmock.RegisterResponder("POST", "https://trunking.twilio.com/v1/Trunks",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := trunksClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create trunk response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of trunks are successfully retrieved", func() {
			pageOptions := &trunks.TrunksPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://trunking.twilio.com/v1/Trunks?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/trunksPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := trunksClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the trunks page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://trunking.twilio.com/v1/Trunks?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://trunking.twilio.com/v1/Trunks?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("trunks"))

				trunkRecording := trunks.PageTrunkRecordingResponse{
					Trim: "do-not-trim",
					Mode: "do-not-record",
				}

				trunks := resp.Trunks
				Expect(trunks).ToNot(BeNil())
				Expect(len(trunks)).To(Equal(1))

				Expect(trunks[0].Sid).To(Equal("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(trunks[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(trunks[0].AuthType).To(Equal(utils.String("")))
				Expect(trunks[0].AuthTypeSet).To(Equal(&[]string{}))
				Expect(trunks[0].TransferMode).To(Equal("disable-all"))
				Expect(trunks[0].Secure).To(Equal(false))
				Expect(trunks[0].CnamLookupEnabled).To(Equal(false))
				Expect(trunks[0].FriendlyName).To(BeNil())
				Expect(trunks[0].DomainName).To(BeNil())
				Expect(trunks[0].Recording).To(Equal(trunkRecording))
				Expect(trunks[0].DisasterRecoveryURL).To(BeNil())
				Expect(trunks[0].DisasterRecoveryMethod).To(BeNil())
				Expect(trunks[0].DateUpdated).To(BeNil())
				Expect(trunks[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(trunks[0].URL).To(Equal("https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))

			})
		})

		Describe("When the page of trunks api returns a 500 response", func() {
			pageOptions := &trunks.TrunksPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://trunking.twilio.com/v1/Trunks?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := trunksClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the trunks page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated trunks are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://trunking.twilio.com/v1/Trunks",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/trunksPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://trunking.twilio.com/v1/Trunks?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/trunksPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := trunksClient.NewTrunksPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated trunks current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated trunks results should be returned", func() {
				Expect(len(paginator.Trunks)).To(Equal(3))
			})
		})

		Describe("When the trunks api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://trunking.twilio.com/v1/Trunks",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/trunksPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://trunking.twilio.com/v1/Trunks?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := trunksClient.NewTrunksPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated trunks current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a Trunk SID", func() {
		trunkClient := trunkingClient.Trunk("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the trunk is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/trunkResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := trunkClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get trunk response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AuthType).To(Equal(utils.String("")))
				Expect(resp.AuthTypeSet).To(Equal(&[]string{}))
				Expect(resp.TransferMode).To(Equal("disable-all"))
				Expect(resp.Secure).To(Equal(false))
				Expect(resp.CnamLookupEnabled).To(Equal(false))
				Expect(resp.FriendlyName).To(BeNil())
				Expect(resp.DomainName).To(BeNil())
				Expect(resp.Recording).To(Equal(trunk.FetchTrunkRecordingResponse{
					Trim: "do-not-trim",
					Mode: "do-not-record",
				}))
				Expect(resp.DisasterRecoveryURL).To(BeNil())
				Expect(resp.DisasterRecoveryMethod).To(BeNil())
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.URL).To(Equal("https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the get trunk response returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://trunking.twilio.com/v1/Trunks/TK71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := trunkingClient.Trunk("TK71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get trunk response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the trunk is successfully updated", func() {

			httpmock.RegisterResponder("POST", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateTrunkResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &trunk.UpdateTrunkInput{}

			resp, err := trunkClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update trunk response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AuthType).To(Equal(utils.String("")))
				Expect(resp.AuthTypeSet).To(Equal(&[]string{}))
				Expect(resp.TransferMode).To(Equal("disable-all"))
				Expect(resp.Secure).To(Equal(false))
				Expect(resp.CnamLookupEnabled).To(Equal(false))
				Expect(resp.FriendlyName).To(BeNil())
				Expect(resp.DomainName).To(BeNil())
				Expect(resp.Recording).To(Equal(trunk.UpdateTrunkRecordingResponse{
					Trim: "do-not-trim",
					Mode: "do-not-record",
				}))
				Expect(resp.DisasterRecoveryURL).To(BeNil())
				Expect(resp.DisasterRecoveryMethod).To(BeNil())
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-27T23:10:00Z"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.URL).To(Equal("https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update trunk response returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://trunking.twilio.com/v1/Trunks/TK71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &trunk.UpdateTrunkInput{}

			resp, err := trunkingClient.Trunk("TK71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update trunk response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the trunk is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := trunkClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the delete trunk response returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://trunking.twilio.com/v1/Trunks/TK71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := trunkingClient.Trunk("TK71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the Origination URL Client", func() {
		originationURLsClient := trunkingClient.Trunk("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").OriginationURLs

		Describe("When the origination url is successfully created", func() {
			createInput := &origination_urls.CreateOriginationURLInput{
				Weight:       0,
				Priority:     0,
				Enabled:      false,
				FriendlyName: "test",
				SipURL:       "sip:test@test.com",
			}

			httpmock.RegisterResponder("POST", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/OriginationUrls",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/originationUrlResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := originationURLsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create origination url response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("OUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.TrunkSid).To(Equal("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Weight).To(Equal(0))
				Expect(resp.FriendlyName).To(Equal("test"))
				Expect(resp.Enabled).To(Equal(true))
				Expect(resp.Priority).To(Equal(0))
				Expect(resp.SipURL).To(Equal("sip:test@test.com"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.URL).To(Equal("https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/OriginationUrls/OUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the origination url does not contain a friendly name", func() {
			createInput := &origination_urls.CreateOriginationURLInput{
				Weight:       0,
				Priority:     0,
				Enabled:      false,
				FriendlyName: "test",
			}

			resp, err := originationURLsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create origination url  response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the origination url does not contain a sip url", func() {
			createInput := &origination_urls.CreateOriginationURLInput{
				Weight:   0,
				Priority: 0,
				Enabled:  false,
				SipURL:   "sip:test@test.com",
			}

			resp, err := originationURLsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create origination url  response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the Create Origination URL API returns a 500 response", func() {
			createInput := &origination_urls.CreateOriginationURLInput{
				Weight:       0,
				Priority:     0,
				Enabled:      false,
				FriendlyName: "test",
				SipURL:       "sip:test@test.com",
			}

			httpmock.RegisterResponder("POST", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/OriginationUrls",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := originationURLsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create origination url response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of origination urls are successfully retrieved", func() {
			pageOptions := &origination_urls.OriginationURLsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/OriginationUrls?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/originationUrlsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := originationURLsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the origination urls  page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/OriginationUrls?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/OriginationUrls?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("origination_urls"))

				originationURLs := resp.OriginationURLs
				Expect(originationURLs).ToNot(BeNil())
				Expect(len(originationURLs)).To(Equal(1))

				Expect(originationURLs[0].Sid).To(Equal("OUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(originationURLs[0].TrunkSid).To(Equal("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(originationURLs[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(originationURLs[0].Weight).To(Equal(0))
				Expect(originationURLs[0].FriendlyName).To(Equal("test"))
				Expect(originationURLs[0].Enabled).To(Equal(true))
				Expect(originationURLs[0].Priority).To(Equal(0))
				Expect(originationURLs[0].SipURL).To(Equal("sip:test@test.com"))
				Expect(originationURLs[0].DateUpdated).To(BeNil())
				Expect(originationURLs[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(originationURLs[0].URL).To(Equal("https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/OriginationUrls/OUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))

			})
		})

		Describe("When the page of origination urls api returns a 500 response", func() {
			pageOptions := &origination_urls.OriginationURLsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/OriginationUrls?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := originationURLsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the origination urls page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated origination urls are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/OriginationUrls",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/originationUrlsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/OriginationUrls?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/originationUrlsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := originationURLsClient.NewOriginationURLsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated origination urls current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated origination urls results should be returned", func() {
				Expect(len(paginator.OriginationURLs)).To(Equal(3))
			})
		})

		Describe("When the origination urls api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/OriginationUrls",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/originationUrlsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/OriginationUrls?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := originationURLsClient.NewOriginationURLsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated origination urls current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a Origination URL SID", func() {
		originationURLClient := trunkingClient.Trunk("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").OriginationURL("OUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the origination url is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/OriginationUrls/OUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/originationUrlResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := originationURLClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get origination url response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("OUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.TrunkSid).To(Equal("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Weight).To(Equal(0))
				Expect(resp.FriendlyName).To(Equal("test"))
				Expect(resp.Enabled).To(Equal(true))
				Expect(resp.Priority).To(Equal(0))
				Expect(resp.SipURL).To(Equal("sip:test@test.com"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.URL).To(Equal("https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/OriginationUrls/OUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the get origination url response returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/OriginationUrls/OU71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := trunkingClient.Trunk("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").OriginationURL("OU71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get origination url response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the origination url is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/OriginationUrls/OUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateOriginationUrlResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &origination_url.UpdateOriginationURLInput{}

			resp, err := originationURLClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update origination url response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("OUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.TrunkSid).To(Equal("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Weight).To(Equal(0))
				Expect(resp.FriendlyName).To(Equal("test"))
				Expect(resp.Enabled).To(Equal(true))
				Expect(resp.Priority).To(Equal(0))
				Expect(resp.SipURL).To(Equal("sip:test@test.com"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-27T23:10:00Z"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.URL).To(Equal("https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/OriginationUrls/OUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update origination url response returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/OriginationUrls/OU71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &origination_url.UpdateOriginationURLInput{}

			resp, err := trunkingClient.Trunk("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").OriginationURL("OU71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update origination url response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the origination url is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/OriginationUrls/OUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := originationURLClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the delete origination url response returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/OriginationUrls/OU71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := trunkingClient.Trunk("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").OriginationURL("OU71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the Phone Number Client", func() {
		phoneNumbersClient := trunkingClient.Trunk("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").PhoneNumbers

		Describe("When the  is successfully created", func() {
			createInput := &phone_numbers.CreatePhoneNumberInput{
				PhoneNumberSid: "PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			}

			httpmock.RegisterResponder("POST", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/PhoneNumbers",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/phoneNumberResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := phoneNumbersClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create  response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal(utils.String("441234567890")))
				Expect(resp.PhoneNumber).To(Equal("+441234567890"))
				Expect(resp.APIVersion).To(Equal("2010-04-01"))
				Expect(resp.SmsFallbackMethod).To(Equal("POST"))
				Expect(resp.SmsFallbackURL).To(BeNil())
				Expect(resp.SmsMethod).To(Equal("POST"))
				Expect(resp.SmsURL).To(BeNil())
				Expect(resp.SmsApplicationSid).To(BeNil())
				Expect(resp.VoiceApplicationSid).To(BeNil())
				Expect(resp.VoiceReceiveMode).To(BeNil())
				Expect(resp.TrunkSid).To(Equal("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Beta).To(Equal(false))
				Expect(resp.Capabilities).To(Equal(phone_numbers.CreatePhoneNumberCapabilitiesResponse{
					Mms:   false,
					Sms:   true,
					Voice: true,
				}))
				Expect(resp.StatusCallback).To(BeNil())
				Expect(resp.StatusCallbackMethod).To(Equal("POST"))
				Expect(resp.VoiceCallerIDLookup).To(Equal(false))
				Expect(resp.VoiceFallbackMethod).To(Equal("POST"))
				Expect(resp.VoiceFallbackURL).To(BeNil())
				Expect(resp.VoiceMethod).To(Equal("POST"))
				Expect(resp.VoiceURL).To(BeNil())
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.URL).To(Equal("https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/PhoneNumbers/PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the  does not contain a sip url", func() {
			createInput := &phone_numbers.CreatePhoneNumberInput{}

			resp, err := phoneNumbersClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create  response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create  API returns a 500 response", func() {
			createInput := &phone_numbers.CreatePhoneNumberInput{
				PhoneNumberSid: "PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			}

			httpmock.RegisterResponder("POST", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/PhoneNumbers",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := phoneNumbersClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create  response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of s are successfully retrieved", func() {
			pageOptions := &phone_numbers.PhoneNumbersPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/PhoneNumbers?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/phoneNumbersPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := phoneNumbersClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the  page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/PhoneNumbers?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/PhoneNumbers?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("phone_numbers"))

				phoneNumbers := resp.PhoneNumbers
				Expect(phoneNumbers).ToNot(BeNil())
				Expect(len(phoneNumbers)).To(Equal(1))

				Expect(phoneNumbers[0].Sid).To(Equal("PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(phoneNumbers[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(phoneNumbers[0].FriendlyName).To(Equal(utils.String("441234567890")))
				Expect(phoneNumbers[0].PhoneNumber).To(Equal("+441234567890"))
				Expect(phoneNumbers[0].APIVersion).To(Equal("2010-04-01"))
				Expect(phoneNumbers[0].SmsFallbackMethod).To(Equal("POST"))
				Expect(phoneNumbers[0].SmsFallbackURL).To(BeNil())
				Expect(phoneNumbers[0].SmsMethod).To(Equal("POST"))
				Expect(phoneNumbers[0].SmsURL).To(BeNil())
				Expect(phoneNumbers[0].SmsApplicationSid).To(BeNil())
				Expect(phoneNumbers[0].VoiceApplicationSid).To(BeNil())
				Expect(phoneNumbers[0].VoiceReceiveMode).To(BeNil())
				Expect(phoneNumbers[0].TrunkSid).To(Equal("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(phoneNumbers[0].Beta).To(Equal(false))
				Expect(phoneNumbers[0].Capabilities).To(Equal(phone_numbers.PagePhoneNumberCapabilitiesResponse{
					Mms:   false,
					Sms:   true,
					Voice: true,
				}))
				Expect(phoneNumbers[0].StatusCallback).To(BeNil())
				Expect(phoneNumbers[0].StatusCallbackMethod).To(Equal("POST"))
				Expect(phoneNumbers[0].VoiceCallerIDLookup).To(Equal(false))
				Expect(phoneNumbers[0].VoiceFallbackMethod).To(Equal("POST"))
				Expect(phoneNumbers[0].VoiceFallbackURL).To(BeNil())
				Expect(phoneNumbers[0].VoiceMethod).To(Equal("POST"))
				Expect(phoneNumbers[0].VoiceURL).To(BeNil())
				Expect(phoneNumbers[0].DateUpdated).To(BeNil())
				Expect(phoneNumbers[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(phoneNumbers[0].URL).To(Equal("https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/PhoneNumbers/PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of s api returns a 500 response", func() {
			pageOptions := &phone_numbers.PhoneNumbersPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/PhoneNumbers?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := phoneNumbersClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the s page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated s are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/PhoneNumbers",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/phoneNumbersPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/PhoneNumbers?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/phoneNumbersPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := phoneNumbersClient.NewPhoneNumbersPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated s current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated s results should be returned", func() {
				Expect(len(paginator.PhoneNumbers)).To(Equal(3))
			})
		})

		Describe("When the s api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/PhoneNumbers",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/phoneNumbersPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/PhoneNumbers?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := phoneNumbersClient.NewPhoneNumbersPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated s current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a Phone Number SID", func() {
		phoneNumberClient := trunkingClient.Trunk("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").PhoneNumber("PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the  is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/PhoneNumbers/PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/phoneNumberResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := phoneNumberClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get  response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal(utils.String("441234567890")))
				Expect(resp.PhoneNumber).To(Equal("+441234567890"))
				Expect(resp.APIVersion).To(Equal("2010-04-01"))
				Expect(resp.SmsFallbackMethod).To(Equal("POST"))
				Expect(resp.SmsFallbackURL).To(BeNil())
				Expect(resp.SmsMethod).To(Equal("POST"))
				Expect(resp.SmsURL).To(BeNil())
				Expect(resp.SmsApplicationSid).To(BeNil())
				Expect(resp.VoiceApplicationSid).To(BeNil())
				Expect(resp.VoiceReceiveMode).To(BeNil())
				Expect(resp.TrunkSid).To(Equal("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Beta).To(Equal(false))
				Expect(resp.Capabilities).To(Equal(phone_number.FetchPhoneNumberCapabilitiesResponse{
					Mms:   false,
					Sms:   true,
					Voice: true,
				}))
				Expect(resp.StatusCallback).To(BeNil())
				Expect(resp.StatusCallbackMethod).To(Equal("POST"))
				Expect(resp.VoiceCallerIDLookup).To(Equal(false))
				Expect(resp.VoiceFallbackMethod).To(Equal("POST"))
				Expect(resp.VoiceFallbackURL).To(BeNil())
				Expect(resp.VoiceMethod).To(Equal("POST"))
				Expect(resp.VoiceURL).To(BeNil())
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.URL).To(Equal("https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/PhoneNumbers/PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the get  response returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/PhoneNumbers/PN71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := trunkingClient.Trunk("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").PhoneNumber("PN71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get  response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the  is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/PhoneNumbers/PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := phoneNumberClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the delete  response returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/PhoneNumbers/PN71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := trunkingClient.Trunk("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").PhoneNumber("PN71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given I have a Recording client", func() {
		recordingClient := trunkingClient.Trunk("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Recording()

		Describe("When the  is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recording",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/recordingResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := recordingClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get recording response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Trim).To(Equal("do-not-trim"))
				Expect(resp.Mode).To(Equal("do-not-record"))
			})
		})

		Describe("When the get recording response returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recording",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := trunkingClient.Trunk("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Recording().Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get recording response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the recording is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recording",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateRecordingResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &recording.UpdateRecordingInput{}

			resp, err := recordingClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update recording response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Trim).To(Equal("trim-silence"))
				Expect(resp.Mode).To(Equal("do-not-record"))
			})
		})

		Describe("When the update recording response returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recording",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &recording.UpdateRecordingInput{}

			resp, err := trunkingClient.Trunk("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Recording().Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update recording response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given the Credential Lists Client", func() {
		credentialListsClient := trunkingClient.Trunk("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").CredentialLists

		Describe("When the credential list is successfully created", func() {
			createInput := &credential_lists.CreateCredentialListInput{
				CredentialListSid: "CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			}

			httpmock.RegisterResponder("POST", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/CredentialLists",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/credentialListResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := credentialListsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create credential list response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.TrunkSid).To(Equal("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.URL).To(Equal("https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/CredentialLists/CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the credential list does not contain a sip url", func() {
			createInput := &credential_lists.CreateCredentialListInput{}

			resp, err := credentialListsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create credential list response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create credential list API returns a 500 response", func() {
			createInput := &credential_lists.CreateCredentialListInput{
				CredentialListSid: "CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			}

			httpmock.RegisterResponder("POST", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/CredentialLists",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := credentialListsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create credential list response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of credential lists are successfully retrieved", func() {
			pageOptions := &credential_lists.CredentialListsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/CredentialLists?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/credentialListsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := credentialListsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the credential list page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/CredentialLists?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/CredentialLists?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("credential_lists"))

				credentialLists := resp.CredentialLists
				Expect(credentialLists).ToNot(BeNil())
				Expect(len(credentialLists)).To(Equal(1))

				Expect(credentialLists[0].Sid).To(Equal("CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(credentialLists[0].TrunkSid).To(Equal("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(credentialLists[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(credentialLists[0].FriendlyName).To(Equal("Test"))
				Expect(credentialLists[0].DateUpdated).To(BeNil())
				Expect(credentialLists[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(credentialLists[0].URL).To(Equal("https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/CredentialLists/CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of credential lists api returns a 500 response", func() {
			pageOptions := &credential_lists.CredentialListsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/CredentialLists?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := credentialListsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the credential lists page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated credential lists are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/CredentialLists",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/credentialListsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/CredentialLists?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/credentialListsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := credentialListsClient.NewCredentialListsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated credential lists current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated credential lists results should be returned", func() {
				Expect(len(paginator.CredentialLists)).To(Equal(3))
			})
		})

		Describe("When the credential lists api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/CredentialLists",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/phoneNumbersPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/CredentialLists?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := credentialListsClient.NewCredentialListsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated credential lists current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a Credential List SID", func() {
		credentialListClient := trunkingClient.Trunk("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").CredentialList("CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the credential list is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/CredentialLists/CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/credentialListResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := credentialListClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get credential list response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.TrunkSid).To(Equal("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.URL).To(Equal("https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/CredentialLists/CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the get credential list response returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/CredentialLists/CL71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := trunkingClient.Trunk("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").CredentialList("CL71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get Credential List response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the credential list is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/CredentialLists/CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := credentialListClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the delete credential list response returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/CredentialLists/CL71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := trunkingClient.Trunk("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").CredentialList("CL71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the IP Access Control Lists Client", func() {
		ipAccessControlListsClient := trunkingClient.Trunk("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").IpAccessControlLists

		Describe("When the IP access control list is successfully created", func() {
			createInput := &ip_access_control_lists.CreateIpAccessControlListInput{
				IpAccessControlListSid: "ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			}

			httpmock.RegisterResponder("POST", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IpAccessControlLists",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/ipAccessControlListResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := ipAccessControlListsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create IP access control list response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.TrunkSid).To(Equal("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.URL).To(Equal("https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IpAccessControlLists/ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the IP access control list does not contain a sip url", func() {
			createInput := &ip_access_control_lists.CreateIpAccessControlListInput{}

			resp, err := ipAccessControlListsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create IP access control list response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create IP access control list API returns a 500 response", func() {
			createInput := &ip_access_control_lists.CreateIpAccessControlListInput{
				IpAccessControlListSid: "ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			}

			httpmock.RegisterResponder("POST", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IpAccessControlLists",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := ipAccessControlListsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create IP access control list response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of IP access control lists are successfully retrieved", func() {
			pageOptions := &ip_access_control_lists.IpAccessControlListsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IpAccessControlLists?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/ipAccessControlListsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := ipAccessControlListsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the IP access control list page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IpAccessControlLists?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IpAccessControlLists?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("ip_access_control_lists"))

				ipAccessControlLists := resp.IpAccessControlLists
				Expect(ipAccessControlLists).ToNot(BeNil())
				Expect(len(ipAccessControlLists)).To(Equal(1))

				Expect(ipAccessControlLists[0].Sid).To(Equal("ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(ipAccessControlLists[0].TrunkSid).To(Equal("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(ipAccessControlLists[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(ipAccessControlLists[0].FriendlyName).To(Equal("Test"))
				Expect(ipAccessControlLists[0].DateUpdated).To(BeNil())
				Expect(ipAccessControlLists[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(ipAccessControlLists[0].URL).To(Equal("https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IpAccessControlLists/ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of IP access control lists api returns a 500 response", func() {
			pageOptions := &ip_access_control_lists.IpAccessControlListsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IpAccessControlLists?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := ipAccessControlListsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the IP access control lists page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated IP access control lists are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IpAccessControlLists",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/ipAccessControlListsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IpAccessControlLists?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/ipAccessControlListsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := ipAccessControlListsClient.NewIpAccessControlListsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated IP access control lists current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated IP access control lists results should be returned", func() {
				Expect(len(paginator.IpAccessControlLists)).To(Equal(3))
			})
		})

		Describe("When the IP access control lists api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IpAccessControlLists",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/phoneNumbersPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IpAccessControlLists?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := ipAccessControlListsClient.NewIpAccessControlListsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated IP access control lists current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a IP Access Control List SID", func() {
		ipAccessControlListClient := trunkingClient.Trunk("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").IpAccessControlList("ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the IP access control list is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IpAccessControlLists/ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/ipAccessControlListResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := ipAccessControlListClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get IP access control list response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.TrunkSid).To(Equal("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.URL).To(Equal("https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IpAccessControlLists/ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the get IP access control list response returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IpAccessControlLists/AL71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := trunkingClient.Trunk("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").IpAccessControlList("AL71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get IP Access Control List response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the IP access control list is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IpAccessControlLists/ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := ipAccessControlListClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the delete IP access control list response returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IpAccessControlLists/AL71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := trunkingClient.Trunk("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").IpAccessControlList("AL71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
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
	Expect(twilioErr.Message).To(Equal("The requested resource /Trunks/TK71 was not found"))

	moreInfo := "https://www.twilio.com/docs/errors/20404"
	Expect(twilioErr.MoreInfo).To(Equal(&moreInfo))
	Expect(twilioErr.Status).To(Equal(404))
}

func ExpectErrorToNotBeATwilioError(err error) {
	Expect(err).ToNot(BeNil())
	_, ok := err.(*utils.TwilioError)
	Expect(ok).To(Equal(false))
}
