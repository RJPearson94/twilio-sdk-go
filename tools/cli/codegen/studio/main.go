package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/RJPearson94/twilio-sdk-go-tools/cli/codegen/studio/widgets"
	"github.com/iancoleman/strcase"
)

func main() {
	var definition string
	var outputPath string

	flag.StringVar(&definition, "definition", "",
		"The full `path` to widget definition",
	)

	flag.StringVar(&outputPath, "target", "",
		"The target `path` to generate the widget into",
	)

	flag.Parse()

	fixture, _ := ioutil.ReadFile(definition)

	var widget map[string]interface{}
	json.Unmarshal(fixture, &widget)

	if err := generateWidget(widget, outputPath); err != nil {
		fmt.Println(err)
		return
	}
}

func generateWidget(widgetMap map[string]interface{}, path string) error {
	bytes, err := json.Marshal(widgetMap)
	if err != nil {
		return err
	}
	translationResp, err := widgets.Translate(bytes)
	if err != nil {
		return err
	}
	contents, err := widgets.Generate(translationResp, false)
	if err != nil {
		return err
	}

	if err := createAndWriteFile(path, fmt.Sprintf("%s.go", strcase.ToSnake(widgetMap["name"].(string))), string(*contents)); err != nil {
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
