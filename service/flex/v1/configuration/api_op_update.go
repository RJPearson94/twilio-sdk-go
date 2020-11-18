// Package configuration contains auto-generated files. DO NOT MODIFY
package configuration

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type UpdateConfigurationInputIntegration struct {
	Active bool    `json:"active"`
	Author *string `json:"author,omitempty"`
	Config string  `validate:"required" json:"config"`
	Logo   *string `json:"logo,omitempty"`
	Name   string  `validate:"required" json:"name"`
	Type   string  `validate:"required" json:"type"`
}

type UpdateConfigurationInputSkill struct {
	Maximum    bool   `json:"maximum"`
	Minimum    int    `validate:"required" json:"minimum"`
	MultiValue bool   `json:"multivalue"`
	Name       string `validate:"required" json:"name"`
}

type UpdateConfigurationInputTaskQueue struct {
	Sid         string `validate:"required" json:"sid"`
	Targettable bool   `json:"targettable"`
}

type UpdateConfigurationInputWorkerChannel struct {
	Availability bool   `json:"availability"`
	Capacity     int    `validate:"required" json:"capacity"`
	Name         string `validate:"required" json:"name"`
}

// UpdateConfigurationInput defines input fields for updating a configuration resource
type UpdateConfigurationInput struct {
	AccountSid                   string                                              `validate:"required" json:"account_sid"`
	Attributes                   *interface{}                                        `json:"attributes,omitempty"`
	CallRecordingEnabled         *bool                                               `json:"call_recording_enabled,omitempty"`
	CallRecordingWebhookURL      *string                                             `json:"call_recording_webhook_url,omitempty"`
	ChatServiceInstanceSid       *string                                             `json:"chat_service_instance_sid,omitempty"`
	CrmAttributes                *interface{}                                        `json:"crm_attributes,omitempty"`
	CrmCallbackURL               *string                                             `json:"crm_callback_url,omitempty"`
	CrmEnabled                   *bool                                               `json:"crm_enabled,omitempty"`
	CrmFallbackURL               *string                                             `json:"crm_fallback_url,omitempty"`
	CrmType                      *string                                             `json:"crm_type,omitempty"`
	Integrations                 *[]UpdateConfigurationInputIntegration              `json:"integrations,omitempty"`
	MessagingServiceInstanceSid  *string                                             `json:"messaging_service_instance_sid,omitempty"`
	OutboundCallFlows            *interface{}                                        `json:"outbound_call_flows,omitempty"`
	PluginServiceAttributes      *interface{}                                        `json:"plugin_service_attributes,omitempty"`
	PluginServiceEnabled         *bool                                               `json:"plugin_service_enabled,omitempty"`
	PublicAttributes             *interface{}                                        `json:"public_attributes,omitempty"`
	QueueStatsConfiguration      *interface{}                                        `json:"queue_stats_configuration,omitempty"`
	ServerlessServiceSids        *[]string                                           `json:"serverless_service_sids,omitempty"`
	TaskRouterSkills             *[]UpdateConfigurationInputSkill                    `json:"taskrouter_skills,omitempty"`
	TaskRouterTargetTaskQueueSid *string                                             `json:"taskrouter_target_taskqueue_sid,omitempty"`
	TaskRouterTargetWorkflowSid  *string                                             `json:"taskrouter_target_workflow_sid,omitempty"`
	TaskRouterTaskQueues         *[]UpdateConfigurationInputTaskQueue                `json:"taskrouter_taskqueues,omitempty"`
	TaskRouterWorkerAttributes   *map[string]interface{}                             `json:"taskrouter_worker_attributes,omitempty"`
	TaskRouterWorkerChannels     *map[string][]UpdateConfigurationInputWorkerChannel `json:"taskrouter_worker_channels,omitempty"`
	UiAttributes                 *interface{}                                        `json:"ui_attributes,omitempty"`
	UiDependencies               *interface{}                                        `json:"ui_dependencies,omitempty"`
	UiLanguage                   *string                                             `json:"ui_language,omitempty"`
	UiVersion                    *string                                             `json:"ui_version,omitempty"`
	WfmIntegrations              *[]UpdateConfigurationInputIntegration              `json:"wfm_integrations,omitempty"`
}

type UpdateConfigurationIntegrationResponse struct {
	Active bool    `json:"active"`
	Author *string `json:"author,omitempty"`
	Config string  `json:"config"`
	Logo   *string `json:"logo,omitempty"`
	Name   string  `json:"name"`
	Type   string  `json:"type"`
}

type UpdateConfigurationSkillResponse struct {
	Maximum    bool   `json:"maximum"`
	Minimum    int    `json:"minimum"`
	MultiValue bool   `json:"multivalue"`
	Name       string `json:"name"`
}

type UpdateConfigurationTaskQueueResponse struct {
	Sid         string `json:"sid"`
	Targettable bool   `json:"targettable"`
}

type UpdateConfigurationWorkerChannelResponse struct {
	Availability bool   `json:"availability"`
	Capacity     int    `json:"capacity"`
	Name         string `json:"name"`
}

// UpdateConfigurationResponse defines the response fields for the updated configuration
type UpdateConfigurationResponse struct {
	AccountSid                   string                                                 `json:"account_sid"`
	Attributes                   *interface{}                                           `json:"attributes,omitempty"`
	CallRecordingEnabled         *bool                                                  `json:"call_recording_enabled,omitempty"`
	CallRecordingWebhookURL      *string                                                `json:"call_recording_webhook_url,omitempty"`
	ChatServiceInstanceSid       *string                                                `json:"chat_service_instance_sid,omitempty"`
	CrmAttributes                *interface{}                                           `json:"crm_attributes,omitempty"`
	CrmCallbackURL               *string                                                `json:"crm_callback_url,omitempty"`
	CrmEnabled                   *bool                                                  `json:"crm_enabled,omitempty"`
	CrmFallbackURL               *string                                                `json:"crm_fallback_url,omitempty"`
	CrmType                      *string                                                `json:"crm_type,omitempty"`
	DateCreated                  time.Time                                              `json:"date_created"`
	DateUpdated                  *time.Time                                             `json:"date_updated,omitempty"`
	FlexServiceInstanceSid       string                                                 `json:"flex_service_instance_sid"`
	Integrations                 *[]UpdateConfigurationIntegrationResponse              `json:"integrations,omitempty"`
	MessagingServiceInstanceSid  *string                                                `json:"messaging_service_instance_sid,omitempty"`
	OutboundCallFlows            *interface{}                                           `json:"outbound_call_flows,omitempty"`
	PluginServiceAttributes      *interface{}                                           `json:"plugin_service_attributes,omitempty"`
	PluginServiceEnabled         *bool                                                  `json:"plugin_service_enabled,omitempty"`
	PublicAttributes             *interface{}                                           `json:"public_attributes,omitempty"`
	QueueStatsConfiguration      *interface{}                                           `json:"queue_stats_configuration,omitempty"`
	RuntimeDomain                string                                                 `json:"runtime_domain"`
	ServerlessServiceSids        *[]string                                              `json:"serverless_service_sids,omitempty"`
	ServiceVersion               *string                                                `json:"service_version,omitempty"`
	Status                       string                                                 `json:"status"`
	TaskRouterOfflineActivitySid string                                                 `json:"taskrouter_offline_activity_sid"`
	TaskRouterSkills             *[]UpdateConfigurationSkillResponse                    `json:"taskrouter_skills,omitempty"`
	TaskRouterTargetTaskQueueSid string                                                 `json:"taskrouter_target_taskqueue_sid"`
	TaskRouterTargetWorkflowSid  string                                                 `json:"taskrouter_target_workflow_sid"`
	TaskRouterTaskQueues         *[]UpdateConfigurationTaskQueueResponse                `json:"taskrouter_taskqueues,omitempty"`
	TaskRouterWorkerAttributes   *map[string]interface{}                                `json:"taskrouter_worker_attributes,omitempty"`
	TaskRouterWorkerChannels     *map[string][]UpdateConfigurationWorkerChannelResponse `json:"taskrouter_worker_channels,omitempty"`
	TaskRouterWorkspaceSid       string                                                 `json:"taskrouter_workspace_sid"`
	URL                          string                                                 `json:"url"`
	UiAttributes                 *interface{}                                           `json:"ui_attributes,omitempty"`
	UiDependencies               *interface{}                                           `json:"ui_dependencies,omitempty"`
	UiLanguage                   *string                                                `json:"ui_language,omitempty"`
	UiVersion                    string                                                 `json:"ui_version"`
	WfmIntegrations              *[]UpdateConfigurationIntegrationResponse              `json:"wfm_integrations,omitempty"`
}

// Update modifies a configuration resource
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateConfigurationInput) (*UpdateConfigurationResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a configuration resource
func (c Client) UpdateWithContext(context context.Context, input *UpdateConfigurationInput) (*UpdateConfigurationResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Configuration",
		ContentType: client.JSON,
	}

	if input == nil {
		input = &UpdateConfigurationInput{}
	}

	response := &UpdateConfigurationResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
