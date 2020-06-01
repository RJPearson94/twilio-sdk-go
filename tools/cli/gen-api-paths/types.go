package main

type pathTemplate struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type paramNames struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type apiPathTemplates struct {
	PathTemplates []pathTemplate `json:"pathTemplates"`
	ParamNames    []paramNames   `json:"paramNames"`
}
