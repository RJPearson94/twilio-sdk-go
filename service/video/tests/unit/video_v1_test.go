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
	"github.com/RJPearson94/twilio-sdk-go/service/video"
	"github.com/RJPearson94/twilio-sdk-go/service/video/v1/room"
	"github.com/RJPearson94/twilio-sdk-go/service/video/v1/rooms"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

var _ = Describe("Video V1", func() {
	creds, err := credentials.New(credentials.Account{
		Sid:       "ACxxx",
		AuthToken: "Test",
	})
	if err != nil {
		log.Panicf("%s", err)
	}

	videoSession := video.New(session.New(creds), &client.Config{
		RetryAttempts: utils.Int(0),
	}).V1

	httpmock.ActivateNonDefault(videoSession.GetClient().GetRestyClient().GetClient())
	defer httpmock.DeactivateAndReset()

	Describe("Given I have a rooms client", func() {
		roomsClient := videoSession.Rooms

		Describe("When the room resource is successfully created", func() {
			createInput := &rooms.CreateRoomInput{}

			httpmock.RegisterResponder("POST", "https://video.twilio.com/v1/Rooms",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/roomResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := roomsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create room response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.UniqueName).To(Equal("RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Status).To(Equal("in-progress"))
				Expect(resp.VideoCodecs).To(Equal(&[]string{"VP8", "H264"}))
				Expect(resp.MaxParticipants).To(Equal(50))
				Expect(resp.RecordParticipantsOnConnect).To(Equal(false))
				Expect(resp.EndTime).To(BeNil())
				Expect(resp.Duration).To(BeNil())
				Expect(resp.MaxConcurrentPublishedTracks).To(BeNil())
				Expect(resp.StatusCallbackMethod).To(BeNil())
				Expect(resp.StatusCallback).To(BeNil())
				Expect(resp.Type).To(Equal("group"))
				Expect(resp.MediaRegion).To(Equal(utils.String("us1")))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2021-02-20T10:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the create rooms api returns a 500 response", func() {
			createInput := &rooms.CreateRoomInput{}

			httpmock.RegisterResponder("POST", "https://video.twilio.com/v1/Rooms",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := roomsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create room response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of rooms are successfully retrieved", func() {
			pageOptions := &rooms.RoomsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Rooms?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/roomsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := roomsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the rooms page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://video.twilio.com/v1/Rooms?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://video.twilio.com/v1/Rooms?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("rooms"))

				rooms := resp.Rooms
				Expect(rooms).ToNot(BeNil())
				Expect(len(rooms)).To(Equal(1))

				Expect(rooms[0].Sid).To(Equal("RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(rooms[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(rooms[0].UniqueName).To(Equal("RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(rooms[0].Status).To(Equal("in-progress"))
				Expect(rooms[0].VideoCodecs).To(Equal(&[]string{"VP8", "H264"}))
				Expect(rooms[0].MaxParticipants).To(Equal(50))
				Expect(rooms[0].RecordParticipantsOnConnect).To(Equal(false))
				Expect(rooms[0].EndTime).To(BeNil())
				Expect(rooms[0].Duration).To(BeNil())
				Expect(rooms[0].MaxConcurrentPublishedTracks).To(BeNil())
				Expect(rooms[0].StatusCallbackMethod).To(BeNil())
				Expect(rooms[0].StatusCallback).To(BeNil())
				Expect(rooms[0].Type).To(Equal("group"))
				Expect(rooms[0].MediaRegion).To(Equal(utils.String("us1")))
				Expect(rooms[0].DateCreated.Format(time.RFC3339)).To(Equal("2021-02-20T10:00:00Z"))
				Expect(rooms[0].DateUpdated).To(BeNil())
				Expect(rooms[0].URL).To(Equal("https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of rooms api returns a 500 response", func() {
			pageOptions := &rooms.RoomsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Rooms?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := roomsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the rooms page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated rooms are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Rooms",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/roomsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Rooms?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/roomsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := roomsClient.NewRoomsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated rooms current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated rooms results should be returned", func() {
				Expect(len(paginator.Rooms)).To(Equal(3))
			})
		})

		Describe("When the rooms api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Rooms",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/roomsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Rooms?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := roomsClient.NewRoomsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated rooms current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a room sid", func() {
		roomClient := videoSession.Room("RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the room resource is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/roomResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := roomClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get room resource response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.UniqueName).To(Equal("RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Status).To(Equal("in-progress"))
				Expect(resp.VideoCodecs).To(Equal(&[]string{"VP8", "H264"}))
				Expect(resp.MaxParticipants).To(Equal(50))
				Expect(resp.RecordParticipantsOnConnect).To(Equal(false))
				Expect(resp.EndTime).To(BeNil())
				Expect(resp.Duration).To(BeNil())
				Expect(resp.MaxConcurrentPublishedTracks).To(BeNil())
				Expect(resp.StatusCallbackMethod).To(BeNil())
				Expect(resp.StatusCallback).To(BeNil())
				Expect(resp.Type).To(Equal("group"))
				Expect(resp.MediaRegion).To(Equal(utils.String("us1")))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2021-02-20T10:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the room resource api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Rooms/RM71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := videoSession.Room("RM71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get room response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the room resource is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateRoomResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &room.UpdateRoomInput{
				Status: "completed",
			}

			resp, err := roomClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update room response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.UniqueName).To(Equal("RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Status).To(Equal("completed"))
				Expect(resp.VideoCodecs).To(Equal(&[]string{"VP8", "H264"}))
				Expect(resp.MaxParticipants).To(Equal(50))
				Expect(resp.RecordParticipantsOnConnect).To(Equal(false))
				Expect(resp.EndTime.Format(time.RFC3339)).To(Equal("2021-02-20T10:05:00Z"))
				Expect(resp.Duration).To(Equal(utils.Int(320)))
				Expect(resp.MaxConcurrentPublishedTracks).To(BeNil())
				Expect(resp.StatusCallbackMethod).To(BeNil())
				Expect(resp.StatusCallback).To(BeNil())
				Expect(resp.Type).To(Equal("group"))
				Expect(resp.MediaRegion).To(Equal(utils.String("us1")))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2021-02-20T10:00:00Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2021-02-20T10:05:00Z"))
				Expect(resp.URL).To(Equal("https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update room request does not contain a status", func() {
			updateInput := &room.UpdateRoomInput{}

			resp, err := roomClient.Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the update service response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the update room resource api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://video.twilio.com/v1/Rooms/RM71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &room.UpdateRoomInput{
				Status: "completed",
			}

			resp, err := videoSession.Room("RM71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update room response should be nil", func() {
				Expect(resp).To(BeNil())
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
	Expect(twilioErr.Message).To(Equal("The requested resource /Rooms/RM71 was not found"))

	moreInfo := "https://www.twilio.com/docs/errors/20404"
	Expect(twilioErr.MoreInfo).To(Equal(&moreInfo))
	Expect(twilioErr.Status).To(Equal(404))
}

func ExpectInternalServerError(err error) {
	Expect(err).ToNot(BeNil())
	twilioErr, ok := err.(*utils.TwilioError)
	Expect(ok).To(Equal(true))
	Expect(twilioErr.Code).To(BeNil())
	Expect(twilioErr.Message).To(Equal("An error occurred"))
	Expect(twilioErr.MoreInfo).To(BeNil())
	Expect(twilioErr.Status).To(Equal(500))
}

func ExpectErrorToNotBeATwilioError(err error) {
	Expect(err).ToNot(BeNil())
	_, ok := err.(*utils.TwilioError)
	Expect(ok).To(Equal(false))
}
