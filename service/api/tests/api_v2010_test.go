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
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/messages"
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

			It("Then the create fax response should be nil", func() {
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

			resp, err := messageClient.Get()
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

			resp, err := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Message("SM71").Get()
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

	Describe("Given I have a media sid", func() {
		mediaClient := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Message("SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Media("MEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the media is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Media/MEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/mediaResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := mediaClient.Get()
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

			resp, err := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Message("SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Media("ME71").Get()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get media response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the media is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Media/MEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json", httpmock.NewStringResponder(204, ""))

			err := mediaClient.Delete()
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

			err := apiSession.Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Message("SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Media("ME71").Delete()
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
