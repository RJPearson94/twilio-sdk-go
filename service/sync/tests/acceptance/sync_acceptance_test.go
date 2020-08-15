package acceptance

import (
	"fmt"
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
		Fail(fmt.Sprintf("Failed to create credentials. Error %s", err.Error()))
	}

	syncSession := twilio.NewWithCredentials(creds).Sync.V1

	Describe("Given the Sync Service clients", func() {
		It("Then the service is created, fetched, updated and deleted", func() {
			servicesClient := syncSession.Services

			createResp, createErr := servicesClient.Create(&services.CreateServiceInput{})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := servicesClient.Page(&services.ServicesPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Services)).Should(BeNumerically(">=", 1))

			paginator := servicesClient.NewServicesPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Services)).Should(BeNumerically(">=", 1))

			serviceClient := syncSession.Service(createResp.Sid)

			fetchResp, fetchErr := serviceClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := serviceClient.Update(&service.UpdateServiceInput{})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := serviceClient.Delete()
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
			documentsClient := syncSession.Service(serviceSid).Documents

			createResp, createErr := documentsClient.Create(&documents.CreateDocumentInput{})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := documentsClient.Page(&documents.DocumentsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Documents)).Should(BeNumerically(">=", 1))

			paginator := documentsClient.NewDocumentsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Documents)).Should(BeNumerically(">=", 1))

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
			syncListsClient := syncSession.Service(serviceSid).SyncLists

			createResp, createErr := syncListsClient.Create(&sync_lists.CreateSyncListInput{})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := syncListsClient.Page(&sync_lists.SyncListsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.SyncLists)).Should(BeNumerically(">=", 1))

			paginator := syncListsClient.NewSyncListsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.SyncLists)).Should(BeNumerically(">=", 1))

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
			syncListItemsClient := syncSession.Service(serviceSid).SyncList(syncListSid).Items

			createResp, createErr := syncListItemsClient.Create(&syncListItems.CreateSyncListItemInput{
				Data: "{}",
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Index).ToNot(BeNil())

			pageResp, pageErr := syncListItemsClient.Page(&syncListItems.SyncListItemsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.SyncListItems)).Should(BeNumerically(">=", 1))

			paginator := syncListItemsClient.NewSyncListItemsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.SyncListItems)).Should(BeNumerically(">=", 1))

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
			syncMapsClient := syncSession.Service(serviceSid).SyncMaps

			createResp, createErr := syncMapsClient.Create(&sync_maps.CreateSyncMapInput{})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := syncMapsClient.Page(&sync_maps.SyncMapsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.SyncMaps)).Should(BeNumerically(">=", 1))

			paginator := syncMapsClient.NewSyncMapsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.SyncMaps)).Should(BeNumerically(">=", 1))

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
			syncMapItemsClient := syncSession.Service(serviceSid).SyncMap(syncMapSid).Items

			createResp, createErr := syncMapItemsClient.Create(&syncMapItems.CreateSyncMapItemInput{
				Data: "{}",
				Key:  "test",
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Key).ToNot(BeNil())

			pageResp, pageErr := syncMapItemsClient.Page(&syncMapItems.SyncMapItemsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.SyncMapItems)).Should(BeNumerically(">=", 1))

			paginator := syncMapItemsClient.NewSyncMapItemsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.SyncMapItems)).Should(BeNumerically(">=", 1))

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
			syncStreamsClient := syncSession.Service(serviceSid).SyncStreams

			createResp, createErr := syncStreamsClient.Create(&sync_streams.CreateSyncStreamInput{})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := syncStreamsClient.Page(&sync_streams.SyncStreamsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.SyncStreams)).Should(BeNumerically(">=", 1))

			paginator := syncStreamsClient.NewSyncStreamsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.SyncStreams)).Should(BeNumerically(">=", 1))

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
