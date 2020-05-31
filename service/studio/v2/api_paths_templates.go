package v2

type templates struct {
	Flow           string
	FlowSID        string
	FlowValidation string
}

var PathTemplates = templates{
	Flow:           "/Flows",
	FlowSID:        "/Flows/{sid}",
	FlowValidation: "/Flows/Validate",
}

type pathParams struct {
	SID string
}

var PathTemplateParamNames = pathParams{
	SID: "sid",
}
