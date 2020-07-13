package apiclient

import (
	"github.com/Jeffail/gabs/v2"
)

func Translate(content []byte) (*interface{}, error) {
	jsonParsed, err := gabs.ParseJSON(content)
	if err != nil {
		return nil, err
	}

	response := gabs.New()
	response.Set(jsonParsed.Path("packageName").Data(), "packageName")
	response.Set(jsonParsed.Path("name").Data(), "name")

	if jsonParsed.Exists("properties") {
		response.Array("properties")

		for key, property := range jsonParsed.S("properties").ChildrenMap() {
			functionParamsResponse := gabs.New()
			functionParamsResponse.Set(property.Path("dataType").Data(), "type")
			functionParamsResponse.Set(key, "name")

			response.ArrayAppend(functionParamsResponse.Data(), "properties")
		}
	}

	if jsonParsed.Exists("subClients") {
		response.Array("subClients")

		for _, subClient := range jsonParsed.S("subClients").Children() {
			subClientResponse := gabs.New()
			subClientResponse.Set(subClient.Path("name").Data(), "name")

			if subClient.Exists("function") {
				subClientResponse.Array("functionParams")

				for key, property := range subClient.S("function", "parameters").ChildrenMap() {
					functionParamsResponse := gabs.New()
					functionParamsResponse.Set(property.Path("dataType").Data(), "type")
					functionParamsResponse.Set(key, "name")

					subClientResponse.ArrayAppend(functionParamsResponse.Data(), "functionParams")
				}
			}

			response.ArrayAppend(subClientResponse.Data(), "subClients")
		}
	}

	data := response.Data()
	return &data, nil
}
