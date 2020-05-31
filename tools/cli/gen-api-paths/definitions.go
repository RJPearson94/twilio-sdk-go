type PathTemplate struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type ParamNames struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type APIPathTemplates struct {
	PathTemplates []PathTemplate `json:"pathTemplates"`
	ParamNames    []ParamNames   `json:"paramNames"`
}