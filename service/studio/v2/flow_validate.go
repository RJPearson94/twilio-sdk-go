package v2

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type FlowValidateService service

type ValidateFlowInput struct {
	FriendlyName  string `validate:"required"`
	Status        string `validate:"required"`
	Definition    string `validate:"required"`
	CommitMessage string `mapstructure:",omitempty"`
}

type ValidateFlowResponse struct {
	Valid bool `json:"valid"`
}

func (service FlowValidateService) Validate(input *ValidateFlowInput) (*ValidateFlowResponse, error) {
	return service.ValidateWithContext(context.Background(), input)
}

func (service FlowValidateService) ValidateWithContext(context context.Context, input *ValidateFlowInput) (*ValidateFlowResponse, error) {
	op := client.Operation{
		HTTPMethod:  http.MethodPost,
		HTTPPath:    PathTemplates.FlowValidation,
		ContentType: client.URLEncoded,
	}

	output := &ValidateFlowResponse{}
	if err := service.client.Send(context, op, input, output); err != nil {
		return nil, err
	}
	return output, nil
}
