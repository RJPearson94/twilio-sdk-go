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
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/key"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/keys"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/message"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/message/feedback"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/message/media_attachments"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/messages"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/queue"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/queues"
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

	Describe("Given the feedback client", func() {
		feedbackClient := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Message("MMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Feedback

		Describe("When the feedback is successfully created", func() {
			createInput := &feedback.CreateFeedbackInput{
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

			resp, err := feedbackClient.Create(createInput)
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

		Describe("When the create feedback api returns a 500 response", func() {
			createInput := &feedback.CreateFeedbackInput{
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

			resp, err := feedbackClient.Create(createInput)
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

		Describe("When the page of queuess api returns a 500 response", func() {
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
