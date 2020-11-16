package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	apiclient "github.com/RJPearson94/twilio-sdk-go/tools/cli/codegen/api_client"
	apioperation "github.com/RJPearson94/twilio-sdk-go/tools/cli/codegen/api_operation"

	"github.com/iancoleman/strcase"
)

func main() {
	var definitionPath string
	var outputPath string

	flag.StringVar(&definitionPath, "definition", "",
		"The `path` to client definition",
	)

	flag.StringVar(&outputPath, "target", "",
		"The target `path` to generate the api into",
	)

	flag.Parse()

	fixture, _ := ioutil.ReadFile(fmt.Sprintf("%s/api.json", definitionPath))
	var api map[string]interface{}
	json.Unmarshal(fixture, &api)

	if err := generateApiClients(api["subClients"].([]interface{}), api["structures"].(map[string]interface{}), outputPath); err != nil {
		fmt.Println(err)
		return
	}
}

func generateApiClients(apiClients []interface{}, structures map[string]interface{}, path string) error {
	for _, apiClient := range apiClients {
		apiClientMap := apiClient.(map[string]interface{})

		filePath := fmt.Sprintf("%s/%s", path, strcase.ToSnake(apiClientMap["name"].(string)))

		if apiClientMap["subClients"] != nil {
			if err := generateApiClients(apiClientMap["subClients"].([]interface{}), structures, filePath); err != nil {
				return err
			}
		}

		if err := translateAndGenerateApiClient(filePath, apiClientMap); err != nil {
			return err
		}

		for _, operation := range apiClientMap["operations"].([]interface{}) {
			operationMap := operation.(map[string]interface{})
			operationMap["packageName"] = apiClientMap["packageName"]
			operationMap["config"] = apiClientMap["config"]
			operationMap["structures"] = structures

			bytes, err := json.Marshal(operation)
			if err != nil {
				return err
			}
			operationResp, err := apioperation.Translate(bytes)
			if err != nil {
				return err
			}

			contents, err := apioperation.Generate(operationResp, false)
			if err != nil {
				return err
			}
			if err := createAndWriteFile(filePath, fmt.Sprintf("api_op_%s.go", strcase.ToSnake(operationMap["name"].(string))), string(*contents)); err != nil {
				return err
			}
		}
	}
	return nil
}

func translateAndGenerateApiClient(path string, content map[string]interface{}) error {
	bytes, err := json.Marshal(content)
	if err != nil {
		return err
	}
	translationResp, err := apiclient.Translate(bytes)
	if err != nil {
		return err
	}
	contents, err := apiclient.Generate(translationResp, false)
	if err != nil {
		return err
	}
	if err := createAndWriteFile(path, "api_op_client.go", string(*contents)); err != nil {
		return err
	}
	return nil
}

func createAndWriteFile(path string, fileName string, content string) error {
	os.MkdirAll(path, os.ModePerm)

	file, err := os.Create(fmt.Sprintf("%s/%s", path, fileName))
	defer file.Close()

	if err != nil {
		return fmt.Errorf("Unable to create file on disk. %s", err)
	}

	if _, err := file.WriteString(content); err != nil {
		return fmt.Errorf("Unable to write to file. %s", err)
	}
	return nil
}
