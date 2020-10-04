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

	"github.com/RJPearson94/twilio-sdk-go/service/api"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/address"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/addresses"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/application"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/applications"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/call"
	callFeedback "github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/call/feedback"
	callFeedbacks "github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/call/feedbacks"
	callRecordings "github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/call/recordings"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/calls"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/calls/feedback_summaries"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/conference"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/conference/participant"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/conference/participants"
	conferenceRecordings "github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/conference/recordings"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/conferences"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/key"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/keys"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/message"
	messageFeedbacks "github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/message/feedbacks"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/message/media_attachments"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/messages"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/queue"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/queue/member"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/queue/members"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/queues"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/recording"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/recordings"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/tokens"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/accounts"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

var _ = Describe("API V2010", func() {
	creds, err := credentials.New(credentials.Account{
		Sid:       "ACxxx",
		AuthToken: "Test",
	})
	if err != nil {
		log.Panicf("%s", err)
	}

	apiSession := api.NewWithCredentials(creds).V2010

	httpmock.ActivateNonDefault(apiSession.GetClient().GetRestyClient().GetClient())
	defer httpmock.DeactivateAndReset()

	Describe("Given the accounts client", func() {
		accountsClient := apiSession.Accounts

		Describe("When the assistant is successfully created", func() {
			createInput := &accounts.CreateAccountInput{}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/accountResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := accountsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create assistant response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test"))
				Expect(resp.Status).To(Equal("active"))
				Expect(resp.Type).To(Equal("Trial"))
				Expect(resp.AuthToken).To(Equal("TestToken"))
				Expect(resp.OwnerAccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the create account api returns a 500 response", func() {
			createInput := &accounts.CreateAccountInput{}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := accountsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create account response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of accounts are successfully retrieved", func() {
			pageOptions := &accounts.AccountsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/accountsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := accountsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the accounts page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				Expect(resp.Page).To(Equal(0))
				Expect(resp.Start).To(Equal(0))
				Expect(resp.End).To(Equal(1))
				Expect(resp.PageSize).To(Equal(50))
				Expect(resp.FirstPageURI).To(Equal("/2010-04-01/Accounts.json?PageSize=50&Page=0"))
				Expect(resp.PreviousPageURI).To(BeNil())
				Expect(resp.URI).To(Equal("/2010-04-01/Accounts.json?PageSize=50&Page=0"))
				Expect(resp.NextPageURI).To(BeNil())

				accounts := resp.Accounts
				Expect(accounts).ToNot(BeNil())
				Expect(len(accounts)).To(Equal(1))

				Expect(accounts[0].Sid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(accounts[0].FriendlyName).To(Equal("Test"))
				Expect(accounts[0].Status).To(Equal("active"))
				Expect(accounts[0].Type).To(Equal("Trial"))
				Expect(accounts[0].AuthToken).To(Equal("TestToken"))
				Expect(accounts[0].OwnerAccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(accounts[0].DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(accounts[0].DateUpdated).To(BeNil())
			})
		})

		Describe("When the page of accounts api returns a 500 response", func() {
			pageOptions := &accounts.AccountsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := accountsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the accounts page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated accounts are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/accountsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/accountsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := accountsClient.NewAccountsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated accounts current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated accounts results should be returned", func() {
				Expect(len(paginator.Accounts)).To(Equal(3))
			})
		})

		Describe("When the accounts api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/accountsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := accountsClient.NewAccountsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated accounts current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a account sid", func() {
		accountClient := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the account is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/accountResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := accountClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get account response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test"))
				Expect(resp.Status).To(Equal("active"))
				Expect(resp.Type).To(Equal("Trial"))
				Expect(resp.AuthToken).To(Equal("TestToken"))
				Expect(resp.OwnerAccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the account api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/AC71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := apiSession.Account("AC71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get account response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the account is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateAccountResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &account.UpdateAccountInput{
				Status: utils.String("closed"),
			}

			resp, err := accountClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update account response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test"))
				Expect(resp.Status).To(Equal("closed"))
				Expect(resp.Type).To(Equal("Trial"))
				Expect(resp.AuthToken).To(Equal("TestToken"))
				Expect(resp.OwnerAccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:10:00Z"))
			})
		})

		Describe("When the account api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/AC71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &account.UpdateAccountInput{
				Status: utils.String("closed"),
			}

			resp, err := apiSession.Account("AC71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update account response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given the keys client", func() {
		keysClient := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Keys

		Describe("When the key is successfully created", func() {
			createInput := &keys.CreateKeyInput{
				FriendlyName: utils.String("Test"),
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Keys.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/keyResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := keysClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create key response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("SKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Secret).To(Equal("SecretValue"))
				Expect(resp.FriendlyName).To(Equal("Test"))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the create key api returns a 500 response", func() {
			createInput := &keys.CreateKeyInput{
				FriendlyName: utils.String("Test"),
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Keys.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := keysClient.Create(createInput)
			It("Then an error should be returned", func() {
				Expect(err).ToNot(BeNil())
				twilioErr, ok := err.(*utils.TwilioError)
				Expect(ok).To(Equal(true))
				Expect(twilioErr.Code).To(BeNil())
				Expect(twilioErr.Message).To(Equal("An error occurred"))
				Expect(twilioErr.MoreInfo).To(BeNil())
				Expect(twilioErr.Status).To(Equal(500))
			})

			It("Then the create key response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of keys are successfully retrieved", func() {
			pageOptions := &keys.KeysPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Keys.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/keysPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := keysClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the keys page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				Expect(resp.Page).To(Equal(0))
				Expect(resp.Start).To(Equal(0))
				Expect(resp.End).To(Equal(1))
				Expect(resp.PageSize).To(Equal(50))
				Expect(resp.FirstPageURI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Keys.json?PageSize=50&Page=0"))
				Expect(resp.PreviousPageURI).To(BeNil())
				Expect(resp.URI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Keys.json?PageSize=50&Page=0"))
				Expect(resp.NextPageURI).To(BeNil())

				keys := resp.Keys
				Expect(keys).ToNot(BeNil())
				Expect(len(keys)).To(Equal(1))

				Expect(keys[0].Sid).To(Equal("SKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(keys[0].FriendlyName).To(Equal("Test"))
				Expect(keys[0].DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(keys[0].DateUpdated).To(BeNil())
			})
		})

		Describe("When the page of keys api returns a 500 response", func() {
			pageOptions := &keys.KeysPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Keys.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := keysClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the keys page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated keys are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Keys.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/keysPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Keys.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/keysPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := keysClient.NewKeysPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated keys current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated keys results should be returned", func() {
				Expect(len(paginator.Keys)).To(Equal(3))
			})
		})

		Describe("When the keys api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Keys.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/keysPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Keys.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := keysClient.NewKeysPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated keys current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a key sid", func() {
		keyClient := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Key("SKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the key is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Keys/SKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/getKeyResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := keyClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get key response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("SKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test"))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the key api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Keys/SK71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Key("SK71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get key response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the key is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Keys/SKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateKeyResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &key.UpdateKeyInput{
				FriendlyName: utils.String("Test 2"),
			}

			resp, err := keyClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update key response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("SKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test 2"))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:10:00Z"))
			})
		})

		Describe("When the key api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Keys/SK71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &key.UpdateKeyInput{
				FriendlyName: utils.String("Test 2"),
			}

			resp, err := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Key("SK71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update key response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the key is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Keys/SKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json", httpmock.NewStringResponder(204, ""))

			err := keyClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the key api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Keys/SK71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Key("SK71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the messages client", func() {
		messagesClient := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Messages

		Describe("When the message is successfully created", func() {
			createInput := &messages.CreateMessageInput{
				To:                  "+10123456789",
				MessagingServiceSid: utils.String("MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
				Body:                utils.String("Hello World"),
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/messageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := messagesClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create message response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Body).To(Equal("Hello World"))
				Expect(resp.NumSegments).To(Equal("1"))
				Expect(resp.Direction).To(Equal("outbound-api"))
				Expect(resp.From).To(Equal(utils.String("")))
				Expect(resp.Price).To(BeNil())
				Expect(resp.ErrorMessage).To(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.NumMedia).To(Equal("0"))
				Expect(resp.To).To(Equal("+10123456789"))
				Expect(resp.Status).To(Equal("failed"))
				Expect(resp.MessagingServiceSid).To(Equal(utils.String("MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.ErrorCode).To(Equal(utils.Int(21704)))
				Expect(resp.PriceUnit).To(Equal("GBP"))
				Expect(resp.APIVersion).To(Equal("2010-04-01"))
				Expect(resp.DateSent.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the message request does not contain a to", func() {
			createInput := &messages.CreateMessageInput{
				MessagingServiceSid: utils.String("MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
				Body:                utils.String("Hello World"),
			}

			resp, err := messagesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create message response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create message api returns a 500 response", func() {
			createInput := &messages.CreateMessageInput{
				To:                  "+10123456789",
				MessagingServiceSid: utils.String("MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
				Body:                utils.String("Hello World"),
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := messagesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create message response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of messages are successfully retrieved", func() {
			pageOptions := &messages.MessagesPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/messagesPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := messagesClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the messages page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				Expect(resp.Page).To(Equal(0))
				Expect(resp.Start).To(Equal(0))
				Expect(resp.End).To(Equal(1))
				Expect(resp.PageSize).To(Equal(50))
				Expect(resp.FirstPageURI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages.json?PageSize=50&Page=0"))
				Expect(resp.PreviousPageURI).To(BeNil())
				Expect(resp.URI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages.json?PageSize=50&Page=0"))
				Expect(resp.NextPageURI).To(BeNil())

				messages := resp.Messages
				Expect(messages).ToNot(BeNil())
				Expect(len(messages)).To(Equal(1))

				Expect(messages[0].Sid).To(Equal("SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(messages[0].Body).To(Equal("Hello World"))
				Expect(messages[0].NumSegments).To(Equal("1"))
				Expect(messages[0].Direction).To(Equal("outbound-api"))
				Expect(messages[0].From).To(Equal(utils.String("")))
				Expect(messages[0].Price).To(BeNil())
				Expect(messages[0].ErrorMessage).To(BeNil())
				Expect(messages[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(messages[0].NumMedia).To(Equal("0"))
				Expect(messages[0].To).To(Equal("+10123456789"))
				Expect(messages[0].Status).To(Equal("failed"))
				Expect(messages[0].MessagingServiceSid).To(Equal(utils.String("MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(messages[0].ErrorCode).To(Equal(utils.Int(21704)))
				Expect(messages[0].PriceUnit).To(Equal("GBP"))
				Expect(messages[0].APIVersion).To(Equal("2010-04-01"))
				Expect(messages[0].DateSent.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(messages[0].DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(messages[0].DateUpdated).To(BeNil())
			})
		})

		Describe("When the page of messages api returns a 500 response", func() {
			pageOptions := &messages.MessagesPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := messagesClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the messages page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated messages are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/messagesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/messagesPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := messagesClient.NewMessagesPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated messages current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated messages results should be returned", func() {
				Expect(len(paginator.Messages)).To(Equal(3))
			})
		})

		Describe("When the messages api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/messagesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := messagesClient.NewMessagesPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated messages current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a message sid", func() {
		messageClient := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Message("SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the message is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/messageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := messageClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get message response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Body).To(Equal("Hello World"))
				Expect(resp.NumSegments).To(Equal("1"))
				Expect(resp.Direction).To(Equal("outbound-api"))
				Expect(resp.From).To(Equal(utils.String("")))
				Expect(resp.Price).To(BeNil())
				Expect(resp.ErrorMessage).To(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.NumMedia).To(Equal("0"))
				Expect(resp.To).To(Equal("+10123456789"))
				Expect(resp.Status).To(Equal("failed"))
				Expect(resp.MessagingServiceSid).To(Equal(utils.String("MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.ErrorCode).To(Equal(utils.Int(21704)))
				Expect(resp.PriceUnit).To(Equal("GBP"))
				Expect(resp.APIVersion).To(Equal("2010-04-01"))
				Expect(resp.DateSent.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the message api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/SM71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Message("SM71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get message response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the message is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateMessageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &message.UpdateMessageInput{
				Body: "Test",
			}

			resp, err := messageClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update message response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Body).To(Equal("Test"))
				Expect(resp.NumSegments).To(Equal("1"))
				Expect(resp.Direction).To(Equal("outbound-api"))
				Expect(resp.From).To(Equal(utils.String("")))
				Expect(resp.Price).To(BeNil())
				Expect(resp.ErrorMessage).To(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.NumMedia).To(Equal("0"))
				Expect(resp.To).To(Equal("+10123456789"))
				Expect(resp.Status).To(Equal("failed"))
				Expect(resp.MessagingServiceSid).To(Equal(utils.String("MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.ErrorCode).To(Equal(utils.Int(21704)))
				Expect(resp.PriceUnit).To(Equal("GBP"))
				Expect(resp.APIVersion).To(Equal("2010-04-01"))
				Expect(resp.DateSent.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:10:00Z"))
			})
		})

		Describe("When the message api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/SM71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &message.UpdateMessageInput{
				Body: "Test",
			}

			resp, err := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Message("SM71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update message response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the message is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json", httpmock.NewStringResponder(204, ""))

			err := messageClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the message api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/SM71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Message("SM71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the feedbacks client", func() {
		feedbacksClient := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Message("MMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Feedbacks

		Describe("When the feedback is successfully created", func() {
			createInput := &messageFeedbacks.CreateFeedbackInput{
				Outcome: utils.String("confirmed"),
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/MMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Feedback.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/feedbackResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := feedbacksClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create feedback response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.MessageSid).To(Equal("MMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Outcome).To(Equal("confirmed"))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the create feedbacks api returns a 500 response", func() {
			createInput := &messageFeedbacks.CreateFeedbackInput{
				Outcome: utils.String("confirmed"),
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/MMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Feedback.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := feedbacksClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create feedback response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given the media attachments client", func() {
		mediaAttachmentsClient := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Message("SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").MediaAttachments

		Describe("When the page of media are successfully retrieved", func() {
			pageOptions := &media_attachments.MediaPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Media.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/mediaPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := mediaAttachmentsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the media page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				Expect(resp.Page).To(Equal(0))
				Expect(resp.Start).To(Equal(0))
				Expect(resp.End).To(Equal(1))
				Expect(resp.PageSize).To(Equal(50))
				Expect(resp.FirstPageURI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Media.json?PageSize=50&Page=0"))
				Expect(resp.PreviousPageURI).To(BeNil())
				Expect(resp.URI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Media.json?PageSize=50&Page=0"))
				Expect(resp.NextPageURI).To(BeNil())

				media := resp.Media
				Expect(media).ToNot(BeNil())
				Expect(len(media)).To(Equal(1))

				Expect(media[0].Sid).To(Equal("MEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(media[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(media[0].ParentSid).To(Equal("SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(media[0].ContentType).To(Equal("image/jpeg"))
				Expect(media[0].DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(media[0].DateUpdated).To(BeNil())
			})
		})

		Describe("When the page of media api returns a 500 response", func() {
			pageOptions := &media_attachments.MediaPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Media.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := mediaAttachmentsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the media page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated media are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Media.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/mediaPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Media.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/mediaPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := mediaAttachmentsClient.NewMediaPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated media current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated media results should be returned", func() {
				Expect(len(paginator.Media)).To(Equal(3))
			})
		})

		Describe("When the media api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Media.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/mediaPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Media.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := mediaAttachmentsClient.NewMediaPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated media current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a media attachment sid", func() {
		mediaAttachmentClient := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Message("SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").MediaAttachment("MEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the media is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Media/MEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/mediaResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := mediaAttachmentClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get media response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("MEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ParentSid).To(Equal("SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ContentType).To(Equal("image/jpeg"))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the media api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Media/ME71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Message("SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").MediaAttachment("ME71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get media response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the media is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Media/MEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json", httpmock.NewStringResponder(204, ""))

			err := mediaAttachmentClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the media api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Media/ME71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Message("SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").MediaAttachment("ME71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})

		Describe("Given I have a balance client", func() {
			balanceClient := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Balance()

			Describe("When the balance is successfully retrieved", func() {
				httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Balance.json",
					func(req *http.Request) (*http.Response, error) {
						fixture, _ := ioutil.ReadFile("testdata/balanceResponse.json")
						resp := make(map[string]interface{})
						json.Unmarshal(fixture, &resp)
						return httpmock.NewJsonResponse(200, resp)
					},
				)

				resp, err := balanceClient.Fetch()
				It("Then no error should be returned", func() {
					Expect(err).To(BeNil())
				})

				It("Then the get balance response should be returned", func() {
					Expect(resp).ToNot(BeNil())
					Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
					Expect(resp.Balance).To(Equal("1.00000"))
					Expect(resp.Currency).To(Equal("GBP"))
				})
			})

			Describe("When the balance api returns a 404", func() {
				httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Balance.json",
					func(req *http.Request) (*http.Response, error) {
						fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
						resp := make(map[string]interface{})
						json.Unmarshal(fixture, &resp)
						return httpmock.NewJsonResponse(500, resp)
					},
				)

				resp, err := balanceClient.Fetch()
				It("Then an error should be returned", func() {
					ExpectInternalServerError(err)
				})

				It("Then the get balance response should be nil", func() {
					Expect(resp).To(BeNil())
				})
			})
		})

		Describe("Given I have a tokens client", func() {
			tokensClient := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Tokens

			Describe("When the token is successfully created", func() {
				httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tokens.json",
					func(req *http.Request) (*http.Response, error) {
						fixture, _ := ioutil.ReadFile("testdata/tokenResponse.json")
						resp := make(map[string]interface{})
						json.Unmarshal(fixture, &resp)
						return httpmock.NewJsonResponse(200, resp)
					},
				)

				resp, err := tokensClient.Create(&tokens.CreateTokenInput{})
				It("Then no error should be returned", func() {
					Expect(err).To(BeNil())
				})

				It("Then the get token response should be returned", func() {
					Expect(resp).ToNot(BeNil())
					Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
					Expect(resp.Username).To(Equal("username"))
					Expect(resp.Password).To(Equal("password"))
					Expect(resp.IceServers).To(Equal([]tokens.CreateIceServerResponse{{
						URL:        "stun:global.stun.twilio.com:3478?transport=udp",
						URLs:       "stun:global.stun.twilio.com:3478?transport=udp",
						Username:   nil,
						Credential: nil,
					}, {
						URL:        "turn:global.turn.twilio.com:3478?transport=udp",
						URLs:       "turn:global.turn.twilio.com:3478?transport=udp",
						Username:   utils.String("username"),
						Credential: utils.String("password"),
					}, {
						URL:        "turn:global.turn.twilio.com:3478?transport=tcp",
						URLs:       "turn:global.turn.twilio.com:3478?transport=tcp",
						Username:   utils.String("username"),
						Credential: utils.String("password"),
					}, {
						URL:        "turn:global.turn.twilio.com:443?transport=tcp",
						URLs:       "turn:global.turn.twilio.com:443?transport=tcp",
						Username:   utils.String("username"),
						Credential: utils.String("password"),
					}}))
					Expect(resp.Ttl).To(Equal("1"))
					Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
					Expect(resp.DateUpdated).To(BeNil())
				})
			})

			Describe("When the token api returns a 404", func() {
				httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tokens.json",
					func(req *http.Request) (*http.Response, error) {
						fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
						resp := make(map[string]interface{})
						json.Unmarshal(fixture, &resp)
						return httpmock.NewJsonResponse(500, resp)
					},
				)

				resp, err := tokensClient.Create(&tokens.CreateTokenInput{})
				It("Then an error should be returned", func() {
					ExpectInternalServerError(err)
				})

				It("Then the get token response should be nil", func() {
					Expect(resp).To(BeNil())
				})
			})
		})
	})

	Describe("Given the queues client", func() {
		queuesClient := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Queues

		Describe("When the queue is successfully created", func() {
			createInput := &queues.CreateQueueInput{
				FriendlyName: "Test",
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/queueResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := queuesClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create queue response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test"))
				Expect(resp.CurrentSize).To(Equal(0))
				Expect(resp.AverageWaitTime).To(Equal(0))
				Expect(resp.MaxSize).To(Equal(100))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the queue request does not contain a to", func() {
			createInput := &queues.CreateQueueInput{}

			resp, err := queuesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create queue response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create queue api returns a 500 response", func() {
			createInput := &queues.CreateQueueInput{
				FriendlyName: "Test",
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := queuesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create queue response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of queues are successfully retrieved", func() {
			pageOptions := &queues.QueuesPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/queuesPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := queuesClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the queues page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				Expect(resp.Page).To(Equal(0))
				Expect(resp.Start).To(Equal(0))
				Expect(resp.End).To(Equal(1))
				Expect(resp.PageSize).To(Equal(50))
				Expect(resp.FirstPageURI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues.json?PageSize=50&Page=0"))
				Expect(resp.PreviousPageURI).To(BeNil())
				Expect(resp.URI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues.json?PageSize=50&Page=0"))
				Expect(resp.NextPageURI).To(BeNil())

				queues := resp.Queues
				Expect(queues).ToNot(BeNil())
				Expect(len(queues)).To(Equal(1))

				Expect(queues[0].Sid).To(Equal("QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(queues[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(queues[0].FriendlyName).To(Equal("Test"))
				Expect(queues[0].CurrentSize).To(Equal(0))
				Expect(queues[0].AverageWaitTime).To(Equal(0))
				Expect(queues[0].MaxSize).To(Equal(100))
				Expect(queues[0].DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(queues[0].DateUpdated).To(BeNil())
			})
		})

		Describe("When the page of queues api returns a 500 response", func() {
			pageOptions := &queues.QueuesPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := queuesClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the queues page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated queues are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/queuesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/queuesPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := queuesClient.NewQueuesPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated queues current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated queues results should be returned", func() {
				Expect(len(paginator.Queues)).To(Equal(3))
			})
		})

		Describe("When the queues api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/queuesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := queuesClient.NewQueuesPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated queue current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a queue sid", func() {
		queueClient := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Queue("QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the queue is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues/QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/queueResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := queueClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get queue response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test"))
				Expect(resp.CurrentSize).To(Equal(0))
				Expect(resp.AverageWaitTime).To(Equal(0))
				Expect(resp.MaxSize).To(Equal(100))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the queue api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues/QU71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Queue("QU71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get queue response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the queue is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues/QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateQueueResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &queue.UpdateQueueInput{
				FriendlyName: utils.String("Test 2"),
			}

			resp, err := queueClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update queue response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.CurrentSize).To(Equal(0))
				Expect(resp.FriendlyName).To(Equal("Test 2"))
				Expect(resp.AverageWaitTime).To(Equal(0))
				Expect(resp.MaxSize).To(Equal(100))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:10:00Z"))
			})
		})

		Describe("When the queue api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues/QU71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &queue.UpdateQueueInput{
				FriendlyName: utils.String("Test 2"),
			}

			resp, err := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Queue("QU71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update queue response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the queue is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues/QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json", httpmock.NewStringResponder(204, ""))

			err := queueClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the queue api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues/QU71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Queue("QU71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the members client", func() {
		membersClient := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Queue("QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Members

		Describe("When the page of members are successfully retrieved", func() {
			pageOptions := &members.MembersPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues/QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Members.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/membersPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := membersClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the members page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				Expect(resp.Page).To(Equal(0))
				Expect(resp.Start).To(Equal(0))
				Expect(resp.End).To(Equal(1))
				Expect(resp.PageSize).To(Equal(50))
				Expect(resp.FirstPageURI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues/QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Members.json?PageSize=50&Page=0"))
				Expect(resp.PreviousPageURI).To(BeNil())
				Expect(resp.URI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues/QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Members.json?PageSize=50&Page=0"))
				Expect(resp.NextPageURI).To(BeNil())

				members := resp.Members
				Expect(members).ToNot(BeNil())
				Expect(len(members)).To(Equal(1))

				Expect(members[0].QueueSid).To(Equal("QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(members[0].CallSid).To(Equal("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(members[0].Position).To(Equal(1))
				Expect(members[0].WaitTime).To(Equal(100))
				Expect(members[0].DateEnqueued.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
			})
		})

		Describe("When the page of members api returns a 500 response", func() {
			pageOptions := &members.MembersPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues/QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Members.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := membersClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the members page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated members are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues/QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Members.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/membersPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues/QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Members.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/membersPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := membersClient.NewMembersPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated members current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated members results should be returned", func() {
				Expect(len(paginator.Members)).To(Equal(3))
			})
		})

		Describe("When the members api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues/QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Members.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/membersPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues/QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Members.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := membersClient.NewMembersPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated member current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a member sid", func() {
		memberClient := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Queue("QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Member("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the member is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues/QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Members/CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/memberResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := memberClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get member response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.QueueSid).To(Equal("QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.CallSid).To(Equal("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Position).To(Equal(1))
				Expect(resp.WaitTime).To(Equal(100))
				Expect(resp.DateEnqueued.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
			})
		})

		Describe("When the member api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues/QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Members/CA71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Queue("QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Member("CA71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get member response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the member is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues/QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Members/CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/memberResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &member.UpdateMemberInput{
				URL: "http://localhost",
			}

			resp, err := memberClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update member response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.QueueSid).To(Equal("QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.CallSid).To(Equal("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Position).To(Equal(1))
				Expect(resp.WaitTime).To(Equal(100))
				Expect(resp.DateEnqueued.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
			})
		})

		Describe("When the member request does not contain a url", func() {
			updateInput := &member.UpdateMemberInput{}

			resp, err := memberClient.Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the update member response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the member api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues/QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Members/CA71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &member.UpdateMemberInput{
				URL: "http://localhost",
			}

			resp, err := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Queue("QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Member("CA71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update member response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given the calls client", func() {
		callsClient := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Calls

		Describe("When the call is successfully created", func() {
			createInput := &calls.CreateCallInput{
				To:    "+10123456789",
				From:  "+19876543210",
				TwiML: utils.String(`<?xml version="1.0" encoding="UTF-8"?><Response><Say>Hello World</Say></Response>`),
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/callResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := callsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create call response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.APIVersion).To(Equal("2010-04-01"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AnsweredBy).To(BeNil())
				Expect(resp.CallerName).To(BeNil())
				Expect(resp.Direction).To(Equal("outbound-api"))
				Expect(resp.Duration).To(Equal("0"))
				Expect(resp.EndTime).To(BeNil())
				Expect(resp.ForwardedFrom).To(BeNil())
				Expect(resp.From).To(Equal("+19876543210"))
				Expect(resp.FromFormatted).To(Equal("+19876543210"))
				Expect(resp.GroupSid).To(BeNil())
				Expect(resp.ParentCallSid).To(BeNil())
				Expect(resp.PhoneNumberSid).To(Equal("PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Price).To(BeNil())
				Expect(resp.PriceUnit).To(Equal(utils.String("GBP")))
				Expect(resp.QueueTime).To(Equal("0"))
				Expect(resp.StartTime).To(BeNil())
				Expect(resp.Status).To(Equal("ringing"))
				Expect(resp.To).To(Equal("+10123456789"))
				Expect(resp.ToFormatted).To(Equal("+10123456789"))
				Expect(resp.TrunkSid).To(BeNil())
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the call request does not contain a to", func() {
			createInput := &calls.CreateCallInput{
				From: "+1987654321",
			}

			resp, err := callsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create call response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the call request does not contain a from", func() {
			createInput := &calls.CreateCallInput{
				To: "+10123456789",
			}

			resp, err := callsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create call response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create call api returns a 500 response", func() {
			createInput := &calls.CreateCallInput{
				To:    "+10123456789",
				From:  "+1987654321",
				TwiML: utils.String(`<?xml version="1.0" encoding="UTF-8"?><Response><Say>Hello World</Say></Response>`),
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := callsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create call response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of calls are successfully retrieved", func() {
			pageOptions := &calls.CallsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/callsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := callsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the calls page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				Expect(resp.Page).To(Equal(0))
				Expect(resp.Start).To(Equal(0))
				Expect(resp.End).To(Equal(1))
				Expect(resp.PageSize).To(Equal(50))
				Expect(resp.FirstPageURI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls.json?PageSize=50&Page=0"))
				Expect(resp.PreviousPageURI).To(BeNil())
				Expect(resp.URI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls.json?PageSize=50&Page=0"))
				Expect(resp.NextPageURI).To(BeNil())

				calls := resp.Calls
				Expect(calls).ToNot(BeNil())
				Expect(len(calls)).To(Equal(1))

				Expect(calls[0].Sid).To(Equal("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(calls[0].APIVersion).To(Equal("2010-04-01"))
				Expect(calls[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(calls[0].AnsweredBy).To(BeNil())
				Expect(calls[0].CallerName).To(BeNil())
				Expect(calls[0].Direction).To(Equal("outbound-api"))
				Expect(calls[0].Duration).To(Equal("0"))
				Expect(calls[0].EndTime).To(BeNil())
				Expect(calls[0].ForwardedFrom).To(BeNil())
				Expect(calls[0].From).To(Equal("+19876543210"))
				Expect(calls[0].FromFormatted).To(Equal("+19876543210"))
				Expect(calls[0].GroupSid).To(BeNil())
				Expect(calls[0].ParentCallSid).To(BeNil())
				Expect(calls[0].PhoneNumberSid).To(Equal("PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(calls[0].Price).To(BeNil())
				Expect(calls[0].PriceUnit).To(Equal(utils.String("GBP")))
				Expect(calls[0].QueueTime).To(Equal("0"))
				Expect(calls[0].StartTime).To(BeNil())
				Expect(calls[0].Status).To(Equal("ringing"))
				Expect(calls[0].To).To(Equal("+10123456789"))
				Expect(calls[0].ToFormatted).To(Equal("+10123456789"))
				Expect(calls[0].TrunkSid).To(BeNil())
				Expect(calls[0].DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(calls[0].DateUpdated).To(BeNil())
			})
		})

		Describe("When the page of calls api returns a 500 response", func() {
			pageOptions := &calls.CallsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := callsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the calls page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated calls are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/callsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/callsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := callsClient.NewCallsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated calls current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated calls results should be returned", func() {
				Expect(len(paginator.Calls)).To(Equal(3))
			})
		})

		Describe("When the calls api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/callsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := callsClient.NewCallsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated calls current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a call sid", func() {
		callClient := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Call("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the call is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/callResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := callClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get call response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.APIVersion).To(Equal("2010-04-01"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AnsweredBy).To(BeNil())
				Expect(resp.CallerName).To(BeNil())
				Expect(resp.Direction).To(Equal("outbound-api"))
				Expect(resp.Duration).To(Equal("0"))
				Expect(resp.EndTime).To(BeNil())
				Expect(resp.ForwardedFrom).To(BeNil())
				Expect(resp.From).To(Equal("+19876543210"))
				Expect(resp.FromFormatted).To(Equal("+19876543210"))
				Expect(resp.GroupSid).To(BeNil())
				Expect(resp.ParentCallSid).To(BeNil())
				Expect(resp.PhoneNumberSid).To(Equal("PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Price).To(BeNil())
				Expect(resp.PriceUnit).To(Equal(utils.String("GBP")))
				Expect(resp.QueueTime).To(Equal("0"))
				Expect(resp.StartTime).To(BeNil())
				Expect(resp.Status).To(Equal("ringing"))
				Expect(resp.To).To(Equal("+10123456789"))
				Expect(resp.ToFormatted).To(Equal("+10123456789"))
				Expect(resp.TrunkSid).To(BeNil())
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the call api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/CA71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Call("CA71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get call response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the call is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateCallResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &call.UpdateCallInput{
				Status: utils.String("Completed"),
			}

			resp, err := callClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update call response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.APIVersion).To(Equal("2010-04-01"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AnsweredBy).To(BeNil())
				Expect(resp.CallerName).To(BeNil())
				Expect(resp.Direction).To(Equal("outbound-api"))
				Expect(resp.Duration).To(Equal("0"))
				Expect(resp.EndTime).To(BeNil())
				Expect(resp.ForwardedFrom).To(BeNil())
				Expect(resp.From).To(Equal("+19876543210"))
				Expect(resp.FromFormatted).To(Equal("+19876543210"))
				Expect(resp.GroupSid).To(BeNil())
				Expect(resp.ParentCallSid).To(BeNil())
				Expect(resp.PhoneNumberSid).To(Equal("PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Price).To(BeNil())
				Expect(resp.PriceUnit).To(Equal(utils.String("GBP")))
				Expect(resp.QueueTime).To(Equal("0"))
				Expect(resp.StartTime).To(BeNil())
				Expect(resp.Status).To(Equal("completed"))
				Expect(resp.To).To(Equal("+10123456789"))
				Expect(resp.ToFormatted).To(Equal("+10123456789"))
				Expect(resp.TrunkSid).To(BeNil())
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:10:00Z"))
			})
		})

		Describe("When the call api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/CA71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &call.UpdateCallInput{
				Status: utils.String("Completed"),
			}

			resp, err := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Call("CA71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update call response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the call is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json", httpmock.NewStringResponder(204, ""))

			err := callClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the call api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/CA71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Call("CA71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the conferences client", func() {
		conferencesClient := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Conferences

		Describe("When the page of conferences are successfully retrieved", func() {
			pageOptions := &conferences.ConferencesPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/conferencesPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := conferencesClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the conferences page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				Expect(resp.Page).To(Equal(0))
				Expect(resp.Start).To(Equal(0))
				Expect(resp.End).To(Equal(1))
				Expect(resp.PageSize).To(Equal(50))
				Expect(resp.FirstPageURI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences.json?PageSize=50&Page=0"))
				Expect(resp.PreviousPageURI).To(BeNil())
				Expect(resp.URI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences.json?PageSize=50&Page=0"))
				Expect(resp.NextPageURI).To(BeNil())

				conferences := resp.Conferences
				Expect(conferences).ToNot(BeNil())
				Expect(len(conferences)).To(Equal(1))

				Expect(conferences[0].Sid).To(Equal("CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(conferences[0].Status).To(Equal("in-progress"))
				Expect(conferences[0].ReasonConferenceEnded).To(BeNil())
				Expect(conferences[0].Region).To(Equal("us1"))
				Expect(conferences[0].FriendlyName).To(Equal("Test"))
				Expect(conferences[0].CallSidEndingConference).To(BeNil())
				Expect(conferences[0].APIVersion).To(Equal("2010-04-01"))
				Expect(conferences[0].DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(conferences[0].DateUpdated).To(BeNil())
			})
		})

		Describe("When the page of conferences api returns a 500 response", func() {
			pageOptions := &conferences.ConferencesPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := conferencesClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the conferences page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated conferences are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/conferencesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/conferencesPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := conferencesClient.NewConferencesPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated conferences current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated conferences results should be returned", func() {
				Expect(len(paginator.Conferences)).To(Equal(3))
			})
		})

		Describe("When the conferences api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/conferencesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := conferencesClient.NewConferencesPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated conferences current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a conference sid", func() {
		conferenceClient := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Conference("CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the conference is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/conferenceResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := conferenceClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get conference response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Status).To(Equal("in-progress"))
				Expect(resp.ReasonConferenceEnded).To(BeNil())
				Expect(resp.Region).To(Equal("us1"))
				Expect(resp.FriendlyName).To(Equal("Test"))
				Expect(resp.CallSidEndingConference).To(BeNil())
				Expect(resp.APIVersion).To(Equal("2010-04-01"))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the conference api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CF71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Conference("CF71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get conference response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the conference is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateConferenceResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &conference.UpdateConferenceInput{
				Status: utils.String("Completed"),
			}

			resp, err := conferenceClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update conference response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Status).To(Equal("completed"))
				Expect(resp.ReasonConferenceEnded).To(BeNil())
				Expect(resp.Region).To(Equal("us1"))
				Expect(resp.FriendlyName).To(Equal("Test"))
				Expect(resp.CallSidEndingConference).To(BeNil())
				Expect(resp.APIVersion).To(Equal("2010-04-01"))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:10:00Z"))
			})
		})

		Describe("When the conference api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CF71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &conference.UpdateConferenceInput{
				Status: utils.String("Completed"),
			}

			resp, err := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Conference("CF71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update conference response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given the participants client", func() {
		participantsClient := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Conference("CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Participants

		Describe("When the participant is successfully created", func() {
			createInput := &participants.CreateParticipantInput{
				From: "+19876543210",
				To:   "+10123456789",
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/participantResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := participantsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create participant response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Status).To(Equal("connected"))
				Expect(resp.ConferenceSid).To(Equal("CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Hold).To(Equal(false))
				Expect(resp.EndConferenceOnExit).To(Equal(false))
				Expect(resp.Label).To(BeNil())
				Expect(resp.Muted).To(Equal(false))
				Expect(resp.Coaching).To(Equal(false))
				Expect(resp.StartConferenceOnEnter).To(Equal(true))
				Expect(resp.CallSid).To(Equal("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.CallSidToCoach).To(BeNil())
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the participant request does not contain a From", func() {
			createInput := &participants.CreateParticipantInput{
				To: "+10123456789",
			}

			resp, err := participantsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create participants response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the participant request does not contain a To", func() {
			createInput := &participants.CreateParticipantInput{
				From: "+19876543210",
			}

			resp, err := participantsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create participants response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create participant api returns a 500 response", func() {
			createInput := &participants.CreateParticipantInput{
				From: "+19876543210",
				To:   "+10123456789",
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := participantsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create participant response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of participants are successfully retrieved", func() {
			pageOptions := &participants.ParticipantsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/participantsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := participantsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the participants page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				Expect(resp.Page).To(Equal(0))
				Expect(resp.Start).To(Equal(0))
				Expect(resp.End).To(Equal(1))
				Expect(resp.PageSize).To(Equal(50))
				Expect(resp.FirstPageURI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants.json?PageSize=50&Page=0"))
				Expect(resp.PreviousPageURI).To(BeNil())
				Expect(resp.URI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants.json?PageSize=50&Page=0"))
				Expect(resp.NextPageURI).To(BeNil())

				participants := resp.Participants
				Expect(participants).ToNot(BeNil())
				Expect(len(participants)).To(Equal(1))

				Expect(participants[0].Status).To(Equal("connected"))
				Expect(participants[0].ConferenceSid).To(Equal("CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(participants[0].Hold).To(Equal(false))
				Expect(participants[0].EndConferenceOnExit).To(Equal(false))
				Expect(participants[0].Label).To(BeNil())
				Expect(participants[0].Muted).To(Equal(false))
				Expect(participants[0].Coaching).To(Equal(false))
				Expect(participants[0].StartConferenceOnEnter).To(Equal(true))
				Expect(participants[0].CallSid).To(Equal("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(participants[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(participants[0].CallSidToCoach).To(BeNil())
				Expect(participants[0].DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(participants[0].DateUpdated).To(BeNil())
			})
		})

		Describe("When the page of participants api returns a 500 response", func() {
			pageOptions := &participants.ParticipantsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := participantsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the participants page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated participants are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/participantsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/participantsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := participantsClient.NewParticipantsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated participants current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated participants results should be returned", func() {
				Expect(len(paginator.Participants)).To(Equal(3))
			})
		})

		Describe("When the participants api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/participantsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := participantsClient.NewParticipantsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated participants current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a participant sid", func() {
		participantClient := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Conference("CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Participant("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the participant is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/participantResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := participantClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get participant response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Status).To(Equal("connected"))
				Expect(resp.ConferenceSid).To(Equal("CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Hold).To(Equal(false))
				Expect(resp.EndConferenceOnExit).To(Equal(false))
				Expect(resp.Label).To(BeNil())
				Expect(resp.Muted).To(Equal(false))
				Expect(resp.Coaching).To(Equal(false))
				Expect(resp.StartConferenceOnEnter).To(Equal(true))
				Expect(resp.CallSid).To(Equal("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.CallSidToCoach).To(BeNil())
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the participant api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/CA71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Conference("CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Participant("CA71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get participant response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the participant is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateParticipantResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &participant.UpdateParticipantInput{
				Muted: utils.Bool(true),
			}

			resp, err := participantClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update participant response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Status).To(Equal("connected"))
				Expect(resp.ConferenceSid).To(Equal("CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Hold).To(Equal(false))
				Expect(resp.EndConferenceOnExit).To(Equal(false))
				Expect(resp.Label).To(BeNil())
				Expect(resp.Muted).To(Equal(true))
				Expect(resp.Coaching).To(Equal(false))
				Expect(resp.StartConferenceOnEnter).To(Equal(true))
				Expect(resp.CallSid).To(Equal("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.CallSidToCoach).To(BeNil())
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:10:00Z"))
			})
		})

		Describe("When the participant api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/CA71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &participant.UpdateParticipantInput{
				Muted: utils.Bool(true),
			}

			resp, err := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Conference("CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Participant("CA71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update participant response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the participant is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json", httpmock.NewStringResponder(204, ""))

			err := participantClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the participant api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/CA71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Conference("CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Participant("CA71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the addresses client", func() {
		addressesClient := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Addresses

		Describe("When the address is successfully created", func() {
			createInput := &addresses.CreateAddressInput{
				CustomerName: "Test User",
				Street:       "123 Fake Street",
				City:         "Fake City",
				Region:       "Fake Region",
				PostalCode:   "AB12DC",
				IsoCountry:   "GB",
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Addresses.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/addressResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := addressesClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create addresses response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("ADXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(BeNil())
				Expect(resp.CustomerName).To(Equal("Test User"))
				Expect(resp.Street).To(Equal("123 Fake Street"))
				Expect(resp.StreetSecondary).To(BeNil())
				Expect(resp.City).To(Equal("Fake City"))
				Expect(resp.Region).To(Equal("Fake Region"))
				Expect(resp.PostalCode).To(Equal("AB12CD"))
				Expect(resp.IsoCountry).To(Equal("GB"))
				Expect(resp.EmergencyEnabled).To(Equal(false))
				Expect(resp.Validated).To(Equal(false))
				Expect(resp.Verified).To(Equal(false))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the address request does not contain a customer name", func() {
			createInput := &addresses.CreateAddressInput{
				Street:     "123 Fake Street",
				City:       "Fake City",
				Region:     "Fake Region",
				PostalCode: "AB12DC",
				IsoCountry: "GB",
			}

			resp, err := addressesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create addresses response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the address request does not contain a street", func() {
			createInput := &addresses.CreateAddressInput{
				CustomerName: "Test User",
				City:         "Fake City",
				Region:       "Fake Region",
				PostalCode:   "AB12DC",
				IsoCountry:   "GB",
			}

			resp, err := addressesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create addresses response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the address request does not contain a city", func() {
			createInput := &addresses.CreateAddressInput{
				CustomerName: "Test User",
				Street:       "123 Fake Street",
				Region:       "Fake Region",
				PostalCode:   "AB12DC",
				IsoCountry:   "GB",
			}

			resp, err := addressesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create addresses response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the address request does not contain a region", func() {
			createInput := &addresses.CreateAddressInput{
				CustomerName: "Test User",
				Street:       "123 Fake Street",
				City:         "Fake City",
				PostalCode:   "AB12DC",
				IsoCountry:   "GB",
			}

			resp, err := addressesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create addresses response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the address request does not contain a postal code", func() {
			createInput := &addresses.CreateAddressInput{
				CustomerName: "Test User",
				Street:       "123 Fake Street",
				City:         "Fake City",
				Region:       "Fake Region",
				IsoCountry:   "GB",
			}

			resp, err := addressesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create addresses response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the address request does not contain a iso country", func() {
			createInput := &addresses.CreateAddressInput{
				CustomerName: "Test User",
				Street:       "123 Fake Street",
				City:         "Fake City",
				Region:       "Fake Region",
				PostalCode:   "AB12DC",
			}

			resp, err := addressesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create addresses response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create address api returns a 500 response", func() {
			createInput := &addresses.CreateAddressInput{
				CustomerName: "Test User",
				Street:       "123 Fake Street",
				City:         "Fake City",
				Region:       "Fake Region",
				PostalCode:   "AB12DC",
				IsoCountry:   "GB",
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Addresses.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := addressesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create address response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of addresses are successfully retrieved", func() {
			pageOptions := &addresses.AddressesPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Addresses.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/addressesPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := addressesClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the addresses page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				Expect(resp.Page).To(Equal(0))
				Expect(resp.Start).To(Equal(0))
				Expect(resp.End).To(Equal(1))
				Expect(resp.PageSize).To(Equal(50))
				Expect(resp.FirstPageURI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Addresses.json?PageSize=50&Page=0"))
				Expect(resp.PreviousPageURI).To(BeNil())
				Expect(resp.URI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Addresses.json?PageSize=50&Page=0"))
				Expect(resp.NextPageURI).To(BeNil())

				addresses := resp.Addresses
				Expect(addresses).ToNot(BeNil())
				Expect(len(addresses)).To(Equal(1))

				Expect(addresses[0].Sid).To(Equal("ADXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(addresses[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(addresses[0].FriendlyName).To(BeNil())
				Expect(addresses[0].CustomerName).To(Equal("Test User"))
				Expect(addresses[0].Street).To(Equal("123 Fake Street"))
				Expect(addresses[0].StreetSecondary).To(BeNil())
				Expect(addresses[0].City).To(Equal("Fake City"))
				Expect(addresses[0].Region).To(Equal("Fake Region"))
				Expect(addresses[0].PostalCode).To(Equal("AB12CD"))
				Expect(addresses[0].IsoCountry).To(Equal("GB"))
				Expect(addresses[0].EmergencyEnabled).To(Equal(false))
				Expect(addresses[0].Validated).To(Equal(false))
				Expect(addresses[0].Verified).To(Equal(false))
				Expect(addresses[0].DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(addresses[0].DateUpdated).To(BeNil())
			})
		})

		Describe("When the page of addresses api returns a 500 response", func() {
			pageOptions := &addresses.AddressesPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Addresses.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := addressesClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the addresses page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated addresses are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Addresses.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/addressesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Addresses.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/addressesPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := addressesClient.NewAddressesPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated addresses current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated addresses results should be returned", func() {
				Expect(len(paginator.Addresses)).To(Equal(3))
			})
		})

		Describe("When the addresses api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Addresses.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/addressesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Addresses.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := addressesClient.NewAddressesPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated addresses current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a address sid", func() {
		addressClient := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Address("ADXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the address is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Addresses/ADXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/addressResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := addressClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get address response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("ADXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(BeNil())
				Expect(resp.CustomerName).To(Equal("Test User"))
				Expect(resp.Street).To(Equal("123 Fake Street"))
				Expect(resp.StreetSecondary).To(BeNil())
				Expect(resp.City).To(Equal("Fake City"))
				Expect(resp.Region).To(Equal("Fake Region"))
				Expect(resp.PostalCode).To(Equal("AB12CD"))
				Expect(resp.IsoCountry).To(Equal("GB"))
				Expect(resp.EmergencyEnabled).To(Equal(false))
				Expect(resp.Validated).To(Equal(false))
				Expect(resp.Verified).To(Equal(false))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the address api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Addresses/AD71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Address("AD71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get address response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the address is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Addresses/ADXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateAddressResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &address.UpdateAddressInput{
				PostalCode: utils.String("Fake Postal Code"),
			}

			resp, err := addressClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update address response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("ADXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(BeNil())
				Expect(resp.CustomerName).To(Equal("Test User"))
				Expect(resp.Street).To(Equal("123 Fake Street"))
				Expect(resp.StreetSecondary).To(BeNil())
				Expect(resp.City).To(Equal("Fake City"))
				Expect(resp.Region).To(Equal("Fake Region"))
				Expect(resp.PostalCode).To(Equal("Fake Postal Code"))
				Expect(resp.IsoCountry).To(Equal("GB"))
				Expect(resp.EmergencyEnabled).To(Equal(false))
				Expect(resp.Validated).To(Equal(false))
				Expect(resp.Verified).To(Equal(false))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:10:00Z"))
			})
		})

		Describe("When the address api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Addresses/AD71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &address.UpdateAddressInput{
				PostalCode: utils.String("Fake Postal Code"),
			}

			resp, err := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Address("AD71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update address response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the address is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Addresses/ADXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json", httpmock.NewStringResponder(204, ""))

			err := addressClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the address api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Addresses/AD71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Address("AD71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the recordings client", func() {
		recordingsClient := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Recordings

		Describe("When the page of recordings are successfully retrieved", func() {
			pageOptions := &recordings.RecordingsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/recordingsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := recordingsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the recordings page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				Expect(resp.Page).To(Equal(0))
				Expect(resp.Start).To(Equal(0))
				Expect(resp.End).To(Equal(1))
				Expect(resp.PageSize).To(Equal(50))
				Expect(resp.FirstPageURI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json?PageSize=50&Page=0"))
				Expect(resp.PreviousPageURI).To(BeNil())
				Expect(resp.URI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json?PageSize=50&Page=0"))
				Expect(resp.NextPageURI).To(BeNil())

				recordings := resp.Recordings
				Expect(recordings).ToNot(BeNil())
				Expect(len(recordings)).To(Equal(1))

				Expect(recordings[0].APIVersion).To(Equal("2010-04-01"))
				Expect(recordings[0].Sid).To(Equal("REXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(recordings[0].CallSid).To(Equal("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(recordings[0].ConferenceSid).To(BeNil())
				Expect(recordings[0].Duration).To(BeNil())
				Expect(recordings[0].Price).To(BeNil())
				Expect(recordings[0].PriceUnit).To(Equal(utils.String("GBP")))
				Expect(recordings[0].Status).To(Equal("in-progress"))
				Expect(recordings[0].Channels).To(Equal(1))
				Expect(recordings[0].Source).To(Equal("OutboundAPI"))
				Expect(recordings[0].ErrorCode).To(BeNil())
				Expect(recordings[0].EncryptionDetails).To(BeNil())
				Expect(recordings[0].StartTime.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(recordings[0].DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(recordings[0].DateUpdated).To(BeNil())
			})
		})

		Describe("When the page of recordings api returns a 500 response", func() {
			pageOptions := &recordings.RecordingsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := recordingsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the recordings page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated recordings are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/recordingsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/recordingsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := recordingsClient.NewRecordingsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated recordings current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated recordings results should be returned", func() {
				Expect(len(paginator.Recordings)).To(Equal(3))
			})
		})

		Describe("When the recordings api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/recordingsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := recordingsClient.NewRecordingsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated recordings current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a recording sid", func() {
		recordingClient := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Recording("REXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the recording is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings/REXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
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
				Expect(resp.APIVersion).To(Equal("2010-04-01"))
				Expect(resp.Sid).To(Equal("REXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.CallSid).To(Equal("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ConferenceSid).To(BeNil())
				Expect(resp.Duration).To(BeNil())
				Expect(resp.Price).To(BeNil())
				Expect(resp.PriceUnit).To(Equal(utils.String("GBP")))
				Expect(resp.Status).To(Equal("in-progress"))
				Expect(resp.Channels).To(Equal(1))
				Expect(resp.Source).To(Equal("OutboundAPI"))
				Expect(resp.ErrorCode).To(BeNil())
				Expect(resp.EncryptionDetails).To(BeNil())
				Expect(resp.StartTime.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the recording api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings/RE71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Recording("RE71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get recording response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the recording is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings/REXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateRecordingResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &recording.UpdateRecordingInput{
				Status: "completed",
			}

			resp, err := recordingClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update recording response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.APIVersion).To(Equal("2010-04-01"))
				Expect(resp.Sid).To(Equal("REXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.CallSid).To(Equal("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ConferenceSid).To(BeNil())
				Expect(resp.Duration).To(BeNil())
				Expect(resp.Price).To(BeNil())
				Expect(resp.PriceUnit).To(Equal(utils.String("GBP")))
				Expect(resp.Status).To(Equal("completed"))
				Expect(resp.Channels).To(Equal(1))
				Expect(resp.Source).To(Equal("OutboundAPI"))
				Expect(resp.ErrorCode).To(BeNil())
				Expect(resp.EncryptionDetails).To(BeNil())
				Expect(resp.StartTime.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:10:00Z"))
			})
		})

		Describe("When the recording update request does not contain a status", func() {
			updateInput := &recording.UpdateRecordingInput{}

			resp, err := recordingClient.Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create addresses response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the recording api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings/RE71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &recording.UpdateRecordingInput{
				Status: "completed",
			}

			resp, err := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Recording("RE71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update recording response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the recording is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings/REXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json", httpmock.NewStringResponder(204, ""))

			err := recordingClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the recording api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings/RE71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Recording("RE71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the call recordings client", func() {
		recordingsClient := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Call("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Recordings

		Describe("When the recordings is successfully created", func() {
			createInput := &callRecordings.CreateRecordingInput{}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/recordingResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := recordingsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create addresses response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.APIVersion).To(Equal("2010-04-01"))
				Expect(resp.Sid).To(Equal("REXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.CallSid).To(Equal("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ConferenceSid).To(BeNil())
				Expect(resp.Duration).To(BeNil())
				Expect(resp.Price).To(BeNil())
				Expect(resp.PriceUnit).To(Equal(utils.String("GBP")))
				Expect(resp.Status).To(Equal("in-progress"))
				Expect(resp.Channels).To(Equal(1))
				Expect(resp.Source).To(Equal("OutboundAPI"))
				Expect(resp.ErrorCode).To(BeNil())
				Expect(resp.EncryptionDetails).To(BeNil())
				Expect(resp.StartTime.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the create recording api returns a 500 response", func() {
			createInput := &callRecordings.CreateRecordingInput{}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := recordingsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create recording response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of recordings are successfully retrieved", func() {
			pageOptions := &callRecordings.RecordingsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/recordingsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := recordingsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the recordings page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				Expect(resp.Page).To(Equal(0))
				Expect(resp.Start).To(Equal(0))
				Expect(resp.End).To(Equal(1))
				Expect(resp.PageSize).To(Equal(50))
				Expect(resp.FirstPageURI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json?PageSize=50&Page=0"))
				Expect(resp.PreviousPageURI).To(BeNil())
				Expect(resp.URI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json?PageSize=50&Page=0"))
				Expect(resp.NextPageURI).To(BeNil())

				recordings := resp.Recordings
				Expect(recordings).ToNot(BeNil())
				Expect(len(recordings)).To(Equal(1))

				Expect(recordings[0].APIVersion).To(Equal("2010-04-01"))
				Expect(recordings[0].Sid).To(Equal("REXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(recordings[0].CallSid).To(Equal("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(recordings[0].ConferenceSid).To(BeNil())
				Expect(recordings[0].Duration).To(BeNil())
				Expect(recordings[0].Price).To(BeNil())
				Expect(recordings[0].PriceUnit).To(Equal(utils.String("GBP")))
				Expect(recordings[0].Status).To(Equal("in-progress"))
				Expect(recordings[0].Channels).To(Equal(1))
				Expect(recordings[0].Source).To(Equal("OutboundAPI"))
				Expect(recordings[0].ErrorCode).To(BeNil())
				Expect(recordings[0].EncryptionDetails).To(BeNil())
				Expect(recordings[0].StartTime.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(recordings[0].DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(recordings[0].DateUpdated).To(BeNil())
			})
		})

		Describe("When the page of recordings api returns a 500 response", func() {
			pageOptions := &callRecordings.RecordingsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := recordingsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the recordings page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated recordings are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/recordingsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/recordingsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := recordingsClient.NewRecordingsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated recordings current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated recordings results should be returned", func() {
				Expect(len(paginator.Recordings)).To(Equal(3))
			})
		})

		Describe("When the recordings api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/recordingsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := recordingsClient.NewRecordingsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated recordings current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a call recording sid", func() {
		recordingClient := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Call("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Recording("REXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the recording is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings/REXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
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
				Expect(resp.APIVersion).To(Equal("2010-04-01"))
				Expect(resp.Sid).To(Equal("REXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.CallSid).To(Equal("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ConferenceSid).To(BeNil())
				Expect(resp.Duration).To(BeNil())
				Expect(resp.Price).To(BeNil())
				Expect(resp.PriceUnit).To(Equal(utils.String("GBP")))
				Expect(resp.Status).To(Equal("in-progress"))
				Expect(resp.Channels).To(Equal(1))
				Expect(resp.Source).To(Equal("OutboundAPI"))
				Expect(resp.ErrorCode).To(BeNil())
				Expect(resp.EncryptionDetails).To(BeNil())
				Expect(resp.StartTime.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the recording api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings/RE71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Call("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Recording("RE71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get recording response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given the conference recordings client", func() {
		recordingsClient := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Conference("CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Recordings

		Describe("When the recordings is successfully created", func() {
			createInput := &conferenceRecordings.CreateRecordingInput{}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/recordingResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := recordingsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create addresses response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.APIVersion).To(Equal("2010-04-01"))
				Expect(resp.Sid).To(Equal("REXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.CallSid).To(Equal("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ConferenceSid).To(BeNil())
				Expect(resp.Duration).To(BeNil())
				Expect(resp.Price).To(BeNil())
				Expect(resp.PriceUnit).To(Equal(utils.String("GBP")))
				Expect(resp.Status).To(Equal("in-progress"))
				Expect(resp.Channels).To(Equal(1))
				Expect(resp.Source).To(Equal("OutboundAPI"))
				Expect(resp.ErrorCode).To(BeNil())
				Expect(resp.EncryptionDetails).To(BeNil())
				Expect(resp.StartTime.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the create recording api returns a 500 response", func() {
			createInput := &conferenceRecordings.CreateRecordingInput{}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := recordingsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create recording response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of recordings are successfully retrieved", func() {
			pageOptions := &conferenceRecordings.RecordingsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/recordingsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := recordingsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the recordings page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				Expect(resp.Page).To(Equal(0))
				Expect(resp.Start).To(Equal(0))
				Expect(resp.End).To(Equal(1))
				Expect(resp.PageSize).To(Equal(50))
				Expect(resp.FirstPageURI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json?PageSize=50&Page=0"))
				Expect(resp.PreviousPageURI).To(BeNil())
				Expect(resp.URI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json?PageSize=50&Page=0"))
				Expect(resp.NextPageURI).To(BeNil())

				recordings := resp.Recordings
				Expect(recordings).ToNot(BeNil())
				Expect(len(recordings)).To(Equal(1))

				Expect(recordings[0].APIVersion).To(Equal("2010-04-01"))
				Expect(recordings[0].Sid).To(Equal("REXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(recordings[0].CallSid).To(Equal("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(recordings[0].ConferenceSid).To(BeNil())
				Expect(recordings[0].Duration).To(BeNil())
				Expect(recordings[0].Price).To(BeNil())
				Expect(recordings[0].PriceUnit).To(Equal(utils.String("GBP")))
				Expect(recordings[0].Status).To(Equal("in-progress"))
				Expect(recordings[0].Channels).To(Equal(1))
				Expect(recordings[0].Source).To(Equal("OutboundAPI"))
				Expect(recordings[0].ErrorCode).To(BeNil())
				Expect(recordings[0].EncryptionDetails).To(BeNil())
				Expect(recordings[0].StartTime.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(recordings[0].DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(recordings[0].DateUpdated).To(BeNil())
			})
		})

		Describe("When the page of recordings api returns a 500 response", func() {
			pageOptions := &conferenceRecordings.RecordingsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := recordingsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the recordings page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated recordings are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/recordingsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/recordingsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := recordingsClient.NewRecordingsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated recordings current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated recordings results should be returned", func() {
				Expect(len(paginator.Recordings)).To(Equal(3))
			})
		})

		Describe("When the recordings api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/recordingsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := recordingsClient.NewRecordingsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated recordings current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a conference recording sid", func() {
		recordingClient := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Conference("CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Recording("REXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the recording is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings/REXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
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
				Expect(resp.APIVersion).To(Equal("2010-04-01"))
				Expect(resp.Sid).To(Equal("REXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.CallSid).To(Equal("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ConferenceSid).To(BeNil())
				Expect(resp.Duration).To(BeNil())
				Expect(resp.Price).To(BeNil())
				Expect(resp.PriceUnit).To(Equal(utils.String("GBP")))
				Expect(resp.Status).To(Equal("in-progress"))
				Expect(resp.Channels).To(Equal(1))
				Expect(resp.Source).To(Equal("OutboundAPI"))
				Expect(resp.ErrorCode).To(BeNil())
				Expect(resp.EncryptionDetails).To(BeNil())
				Expect(resp.StartTime.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the recording api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conferences/CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings/RE71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Conference("CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Recording("RE71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get recording response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given the feedback summaries client", func() {
		feedbackSummariesClient := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Calls.FeedbackSummaries

		Describe("When the feedback summary is successfully created", func() {
			createInput := &feedback_summaries.CreateFeedbackSummaryInput{
				StartDate: "2019-10-03",
				EndDate:   "2020-10-03",
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/FeedbackSummary.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/feedbackSummaryResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := feedbackSummariesClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create feedback summary response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("FSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.StartDate).To(Equal("2019-10-03"))
				Expect(resp.EndDate).To(Equal("2020-10-03"))
				Expect(resp.IncludeSubaccounts).To(Equal(false))
				Expect(resp.Status).To(Equal("queued"))
				Expect(resp.CallCount).To(Equal(0))
				Expect(resp.CallFeedbackCount).To(Equal(0))
				Expect(resp.QualityScoreAverage).To(BeNil())
				Expect(resp.QualityScoreMedian).To(BeNil())
				Expect(resp.QualityScoreStandardDeviation).To(BeNil())
				Expect(resp.Issues).To(BeNil())
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the feedback summary request does not contain a start date", func() {
			createInput := &feedback_summaries.CreateFeedbackSummaryInput{
				EndDate: "2020-10-03",
			}

			resp, err := feedbackSummariesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create feedback summary  response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the feedback summary request does not contain a end date", func() {
			createInput := &feedback_summaries.CreateFeedbackSummaryInput{
				StartDate: "2019-10-03",
			}

			resp, err := feedbackSummariesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create feedback summary response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create feedback summary api returns a 500 response", func() {
			createInput := &feedback_summaries.CreateFeedbackSummaryInput{
				StartDate: "2019-10-03",
				EndDate:   "2020-10-03",
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/FeedbackSummary.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := feedbackSummariesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create feedback summary response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a feedback summary sid", func() {
		feedbackSummaryClient := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Calls.FeedbackSummary("FSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the feedback summary is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/FeedbackSummary/FSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/feedbackSummaryResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := feedbackSummaryClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get feedback summary response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("FSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.StartDate).To(Equal("2019-10-03"))
				Expect(resp.EndDate).To(Equal("2020-10-03"))
				Expect(resp.IncludeSubaccounts).To(Equal(false))
				Expect(resp.Status).To(Equal("queued"))
				Expect(resp.CallCount).To(Equal(0))
				Expect(resp.CallFeedbackCount).To(Equal(0))
				Expect(resp.QualityScoreAverage).To(BeNil())
				Expect(resp.QualityScoreMedian).To(BeNil())
				Expect(resp.QualityScoreStandardDeviation).To(BeNil())
				Expect(resp.Issues).To(BeNil())
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the feedback summary api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/FeedbackSummary/FS71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Calls.FeedbackSummary("FS71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get feedback summary response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the feedback summary is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/FeedbackSummary/FSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json", httpmock.NewStringResponder(204, ""))

			err := feedbackSummaryClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the feedback summary api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/FeedbackSummary/FS71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Calls.FeedbackSummary("FS71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the feedbacks client", func() {
		feedbacksClient := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Call("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Feedbacks

		Describe("When the feedback is successfully created", func() {
			createInput := &callFeedbacks.CreateFeedbackInput{
				QualityScore: 5,
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Feedback.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/callFeedbackResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := feedbacksClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create feedback response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.QualityScore).To(Equal(5))
				Expect(resp.Issues).To(Equal([]string{}))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the feedback request does not contain a quality score", func() {
			createInput := &callFeedbacks.CreateFeedbackInput{}

			resp, err := feedbacksClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create feedback response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create feedback api returns a 500 response", func() {
			createInput := &callFeedbacks.CreateFeedbackInput{
				QualityScore: 5,
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Feedback.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := feedbacksClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create feedback response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a feedback sid", func() {
		feedbackClient := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Call("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Feedback()

		Describe("When the feedback is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Feedback.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/callFeedbackResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := feedbackClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get feedback response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.QualityScore).To(Equal(5))
				Expect(resp.Issues).To(Equal([]string{}))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the feedback api returns a 500", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Feedback.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := feedbackClient.Fetch()
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the get feedback response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the feedback is successfully updated", func() {
			updateInput := &callFeedback.UpdateFeedbackInput{
				QualityScore: 4,
				Issues:       &[]string{"audio-latency"},
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Feedback.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateCallFeedbackResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := feedbackClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update feedback response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.QualityScore).To(Equal(4))
				Expect(resp.Issues).To(Equal([]string{"audio-latency"}))
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:10:00Z"))
			})
		})

		Describe("When the feedback request does not contain a quality score", func() {
			updateInput := &callFeedback.UpdateFeedbackInput{}

			resp, err := feedbackClient.Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the update feedback response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the update feedback api returns a 500 response", func() {
			updateInput := &callFeedback.UpdateFeedbackInput{
				QualityScore: 5,
			}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Calls/CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Feedback.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := feedbackClient.Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the update feedback response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given the applications client", func() {
		applicationsClient := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Applications

		Describe("When the application is successfully created", func() {
			createInput := &applications.CreateApplicationInput{}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Applications.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/applicationResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := applicationsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create applications response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("APXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.APIVersion).To(Equal("2010-04-01"))
				Expect(resp.FriendlyName).To(BeNil())
				Expect(resp.MessageStatusCallback).To(BeNil())
				Expect(resp.SmsFallbackMethod).To(Equal("POST"))
				Expect(resp.SmsFallbackURL).To(BeNil())
				Expect(resp.SmsMethod).To(Equal("POST"))
				Expect(resp.SmsStatusCallback).To(BeNil())
				Expect(resp.SmsURL).To(BeNil())
				Expect(resp.StatusCallback).To(BeNil())
				Expect(resp.StatusCallbackMethod).To(Equal("POST"))
				Expect(resp.VoiceCallerIDLookup).To(Equal(false))
				Expect(resp.VoiceFallbackMethod).To(Equal("POST"))
				Expect(resp.VoiceFallbackURL).To(BeNil())
				Expect(resp.VoiceMethod).To(Equal("POST"))
				Expect(resp.VoiceURL).To(BeNil())
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the create applications api returns a 500 response", func() {
			createInput := &applications.CreateApplicationInput{}

			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Applications.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := applicationsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create application response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of applications are successfully retrieved", func() {
			pageOptions := &applications.ApplicationsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Applications.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/applicationsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := applicationsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the addresses page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				Expect(resp.Page).To(Equal(0))
				Expect(resp.Start).To(Equal(0))
				Expect(resp.End).To(Equal(1))
				Expect(resp.PageSize).To(Equal(50))
				Expect(resp.FirstPageURI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Applications.json?PageSize=50&Page=0"))
				Expect(resp.PreviousPageURI).To(BeNil())
				Expect(resp.URI).To(Equal("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Applications.json?PageSize=50&Page=0"))
				Expect(resp.NextPageURI).To(BeNil())

				applications := resp.Applications
				Expect(applications).ToNot(BeNil())
				Expect(len(applications)).To(Equal(1))

				Expect(applications[0].Sid).To(Equal("APXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(applications[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(applications[0].APIVersion).To(Equal("2010-04-01"))
				Expect(applications[0].FriendlyName).To(BeNil())
				Expect(applications[0].MessageStatusCallback).To(BeNil())
				Expect(applications[0].SmsFallbackMethod).To(Equal("POST"))
				Expect(applications[0].SmsFallbackURL).To(BeNil())
				Expect(applications[0].SmsMethod).To(Equal("POST"))
				Expect(applications[0].SmsStatusCallback).To(BeNil())
				Expect(applications[0].SmsURL).To(BeNil())
				Expect(applications[0].StatusCallback).To(BeNil())
				Expect(applications[0].StatusCallbackMethod).To(Equal("POST"))
				Expect(applications[0].VoiceCallerIDLookup).To(Equal(false))
				Expect(applications[0].VoiceFallbackMethod).To(Equal("POST"))
				Expect(applications[0].VoiceFallbackURL).To(BeNil())
				Expect(applications[0].VoiceMethod).To(Equal("POST"))
				Expect(applications[0].VoiceURL).To(BeNil())
				Expect(applications[0].DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(applications[0].DateUpdated).To(BeNil())
			})
		})

		Describe("When the page of applications api returns a 500 response", func() {
			pageOptions := &applications.ApplicationsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Applications.json?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := applicationsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the applications page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated applications are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Applications.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/applicationsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Applications.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/applicationsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := applicationsClient.NewApplicationsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated applications current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated applications results should be returned", func() {
				Expect(len(paginator.Applications)).To(Equal(3))
			})
		})

		Describe("When the applications api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Applications.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/applicationsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Applications.json?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := applicationsClient.NewApplicationsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated applications current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a application sid", func() {
		applicationClient := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Application("APXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the application is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Applications/APXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/applicationResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := applicationClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get application response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("APXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.APIVersion).To(Equal("2010-04-01"))
				Expect(resp.FriendlyName).To(BeNil())
				Expect(resp.MessageStatusCallback).To(BeNil())
				Expect(resp.SmsFallbackMethod).To(Equal("POST"))
				Expect(resp.SmsFallbackURL).To(BeNil())
				Expect(resp.SmsMethod).To(Equal("POST"))
				Expect(resp.SmsStatusCallback).To(BeNil())
				Expect(resp.SmsURL).To(BeNil())
				Expect(resp.StatusCallback).To(BeNil())
				Expect(resp.StatusCallbackMethod).To(Equal("POST"))
				Expect(resp.VoiceCallerIDLookup).To(Equal(false))
				Expect(resp.VoiceFallbackMethod).To(Equal("POST"))
				Expect(resp.VoiceFallbackURL).To(BeNil())
				Expect(resp.VoiceMethod).To(Equal("POST"))
				Expect(resp.VoiceURL).To(BeNil())
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the application api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Applications/AP71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Application("AP71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get application response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the application is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Applications/APXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateApplicationResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &application.UpdateApplicationInput{
				FriendlyName: utils.String("Test"),
			}

			resp, err := applicationClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update application response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("APXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.APIVersion).To(Equal("2010-04-01"))
				Expect(resp.FriendlyName).To(Equal(utils.String("Test")))
				Expect(resp.MessageStatusCallback).To(BeNil())
				Expect(resp.SmsFallbackMethod).To(Equal("POST"))
				Expect(resp.SmsFallbackURL).To(BeNil())
				Expect(resp.SmsMethod).To(Equal("POST"))
				Expect(resp.SmsStatusCallback).To(BeNil())
				Expect(resp.SmsURL).To(BeNil())
				Expect(resp.StatusCallback).To(BeNil())
				Expect(resp.StatusCallbackMethod).To(Equal("POST"))
				Expect(resp.VoiceCallerIDLookup).To(Equal(false))
				Expect(resp.VoiceFallbackMethod).To(Equal("POST"))
				Expect(resp.VoiceFallbackURL).To(BeNil())
				Expect(resp.VoiceMethod).To(Equal("POST"))
				Expect(resp.VoiceURL).To(BeNil())
				Expect(resp.DateCreated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated.Time.Format(time.RFC3339)).To(Equal("2020-06-27T23:10:00Z"))
			})
		})

		Describe("When the application api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Applications/AP71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &application.UpdateApplicationInput{
				FriendlyName: utils.String("Test"),
			}

			resp, err := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Application("AP71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update application response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the application is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Applications/APXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json", httpmock.NewStringResponder(204, ""))

			err := applicationClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the application api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Applications/AP71.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Application("AP71").Delete()
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
	Expect(twilioErr.Message).To(Equal("The requested resource /Account/AC71.json was not found"))

	moreInfo := "https://www.twilio.com/docs/errors/20404"
	Expect(twilioErr.MoreInfo).To(Equal(&moreInfo))
	Expect(twilioErr.Status).To(Equal(404))
}

func ExpectErrorToNotBeATwilioError(err error) {
	Expect(err).ToNot(BeNil())
	_, ok := err.(*utils.TwilioError)
	Expect(ok).To(Equal(false))
}
