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
