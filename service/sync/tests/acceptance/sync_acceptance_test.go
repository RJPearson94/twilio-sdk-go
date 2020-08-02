package acceptance

import (
	"fmt"
	"log"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/RJPearson94/twilio-sdk-go"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/document"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/documents"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/sync_list"
	syncListItem "github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/sync_list/item"
	syncListItems "github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/sync_list/items"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/sync_lists"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/sync_map"
	syncMapItem "github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/sync_map/item"
	syncMapItems "github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/sync_map/items"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/sync_maps"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/sync_stream"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/sync_stream/messages"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service/sync_streams"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/services"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

var _ = Describe("Sync Acceptance Tests", func() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err)
	}

	syncSession := twilio.NewWithCredentials(creds).Sync.V1

	Describe("Given the Sync Service clients", func() {
		It("Then the service is created, fetched, updated and deleted", func() {
			createResp, createErr := syncSession.Services.Create(&services.CreateServiceInput{})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			syncClient := syncSession.Service(createResp.Sid)

			fetchResp, fetchErr := syncClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := syncClient.Update(&service.UpdateServiceInput{})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := syncClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the Sync Document clients", func() {

		var serviceSid string

		BeforeEach(func() {
			resp, err := syncSession.Services.Create(&services.CreateServiceInput{})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create service. Error %s", err.Error()))
			}
			serviceSid = resp.Sid
		})

		AfterEach(func() {
			if err := syncSession.Service(serviceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete service. Error %s", err.Error()))
			}
		})

		It("Then the document is created, fetched, updated and deleted", func() {
			createResp, createErr := syncSession.Service(serviceSid).Documents.Create(&documents.CreateDocumentInput{})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			documentClient := syncSession.Service(serviceSid).Document(createResp.Sid)

			fetchResp, fetchErr := documentClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := documentClient.Update(&document.UpdateDocumentInput{
				Data: utils.String("{\"message\":\"Hello World\"}"),
			})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := documentClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the Sync List clients", func() {

		var serviceSid string

		BeforeEach(func() {
			resp, err := syncSession.Services.Create(&services.CreateServiceInput{})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create service. Error %s", err.Error()))
			}
			serviceSid = resp.Sid
		})

		AfterEach(func() {
			if err := syncSession.Service(serviceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete service. Error %s", err.Error()))
			}
		})

		It("Then the sync list is created, fetched, updated and deleted", func() {
			createResp, createErr := syncSession.Service(serviceSid).SyncLists.Create(&sync_lists.CreateSyncListInput{})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			syncListClient := syncSession.Service(serviceSid).SyncList(createResp.Sid)

			fetchResp, fetchErr := syncListClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := syncListClient.Update(&sync_list.UpdateSyncListInput{
				CollectionTtl: utils.Int(1),
			})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := syncListClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the Sync List Item clients", func() {

		var serviceSid string
		var syncListSid string

		BeforeEach(func() {
			resp, err := syncSession.Services.Create(&services.CreateServiceInput{})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create service. Error %s", err.Error()))
			}
			serviceSid = resp.Sid

			syncListResp, syncListErr := syncSession.Service(serviceSid).SyncLists.Create(&sync_lists.CreateSyncListInput{})
			if syncListErr != nil {
				Fail(fmt.Sprintf("Failed to create sync list. Error %s", syncListErr.Error()))
			}
			syncListSid = syncListResp.Sid
		})

		AfterEach(func() {
			if err := syncSession.Service(serviceSid).SyncList(syncListSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete sync list. Error %s", err.Error()))
			}

			if err := syncSession.Service(serviceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete service. Error %s", err.Error()))
			}
		})

		It("Then the sync list item is created, fetched, updated and deleted", func() {
			createResp, createErr := syncSession.Service(serviceSid).SyncList(syncListSid).Items.Create(&syncListItems.CreateSyncListItemInput{
				Data: "{}",
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Index).ToNot(BeNil())

			syncListItemClient := syncSession.Service(serviceSid).SyncList(syncListSid).Item(createResp.Index)

			fetchResp, fetchErr := syncListItemClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := syncListItemClient.Update(&syncListItem.UpdateSyncListItemInput{
				Data: utils.String("{\"message\":\"Hello World\"}"),
			})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := syncListItemClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the Sync Map clients", func() {

		var serviceSid string

		BeforeEach(func() {
			resp, err := syncSession.Services.Create(&services.CreateServiceInput{})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create service. Error %s", err.Error()))
			}
			serviceSid = resp.Sid
		})

		AfterEach(func() {
			if err := syncSession.Service(serviceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete service. Error %s", err.Error()))
			}
		})

		It("Then the sync map is created, fetched, updated and deleted", func() {
			createResp, createErr := syncSession.Service(serviceSid).SyncMaps.Create(&sync_maps.CreateSyncMapInput{})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			syncMapClient := syncSession.Service(serviceSid).SyncMap(createResp.Sid)

			fetchResp, fetchErr := syncMapClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := syncMapClient.Update(&sync_map.UpdateSyncMapInput{
				CollectionTtl: utils.Int(1),
			})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := syncMapClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the Sync Map Item clients", func() {

		var serviceSid string
		var syncMapSid string

		BeforeEach(func() {
			resp, err := syncSession.Services.Create(&services.CreateServiceInput{})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create service. Error %s", err.Error()))
			}
			serviceSid = resp.Sid

			syncMapResp, syncMapErr := syncSession.Service(serviceSid).SyncMaps.Create(&sync_maps.CreateSyncMapInput{})
			if syncMapErr != nil {
				Fail(fmt.Sprintf("Failed to create sync map. Error %s", syncMapErr.Error()))
			}
			syncMapSid = syncMapResp.Sid
		})

		AfterEach(func() {
			if err := syncSession.Service(serviceSid).SyncMap(syncMapSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete sync map. Error %s", err.Error()))
			}

			if err := syncSession.Service(serviceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete service. Error %s", err.Error()))
			}
		})

		It("Then the sync map item is created, fetched, updated and deleted", func() {
			createResp, createErr := syncSession.Service(serviceSid).SyncMap(syncMapSid).Items.Create(&syncMapItems.CreateSyncMapItemInput{
				Data: "{}",
				Key:  "test",
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Key).ToNot(BeNil())

			syncMapItemClient := syncSession.Service(serviceSid).SyncMap(syncMapSid).Item(createResp.Key)

			fetchResp, fetchErr := syncMapItemClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := syncMapItemClient.Update(&syncMapItem.UpdateSyncMapItemInput{
				Data: utils.String("{\"message\":\"Hello World\"}"),
			})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := syncMapItemClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the Sync Stream clients", func() {

		var serviceSid string

		BeforeEach(func() {
			resp, err := syncSession.Services.Create(&services.CreateServiceInput{})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create service. Error %s", err.Error()))
			}
			serviceSid = resp.Sid
		})

		AfterEach(func() {
			if err := syncSession.Service(serviceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete service. Error %s", err.Error()))
			}
		})

		It("Then the sync stream is created, fetched, updated and deleted", func() {
			createResp, createErr := syncSession.Service(serviceSid).SyncStreams.Create(&sync_streams.CreateSyncStreamInput{})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			syncStreamClient := syncSession.Service(serviceSid).SyncStream(createResp.Sid)

			fetchResp, fetchErr := syncStreamClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := syncStreamClient.Update(&sync_stream.UpdateSyncStreamInput{
				Ttl: utils.Int(1),
			})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := syncStreamClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the Sync Stream Message client", func() {

		var serviceSid string
		var syncStreamSid string

		BeforeEach(func() {
			resp, err := syncSession.Services.Create(&services.CreateServiceInput{})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create service. Error %s", err.Error()))
			}
			serviceSid = resp.Sid

			syncStreamResp, syncStreamErr := syncSession.Service(serviceSid).SyncStreams.Create(&sync_streams.CreateSyncStreamInput{})
			if syncStreamErr != nil {
				Fail(fmt.Sprintf("Failed to create sync stream. Error %s", syncStreamErr.Error()))
			}
			syncStreamSid = syncStreamResp.Sid
		})

		AfterEach(func() {
			if err := syncSession.Service(serviceSid).SyncStream(syncStreamSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete sync stream. Error %s", err.Error()))
			}

			if err := syncSession.Service(serviceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete service. Error %s", err.Error()))
			}
		})

		It("Then the sync stream message is created", func() {
			createResp, createErr := syncSession.Service(serviceSid).SyncStream(syncStreamSid).Messages.Create(&messages.CreateSyncStreamMessageInput{
				Data: "{}",
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())
		})
	})

	// TODO add sync map, sync list & document permissions
})
