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
	Service    string       `json:"service"`
	Path       string       `json:"path"`
	HTTPMethod string       `json:"httpMethod"`
	PathParams *[]pathParam `json:"pathParams,omitempty"`
	Input      *apiSchema   `json:"input"`
	Response   *apiSchema   `json:"response"`
}
