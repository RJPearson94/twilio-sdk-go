package apiclient

import (
	"sort"

	"github.com/Jeffail/gabs/v2"
)

func Translate(content []byte) (*interface{}, error) {
	jsonParsed, err := gabs.ParseJSON(content)
	if err != nil {
		return nil, err
	}

	response := gabs.New()
	response.Set(jsonParsed.Path("packageName").Data(), "packageName")
	if jsonParsed.Exists("config") {
		response.Set(jsonParsed.Path("config").Data(), "config")
	}
	if jsonParsed.Exists("documentation") {
		response.Set(jsonParsed.Path("documentation").Data(), "documentation")
	}
	response.Set(jsonParsed.Path("name").Data(), "name")

	if jsonParsed.Exists("properties") {
		response.Array("properties")

		propertiesMap := jsonParsed.S("properties").ChildrenMap()
		propertiesMapKeys := sortMapKeys(propertiesMap)

		for _, key := range propertiesMapKeys {
			functionParamsResponse := gabs.New()
			functionParamsResponse.Set(propertiesMap[key].Path("dataType").Data(), "type")
			functionParamsResponse.Set(key, "name")

			response.ArrayAppend(functionParamsResponse.Data(), "properties")
		}
	}

	if jsonParsed.Exists("subClients") {
		response.Array("subClients")

		subClients := jsonParsed.S("subClients").Children()
		sortArrayByName(subClients)

		for _, subClient := range subClients {
			subClientResponse := gabs.New()
			subClientResponse.Set(subClient.Path("name").Data(), "name")
			subClientResponse.Set(subClient.Path("packageName").Data(), "packageName")

			if subClient.Exists("function") {
				subClientResponse.Array("functionParams")

				parameterMap := subClient.S("function", "parameters").ChildrenMap()
				parameterMapKeys := sortMapKeys(parameterMap)

				for _, key := range parameterMapKeys {
					parameter := parameterMap[key]
					functionParamsResponse := gabs.New()
					functionParamsResponse.Set(parameter.Path("dataType").Data(), "type")
					functionParamsResponse.Set(key, "name")

					subClientResponse.ArrayAppend(functionParamsResponse.Data(), "functionParams")
				}
			}

			if subClient.Exists("properties") {
				subClientResponse.Array("properties")

				propertiesMap := subClient.S("properties").ChildrenMap()
				propertiesMapKeys := sortMapKeys(propertiesMap)

				for _, key := range propertiesMapKeys {
					property := propertiesMap[key]
					propertiesResponse := gabs.New()
					propertiesResponse.Set(property.Path("dataType").Data(), "type")
					propertiesResponse.Set(key, "name")

					if property.Exists("parentProperty") {
						propertiesResponse.Set(property.Path("parentProperty").Data(), "parentProperty")
					}

					if property.Exists("functionParameter") {
						propertiesResponse.Set(property.Path("functionParameter").Data(), "functionParameter")
					}

					subClientResponse.ArrayAppend(propertiesResponse.Data(), "properties")
				}
			}

			response.ArrayAppend(subClientResponse.Data(), "subClients")
		}
	}

	data := response.Data()
	return &data, nil
}

func sortMapKeys(data map[string]*gabs.Container) []string {
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func sortArrayByName(array []*gabs.Container) {
	sort.Slice(array[:], func(i, j int) bool {
		return array[i].Path("name").Data().(string) < array[j].Path("name").Data().(string)
	})
}
