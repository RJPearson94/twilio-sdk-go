package apioperation

import (
	"github.com/Jeffail/gabs/v2"
)

func Translate(content []byte) (*interface{}, error) {
	jsonParsed, err := gabs.ParseJSON(content)
	if err != nil {
		return nil, err
	}

	apiOperationName := jsonParsed.Path("name").Data().(string)

	response := gabs.New()
	response.Set(jsonParsed.Path("packageName").Data(), "packageName")
	response.Set(jsonParsed.Path("imports").Data(), "imports")
	response.Set(apiOperationName, "name")
	response.Set(jsonParsed.Path("http").Data(), "http")

	structures := jsonParsed.S("structures").ChildrenMap()

	if jsonParsed.Exists("input") {
		inputStructure := mapStructure(jsonParsed.Path("input"), jsonParsed.Path("name").Data().(string), apiOperationName, structures)
		response.Set(inputStructure.Data(), "input")
	}

	if jsonParsed.Exists("response") {
		responseStructure := mapStructure(jsonParsed.Path("response"), jsonParsed.Path("name").Data().(string), apiOperationName, structures)
		response.Set(responseStructure.Data(), "response")
	}

	data := response.Data()
	return &data, nil
}

func mapStructure(structure *gabs.Container, structureName string, apiOperationName string, structures map[string]*gabs.Container) *gabs.Container {
	structureResponse := gabs.New()
	nestedStructure := structure.Path("structure").Data().(string)

	responseStructure := structures[nestedStructure]
	var name string
	if structure.Exists("name") {
		name = structure.Path("name").Data().(string)
	} else {
		name = structureName + nestedStructure
	}

	structureResponse.Set(name, "name")
	structureResponse.Set(responseStructure.Path("type").Data(), "type")

	dataType := responseStructure.Path("type").Data().(string)
	structureResponse.Set(dataType, "type")

	properties, additionalStructs := mapProperties(responseStructure, dataType, apiOperationName, structures)
	structureResponse.Set(properties, "properties")

	if len(additionalStructs) > 0 {
		structureResponse.Set(additionalStructs, "additionalStructs")
	}

	return structureResponse
}

func mapProperties(structure *gabs.Container, dataType string, apiOperationName string, structures map[string]*gabs.Container) ([]interface{}, []interface{}) {
	properties := make([]interface{}, 0)
	additionalStructs := make([]interface{}, 0)

	for _, property := range structure.S("properties").Children() {
		propertiesResponse := gabs.New()
		propertiesResponse.Set(property.Path("value").Data(), "value")
		propertiesResponse.Set(property.Path("name").Data(), "name")
		propertiesResponse.Set(property.Path("required").Data(), "required")

		typeName, nestedAdditionalStructs := mapType(property, dataType, apiOperationName, structures)
		propertiesResponse.Set(typeName, "type")

		additionalStructs = append(additionalStructs, nestedAdditionalStructs...)

		properties = append(properties, propertiesResponse.Data())
	}

	return properties, additionalStructs
}

func mapType(property *gabs.Container, dataType string, apiOperationName string, structures map[string]*gabs.Container) (string, []interface{}) {
	var typeName string
	additionalStructs := make([]interface{}, 0)

	if property.Exists("type") {
		if property.Path("type").Data().(string) == "array" {
			mappedTypeName, typeAdditionalStructs := mapType(property.Path("items"), dataType, apiOperationName, structures)
			return "[]" + mappedTypeName, typeAdditionalStructs
		}

		typeName = property.Path("type").Data().(string)
	} else if property.Exists("structure") {
		structureName := property.Path("structure").Data().(string)
		structName := apiOperationName + property.Path("structure").Data().(string)
		typeName = structName

		propertyStructure := structures[structureName]

		structureResponse := gabs.New()
		structureResponse.Set(structName, "name")
		structureResponse.Set(dataType, "type")

		nestedProperties, nestedAdditionalStructs := mapProperties(propertyStructure, dataType, apiOperationName, structures)
		structureResponse.Set(nestedProperties, "properties")

		if len(nestedAdditionalStructs) > 0 {
			additionalStructs = append(additionalStructs, nestedAdditionalStructs)
		}
		additionalStructs = append(additionalStructs, structureResponse.Data())
	}

	return typeName, additionalStructs
}
