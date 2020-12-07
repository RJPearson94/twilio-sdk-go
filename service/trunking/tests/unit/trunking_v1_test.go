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

	"github.com/RJPearson94/twilio-sdk-go/service/trunking"
	"github.com/RJPearson94/twilio-sdk-go/service/trunking/v1/trunk"
	"github.com/RJPearson94/twilio-sdk-go/service/trunking/v1/trunks"
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

	trunkingSession := trunking.NewWithCredentials(creds).V1

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
