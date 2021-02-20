package acceptance

import (
	"fmt"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/RJPearson94/twilio-sdk-go"
	"github.com/RJPearson94/twilio-sdk-go/service/video/v1/room"
	"github.com/RJPearson94/twilio-sdk-go/service/video/v1/rooms"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

var _ = Describe("Video Acceptance Tests", func() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		Fail(fmt.Sprintf("Failed to create credentials. Error %s", err.Error()))
	}

	videoSession := twilio.NewWithCredentials(creds).Video.V1

	Describe("Given the video room clients", func() {
		It("Then the room is created, fetched and updated", func() {
			roomsClient := videoSession.Rooms

			createResp, createErr := roomsClient.Create(&rooms.CreateRoomInput{
				Type: utils.String("go"),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := roomsClient.Page(&rooms.RoomsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Rooms)).Should(BeNumerically(">=", 1))

			paginator := roomsClient.NewRoomsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Rooms)).Should(BeNumerically(">=", 1))

			roomClient := videoSession.Room(createResp.Sid)

			fetchResp, fetchErr := roomClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := roomClient.Update(&room.UpdateRoomInput{
				Status: "completed",
			})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

		})
	})
})
