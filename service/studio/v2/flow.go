package v2

import (
	"context"
	"net/http"
	"time"

	"github.com/mitchellh/mapstructure"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type FlowService service

type Flow struct {
	client *client.Client
	sid    string
}

type MetaResponse struct {
	Page            int     `json:"page"`
	PageSize        int     `json:"page_size"`
	PreviousPageURL *string `json:"first_page_url,omitempty"`
	URL             string  `json:"url"`
	NextPageURL     *string `json:"next_page_url"`
	Key             string  `json:"key"`
}

type FlowResponse struct {
	Sid           string         `json:"sid"`
	AccountSid    string         `json:"account_sid"`
	FriendlyName  string         `json:"friendly_name"`
	Definition    interface{}    `json:"definition"`
	Status        string         `json:"status"`
	Revision      int            `json:"revision"`
	CommitMessage *string        `json:"commit_message,omitempty"`
	Valid         bool           `json:"valid"`
	Errors        *[]interface{} `json:"errors,omitempty"`
	Warnings      *[]interface{} `json:"warnings,omitempty"`
	DateCreated   time.Time      `json:"date_created,string"`
	DateUpdated   *time.Time     `json:"date_updated,string"`
	WebhookURL    string         `json:"webhook_url"`
	URL           string         `json:"url"`
}

type GetFlowPageRequest struct {
	PageSize string `mapstructure:"PageSize,omitempty"`
	Page     string `mapstructure:"Page,omitempty"`
}

type GetFlowPageResponse struct {
	Meta  MetaResponse   `json:"meta"`
	Flows []FlowResponse `json:"flows"`
}

func (service FlowService) GetPage(input *GetFlowPageRequest) (*GetFlowPageResponse, error) {
	return service.GetPageWithContext(context.Background(), input)
}

func (service FlowService) GetPageWithContext(context context.Context, input *GetFlowPageRequest) (*GetFlowPageResponse, error) {
	queryParams := map[string]string{}
	if err := mapstructure.Decode(input, &queryParams); err != nil {
		return nil, err
	}

	op := client.Operation{
		HTTPMethod:  http.MethodGet,
		HTTPPath:    PathTemplates.Flow,
		QueryParams: queryParams,
	}

	output := &GetFlowPageResponse{}
	if err := service.client.Send(context, op, nil, output); err != nil {
		return nil, err
	}
	return output, nil
}

type CreateFlowInput struct {
	FriendlyName  string `validate:"required"`
	Status        string `validate:"required"`
	Definition    string `validate:"required"`
	CommitMessage string `mapstructure:"CommitMessage,omitempty"`
}

type CreateFlowResponse struct {
	*FlowResponse
}

func (service FlowService) Create(input *CreateFlowInput) (*CreateFlowResponse, error) {
	return service.CreateWithContext(context.Background(), input)
}

func (service FlowService) CreateWithContext(context context.Context, input *CreateFlowInput) (*CreateFlowResponse, error) {
	op := client.Operation{
		HTTPMethod:  http.MethodPost,
		HTTPPath:    PathTemplates.Flow,
		ContentType: client.URLEncoded,
	}

	output := &CreateFlowResponse{}
	if err := service.client.Send(context, op, input, output); err != nil {
		return nil, err
	}
	return output, nil
}

type GetFlowResponse struct {
	*FlowResponse
}

func (flow Flow) Get() (*GetFlowResponse, error) {
	return flow.GetWithContext(context.Background())
}

func (flow Flow) GetWithContext(context context.Context) (*GetFlowResponse, error) {
	op := client.Operation{
		HTTPMethod: http.MethodGet,
		HTTPPath:   PathTemplates.FlowSID,
		PathParams: map[string]string{
			PathTemplateParamNames.SID: flow.sid,
		},
	}

	output := &GetFlowResponse{}
	if err := flow.client.Send(context, op, nil, output); err != nil {
		return nil, err
	}
	return output, nil
}

type UpdateFlowInput struct {
	FriendlyName  string `mapstructure:"FriendlyName,omitempty"`
	Status        string `validate:"required"`
	Definition    string `mapstructure:"Definition,omitempty"`
	CommitMessage string `mapstructure:"CommitMessage,omitempty"`
}

type UpdateFlowResponse struct {
	*FlowResponse
}

func (flow Flow) Update(input *UpdateFlowInput) (*UpdateFlowResponse, error) {
	return flow.UpdateWithContext(context.Background(), input)
}

func (flow Flow) UpdateWithContext(context context.Context, input *UpdateFlowInput) (*UpdateFlowResponse, error) {
	op := client.Operation{
		HTTPMethod:  http.MethodPost,
		HTTPPath:    PathTemplates.FlowSID,
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			PathTemplateParamNames.SID: flow.sid,
		},
	}

	output := &UpdateFlowResponse{}
	if err := flow.client.Send(context, op, input, output); err != nil {
		return nil, err
	}
	return output, nil
}

func (flow Flow) Delete() error {
	return flow.DeleteWithContext(context.Background())
}

func (flow Flow) DeleteWithContext(context context.Context) error {
	op := client.Operation{
		HTTPMethod: http.MethodDelete,
		HTTPPath:   PathTemplates.FlowSID,
		PathParams: map[string]string{
			PathTemplateParamNames.SID: flow.sid,
		},
	}

	if err := flow.client.Send(context, op, nil, nil); err != nil {
		return err
	}
	return nil
}
