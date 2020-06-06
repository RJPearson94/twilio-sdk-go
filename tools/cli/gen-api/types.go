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
	Name     string `json:"name"`
	Type     string `json:"type"`
	Value    string `json:"value"`
	Required bool   `json:"required"`
}

type apiSchema struct {
	Name       string          `json:"name"`
	Type       string          `json:"type"`
	Properties []apiProperties `json:"properties"`
}

type apiOperation struct {
	Name       string       `json:"name"`
	Path       string       `json:"path"`
	HTTPMethod string       `json:"httpMethod"`
	PathParams *[]pathParam `json:"pathParams,omitempty"`
	Input      *apiSchema   `json:"input"`
	Response   *apiSchema   `json:"response"`
	Service    string
}

type SubClientProperties struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type client struct {
	Name       string                `json:"name"`
	Properties []SubClientProperties `json:"properties"`
	Operations []apiOperation        `json:"operations"`
}
