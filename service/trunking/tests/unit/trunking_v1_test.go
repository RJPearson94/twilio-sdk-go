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

	trunkingSession := trunking.New(session.New(creds), &client.Config{
		RetryAttempts: utils.Int(0),
	}).V1

	httpmock.ActivateNonDefault(trunkingSession.GetClient().GetRestyClient().GetClient())
	defer httpmock.DeactivateAndReset()

	Describe("Given the Elastic SIP Trunk Client", func() {
		trunksClient := trunkingSession.Trunks

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
		trunkClient := trunkingSession.Trunk("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

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

			resp, err := trunkingSession.Trunk("TK71").Fetch()
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

			resp, err := trunkingSession.Trunk("TK71").Update(updateInput)
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

			err := trunkingSession.Trunk("TK71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the Origination URL Client", func() {
		originationURLsClient := trunkingSession.Trunk("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").OriginationURLs

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
		originationURLClient := trunkingSession.Trunk("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").OriginationURL("OUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

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

			resp, err := trunkingSession.Trunk("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").OriginationURL("OU71").Fetch()
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

			resp, err := trunkingSession.Trunk("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").OriginationURL("OU71").Update(updateInput)
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

			err := trunkingSession.Trunk("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").OriginationURL("OU71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the Phone Number Client", func() {
		phoneNumbersClient := trunkingSession.Trunk("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").PhoneNumbers

		Describe("When the phone number is successfully created", func() {
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

			It("Then the create phone number response should be returned", func() {
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

		Describe("When the phone number does not contain a sip url", func() {
			createInput := &phone_numbers.CreatePhoneNumberInput{}

			resp, err := phoneNumbersClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create phone number response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create phone number API returns a 500 response", func() {
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

			It("Then the create phone number response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of phone numbers are successfully retrieved", func() {
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

			It("Then the phone number page response should be returned", func() {
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

		Describe("When the page of phone numbers api returns a 500 response", func() {
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

			It("Then the phone numbers page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated phone numbers are successfully retrieved", func() {
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

			It("Then the paginated phone numbers current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated phone numbers results should be returned", func() {
				Expect(len(paginator.PhoneNumbers)).To(Equal(3))
			})
		})

		Describe("When the phone numbers api returns a 500 response when making a paginated request", func() {
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

			It("Then the paginated phone numbers current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a Phone Number SID", func() {
		phoneNumberClient := trunkingSession.Trunk("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").PhoneNumber("PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the phone number is successfully retrieved", func() {
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

			It("Then the get phone number response should be returned", func() {
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

		Describe("When the get phone number response returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/PhoneNumbers/PN71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := trunkingSession.Trunk("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").PhoneNumber("PN71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get phone number response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the phone number is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/PhoneNumbers/PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := phoneNumberClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the delete phone number response returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://trunking.twilio.com/v1/Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/PhoneNumbers/PN71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := trunkingSession.Trunk("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").PhoneNumber("PN71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given I have a Recording client", func() {
		recordingClient := trunkingSession.Trunk("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Recording()

		Describe("When the phone number is successfully retrieved", func() {
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

			resp, err := trunkingSession.Trunk("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Recording().Fetch()
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

			resp, err := trunkingSession.Trunk("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Recording().Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update recording response should be nil", func() {
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
