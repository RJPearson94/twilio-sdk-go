package main

type pathParamValue struct {
	OnService bool   `json:"onService"`
	Property  string `json:"property"`
}

type pathParam struct {
	PathParamName string         `json:"pathParamName"`
	Value         pathParamValue `json:"value"`
}

type apiProperties struct {
	Name             string `json:"name"`
	Type             string `json:"type"`
	Value            string `json:"value"`
	OverrideDataType string `json:"overrideDataType"`
	Required         bool   `json:"required"`
}

type apiSchema struct {
	Name       string          `json:"name"`
	Type       string          `json:"type"`
	Properties []apiProperties `json:"properties"`
}

type property struct {
	DataType          string  `json:"dataType"`
	ParentProperty    *string `json:"parentProperty"`
	FunctionParameter *string `json:"functionParameter"`
}

type apiOperation struct {
	Name       string       `json:"name"`
	Path       string       `json:"path"`
	HTTPMethod string       `json:"httpMethod"`
	PathParams *[]pathParam `json:"pathParams,omitempty"`
	Input      *apiSchema   `json:"input"`
	Response   *apiSchema   `json:"response"`
	Service    string
	Properties map[string]property
}

type clientFunctionParameter struct {
	DataType string `json:"dataType"`
}

type clientFunction struct {
	Parameters map[string]clientFunctionParameter `json:"parameters"`
}

type client struct {
	Name       string              `json:"name"`
	Function   *clientFunction     `json:"function"`
	Properties map[string]property `json:"properties"`
	SubClients []client            `json:"subClients"`
	Operations []apiOperation      `json:"operations"`
}
