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
	"github.com/RJPearson94/twilio-sdk-go/service/video/v1/composition_hook"
	"github.com/RJPearson94/twilio-sdk-go/service/video/v1/composition_hooks"
	"github.com/RJPearson94/twilio-sdk-go/service/video/v1/compositions"
	"github.com/RJPearson94/twilio-sdk-go/service/video/v1/recording"
	"github.com/RJPearson94/twilio-sdk-go/service/video/v1/recordings"
	"github.com/RJPearson94/twilio-sdk-go/service/video/v1/room"
	"github.com/RJPearson94/twilio-sdk-go/service/video/v1/room/participant"
	"github.com/RJPearson94/twilio-sdk-go/service/video/v1/room/participants"
	roomRecording "github.com/RJPearson94/twilio-sdk-go/service/video/v1/room/recording"
	roomRecordings "github.com/RJPearson94/twilio-sdk-go/service/video/v1/room/recordings"
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
				Expect(resp.StatusCallbackMethod).To(Equal("POST"))
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
				Expect(rooms[0].StatusCallbackMethod).To(Equal("POST"))
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
				Expect(resp.StatusCallbackMethod).To(Equal("POST"))
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
				Expect(resp.StatusCallbackMethod).To(Equal("POST"))
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

			It("Then the update room response should be nil", func() {
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

	Describe("Given I have a recordings client", func() {
		recordingsClient := videoSession.Recordings

		Describe("When the page of recordings are successfully retrieved", func() {
			pageOptions := &recordings.RecordingsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Recordings?Page=0&PageSize=50",
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

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://video.twilio.com/v1/Recordings?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://video.twilio.com/v1/Recordings?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("recordings"))

				recordingGroupingSidsResponse := recordings.PageRecordingGroupingSidsResponse{
					RoomSid:        "RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
					ParticipantSid: "PAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				}

				recordings := resp.Recordings
				Expect(recordings).ToNot(BeNil())
				Expect(len(recordings)).To(Equal(1))

				Expect(recordings[0].Sid).To(Equal("RTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(recordings[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(recordings[0].Status).To(Equal("completed"))
				Expect(recordings[0].GroupingSids).To(Equal(recordingGroupingSidsResponse))
				Expect(recordings[0].ContainerFormat).To(Equal("mka"))
				Expect(recordings[0].TrackName).To(Equal("test"))
				Expect(recordings[0].Offset).To(Equal(171092213859))
				Expect(recordings[0].Codec).To(Equal("opus"))
				Expect(recordings[0].SourceSid).To(Equal("MTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(recordings[0].Duration).To(Equal(15))
				Expect(recordings[0].Type).To(Equal("audio"))
				Expect(recordings[0].RoomSid).To(Equal("RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(recordings[0].Size).To(Equal(4234))
				Expect(recordings[0].DateCreated.Format(time.RFC3339)).To(Equal("2021-02-20T10:00:00Z"))
				Expect(recordings[0].URL).To(Equal("https://video.twilio.com/v1/Recordings/RTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of recordings api returns a 500 response", func() {
			pageOptions := &recordings.RecordingsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Recordings?Page=0&PageSize=50",
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
			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Recordings",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/recordingsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Recordings?Page=1&PageSize=50&PageToken=abc1234",
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
			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Recordings",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/recordingsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Recordings?Page=1&PageSize=50&PageToken=abc1234",
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
		recordingClient := videoSession.Recording("RTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the recording resource is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Recordings/RTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
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

			It("Then the get recording resource response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("RTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Status).To(Equal("completed"))
				Expect(resp.GroupingSids).To(Equal(recording.FetchRecordingGroupingSidsResponse{
					RoomSid:        "RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
					ParticipantSid: "PAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				}))
				Expect(resp.ContainerFormat).To(Equal("mka"))
				Expect(resp.TrackName).To(Equal("test"))
				Expect(resp.Offset).To(Equal(171092213859))
				Expect(resp.Codec).To(Equal("opus"))
				Expect(resp.SourceSid).To(Equal("MTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Duration).To(Equal(15))
				Expect(resp.Type).To(Equal("audio"))
				Expect(resp.RoomSid).To(Equal("RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Size).To(Equal(4234))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2021-02-20T10:00:00Z"))
				Expect(resp.URL).To(Equal("https://video.twilio.com/v1/Recordings/RTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the recording resource api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Recordings/RT71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := videoSession.Recording("RT71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get recording response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the recording resource is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://video.twilio.com/v1/Recordings/RTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := recordingClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the recording resource api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://video.twilio.com/v1/Recordings/RT71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := videoSession.Recording("RT71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given I have a room recordings client", func() {
		roomRecordingsClient := videoSession.Room("RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Recordings

		Describe("When the page of room recordings are successfully retrieved", func() {
			pageOptions := &roomRecordings.RecordingsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/roomRecordingsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := roomRecordingsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the room recordings page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("recordings"))

				recordingGroupingSidsResponse := roomRecordings.PageRecordingGroupingSidsResponse{
					RoomSid:        "RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
					ParticipantSid: "PAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				}

				recordings := resp.Recordings
				Expect(recordings).ToNot(BeNil())
				Expect(len(recordings)).To(Equal(1))

				Expect(recordings[0].Sid).To(Equal("RTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(recordings[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(recordings[0].Status).To(Equal("completed"))
				Expect(recordings[0].GroupingSids).To(Equal(recordingGroupingSidsResponse))
				Expect(recordings[0].ContainerFormat).To(Equal("mka"))
				Expect(recordings[0].TrackName).To(Equal("test"))
				Expect(recordings[0].Offset).To(Equal(171092213859))
				Expect(recordings[0].Codec).To(Equal("opus"))
				Expect(recordings[0].SourceSid).To(Equal("MTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(recordings[0].Duration).To(Equal(15))
				Expect(recordings[0].Type).To(Equal("audio"))
				Expect(recordings[0].RoomSid).To(Equal("RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(recordings[0].Size).To(Equal(4234))
				Expect(recordings[0].DateCreated.Format(time.RFC3339)).To(Equal("2021-02-20T10:00:00Z"))
				Expect(recordings[0].URL).To(Equal("https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings/RTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of room recordings api returns a 500 response", func() {
			pageOptions := &roomRecordings.RecordingsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := roomRecordingsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the room recordings page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated room recordings are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/roomRecordingsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/roomRecordingsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := roomRecordingsClient.NewRecordingsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated room recordings current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated room recordings results should be returned", func() {
				Expect(len(paginator.Recordings)).To(Equal(3))
			})
		})

		Describe("When the room recordings api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/roomRecordingsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := roomRecordingsClient.NewRecordingsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated room recordings current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a room recording sid", func() {
		roomRecordingClient := videoSession.Room("RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Recording("RTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the room recording resource is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings/RTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/roomRecordingResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := roomRecordingClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get room recording resource response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("RTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Status).To(Equal("completed"))
				Expect(resp.GroupingSids).To(Equal(roomRecording.FetchRecordingGroupingSidsResponse{
					RoomSid:        "RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
					ParticipantSid: "PAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				}))
				Expect(resp.ContainerFormat).To(Equal("mka"))
				Expect(resp.TrackName).To(Equal("test"))
				Expect(resp.Offset).To(Equal(171092213859))
				Expect(resp.Codec).To(Equal("opus"))
				Expect(resp.SourceSid).To(Equal("MTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Duration).To(Equal(15))
				Expect(resp.Type).To(Equal("audio"))
				Expect(resp.RoomSid).To(Equal("RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Size).To(Equal(4234))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2021-02-20T10:00:00Z"))
				Expect(resp.URL).To(Equal("https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings/RTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the room recording resource api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings/RT71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := videoSession.Room("RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Recording("RT71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get room recording response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the room recording resource is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings/RTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := roomRecordingClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the room recording resource api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings/RT71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := videoSession.Room("RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Recording("RT71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given I have a compositions client", func() {
		compositionsClient := videoSession.Compositions

		Describe("When the composition resource is successfully created", func() {
			createInput := &compositions.CreateCompositionInput{
				RoomSid: "RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			}

			httpmock.RegisterResponder("POST", "https://video.twilio.com/v1/Compositions",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/compositionResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := compositionsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create composition response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("CJXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Status).To(Equal("enqueued"))
				Expect(resp.Trim).To(Equal(true))
				Expect(resp.VideoLayout).To(Equal(map[string]interface{}{}))
				Expect(resp.DateCompleted).To(BeNil())
				Expect(resp.Format).To(Equal("mp4"))
				Expect(resp.DateDeleted).To(BeNil())
				Expect(resp.Duration).To(Equal(0))
				Expect(resp.Bitrate).To(Equal(0))
				Expect(resp.AudioSources).To(Equal([]string{"RTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"}))
				Expect(resp.RoomSid).To(Equal("RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Size).To(Equal(0))
				Expect(resp.Resolution).To(Equal("640x480"))
				Expect(resp.AudioSourcesExcluded).To(Equal([]string{}))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2021-02-20T10:00:00Z"))
				Expect(resp.URL).To(Equal("https://video.twilio.com/v1/Compositions/CJXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the create composition request does not contain a room sid", func() {
			createInput := &compositions.CreateCompositionInput{}

			resp, err := compositionsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create composition response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create compositions api returns a 500 response", func() {
			createInput := &compositions.CreateCompositionInput{
				RoomSid: "RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			}

			httpmock.RegisterResponder("POST", "https://video.twilio.com/v1/Compositions",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := compositionsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create composition response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of compositions are successfully retrieved", func() {
			pageOptions := &compositions.CompositionsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Compositions?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/compositionsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := compositionsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the compositions page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://video.twilio.com/v1/Compositions?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://video.twilio.com/v1/Compositions?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("compositions"))

				compositions := resp.Compositions
				Expect(compositions).ToNot(BeNil())
				Expect(len(compositions)).To(Equal(1))

				Expect(compositions[0].Sid).To(Equal("CJXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(compositions[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(compositions[0].Status).To(Equal("enqueued"))
				Expect(compositions[0].Trim).To(Equal(true))
				Expect(compositions[0].VideoLayout).To(Equal(map[string]interface{}{}))
				Expect(compositions[0].DateCompleted).To(BeNil())
				Expect(compositions[0].Format).To(Equal("mp4"))
				Expect(compositions[0].DateDeleted).To(BeNil())
				Expect(compositions[0].Duration).To(Equal(0))
				Expect(compositions[0].AudioSources).To(Equal([]string{"RTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"}))
				Expect(compositions[0].RoomSid).To(Equal("RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(compositions[0].Size).To(Equal(0))
				Expect(compositions[0].Resolution).To(Equal("640x480"))
				Expect(compositions[0].AudioSourcesExcluded).To(Equal([]string{}))
				Expect(compositions[0].DateCreated.Format(time.RFC3339)).To(Equal("2021-02-20T10:00:00Z"))
				Expect(compositions[0].URL).To(Equal("https://video.twilio.com/v1/Compositions/CJXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of compositions api returns a 500 response", func() {
			pageOptions := &compositions.CompositionsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Compositions?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := compositionsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the compositions page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated compositions are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Compositions",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/compositionsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Compositions?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/compositionsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := compositionsClient.NewCompositionsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated compositions current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated compositions results should be returned", func() {
				Expect(len(paginator.Compositions)).To(Equal(3))
			})
		})

		Describe("When the compositions api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Compositions",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/compositionsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Compositions?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := compositionsClient.NewCompositionsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated compositions current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a composition sid", func() {
		compositionClient := videoSession.Composition("CJXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the composition resource is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Compositions/CJXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/compositionResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := compositionClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get composition resource response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("CJXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Status).To(Equal("enqueued"))
				Expect(resp.Trim).To(Equal(true))
				Expect(resp.VideoLayout).To(Equal(map[string]interface{}{}))
				Expect(resp.DateCompleted).To(BeNil())
				Expect(resp.Format).To(Equal("mp4"))
				Expect(resp.DateDeleted).To(BeNil())
				Expect(resp.Duration).To(Equal(0))
				Expect(resp.Bitrate).To(Equal(0))
				Expect(resp.AudioSources).To(Equal([]string{"RTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"}))
				Expect(resp.RoomSid).To(Equal("RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Size).To(Equal(0))
				Expect(resp.Resolution).To(Equal("640x480"))
				Expect(resp.AudioSourcesExcluded).To(Equal([]string{}))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2021-02-20T10:00:00Z"))
				Expect(resp.URL).To(Equal("https://video.twilio.com/v1/Compositions/CJXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the composition resource api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Compositions/CJ71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := videoSession.Composition("CJ71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get composition response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the composition resource is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://video.twilio.com/v1/Compositions/CJXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := compositionClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the composition resource api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://video.twilio.com/v1/Compositions/CJ71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := videoSession.Composition("CJ71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given I have a composition hooks client", func() {
		compositionHooksClient := videoSession.CompositionHooks

		Describe("When the composition hooks resource is successfully created", func() {
			createInput := &composition_hooks.CreateCompositionHookInput{
				FriendlyName: "Test",
			}

			httpmock.RegisterResponder("POST", "https://video.twilio.com/v1/CompositionHooks",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/compositionHookResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := compositionHooksClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create composition hooks response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("HKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("test"))
				Expect(resp.Trim).To(Equal(true))
				Expect(resp.VideoLayout).To(Equal(map[string]interface{}{}))
				Expect(resp.Format).To(Equal("mp4"))
				Expect(resp.AudioSources).To(Equal([]string{"*"}))
				Expect(resp.Resolution).To(Equal("640x480"))
				Expect(resp.Enabled).To(Equal(true))
				Expect(resp.StatusCallback).To(BeNil())
				Expect(resp.StatusCallbackMethod).To(Equal("POST"))
				Expect(resp.AudioSourcesExcluded).To(Equal([]string{}))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2021-02-20T10:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://video.twilio.com/v1/CompositionHooks/HKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the create composition hooks request does not contain a friendly name", func() {
			createInput := &composition_hooks.CreateCompositionHookInput{}

			resp, err := compositionHooksClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create composition hook response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create composition hooks api returns a 500 response", func() {
			createInput := &composition_hooks.CreateCompositionHookInput{
				FriendlyName: "Test",
			}

			httpmock.RegisterResponder("POST", "https://video.twilio.com/v1/CompositionHooks",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := compositionHooksClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create composition hook response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of composition hooks are successfully retrieved", func() {
			pageOptions := &composition_hooks.CompositionHooksPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/CompositionHooks?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/compositionHooksPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := compositionHooksClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the composition hooks page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://video.twilio.com/v1/CompositionHooks?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://video.twilio.com/v1/CompositionHooks?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("composition_hooks"))

				compositionHooks := resp.CompositionHooks
				Expect(compositionHooks).ToNot(BeNil())
				Expect(len(compositionHooks)).To(Equal(1))

				Expect(compositionHooks[0].Sid).To(Equal("HKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(compositionHooks[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(compositionHooks[0].FriendlyName).To(Equal("test"))
				Expect(compositionHooks[0].Trim).To(Equal(true))
				Expect(compositionHooks[0].VideoLayout).To(Equal(map[string]interface{}{}))
				Expect(compositionHooks[0].Format).To(Equal("mp4"))
				Expect(compositionHooks[0].AudioSources).To(Equal([]string{"*"}))
				Expect(compositionHooks[0].Resolution).To(Equal("640x480"))
				Expect(compositionHooks[0].Enabled).To(Equal(true))
				Expect(compositionHooks[0].StatusCallback).To(BeNil())
				Expect(compositionHooks[0].StatusCallbackMethod).To(Equal("POST"))
				Expect(compositionHooks[0].AudioSourcesExcluded).To(Equal([]string{}))
				Expect(compositionHooks[0].DateCreated.Format(time.RFC3339)).To(Equal("2021-02-20T10:00:00Z"))
				Expect(compositionHooks[0].DateUpdated).To(BeNil())
				Expect(compositionHooks[0].URL).To(Equal("https://video.twilio.com/v1/CompositionHooks/HKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of composition hooks api returns a 500 response", func() {
			pageOptions := &composition_hooks.CompositionHooksPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/CompositionHooks?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := compositionHooksClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the composition hooks page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated compositionHooks are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/CompositionHooks",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/compositionHooksPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/CompositionHooks?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/compositionHooksPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := compositionHooksClient.NewCompositionHooksPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated composition hooks current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated composition hooks results should be returned", func() {
				Expect(len(paginator.CompositionHooks)).To(Equal(3))
			})
		})

		Describe("When the composition hooks api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/CompositionHooks",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/compositionHooksPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/CompositionHooks?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := compositionHooksClient.NewCompositionHooksPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated composition hooks current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a composition hook sid", func() {
		compositionHookClient := videoSession.CompositionHook("HKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the composition hook resource is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/CompositionHooks/HKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/compositionHookResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := compositionHookClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get composition hook resource response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("HKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("test"))
				Expect(resp.Trim).To(Equal(true))
				Expect(resp.VideoLayout).To(Equal(map[string]interface{}{}))
				Expect(resp.Format).To(Equal("mp4"))
				Expect(resp.AudioSources).To(Equal([]string{"*"}))
				Expect(resp.Resolution).To(Equal("640x480"))
				Expect(resp.Enabled).To(Equal(true))
				Expect(resp.StatusCallback).To(BeNil())
				Expect(resp.StatusCallbackMethod).To(Equal("POST"))
				Expect(resp.AudioSourcesExcluded).To(Equal([]string{}))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2021-02-20T10:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://video.twilio.com/v1/CompositionHooks/HKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the composition hook resource api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/CompositionHooks/HK71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := videoSession.CompositionHook("HK71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get composition hook response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the composition hook resource is successfully updated", func() {
			updateInput := &composition_hook.UpdateCompositionHookInput{
				FriendlyName: "Test",
			}

			httpmock.RegisterResponder("POST", "https://video.twilio.com/v1/CompositionHooks/HKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateCompositionHookResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := compositionHookClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update composition hook resource response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("HKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("new test"))
				Expect(resp.Trim).To(Equal(true))
				Expect(resp.VideoLayout).To(Equal(map[string]interface{}{}))
				Expect(resp.Format).To(Equal("mp4"))
				Expect(resp.AudioSources).To(Equal([]string{"*"}))
				Expect(resp.Resolution).To(Equal("640x480"))
				Expect(resp.Enabled).To(Equal(true))
				Expect(resp.StatusCallback).To(BeNil())
				Expect(resp.StatusCallbackMethod).To(Equal("POST"))
				Expect(resp.AudioSourcesExcluded).To(Equal([]string{}))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2021-02-20T10:00:00Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2021-02-20T10:05:00Z"))
				Expect(resp.URL).To(Equal("https://video.twilio.com/v1/CompositionHooks/HKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update composition hooks request does not contain a friendly name", func() {
			updateInput := &composition_hook.UpdateCompositionHookInput{}

			resp, err := compositionHookClient.Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the update composition hook response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the composition hook resource api returns a 404", func() {
			updateInput := &composition_hook.UpdateCompositionHookInput{
				FriendlyName: "Test",
			}

			httpmock.RegisterResponder("POST", "https://video.twilio.com/v1/CompositionHooks/HK71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := videoSession.CompositionHook("HK71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update composition hook response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the composition hook resource is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://video.twilio.com/v1/CompositionHooks/HKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := compositionHookClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the composition hook resource api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://video.twilio.com/v1/CompositionHooks/HK71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := videoSession.CompositionHook("HK71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given I have a participant client", func() {
		participantsClient := videoSession.Room("RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Participants

		Describe("When the page of participants are successfully retrieved", func() {
			pageOptions := &participants.ParticipantsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants?Page=0&PageSize=50",
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

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("participants"))

				participants := resp.Participants
				Expect(participants).ToNot(BeNil())
				Expect(len(participants)).To(Equal(1))

				Expect(participants[0].Sid).To(Equal("PAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(participants[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(participants[0].Status).To(Equal("in-progress"))
				Expect(participants[0].StartTime.Format(time.RFC3339)).To(Equal("2021-02-20T10:00:00Z"))
				Expect(participants[0].Duration).To(BeNil())
				Expect(participants[0].EndTime).To(BeNil())
				Expect(participants[0].RoomSid).To(Equal("RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(participants[0].Identity).To(Equal("Test"))
				Expect(participants[0].DateCreated.Format(time.RFC3339)).To(Equal("2021-02-20T10:00:00Z"))
				Expect(participants[0].DateUpdated).To(BeNil())
				Expect(participants[0].URL).To(Equal("https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/PAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of participants api returns a 500 response", func() {
			pageOptions := &participants.ParticipantsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants?Page=0&PageSize=50",
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
			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/participantsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants?Page=1&PageSize=50&PageToken=abc1234",
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
			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/participantsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants?Page=1&PageSize=50&PageToken=abc1234",
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
		participantClient := videoSession.Room("RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Participant("PAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the participant resource is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/PAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
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

			It("Then the get participant resource response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("PAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Status).To(Equal("in-progress"))
				Expect(resp.StartTime.Format(time.RFC3339)).To(Equal("2021-02-20T10:00:00Z"))
				Expect(resp.Duration).To(BeNil())
				Expect(resp.EndTime).To(BeNil())
				Expect(resp.RoomSid).To(Equal("RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Identity).To(Equal("Test"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2021-02-20T10:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/PAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the participant resource api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/PA71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := videoSession.Room("RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Participant("PA71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get participant response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the participant resource is successfully updated", func() {
			updateInput := &participant.UpdateParticipantInput{
				Status: utils.String("completed"),
			}

			httpmock.RegisterResponder("POST", "https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/PAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateParticipantResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := participantClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update participant resource response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("PAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Status).To(Equal("disconnected"))
				Expect(resp.StartTime.Format(time.RFC3339)).To(Equal("2021-02-20T10:00:00Z"))
				Expect(resp.Duration).To(Equal(utils.Int(300)))
				Expect(resp.EndTime.Format(time.RFC3339)).To(Equal("2021-02-20T10:05:00Z"))
				Expect(resp.RoomSid).To(Equal("RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Identity).To(Equal("Test"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2021-02-20T10:00:00Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2021-02-20T10:05:00Z"))
				Expect(resp.URL).To(Equal("https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/PAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the participant resource api returns a 404", func() {
			updateInput := &participant.UpdateParticipantInput{
				Status: utils.String("completed"),
			}

			httpmock.RegisterResponder("POST", "https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/PA71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := videoSession.Room("RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Participant("PA71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update participant response should be nil", func() {
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
