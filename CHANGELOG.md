## v0.14.0 (2021-02-18)

NOTES

Add missing beta tag to Flex Plugin Service resources

FIXES

### Flex

- Replace `Disabled` field with an `Archived` field on `CreatePluginResponse`, `PagePluginResponse`, `FetchPluginResponse` and `UpdatePluginResponse` structs, this matches the updated Twilio API response **Breaking Change**
- Replace `Disabled` field with an `Archived` field on `CreateVersionResponse`, `PageVersionResponse` and `FetchVersionResponse` structs, this matches the updated Twilio API response **Breaking Change**
- Replace `Disabled` field with an `Archived` field on `CreateConfigurationResponse`, `PageConfigurationResponse` and `FetchConfigurationResponse` structs, this matches the updated Twilio API response **Breaking Change**

## v0.13.0 (2021-02-16)

FIXES

### TwiML

- Correct StatusCallback xml attribute name `statusCallback` on voice response stream verb
- Correct StatusCallbackMethod xml attribute name `statusCallbackMethod` on voice response stream verb

### API

- Rename `BYOC` field to `Byoc` on CreateCallInput struct to keep consistency with the rest of the SDK **Breaking Change**
- Rename `BYOC` field to `Byoc` on CreateParticipantInput struct to keep consistency with the rest of the SDK **Breaking Change**

### Monitor

- Rename `SourceIPAddress` field to `SourceIpAddress` on Monitor response structs to keep consistency with the rest of the SDK **Breaking Change**
- Correct Source IP Address query param name on EventsPageOptions, was `SourceIPAddress`, now `SourceIpAddress`

FEATURES

### API

- **New Resource:** SIP Credential Lists
- **New Resource:** SIP Credentials
- **New Resource:** SIP IP Access Control Lists
- **New Resource:** SIP IP Addresses
- **New Resource:** SIP Domain
- **New Resource:** SIP Credential List Mapping
- **New Resource:** SIP Registration Credential List Mapping
- **New Resource:** SIP IP Access Control List Mapping

### Serverless

- Add `Runtime` field to serverless build input and response structs

### Trunking

- **New Resource:** Credential Lists
- **New Resource:** IP Access Control Lists

### TwiML

- Add `referMethod` and `referUrl` attributes to the voice response dial verb

## v0.12.0 (2021-02-08)

NOTES

New notify resources do not support Twilio deprecated channels (Alexa & Facebook Messenger), Twilio deprecated resources i.e. segments, etc. and Google Cloud Messaging (GCM) which is a deprecated by Google

FEATURES

### Notify

- **New Resource:** Service
- **New Resource:** Credentials
- **New Resource:** Bindings
- **New Resource:** Notification

### TwiML

- Add `status_callback` and `status_callback_method` attributes to the voice response stream noun

## v0.11.0 (2021-01-23)

### Conversations

- Add `ReachabilityEnabled` response field to service configuration fetch and update response structs
- Add `ReachabilityEnabled` input field to service configuration update input structs
- Add `UniqueName` response field to conversation & service conversation create, fetch and update response structs
- Add `UniqueName` input field to conversation & service conversation create and update input structs

## v0.10.0 (2020-12-23)

NOTES

- Refactor of internal clients to allow user defined config i.e. enabling debug mode to be more easily passed into the SDK.
- Remove `NewWithCredentials` functions from api clients **Breaking Change**
- Add `Config` parameter to New function of each api client **Breaking Change**

FEATURES

- Add support for edge and region configuration for all api operations except Serverless create asset version and Serverless create function version
- Default region to `us1` if an edge is specified but the region is nil

## v0.9.0 (2020-12-20)

FIXES

### Conversations

- Remove `Type` response field from user and service user create, fetch and update response structs as this value is not returned from the API **Breaking Change**

### TwiML

- Rename `DialRecord` attribute to `Record` on dial verb to keep consistency with other Twilio SDKs **Breaking Change**

FEATURES

### API

- Add `RecordingTrack` to call, call recording, conference recording and conference participants input structs

### Autopilot

- **New Resource:** Dialogue

### Conversations

- Add `LastReadMessageIndex` and `LastReadTimestamp` to participant and service participant create, fetch and update response structs
- Add `IsNotifiable` response field to user and service user create, fetch and update response structs

### Flex

- **New Resource:** Plugin
- **New Resource:** Plugin Version
- **New Resource:** Plugin Configuration
- **New Resource:** Configuration Plugins
- **New Resource:** Plugin Release

### Sync

- Add HideExpired query parameter to sync list, sync list item, sync map, sync map item, sync stream and document page requests

### TwiML

- **New Attribute:** Add recording track to dial verb
- **New Attribute:** Add sequential to dial verb

## v0.8.0 (2020-12-10)

FEATURES

### Trunking

- **New Resource:** Recording
- **New Resource:** Phone Number (Fetch API Operation)

## v0.7.0 (2020-12-09)

FIXES

### API

- **Fix:** Update Voice Receive Mode and Fax Capability to be optional in the IncomingPhoneNumbers response payload as these are no longer returned from the API call since Programmable Fax has been disabled for some accounts **Breaking Change**
- **Fix:** Update Voice Receive Mode and Fax Capability to be optional in the Local AvailablePhoneNumber response payload as these are no longer returned from the API call since Programmable Fax has been disabled for some accounts **Breaking Change**
- **Fix:** Update Voice Receive Mode and Fax Capability to be optional in the Mobile AvailablePhoneNumber response payload as these are no longer returned from the API call since Programmable Fax has been disabled for some accounts **Breaking Change**
- **Fix:** Update Voice Receive Mode and Fax Capability to be optional in the Toll Free AvailablePhoneNumber response payload as these are no longer returned from the API call since Programmable Fax has been disabled for some accounts **Breaking Change**

FEATURES

### Trunking

- **New Resource:** Trunk
- **New Resource:** Origination URL
- **New Resource:** Phone Number

## v0.6.0 (2020-11-22)

FIXES

### API

- **Refactor:** Rename nested response structs from Response... to ...Response to introduce a consistent naming convention across the SDK **Breaking Change**

### Autopilot

- **Refactor:** Rename nested response structs from Response... to ...Response to introduce a consistent naming convention across the SDK **Breaking Change**

### Chat

- **Fix:** Update Chat Role Permissions field from `Permission` to `Permissions`. **Breaking Change**
- **Refactor:** Rename nested response structs from Response... to ...Response to introduce a consistent naming convention across the SDK **Breaking Change**
- **Fix:** Replace map[string]interface{} with structs for Service field(s) - Limits, Media & Notifications **Breaking Change**
- **Fix:** Replace map[string]interface{} with structs for Channel Messages field(s) - Media **Breaking Change**
- **Fix:** Expand flattened service inputs fields to structs **Breaking Change**
- **Fix:** Expand flattened webhook inputs fields to structs **Breaking Change**
- **Fix:** Expand flattened channel webhook inputs fields to structs **Breaking Change**

### Conversation

- **Fix:** Expand flattened conversation inputs fields to structs **Breaking Change**
- **Fix:** Expand flattened participant inputs fields to structs **Breaking Change**
- **Fix:** Expand flattened webhook inputs fields to structs **Breaking Change**
- **Fix:** Expand flattened notification inputs fields to structs **Breaking Change**
- **Refactor:** Rename nested response structs from Response... to ...Response to introduce a consistent naming convention across the SDK **Breaking Change**

### Flex

- **Fix:** Expand flattened flex flow inputs fields to structs **Breaking Change**
- **Fix:** Remove wfm integration input and response fields from configuration operations as this cannot be configured via the api anymore **Breaking Change**
- **Refactor:** Rename nested response structs from Response... to ...Response to introduce a consistent naming convention across the SDK **Breaking Change**
- **Refactor:** Rename nested input structs from Input... to ...Input to introduce a consistent naming convention across the SDK **Breaking Change**

### Proxy

- **Refactor:** Rename nested response structs from Response... to ...Response to introduce a consistent naming convention across the SDK **Breaking Change**

### Verify

- **Fix:** Expand flattened service inputs fields to structs **Breaking Change**
- **Fix:** Expand flattened challenge inputs fields to structs **Breaking Change**
- **Fix:** Expand flattened factor inputs fields to structs **Breaking Change**
- **Refactor:** Rename nested response structs from Response... to ...Response to introduce a consistent naming convention across the SDK **Breaking Change**

FEATURES

### Conversation

- **New Resource:** Configuration
- **New Resource:** Configuration Webhook

### Verify

- **New Resource:** Service
- **New Resource:** Service Rate Limits
- **New Resource:** Service Rate Limit Bucket
- **New Resource:** Verification
- **New Resource:** Verification Check
- **New Resource:** Access Tokens
- **New Resource:** Entity
- **New Resource:** Webhook
- **New Resource:** Factor
- **New Resource:** Challenge

## v0.5.0 (2020-10-16)

FEATURES

### API

- **New Resource:** Incoming Phone Number
- **New Resource:** Available Phone Number
- **New Resource:** Available Phone Number - Toll Free
- **New Resource:** Available Phone Number - Mobile
- **New Resource:** Available Phone Number - Local

## v0.4.0 (2020-10-13)

FIXES

### Conversations

- **Fix:** Attributes json tag (on message input) was incorrectly labelled `Attributes.Filters`, this has now been corrected to `Attributes`
- **Refactor:** Rename Conversation Webhook to Webhook, this additional name provided little to no benefit so it has been removed. **Breaking Change**

FEATURES

### API

- **New Resource:** Application

### Conversations

- Add Delivery response field and struct to message resources
- **New Resource:** Roles
- **New Resource:** Users
- **New Resource:** Service
- **New Resource:** Service Configuration
- **New Resource:** Service Notification
- **New Resource:** Service Binding
- **New Resource:** Service Users
- **New Resource:** Service Role
- **New Resource:** Service Conversation
- **New Resource:** Service Conversation Webhook
- **New Resource:** Service Conversation Participant
- **New Resource:** Service Conversation Message
- **New Resource:** Service Conversation Message Delivery Receipt
- **New Resource:** Credential

### Serverless

- **New Resource:** Logs

## v0.3.1 (2020-10-03)

FIXES

### API

- **Fix:** City json tag (on response) was incorrectly labelled `City`, this has now been corrected to `city`

## v0.3.0 (2020-10-03)

Add Additional Core API clients including calls, queues conference, address and recording. Add lookup and add new build status endpoints

FIXES

### API

- **Fix:** Update message feedback client to be feedbacks to ensure consistency with the call feedback client **breaking change**

FEATURES

### TwiML

Add ToString() method to voice, messaging and fax responses to generate the TwiML string

### API

- **New Resource:** Call
- **New Resource:** Queue
- **New Resource:** Queue Member
- **New Resource:** Conference
- **New Resource:** Conference Participants
- **New Resource:** Address
- **New Resource:** Recordings
- **New Resource:** Call Recordings
- **New Resource:** Conference Recordings
- **New Resource:** Call Feedback Summary
- **New Resource:** Call Feedback

### Lookups

- **New Resource:** Phone Number

### Serverless

- **New Resource:** Build Status

## v0.2.0 (2020-09-27)

FIXES

### Flex

- **Fix:** CRM_Enabled json tag (on both input and output) was incorrectly labelled `crm_type`, this has now been corrected to `crm_enabled`

FEATURES

### TaskRouter

- **New Resource:** Worker Channel
- **New Resource:** Worker Reservation
- **New Resource:** Task Reservation
- **New Resource:** Workspace Statistics
- **New Resource:** Workflow Statistics
- **New Resource:** Worker Statistics
- **New Resource:** Task Queue Statistics

### Monitor

- **New Resource:** Alert
- **New Resource:** Events

## v0.1.0 (2020-09-06)

Add first release of the Twilio Go SDK. The list of supported services can be seen in the features section below.

FEATURES

### API

- **New Resource:** Account
- **New Resource:** Balance
- **New Resource:** Message

### Autopilot

- **New Resource:** Assistant
- **New Resource:** Defaults
- **New Resource:** StyleSheet
- **New Resource:** Field Type
- **New Resource:** Field Value
- **New Resource:** Model Build
- **New Resource:** Query
- **New Resource:** Task
- **New Resource:** Task Action
- **New Resource:** Task Field
- **New Resource:** Task Sample
- **New Resource:** Task Statistics
- **New Resource:** Webhook

### Programmable Chat

- **New Resource:** Channel
- **New Resource:** Channel Invite
- **New Resource:** Channel Member
- **New Resource:** Channel Message
- **New Resource:** Channel Webhook
- **New Resource:** Role
- **New Resource:** Credentials
- **New Resource:** Service Binding
- **New Resource:** User
- **New Resource:** User Binding
- **New Resource:** User Channel

### Conversations

- **New Resource:** Conversation
- **New Resource:** Conversation Message
- **New Resource:** Conversation Participant
- **New Resource:** Conversation Webhook
- **New Resource:** Webhook

### Programmable Fax

- **New Resource:** Fax
- **New Resource:** Media Files

### Flex

- **New Resource:** Channel
- **New Resource:** Configuration
- **New Resource:** Flex Flow
- **New Resource:** Web Channels

### Programmable Messaging

- **New Resource:** Alpha Sender
- **New Resource:** Phone Number
- **New Resource:** Service
- **New Resource:** Short Code

### Proxy

- **New Resource:** Phone Number
- **New Resource:** Service
- **New Resource:** Session
- **New Resource:** Session Interaction
- **New Resource:** Session Message Interaction
- **New Resource:** Session Participant
- **New Resource:** Short Code

### Serverless (also known as Runtime)

- **New Resource:** Asset
- **New Resource:** Asset Version
- **New Resource:** Build
- **New Resource:** Deployment
- **New Resource:** Environment
- **New Resource:** Environment Variable
- **New Resource:** Function
- **New Resource:** Function Version
- **New Resource:** Service

### Studio

- **New Resource:** Execution
- **New Resource:** Execution Context
- **New Resource:** Flow
- **New Resource:** Flow Validation
- **New Resource:** Revision
- **New Resource:** Step
- **New Resource:** Step Context
- **New Resource:** Test Users

### Sync

- **New Resource:** Document
- **New Resource:** Document Permission
- **New Resource:** Service
- **New Resource:** Sync List
- **New Resource:** Sync List Item
- **New Resource:** Sync List Permission
- **New Resource:** Sync Map
- **New Resource:** Sync Map Item
- **New Resource:** Sync Map Permission
- **New Resource:** Sync Stream
- **New Resource:** Sync Stream Message

### TaskRouter

- **New Resource:** Activity
- **New Resource:** Task
- **New Resource:** Task Channel
- **New Resource:** Task Queue
- **New Resource:** Worker
- **New Resource:** Workflow
- **New Resource:** Workspace

### TwiML

- **New Resource:** Fax Response
- **New Resource:** Messaging Response
- **New Resource:** Voice Response
