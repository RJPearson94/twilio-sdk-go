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

	"github.com/RJPearson94/twilio-sdk-go/service/monitor"
	"github.com/RJPearson94/twilio-sdk-go/service/monitor/v1/alerts"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

var _ = Describe("Monitor V1", func() {
	creds, err := credentials.New(credentials.Account{
		Sid:       "ACxxx",
		AuthToken: "Test",
	})
	if err != nil {
		log.Panicf("%s", err)
	}

	monitorSession := monitor.NewWithCredentials(creds).V1

	httpmock.ActivateNonDefault(monitorSession.GetClient().GetRestyClient().GetClient())
	defer httpmock.DeactivateAndReset()

	Describe("Given the alerts client", func() {
		alertsClient := monitorSession.Alerts

		Describe("When the page of alerts are successfully retrieved", func() {
			pageOptions := &alerts.AlertsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://monitor.twilio.com/v1/Alerts?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/alertsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := alertsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the alerts page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://monitor.twilio.com/v1/Alerts?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://monitor.twilio.com/v1/Alerts?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("alerts"))

				alerts := resp.Alerts
				Expect(alerts).ToNot(BeNil())
				Expect(len(alerts)).To(Equal(1))

				Expect(alerts[0]).ToNot(BeNil())
				Expect(alerts[0].Sid).To(Equal("NOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(alerts[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(alerts[0].APIVersion).To(BeNil())
				Expect(alerts[0].AlertText).To(BeNil())
				Expect(alerts[0].ErrorCode).To(Equal("80408"))
				Expect(alerts[0].LogLevel).To(Equal("error"))
				Expect(alerts[0].MoreInfo).To(Equal("https://www.twilio.com/docs/errors/80408"))
				Expect(alerts[0].RequestHeaders).To(BeNil())
				Expect(alerts[0].RequestMethod).To(BeNil())
				Expect(alerts[0].RequestURL).To(BeNil())
				Expect(alerts[0].RequestVariables).To(BeNil())
				Expect(alerts[0].ResourceSid).To(Equal("KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(alerts[0].ResponseBody).To(BeNil())
				Expect(alerts[0].ResponseHeaders).To(BeNil())
				Expect(alerts[0].ServiceSid).To(Equal("KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(alerts[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(alerts[0].DateGenerated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(alerts[0].DateUpdated).To(BeNil())
				Expect(alerts[0].URL).To(Equal("https://monitor.twilio.com/v1/Alerts/NOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of alerts api returns a 500 response", func() {
			pageOptions := &alerts.AlertsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://monitor.twilio.com/v1/Alerts?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := alertsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the alerts page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated alerts are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://monitor.twilio.com/v1/Alerts",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/alertsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://monitor.twilio.com/v1/Alerts?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/alertsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := alertsClient.NewAlertsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated alerts current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated alerts results should be returned", func() {
				Expect(len(paginator.Alerts)).To(Equal(3))
			})
		})

		Describe("When the alerts api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://monitor.twilio.com/v1/Alerts",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/alertsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://monitor.twilio.com/v1/Alerts?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := alertsClient.NewAlertsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated alerts current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a alert sid", func() {
		alertClient := monitorSession.Alert("NOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the alert is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://monitor.twilio.com/v1/Alerts/NOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/alertResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := alertClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get alert response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("NOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.APIVersion).To(BeNil())
				Expect(resp.AlertText).To(BeNil())
				Expect(resp.ErrorCode).To(Equal("80408"))
				Expect(resp.LogLevel).To(Equal("error"))
				Expect(resp.MoreInfo).To(Equal("https://www.twilio.com/docs/errors/80408"))
				Expect(resp.RequestHeaders).To(BeNil())
				Expect(resp.RequestMethod).To(BeNil())
				Expect(resp.RequestURL).To(BeNil())
				Expect(resp.RequestVariables).To(BeNil())
				Expect(resp.ResourceSid).To(Equal("KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ResponseBody).To(BeNil())
				Expect(resp.ResponseHeaders).To(BeNil())
				Expect(resp.ServiceSid).To(Equal("KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateGenerated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://monitor.twilio.com/v1/Alerts/NOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the alert api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://monitor.twilio.com/v1/Alerts/NO71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := monitorSession.Alert("NO71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get alert response should be nil", func() {
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
	Expect(twilioErr.Message).To(Equal("The requested resource /Alerts/NO71 was not found"))
	Expect(twilioErr.MoreInfo).To(Equal(utils.String("https://www.twilio.com/docs/errors/20404")))
	Expect(twilioErr.Status).To(Equal(404))
}

func ExpectErrorToNotBeATwilioError(err error) {
	Expect(err).ToNot(BeNil())
	_, ok := err.(*utils.TwilioError)
	Expect(ok).To(Equal(false))
}
