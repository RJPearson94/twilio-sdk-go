package acceptance

import (
	"fmt"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/RJPearson94/twilio-sdk-go"
	"github.com/RJPearson94/twilio-sdk-go/service/monitor/v1/alerts"
	"github.com/RJPearson94/twilio-sdk-go/service/monitor/v1/events"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

var _ = Describe("Monitor Acceptance Tests", func() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		Fail(fmt.Sprintf("Failed to create credentials. Error %s", err.Error()))
	}

	monitorSession := twilio.NewWithCredentials(creds).Monitor.V1

	Describe("Given the monitor alert clients", func() {
		It("Then the alerts are fetched", func() {
			alertsClient := monitorSession.Alerts

			pageResp, pageErr := alertsClient.Page(&alerts.AlertsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Alerts)).Should(BeNumerically(">=", 1))

			paginator := alertsClient.NewAlertsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Alerts)).Should(BeNumerically(">=", 1))

			alertClient := monitorSession.Alert(paginator.Alerts[0].Sid)

			fetchResp, fetchErr := alertClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())
		})
	})

	Describe("Given the monitor event clients", func() {
		It("Then the events are fetched", func() {
			eventsClient := monitorSession.Events

			pageResp, pageErr := eventsClient.Page(&events.EventsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Events)).Should(BeNumerically(">=", 1))

			paginator := eventsClient.NewEventsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Events)).Should(BeNumerically(">=", 1))

			eventClient := monitorSession.Event(paginator.Events[0].Sid)

			fetchResp, fetchErr := eventClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())
		})
	})
})
