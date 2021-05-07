package widgets

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

	name := strcase.ToCamel(jsonParsed.Path("name").Data().(string))

	response := gabs.New()
	response.Set(name, "name")
	response.Set(jsonParsed.Path("type").Data().(string), "type")
	if jsonParsed.Exists("documentation") {
		response.Set(jsonParsed.Path("documentation").Data(), "documentation")
	}

	transitions := jsonParsed.Path("transitions").Data().([]interface{})
	sortArrayByName(transitions)
	response.Set(transitions, "transitions")

	if jsonParsed.Exists("conditionalTransitions") {
		conditionalTransitions := jsonParsed.Path("conditionalTransitions").Data().([]interface{})
		sortArrayByName(conditionalTransitions)
		response.Set(conditionalTransitions, "conditionalTransitions")
	}

	structures := jsonParsed.S("structures").ChildrenMap()
	suppliedProperties := jsonParsed.Path("properties").Children()
	properties, additionalStructs := mapProperties(suppliedProperties, name, structures)
	response.Set(properties, "properties")

	if len(additionalStructs) > 0 {
		response.Set(additionalStructs, "additionalStructs")
	}

	data := response.Data()
	return &data, nil
}

func mapProperties(suppliedProperties []*gabs.Container, name string, structures map[string]*gabs.Container) ([]interface{}, []interface{}) {
	properties := []interface{}{}
	additionalStructs := make([]interface{}, 0)

	for _, property := range suppliedProperties {
		if property != nil {
			propertiesResponse := gabs.New()
			propertiesResponse.Set(property.Path("value").Data(), "value")
			propertiesResponse.Set(property.Path("name").Data(), "name")
			propertiesResponse.Set(property.Path("required").Data(), "required")

			if property.Exists("validation") {
				propertiesResponse.Set(property.Path("validation").Data(), "validation")
			}

			typeName, nestedAdditionalStructs := mapType(property, property.Path("type").Data().(string), name, structures)
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

func mapType(property *gabs.Container, dataType string, name string, structures map[string]*gabs.Container) (string, []interface{}) {
	var typeName string
	additionalStructs := make([]interface{}, 0)

	if property.Exists("type") {
		typeName = property.Path("type").Data().(string)
		if typeName == "array" {
			mappedTypeName, typeAdditionalStructs := mapType(property.Path("items"), dataType, name, structures)
			return "[]" + mappedTypeName, typeAdditionalStructs
		}
	} else if property.Exists("structure") {
		structureName := property.Path("structure").Data().(string)
		structName := name + structureName
		typeName = structName

		propertyStructure := structures[structureName]
		if propertyStructure == nil {
			panic(structureName + " cannot be found in structures")
		}

		structureResponse := gabs.New()
		structureResponse.Set(structName, "name")
		structureResponse.Set(dataType, "type")

		nestedProperties, nestedAdditionalStructs := mapProperties(propertyStructure.Path("properties").Children(), name, structures)

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
