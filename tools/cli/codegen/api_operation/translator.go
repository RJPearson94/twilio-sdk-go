package apioperation

import (
	"sort"

	"github.com/Jeffail/gabs/v2"
	"github.com/iancoleman/strcase"
)

func Translate(content []byte) (*interface{}, error) {
	jsonParsed, err := gabs.ParseJSON(content)
	if err != nil {
		return nil, err
	}

	apiOperationName := jsonParsed.Path("name").Data().(string)
	packageName := jsonParsed.Path("packageName").Data().(string)

	response := gabs.New()
	response.Set(packageName, "packageName")
	response.Set(jsonParsed.Path("imports").Data(), "imports")
	if jsonParsed.Exists("documentation") {
		response.Set(jsonParsed.Path("documentation").Data(), "documentation")
	}
	response.Set(apiOperationName, "name")
	response.Set(jsonParsed.Path("http").Data(), "http")

	structures := jsonParsed.S("structures").ChildrenMap()

	if jsonParsed.Exists("http", "queryParams") {
		optionStructure := mapOptions(jsonParsed.Path("http.queryParams"), apiOperationName, packageName)
		response.Set(optionStructure.Data(), "options")
	}

	if jsonParsed.Exists("input") {
		inputStructure := mapStructure(jsonParsed.Path("input"), apiOperationName, structures)
		response.Set(inputStructure.Data(), "input")
	}

	if jsonParsed.Exists("response") {
		responseStructure := mapStructure(jsonParsed.Path("response"), apiOperationName, structures)
		response.Set(responseStructure.Data(), "response")
	}

	if jsonParsed.Exists("pagination") {
		response.Set(jsonParsed.Path("pagination").Data(), "pagination")
	}

	data := response.Data()
	return &data, nil
}

func mapStructure(structure *gabs.Container, apiOperationName string, structures map[string]*gabs.Container) *gabs.Container {
	structureResponse := gabs.New()
	nestedStructureName := structure.Path("structure").Data().(string)

	nestedStructure := structures[nestedStructureName]
	var name string
	if structure.Exists("name") {
		name = structure.Path("name").Data().(string)
	} else {
		name = nestedStructureName
	}

	structureResponse.Set(name, "name")
	structureResponse.Set(nestedStructure.Path("type").Data(), "type")
	if structure.Exists("documentation") {
		structureResponse.Set(structure.Path("documentation").Data(), "documentation")
	}

	dataType := nestedStructure.Path("type").Data().(string)
	structureResponse.Set(dataType, "type")

	properties, additionalStructs := mapProperties(nestedStructure, dataType, apiOperationName, structures)
	structureResponse.Set(properties, "properties")

	if len(additionalStructs) > 0 {
		structureResponse.Set(additionalStructs, "additionalStructs")
	}

	return structureResponse
}

func getProperties(structure *gabs.Container, structures map[string]*gabs.Container) []interface{} {
	structureProperties := structure.Path("properties").Data().([]interface{})
	properties := make([]interface{}, 0)

	properties = append(properties, structureProperties...)

	if structure.Exists("extends") {
		parentStructure := structures[structure.Path("extends").Data().(string)]
		properties = append(properties, getProperties(parentStructure, structures)...)
	}

	return properties
}

func mapOptions(structure *gabs.Container, name string, packageName string) *gabs.Container {
	structureResponse := gabs.New()
	structureResponse.Set(strcase.ToCamel(packageName)+name+"Options", "name")
	structureResponse.Set("defines the query options for the api operation", "documentation", "description")

	properties := make([]interface{}, 0)
	for _, property := range structure.Children() {
		propertiesResponse := gabs.New()
		propertiesResponse.Set(property.Path("name").Data(), "name")
		propertiesResponse.Set(property.Path("type").Data(), "type")
		propertiesResponse.Set(property.Path("required").Data(), "required")

		properties = append(properties, propertiesResponse.Data())
	}

	structureResponse.Set(properties, "properties")
	return structureResponse
}

func mapProperties(structure *gabs.Container, dataType string, apiOperationName string, structures map[string]*gabs.Container) ([]interface{}, []interface{}) {
	properties := make([]interface{}, 0)
	additionalStructs := make([]interface{}, 0)

	structureProperties := getProperties(structure, structures)
	structure.Set(structureProperties, "properties")

	for _, property := range structure.S("properties").Children() {
		if property != nil {
			propertiesResponse := gabs.New()
			propertiesResponse.Set(property.Path("value").Data(), "value")
			if property.Exists("documentation") {
				propertiesResponse.Set(property.Path("documentation").Data(), "documentation")
			}
			propertiesResponse.Set(property.Path("name").Data(), "name")
			propertiesResponse.Set(property.Path("required").Data(), "required")

			if property.Exists("validation") {
				propertiesResponse.Set(property.Path("validation").Data(), "validation")
			}

			typeName, nestedAdditionalStructs := mapType(property, dataType, apiOperationName, structures)
			propertiesResponse.Set(typeName, "type")

			for _, nestedAdditionalStruct := range nestedAdditionalStructs {
				additionalStructs = appendIfMissing(additionalStructs, nestedAdditionalStruct)
			}

			properties = append(properties, propertiesResponse.Data())
		}
	}

	sortArrayByName(properties)
	sortArrayByName(additionalStructs)

	return properties, additionalStructs
}

func mapType(property *gabs.Container, dataType string, apiOperationName string, structures map[string]*gabs.Container) (string, []interface{}) {
	var typeName string
	additionalStructs := make([]interface{}, 0)

	if property.Exists("type") {
		typeName = property.Path("type").Data().(string)
		if typeName == "array" {
			mappedTypeName, typeAdditionalStructs := mapType(property.Path("items"), dataType, apiOperationName, structures)
			return "[]" + mappedTypeName, typeAdditionalStructs
		}
		if typeName == "map" {
			mappedTypeName, typeAdditionalStructs := mapType(property.Path("items"), dataType, apiOperationName, structures)
			return "map[string]" + mappedTypeName, typeAdditionalStructs
		}
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
			additionalStructs = append(additionalStructs, nestedAdditionalStructs...)
		}
		additionalStructs = append(additionalStructs, structureResponse.Data())
	}

	return typeName, additionalStructs
}

func appendIfMissing(slice []interface{}, newValue interface{}) []interface{} {
	for _, value := range slice {
		if value.(map[string]interface{})["name"] == newValue.(map[string]interface{})["name"] {
			return slice
		}
	}
	return append(slice, newValue)
}

func sortArrayByName(array []interface{}) {
	sort.Slice(array[:], func(i, j int) bool {
		return array[i].(map[string]interface{})["name"].(string) < array[j].(map[string]interface{})["name"].(string)
	})
}
