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
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/key"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/keys"
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

			resp, err := accountClient.Get()
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

			resp, err := apiSession.Account("AC71").Get()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get account response should be nil", func() {
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

			resp, err := keyClient.Get()
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

			resp, err := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Key("SK71").Get()
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
